package datadog

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Counter = (*counter)(nil)

type counter struct {
	client     *statsd.Client
	logHelper  *log.Helper
	metricName string
	tempTags   []string
}

// NewDDCounter new a DataDog counter and returns Counter.
func NewDDCounter(metricName string, logger log.Logger) metrics.Counter {
	return &counter{
		client:     ddClient,
		logHelper:  log.NewHelper("metrics/counter", logger),
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
	if err := d.client.Incr(d.metricName, d.tempTags, defaultRate); err != nil {
		d.logHelper.Warnf("inc %+v error %+v", d.tempTags, err)
	}
}

func (d *counter) Add(delta float64) {
	if err := d.client.Count(d.metricName, int64(delta), d.tempTags, defaultRate); err != nil {
		d.logHelper.Warnf("add %+v error %+v", d.tempTags, err)
	}
}
