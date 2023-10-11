package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	CPU     prometheus.HistogramVec
	Memory  prometheus.HistogramVec
	Pods    prometheus.HistogramVec
	Storage prometheus.HistogramVec
}

func (m *Metrics) ObserveCPU(label string, value int64) {
	m.CPU.With(prometheus.Labels{"name": label}).Observe(float64(value))
}

func (m *Metrics) ObserveMemory(label string, value int64) {
	m.Memory.With(prometheus.Labels{"name": label}).Observe(float64(value))
}

func (m *Metrics) ObservePods(label string, value int64) {
	m.Pods.With(prometheus.Labels{"name": label}).Observe(float64(value))
}

func (m *Metrics) ObserveStorage(label string, value int64) {
	m.Storage.With(prometheus.Labels{"name": label}).Observe(float64(value))
}
