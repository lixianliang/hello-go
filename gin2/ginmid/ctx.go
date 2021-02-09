package main

import (
	"context"
	"fmt"
)

// 定义全局上下文中的key
type (
	transCtx    struct{}
	userIDCtxt  struct{}
	traceIDCtxt struct{}
)

// 穿件事务的上下文
func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// 从上下文中获取事务
func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx)
	return v, v != nil
}

// 创建用户ID的上下文
func NewUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDCtxt{}, userID)
}

// 从上下文中获取用户ID
func FromUserID(ctx context.Context) (string, bool) {
	v := ctx.Value(userIDCtxt{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s, s != ""
		}
	}
	return "", false
}

// 创建traceID的上下文
func NewTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtxt{}, traceID)
}

// 从上下文中获取追踪ID
func FromTraceID(ctx context.Context) (string, bool) {
	v := ctx.Value(traceIDCtxt{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s, s != ""
		}
	}
	return "", false
}

func main() {

}
