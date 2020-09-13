package middleware

import (
	"context"
	"net/http"

	"github.com/jaegertracing/jaeger-client-go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func TraceSpan(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()
		if tracer == nil {
			// Tracer not found, just skip
			next.ServeHTTP(w, f)
		}

		// Start span
		spanName := "HTTP " + r.Method
		span := opentracing.StartSpan(spanName)
		//ext.SpanKindRPCServer
		rc := opentracing.ContextWithSpan(r.Context(), span)
		// Set request ID for context
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			//rc = context.WithValue(rc, constants.RequestID, sc.TraceID().String())
			rc = context.WithValue(rc, "X-Request-ID", sc.TraceID().String())
		}

		next.ServeHTTP(w, r.WithContext(rc))

		// Finish span
		wrapper, ok := w.(WrapResponseWriter)
		if ok {
			ext.HTTPStatusCode.Set(span, uint16(wrapper.Status()))
		}
		span.Finish()
	}

	return http.HandlerFunc(fn)
}

func SetUp() gin.HandlerFunc {
	// gin http拦截
	return func(c *gin.Context) {
		tracer, closer := jaeger_trace.NewJaegerTracer(config.AppName, config.JaegerHostPort)
		defer closer.Close()

		var parentSpan opentracing.Span

		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}

		// 感觉没必要存储起来
		c.Set("Tracer", tracer)
		c.Set("ParentSpanContext", parentSpan.Context())

		c.Next()
	}
}

// http client注入
func x() {
	injectErr := jaeger.Tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	if injectErr != nil {
		log.Fatalf("%s: Couldn't inject headers", err)
	}
}
