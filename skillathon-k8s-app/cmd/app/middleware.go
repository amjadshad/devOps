package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

var requestsTotal = metrics.NewCounter("requests_total")

// nonProxiedEndpointsMiddleware is a workaround to make various service endpoints not to raise requestsTotal metrics
func (app *application) metricsAndHealthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/health/liveness":
			if app.FailLiveness {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintln(w, "The app is stuck")
			} else {
				fmt.Fprintln(w, "OK")
			}
			return
		case "/health/readiness":
			if app.FailReadiness {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintln(w, "The app is not currently responsive, but will be fine")
			} else {
				fmt.Fprintln(w, "OK")
			}
			return
		case "/metrics":
			metrics.WritePrometheus(w, true)
			return
		case "/shutdown":
			time.Sleep(app.ShutdownDelay)
			return
		default:
			requestsTotal.Inc()
			next.ServeHTTP(w, r)
		}
	})
}
