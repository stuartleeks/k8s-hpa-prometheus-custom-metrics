package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO - allow for customising the metric name
var (
	value      = 0.0
	valueGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "prom_fake_value",
		Help: "prom-fake value",
	}) // TODO - consider NewGaugeFunc to only set value in a single way?
)

func main() {
	// initialize metric
	valueGauge.Set(value)

	// set up server
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/value", valueHandler)

	// start server
	log.Printf("Starting...\n")
	http.ListenAndServe(":8080", nil)
}

func valueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, fmt.Sprintf("%f", value))
		return
	}
	if r.Method == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		stringValue := buf.String()
		intValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			log.Printf("Error in /value. %s, body=%s\n", err, stringValue)
			http.Error(w, err.Error(), 400)
			return
		}
		value = intValue
		valueGauge.Set(value)
		return
	}
	http.Error(w, fmt.Sprintf("Unhandled Method '%s'", r.Method), 400)
}
