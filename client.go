package datadog

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
)

const defaultRate = 1

var ddClient *statsd.Client

func init() {
	var err error
	ddClient, err = statsd.New("")
	if err != nil {
		panic(fmt.Sprintf("DataDog Client initialize error %+v\n", err))
	}
}
