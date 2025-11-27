package handler

import (
	"github.com/damianino/practicum-metrics-collector/internal/service/memstorage"
	"net/http"
	"strconv"
)

type Handler struct {
	memstorage memstorage.MemStorage
}

func NewHandler(storage memstorage.MemStorage) *Handler {
	return &Handler{
		memstorage: storage,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metricType := r.PathValue("metricType")
	metricName := r.PathValue("metricName")
	metricValue := r.PathValue("metricValue")

	if metricName == "" || metricValue == "" {
		http.Error(w, "invalid params", http.StatusBadRequest)
	}

	switch metricType {
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
		}

		h.memstorage.SetGauge(metricName, value)
	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "invalid inc value", http.StatusBadRequest)
		}

		h.memstorage.IncCounter(metricName, value)
	default:
		http.Error(w, "invalid metric type", http.StatusBadRequest)
	}
}
