package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	CPU     *prometheus.HistogramVec
	Memory  *prometheus.HistogramVec
	Pods    *prometheus.GaugeVec
	Storage *prometheus.HistogramVec
}

func New() *Metrics {
	return &Metrics{
		CPU: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "cpu_usage",
			},
			[]string{"name"},
		),
		Memory: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "memory_usage",
			},
			[]string{"name"},
		),
		Pods: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "total_pods",
			},
			[]string{"name"},
		),
		Storage: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "storage_usage",
			},
			[]string{"name"},
		),
	}
}

func (m *Metrics) ObserveCPU(label string, value int64) {
	m.CPU.With(prometheus.Labels{"name": label}).Observe(float64(value))
}

func (m *Metrics) ObserveMemory(label string, value int64) {
	m.Memory.With(prometheus.Labels{"name": label}).Observe(float64(value))
}

func (m *Metrics) ObservePods(label string, value int64) {
	m.Pods.With(prometheus.Labels{"name": label}).Set(float64(value))
}

func (m *Metrics) ObserveStorage(label string, value int64) {
	m.Storage.With(prometheus.Labels{"name": label}).Observe(float64(value))
}
