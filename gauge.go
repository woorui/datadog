package datadog

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metrics"
)

var _ metrics.Gauge = (*gauge)(nil)

type gauge struct {
	client     *statsd.Client
	logHelper  *log.Helper
	metricName string
	tempTags   []string
}

// NewDDGauge new a DataDog gauge and returns Gauge.
func NewDDGauge(metricName string, logger log.Logger) metrics.Gauge {
	return &gauge{
		client:     ddClient,
		logHelper:  log.NewHelper("metrics/gauge", logger),
		metricName: metricName,
	}
}

func (d *gauge) With(lvs ...string) metrics.Gauge {
	tags := make([]string, 0, len(lvs))
	d.tempTags = append(tags, lvs...)
	return d
}

func (d *gauge) Set(value float64) {
	if err := d.client.Gauge(d.metricName, value, d.tempTags, defaultRate); err != nil {
		d.logHelper.Warnf("set %+v error %+v", d.tempTags, err)
	}
}

func (d *gauge) Add(delta float64) {
}

func (d *gauge) Sub(delta float64) {
}
