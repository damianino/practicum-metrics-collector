package handler

import (
	"github.com/damianino/practicum-metrics-collector/internal/service/memstorage"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	memstorage memstorage.MemStorage
}

func NewHandler(storage memstorage.MemStorage) *Handler {
	return &Handler{
		memstorage: storage,
	}
}

func (h *Handler) SetGauge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// path should look like this: /update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) != 4 {
		http.Error(w, "invalid path", http.StatusBadRequest)
	}

	if r.PathValue("metricName") == "" || r.PathValue("metricValue") == "" {
		http.Error(w, "invalid params", http.StatusBadRequest)
	}

	value, err := strconv.ParseFloat(r.PathValue("metricValue"), 64)
	if err != nil {
		http.Error(w, "invalid gauge value", http.StatusBadRequest)
	}

	h.memstorage.SetGauge(segments[2], value)
}

func (h *Handler) IncCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// path should look like this: /update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) != 4 {
		http.Error(w, "invalid path", http.StatusBadRequest)
	}

	if r.PathValue("metricName") == "" || r.PathValue("metricValue") == "" {
		http.Error(w, "invalid params", http.StatusBadRequest)
	}

	value, err := strconv.ParseInt(r.PathValue("metricValue"), 10, 64)
	if err != nil {
		http.Error(w, "invalid inc value", http.StatusBadRequest)
	}

	h.memstorage.IncCounter(segments[2], value)
}
