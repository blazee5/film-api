package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var totalRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "grpc_requests_total",
		Help: "Number of grpc requests.",
	},
)

func IncRequests() {
	totalRequests.Inc()
}

func NewMetrics() error {
	err := prometheus.Register(totalRequests)

	if err != nil {
		return err
	}

	return nil
}

func ListenMetrics(address string) error {
	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(address, nil)
}
