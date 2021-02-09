package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	//"github.com/ru-rocker/gokit-playground/lorem-logging"
	"github.com/lixianliang/hello-go/gokit-playground/lorem"
	"github.com/lixianliang/hello-go/gokit-playground/lorem-logging"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	logfile, err := os.OpenFile("./log/golorem.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	// Logging domain
	var logger log.Logger
	{
		w := log.NewSyncWriter(logfile)
		logger = log.NewLogfmtLogger(w)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc lorem.Service
	svc = lorem.LoremService{}
	svc = logging.LoggingMiddleware(logger)(svc)
	endpoint := lorem.Endpoints{
		LoremEndpoint: lorem.MakeLoremEndpoint(svc),
	}

	r := lorem.MakeHttpHandler(ctx, endpoint, logger)

	// HTTP transport
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
