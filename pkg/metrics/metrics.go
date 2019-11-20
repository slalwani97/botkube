package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)
// ServeMetrics expose metrics 
func ServeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
