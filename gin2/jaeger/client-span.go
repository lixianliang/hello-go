package main

import (
	"fmt"
	"io"
	//"log"
	"time"

	//"github.com/jaegertracing/jaeger-client-go"
	//jaegercfg "github.com/jaegertracing/jaeger-client-go/config"
	//"github.com/opentracing/opentracing-go"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(svc string) (opentracing.Tracer, io.Closer) {
	cfg, err := jaegercfg.FromEnv()

	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LocalAgentHostPort = "127.0.0.1:6831"
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.New(svc, jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {
	tracer, closer := InitJaeger("hello-world")
	defer closer.Close()
	opentracing.InitGlobalTracer(tracer) // 初始化，后面其它的地方可以用

	helloStr := "hello jaeger"
	span := tracer.StartSpan("say-hello")
	fmt.Println(span)
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println(helloStr)
	span.SetTag("value", helloStr) // 类似打tag的方式
	//span.LogFields(
	//	log.String("event", "sayhello"),
	//	log.String("value", helloStr),
	//)

	span.Finish()
}
