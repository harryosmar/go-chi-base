package middlewares

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"time"
)

var (
	httpRequestsTotal *prometheus.CounterVec
	httpDuration      *prometheus.HistogramVec
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(rw, r)

		status := strconv.Itoa(rw.Status())
		duration := time.Since(start).Seconds()

		// Increment request count
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()

		// Observe request duration
		httpDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}

func InitPrometheusMetrics() {
	// Initialize counters to track the total number of requests
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "route", "status"},
	)
	prometheus.MustRegister(httpRequestsTotal)

	// Initialize histogram to measure request duration
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)
	prometheus.MustRegister(httpDuration)
}
