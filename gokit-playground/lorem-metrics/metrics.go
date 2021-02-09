package metrics

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/lixianliang/hello-go/gokit-playground/lorem"
)

func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) lorem.ServiceMiddleware {
	return func(next lorem.Service) lorem.Service {
		return metricsMiddleware{
			next,
			requestCount,
			requestLatency,
		}
	}
}

// Make a new type and wrap into Service interface
// Add expected metrics property to this type
type metricsMiddleware struct {
	lorem.Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Implement service functions and add label method for our metrics
func (mw metricsMiddleware) Word(min, max int) (output string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Word"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output = mw.Service.Word(min, max)
	return
}

func (mw metricsMiddleware) Sentence(min, max int) (output string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Sentence"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output = mw.Service.Sentence(min, max)
	return
}

func (mw metricsMiddleware) Paragraph(min, max int) (output string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Paragraph"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output = mw.Service.Paragraph(min, max)
	return
}
