package worker

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Worker struct{}

func (w Worker) Work(cfg *rest.Config) error {
	// create new metrics server
	mc, err := metrics.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to get metrics erorr=%w", err)
	}

	go func() {
		ctx := context.Background()

		for {
			mc.MetricsV1beta1().NodeMetricses().List(ctx, v1.ListOptions{})
		}
	}()

	return nil
}
