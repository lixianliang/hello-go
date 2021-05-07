package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	//jlog "github.com/uber/jaeger-client-go/log"
	zipkin "github.com/uber/jaeger-client-go/zipkin"
	//config "github.com/uber/jaeger-client-go/config"
	xzipkin "github.com/uber/jaeger-client-go/transport/zipkin"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	injector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
	extractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)

	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinSharedRPCSpan := jaeger.TracerOptions.ZipkinSharedRPCSpan(true)

	sender, _ := jaeger.NewUDPTransport("jaeger-agent.istio-system:5775", 0)
	/*sender, err := xzipkin.NewHTTPTransport(
		"http://106.14.240.73:9411/api/v2/spans",
		//xzipkin.HTTPBatchSize(10),
		//xzipkin.HTTPLogger(jlog.StdLogger),
	)
	if err != nil {
		log.Fatalf("Cannot initialize Zipkin HTTP transport: %v", err)
	}

	tracer, closer := jaeger.NewTracer(
		"myhelloworld",
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(
			sender,
			jaeger.ReporterOptions.BufferFlushInterval(1*time.Second)),
		injector,
		extractor,
		zipkinSharedRPCSpan,
	)*/
	/*transport, err := zipkin.NewHTTPTransport(
		"http://106.14.240.73:9411/api/v2/spans",
	//	zipkin.HTTPBatchSize(10),
	//	zipkin.HTTPLogger(jlog.StdLogger),
	)
	if err != nil {
		log.Fatalf("Cannot initialize Zipkin HTTP transport: %v", err)
	}
	tracer, closer := jaeger.NewTracer(
		service,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(transport, nil),
	)*/
	/*cfg := &config.Configuration{
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

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}*/
	//tracer, closer := jaeger.NewTracer(service, jaeger.NewConstSampler(true), jaeger.NewNullReporter(), jaeger.TracerOptions.ZipkinSharedRPCSpan(true))

	return tracer, closer
}

func main() {
	tracer, closer := Init("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	{
		helloTo := "lxl"
		span := tracer.StartSpan("say-hello")
		span.SetTag("hello-to", helloTo)
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(), span)

		helloStr := formatString(ctx, helloTo)
		log.Printf("last response: %v", helloStr)
	}

	time.Sleep(time.Second * 2)

}

func formatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
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
	// "zipkin-span-format"
	span.Tracer().Inject(
		span.Context(),
		//"zipkin-span-format",
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

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
