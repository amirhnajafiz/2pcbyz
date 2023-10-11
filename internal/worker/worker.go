package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	handler "github.com/amirhnajafiz/node-exporter/internal/metrics"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Worker struct {
	Metrics  *handler.Metrics
	Interval time.Duration
}

func (w Worker) Work(cfg *rest.Config) error {
	// create new metrics server
	mc, err := metrics.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to get metrics erorr=%w", err)
	}

	go func() {
		ctx := context.Background()

		for {
			// get nodes
			list, er := mc.MetricsV1beta1().NodeMetricses().List(ctx, v1.ListOptions{})
			if er != nil {
				log.Println(fmt.Errorf("failed to get metrics error=%w", er))

				continue
			}

			// get items resources
			for _, item := range list.Items {
				name := item.GetGenerateName()
				cpu, _ := item.Usage.Cpu().AsInt64()
				memory, _ := item.Usage.Memory().AsInt64()
				pods, _ := item.Usage.Pods().AsInt64()
				storage, _ := item.Usage.Storage().AsInt64()

				w.Metrics.ObserveCPU(name, cpu)
				w.Metrics.ObserveMemory(name, memory)
				w.Metrics.ObservePods(name, pods)
				w.Metrics.ObserveStorage(name, storage)
			}

			time.Sleep(w.Interval)
		}
	}()

	return nil
}
