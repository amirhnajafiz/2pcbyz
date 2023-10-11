package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct{}

func (h Handler) Metrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
