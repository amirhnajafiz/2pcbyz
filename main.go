package main

import (
	"github.com/amirhnajafiz/node-exporter/internal/metrics"
	"github.com/amirhnajafiz/node-exporter/internal/worker"
	"time"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// todo: convert it to read from kube config file
	var kubeconfig, master string // empty, assuming inClusterConfig
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
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
