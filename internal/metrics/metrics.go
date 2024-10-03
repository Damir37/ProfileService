package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "profile_service"

var ProfileHTTPRequest = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "requests_total",
		Help:      "Количество HTTP запросов",
	},
	[]string{"Method", "Path"},
)

var Profile5xxRequest = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "requests_error_total",
		Help:      "Количество HTTP запросов с ответом 500-х ошибок",
	},
	[]string{"Method", "Path"},
)

var ProfileRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "request_duration",
		Help:      "Время выполнения HTTP запросов в секундах",
		Buckets:   prometheus.DefBuckets,
	},
	[]string{"Method", "Path", "Status"},
)

var ProfileCPUUsage = promauto.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "cpu_usage",
		Help:      "Процент использования CPU",
	},
)

var ProfileMemoryUsage = promauto.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "memory_usage",
		Help:      "Процент использования памяти",
	},
)

var ProfileTrafficInbound = promauto.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "traffic_inbound_bytes",
		Help:      "Распределение байтов входящего трафика",
		Buckets:   prometheus.ExponentialBuckets(10, 2, 10),
	},
)

var ProfileTrafficOutbound = promauto.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "traffic_outbound_bytes",
		Help:      "Распределение байтов исходящего трафика",
		Buckets:   prometheus.ExponentialBuckets(10, 2, 10),
	},
)
