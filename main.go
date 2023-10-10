package main

import (
	"github.com/amirhnajafiz/node-exporter/internal/worker"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// todo: convert it to read from kube config file
	var kubeconfig, master string // empty, assuming inClusterConfig
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		panic(err)
	}

	// create a new worker
	w := worker.Worker{}
	w.Work(config)
}
