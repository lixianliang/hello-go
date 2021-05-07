package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	//"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	//ZIPKIN_HTTP_ENDPOINT      = "http://127.0.0.1:9411/api/v2/spans"
	ZIPKIN_HTTP_ENDPOINT = "http://106.14.240.73:9411/api/v2/spans"
	//ZIPKIN_RECORDER_HOST_PORT = "127.0.0.1:9000"
	ZIPKIN_RECORDER_HOST_PORT = ""
)

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
	tracer := Init("printserver")
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("header: %v", r.Header)
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Printf("Extract failed: %v", err)
			return
		}
		span := tracer.StartSpan("publish", opentracing.ChildOf(spanCtx))
		//	span := tracer.StartSpan("publish", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		helloTo := r.FormValue("helloStr")
		helloStr := fmt.Sprintf("print hello %s", helloTo)
		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", helloStr),
		)
		time.Sleep(10 * time.Millisecond)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
