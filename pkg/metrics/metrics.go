package metrics

import(
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServeMetrics(){
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

