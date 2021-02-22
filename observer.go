package datadog

import (
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Observer = (*timer)(nil)

type timer struct {
	client     *statsd.Client
	logHelper  *log.Helper
	metricName string
	tempTags   []string
}

// NewTimer new a DataDog timer and returns Observer.
func NewTimer(metricName string, logger log.Logger) metrics.Observer {
	return &timer{
		client:     ddClient,
		logHelper:  log.NewHelper("metrics/timer", logger),
		metricName: metricName,
	}
}

// With is applied in kratos/middleware/metrics/metrics.go (method,path)
func (d *timer) With(lvs ...string) metrics.Observer {
	tags := make([]string, 0, 2)
	if len(lvs) >= 2 {
		tags = []string{"method:" + lvs[0], "path:" + lvs[1]}
	}
	d.tempTags = tags
	return d
}

func (d *timer) Observe(value float64) {
	if err := d.client.Timing(d.metricName, time.Duration(value)*time.Second, d.tempTags, defaultRate); err != nil {
		d.logHelper.Warnf("observe %+v error %+v", d.tempTags, err)
	}
}
