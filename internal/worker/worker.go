package worker

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Worker struct {
}

func (w Worker) Work(cfg *rest.Config) {
	mc, err := metrics.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	mc.MetricsV1beta1().NodeMetricses().List(metav1.ListOptions{})
}
