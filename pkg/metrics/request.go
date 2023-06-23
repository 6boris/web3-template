package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Version = "v1.0.0"

	RequestMetricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	RequestMetricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	RedisMetricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "redis",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "redis requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"name", "service_id", "service_instance", "cmd", "status"})
	RedisMetricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "redis",
		Subsystem: "requests",
		Name:      "cmd_total",
		Help:      "The total number of processed requests",
	}, []string{"name", "service_id", "service_instance", "cmd", "status"})

	MySQLMetricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "mysql",
		Subsystem: "requests",
		Name:      "cmd_total",
		Help:      "The total number of processed requests",
	}, []string{"name", "service_id", "service_instance", "table", "cmd"})
)

func init() {
	prometheus.MustRegister(
		RequestMetricSeconds,
		RequestMetricRequests,
		RedisMetricRequests,
		RedisMetricSeconds,
		MySQLMetricRequests,
	)
}
