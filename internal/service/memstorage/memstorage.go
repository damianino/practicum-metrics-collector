package memstorage

import "sync"

type MemStorage interface {
	SetGauge(string, float64)
	IncCounter(string, int64)
}

var _ MemStorage = (*memStorage)(nil)

type memStorage struct {
	gauges   map[string]float64
	counters map[string]int64

	gaugeMu   sync.Mutex
	counterMu sync.Mutex
}

func NewMemStorage() MemStorage {
	return &memStorage{
		gauges:    make(map[string]float64),
		counters:  make(map[string]int64),
		gaugeMu:   sync.Mutex{},
		counterMu: sync.Mutex{},
	}
}

func (m *memStorage) SetGauge(name string, v float64) {
	m.gaugeMu.Lock()
	m.gauges[name] = v
	m.gaugeMu.Unlock()
}

func (m *memStorage) IncCounter(name string, v int64) {
	m.counterMu.Lock()
	defer m.counterMu.Unlock()

	prev, ok := m.counters[name]
	if !ok {
		m.counters[name] = v
		return
	}
	m.counters[name] = prev + v
}
