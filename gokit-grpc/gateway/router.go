package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
	"github.com/openzipkin/zipkin-go"
	zipkinhttpsvr "github.com/openzipkin/zipkin-go/middleware/http"
)

// HystrixRouter hystrix路由
type HystrixRouter struct {
	svcMap       *sync.Map      // 服务实例，存储已经通过hystrix监控服务列表
	logger       log.Logger     // 日志工具
	fallbackMsg  string         // 回调消息
	consulClient *api.Client    //consul客户端对象
	tracer       *zipkin.Tracer //服务追踪对象
}

func Routes(client *api.Client, zipkinTracker *zipkin.Tracer, fbMsg string, logger log.Logger) http.Handler {
	return HystrixRouter{
		svcMap:       &sync.Map{},
		logger:       logger,
		fallbackMsg:  fbMsg,
		consulClient: client,
		tracer:       zipkinTracker,
	}
}

func (router HystrixRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path
	if reqPath == "" {
		return
	}

	// 按照分隔符/对路径进行分解，获取服务名称service nanme
	pathArray := strings.Split(reqPath, "/")
	serviceName := pathArray[1]

	// 经检查是否已经加入监控
	if _, ok := router.svcMap.Load(serviceName); !ok {
		// 把serviceName作为命令对象 设置参数
		hystrix.ConfigureCommand(serviceName, hystrix.CommandConfig{Timeout: 1000})
		router.svcMap.Store(serviceName, serviceName)
	}

	// 执行命令
	err := hystrix.Do(serviceName, func() (err error) {
		// consul api查询serviceName
		result, _, err := router.consulClient.Catalog().Service(serviceName, "", nil)
		if err != nil {
			router.logger.Log("ReverseProxy failed", "query service instance error", err.Error())
			return
		}

		if len(result) == 0 {
			router.logger.Log("ReverseProxy failed", "no such service instance", serviceName)
			return errors.New("no such service instance")
		}

		director := func(req *http.Request) {
			destPath := strings.Join(pathArray[2:], "/")
			// 随机选择一个服务实例
			tgt := result[rand.Int()%len(result)]
			router.logger.Log("service id", tgt.ServiceID)

			// 设置代理服务地址信息
			req.URL.Scheme = "http"
			req.URL.Host = fmt.Sprintf("%s:%d", tgt.ServiceAddress, tgt.ServicePort)
			req.URL.Path = "/" + destPath
		}

		var proxyError error = nil
		// 为反向代理增加追踪逻辑，使用RoundTrip代替默认Transport
		roundTrip, _ := zipkinhttpsvr.NewTransport(router.tracer, zipkinhttpsvr.TransportTrace(true))
		errorHandler := func(ew http.ResponseWriter, er *http.Request, err error) {
			proxyError = err
		}

		proxy := &httputil.ReverseProxy{
			Director:     director,
			Transport:    roundTrip,
			ErrorHandler: errorHandler,
		}
		proxy.ServeHTTP(w, r)
		return proxyError
	}, func(err error) error {
		router.logger.Log("fallback error description", err.Error())
		return errors.New(router.fallbackMsg)
	})

	// Do方法执行失败，响应错误信息
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
