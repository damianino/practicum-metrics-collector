package main

import (
	"github.com/damianino/practicum-metrics-collector/internal/handler"
	"github.com/damianino/practicum-metrics-collector/internal/service/memstorage"
	"log"
	"net/http"
)

func main() {
	ms := memstorage.NewMemStorage()

	h := handler.NewHandler(ms)

	mux := http.NewServeMux()
	mux.HandleFunc("/update/gauge/{metricName}/{metricValue}", h.SetGauge)
	mux.HandleFunc("/update/counter/{metricName}/{metricValue}", h.IncCounter)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusBadRequest) })

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
