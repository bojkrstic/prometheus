package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// COUNTER
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_requests_total",
			Help: "Total number of requests",
		},
	)

	// HISTOGRAM
	requestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "my_request_duration_seconds",
			Help:    "Request duration",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// simulacija rada
	sleep := rand.Intn(500)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	requestsTotal.Inc()
	requestDuration.Observe(time.Since(start).Seconds())

	w.Write([]byte("ok"))
}

func main() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)

	http.HandleFunc("/", handler)

	// METRICS endpoint (ovo Prometheus čita)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
