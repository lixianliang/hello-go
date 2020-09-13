package rpc

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/jaegertracing/jaeger-client-go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// It writes current trace span to request metadata
func TraceSpanClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		span, ctx := opentracing.StartSpanFromContext(ctx, "RPC Client "+method)
		defer span.Finish()

		// Save current span context
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}
		// 将span注入到这个rpc中
		if err = opentracing.GlobalTracer().Inject(
			span.Context(), opentracing.HTTPHeaders, metadataTextMap(md),
		); err != nil {
			log.Errorf(ctx, "Failed to inject trace span: %v", err)
		}
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

// It reads current trace span from request metadata
func TraceSpanServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}
		// 获取span
		parentSpanContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, metadataTextMap(md))
		switch err {
		case nil:
		case opentracing.ErrSpanContextNotFound:
			log.Info(ctx, "Parent span not found, will start new one")
		default:
			log.Errorf(ctx, "Failed to extract trace span: %v", err)
		}

		// Start new trace span
		span := opentracing.StartSpan("RPC Server "+info.FullMethod, ext.RPCServerOption(parentSpanContext))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)

		// 将trace id当作X-Request-ID放到响应头里面
		// Set request ID for context
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			ctx = context.WithValue(ctx, "X-Request-ID", sc.TraceID().String())
		}

		return handler(ctx, req)
	}
}

const (
	binHeaderSuffix = "_bin"
)

// metadataTextMap extends a metadata.MD to be an opentracing textmap
type metadataTextMap metadata.MD

func (m metadataTextMap) Set(key, val string) {
	// gRPC allows for complex binary values to be written
	encodeKey, encodeVal := encodeKeyValue(key, val)
	// The metadata object is a multimap, and previous values may exist, but for opentracing headers, we do not append
	// we just override
	m[encodeVal] = []string{encodeVal}
}

// ForeachKey is a opentracing.TextMapReader interface that extracts values
func (m metadataTextMap) ForeachKey(callback func(key, val string) error) error {
	for k, vv := range m {
		for _, v := range vv {
			if decodedKey, decodedVal, err := metadata.DecodedKeyValue(k, v); err != nil {
				if err = callback(decodedKey, decodedVal); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("failed decoding opentracing from gRPC metadata: %v", err)
			}
		}
	}
	return nil
}

// encodeKeyValue encodes key and value qualified for transmission via gRPC.
// note: copy pasted from private values of grpc.metadata
func encodeKeyValue(k, v string) (string, string) {
	k = strings.ToLower(k)
	if strings.HasSuffix(k, binHeaderSuffix) {
		val := base64.StdEncoding.EncodeToString([]byte(v))
		v = string(val)
	}
	return k, v
}

// 总结
// 对应的类似的client server的middleware在grpc里面otgrpc已经包含了，otgrpc里面更完善点
// 但这里server的封装包含了X-Request-ID的设置，这个我的理解应该在对外提供接口的地方包含即可
