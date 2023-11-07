package main

import (
	"time"

	"github.com/amirhnajafiz/node-exporter/internal/metrics"
	"github.com/amirhnajafiz/node-exporter/internal/worker"

	"k8s.io/client-go/rest"
)

func main() {
	// cluster client configs
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	// create prometheus metrics
	m := metrics.New()

	// create a new worker
	w := worker.Worker{
		Metrics:  m,
		Interval: 5 * time.Second,
	}
	if er := w.Work(config); er != nil {
		panic(er)
	}

	// create handler
	metrics.Handler{}.Metrics()
}
