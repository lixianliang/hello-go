package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/raysonxin/gokit-article-demo/arithmetic_grpc_demo/pb"
	"google.golang.org/grpc"

	//"github.com/micro/util/go/lib/net"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/raysonxin/gokit-article-demo/arithmetic_grpc_demo/service"
	"github.com/raysonxin/gokit-article-demo/arithmetic_grpc_demo/service/transport"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		consulHost  = flag.String("consul.host", "localhost", "consul ip address")
		consulPort  = flag.String("consul.port", "8500", "consul port")
		serviceHost = flag.String("service.host", "192.168.192.146", "service ip address")
		servicePort = flag.String("service.port", "9000", "service port")
		zipkinURL   = flag.String("zipkin.url", "http://localhost:9411/api/v2/spans", "Zipkin server url")
		grpcAddr    = flag.String("grpc", ":9001", "gRPC listen address.")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "raysonxin",
		Subsystem: "arithmetic_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "raysonxin",
		Subsystem: "arithemetic_service",
		Name:      "request_latency",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var zipkinTracer *zipkin.Tracer
	{
		var (
			err           error
			hostPort      = *serviceHost + ":" + *servicePort
			serviceName   = "arithmetic-service"
			useNoopTracer = (*zipkinURL == "")
			reporter      = zipkinhttp.NewReporter(*zipkinURL)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracer, err = zipkin.NewTracer(
			reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(useNoopTracer),
		)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if !useNoopTracer {
			logger.Log("tracer", "Zipkin", "type", "Native", "URL", *zipkinURL)
		}
	}

	ratebucket := rate.NewLimiter(rate.Every(time.Second*1), 100)

	var svc service.Service
	svc = service.ArithmeticService{}

	svc = service.LoggingMiddleware(logger)(svc)
	svc = service.Metrics(requestCount, requestLatency)(svc)

	calEndpoint := service.MakeArithmeticEndpoint(svc)
	calEndpoint = service.NewTokenBucketLimitterWithBuildIn(ratebucket)(calEndpoint)
	calEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "calculate-endpoint")(calEndpoint)

	healthEndpoint := service.MakeHealthCheckEndpoint(svc)
	healthEndpoint = service.NewTokenBucketLimitterWithBuildIn(ratebucket)(healthEndpoint)
	healthEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "health-endpoint")(healthEndpoint)

	//身份认证Endpoint
	authEndpoint := service.MakeAuthEndpoint(svc)
	authEndpoint = service.NewTokenBucketLimitterWithBuildIn(ratebucket)(authEndpoint)
	authEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "login-endpoint")(authEndpoint)

	endpts := service.ArithmeticEndpoints{
		CalculateEndpoint:   calEndpoint,
		HealthCheckEndpoint: healthEndpoint,
		AuthEndpoint:        authEndpoint,
	}

	register := service.Register(*consulHost, *consulPort, *serviceHost, *servicePort, logger)

	go func() {
		fmt.Println("http server start at port:", *servicePort)
		handler := transport.MakeHttpHandler(ctx, endpts, zipkinTracer, logger)
		register.Register()
		errChan <- http.ListenAndServe(":"+*servicePort, handler)
	}()

	go func() {
		lis, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errChan <- err
			return
		}

		handler := transport.NewGRPCServer(ctx, endpts)
		gRPCServer := grpc.NewServer()
		pb.RegisterArithmeticServiceServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(lis)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	register.Deregister()
	fmt.Println(error)
}
