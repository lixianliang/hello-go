package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	ratelimitkit "github.com/go-kit/kit/ratelimit"
	//"github.com/juju/ratelimit"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/lixianliang/hello-go/gokit-playground/lorem"
	"github.com/lixianliang/hello-go/gokit-playground/lorem-ratelimit"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	// Logging domain
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// delclare metrics
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "ru_rocker",
		Subsystem: "lorem_service",
		Name:      "request_count",
		Help:      "Number of requests received",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "ru_rocker",
		Subsystem: "lorem_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds",
	}, fieldKeys)

	var svc lorem.Service
	svc = lorem.LoremService{}
	svc = ratelimit.LoggingMiddleware(logger)(svc)
	svc = ratelimit.Metrics(requestCount, requestLatency)(svc)

	e := lorem.MakeLoremEndpoint(svc)
	limit := rate.NewLimiter(rate.Every(35*time.Millisecond), 100)
	e = ratelimitkit.NewErroringLimiter(limit)(e)
	//rlbucket := ratelimit.NewBucket(1*time.Second, 5)
	//e = ratelimitkit.NewTokenBucketThrottler(rlbucket, time.Sleep)(e)

	// 添加限流中间件, 1s间隔, 桶中3个令牌
	//limiter := rate.NewLimiter(1, 3)
	//e = ratelimitkit.NewErroringLimiter(limiter)(e)
	endpoint := lorem.Endpoints{
		LoremEndpoint: e,
	}

	r := lorem.MakeHttpHandler(ctx, endpoint, logger)

	// HTTp transport
	go func() {
		fmt.Println("Starting server at port 8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
