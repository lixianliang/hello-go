package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	//	"strconv"
	"time"

	_ "github.com/opentracing/basictracer-go"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	//"github.com/openzipkin/zipkin-go/model"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

const (
	//ZIPKIN_HTTP_ENDPOINT      = "http://127.0.0.1:9411/api/v2/spans"
	ZIPKIN_HTTP_ENDPOINT = "http://106.14.240.73:9411/api/v2/spans"
	//ZIPKIN_RECORDER_HOST_PORT = "127.0.0.1:9000"
	ZIPKIN_RECORDER_HOST_PORT = ""
)

func InitJager(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			//Type: "ratelimiting",
			//Param: 1,
			Param: 1,
		},
		/*Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "139.224.83.59:6831",
			LogSpans:           true,
			//BufferFlushInterval: 1 * time.Second,
			//CollectorEndpoint:   "http://139.196.20.175:14268/api/v2/spans",
		},*/
	}
	addr := "139.224.83.59:6831"
	//addr := "jaeger-agent.istio-system:6831"
	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, _, err := cfg.NewTracer(
		config.Reporter(reporter),
	)
	if err != nil {
		panic(fmt.Sprintf("falied NewTracer: %v\n", err))
	}

	/*tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}*/
	/*tracer, _, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}*/
	return tracer, nil
}

func Init(service string) opentracing.Tracer {

	reporter := zipkinhttp.NewReporter(ZIPKIN_HTTP_ENDPOINT)
	//defer reporter.Close()
	endpoint, err := zipkin.NewEndpoint(service, ZIPKIN_RECORDER_HOST_PORT)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}
	tracer := zipkinot.Wrap(nativeTracer)

	return tracer
}

func main() {
	//tracer := Init("testclient")
	tracer, _ := InitJager("testclient")
	// defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handler header: %v", r.Header)
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Printf("Extract failed: %v", err)
			//return
		}
		span := tracer.StartSpan("format-server", ext.RPCServerOption(spanCtx))
		//span := tracer.StartSpan("format-server", ext.RPCServerOption(spanCtx), opentracing.ChildOf(spanCtx)) // zipkin
		//span := tracer.StartSpan("format-server", opentracing.ChildOf(spanCtx)) // ok
		defer span.Finish()

		// spanCtx.Get(traceIdKey)
		//sc, ok := span.Context().(basictracer.SpanContext) // 不起作用 Uber-Trace-Id的方式也是不行
		// sc, ok := span.Context().(model.SpanContext) 编译不过
		//log.Printf("get SpanContext: %v  %v %v", ok, sc, sc.TraceID.String())
		sc, ok := span.Context().(jaeger.SpanContext) // uber的方式会报错Extract failed: opentracing: SpanContext not found in Extract carrier
		log.Printf("get SpanContext: %v %v", sc, ok)  //  model.SpanContext does not implement opentracing.SpanContext (missing ForeachBaggageItem method) 没有实现
		if ok {
			//	_ = strconv.FormatUint(sc.TraceID, 10)
			log.Printf("trace: %v", sc.TraceID())
		} else {
			log.Printf("not get SpanContext")
		}

		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("hello ++++ %s", helloTo)
		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", helloStr),
		)
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		accessWeb(ctx)
		//time.Sleep(1000 * time.Millisecond)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func accessWeb(ctx context.Context) string {
	// 通过抓包有数据上传 access-web 当做name serive还是test clie
	span, _ := opentracing.StartSpanFromContext(ctx, "access-web")
	defer span.Finish()

	url := "http://172.16.3.9/algo/jimu/v1beta/task?id=578472881021747201"
	//url := "http://139.196.28.4/algo/jimu/v1beta/task?id=578472881021747201"
	//url := "https://baidu.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(100 * time.Millisecond)
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	log.Printf("access header: %v", req.Header)

	_, err = Do(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err.Error())
	}

	helloStr := string("ok===ok")

	span.LogFields(
		otlog.String("event", "string-format"),
		otlog.String("value", helloStr),
	)

	return helloStr
}

// Do executes an HTTP request and returns the response body.
// Any errors or non-200 status code result in an error.
func Do(req *http.Request) ([]byte, error) {
	time.Sleep(10 * time.Millisecond)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	return body, nil
}
