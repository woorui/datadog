package datadog

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Counter = (*counter)(nil)

type counter struct {
	client     *statsd.Client
	metricName string
	tempTags   []string
}

// NewDDCounter new a DataDog counter and returns Counter.
func NewDDCounter(metricName string) metrics.Counter {
	return &counter{
		client:     ddClient,
		metricName: metricName,
	}
}

// With is applied in kratos/middleware/metrics/metrics.go (method,path,code)
func (d *counter) With(lvs ...string) metrics.Counter {
	tags := make([]string, 0, 3)
	if len(lvs) >= 3 {
		tags = []string{"method:" + lvs[0], "path:" + lvs[1], "code:" + lvs[2]}
	}
	d.tempTags = tags
	return d
}

func (d *counter) Inc() {
	d.client.Incr(d.metricName, d.tempTags, defaultRate)
}

func (d *counter) Add(delta float64) {
	d.client.Count(d.metricName, int64(delta), d.tempTags, defaultRate)
}
