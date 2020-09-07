package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	Service
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return loggingMiddleware{next, logger}
	}
}

func (mw loggingMiddleware) Calculate(ctx context.Context, reqType string, a, b int) (ret int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Calculate",
			"request_type", reqType,
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(begin),
		)
	}(time.Now())

	ret, err = mw.Service.Calculate(ctx, reqType, a, b)
	return
}

func (mw loggingMiddleware) HealthCheck() (result bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "HealthChcek",
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.Service.HealthCheck()
	return
}

func (mw loggingMiddleware) Login(name, pwd string) (token string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Login",
			"result", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	token, err = mw.Service.Login(name, pwd)
	return
}
