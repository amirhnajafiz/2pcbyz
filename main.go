package main

import (
	"os"
	"strconv"
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

	port, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	interval, _ := strconv.Atoi(os.Getenv("INTERVAL"))

	// create prometheus metrics
	m := metrics.New()

	// create a new worker
	w := worker.Worker{
		Metrics:  m,
		Interval: time.Duration(interval) * time.Second,
	}
	if er := w.Work(config); er != nil {
		panic(er)
	}

	// create handler
	metrics.Handler{}.Metrics(port)
}
