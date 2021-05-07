package main

import (
	"context"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"

	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
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

func InitJager(service string) opentracing.Tracer {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "139.224.83.59:6831",
			LogSpans:           true,
		},
	}

	tracer, _, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer
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
	//tracer := Init("hello-world")
	tracer := InitJager("hello-world")
	opentracing.SetGlobalTracer(tracer)

	Abc()
	time.Sleep(time.Second * 1)

}

func Abc() {
	helloTo := "lxl"
	//span := tracer.StartSpan("rrrr") // root span
	span := opentracing.StartSpan("rrrr")
	defer span.Finish()
	//span := tracer.StartSpan("rrrr", opentracing.ChildOf(nil)) // root span
	// span, ctx := tracing.StartRootSpanByContext(context.Background(), "apiHandler")
	span.SetTag("hello-to", helloTo)
	span.LogFields(
		otlog.String("event-1", "string-format"),
		otlog.String("value-1", "xx"),
	)

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	helloStr := formatString(ctx, helloTo)
	log.Printf("last response: %v", helloStr)

	printHello(ctx, "testprint")
}

func formatString(ctx context.Context, helloTo string) string {
	//span, _ := opentracing.StartSpanFromContext(ctx, "formatString-client")
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString-client")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloTo", helloTo)
	url := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	log.Printf("access server header: %v", req.Header)

	resp, err := Do(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err.Error())
	}

	helloStr := string(resp)

	span.LogFields(
		otlog.String("event", "string-format"),
		otlog.String("value", helloStr),
	)

	return helloStr
}

// Do executes an HTTP request and returns the response body.
// Any errors or non-200 status code result in an error.
func Do(req *http.Request) ([]byte, error) {
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

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello-client")
	//span, _ := opentracing.StartSpanFromContext(opentracing.ChildOf(ctx), "printHello-client")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloStr", helloStr)
	url := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	log.Printf("access server header: %v", req.Header)

	if _, err := Do(req); err != nil {
		ext.LogError(span, err)
		panic(err.Error())
	}
}
