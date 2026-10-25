package main

import (
	"fmt"
	"github.com/collins994/redis_clone/code/store"
	"time"
	// "log"
	// "os"
)

type MetricsMiddleware struct {
	inner       store.Storer
	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int

	getMisses       int
	totalGetLatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{
		inner: inner,
	}
}

func (m *MetricsMiddleware) Set(key, value string) error {
	m.setCalls++
	startTime := time.Now()
	m.totalSetLatency += time.Since(startTime)
	return m.inner.Set(key, value)
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	m.getCalls++
	startTime := time.Now()
	val, err := m.inner.Get(key)
	if err != nil {
		m.getMisses++
	}
	m.totalGetLatency += time.Since(startTime)
	return val, err
}

func (m *MetricsMiddleware) Delete(key string) {
	m.deleteCalls++
	m.inner.Delete(key)
}

func (m *MetricsMiddleware) Keys() []string { return m.inner.Keys() }
func (m *MetricsMiddleware) Len() int       { m.lenCalls++; return m.inner.Len() }

func (m *MetricsMiddleware) Report() {
	fmt.Printf("Get calls: %d. (misses %d)\n", m.getCalls, m.getMisses)
	fmt.Printf("Set calls: %d\n", m.setCalls)
	fmt.Printf("Delete calls: %d\n", m.deleteCalls)
	fmt.Printf("len calls: %d\n", m.lenCalls)

	if m.getCalls > 0 {
		fmt.Printf("Average Get latency: %v\n", m.totalGetLatency/time.Duration(m.getCalls))
	}

	if m.setCalls > 0 {
		fmt.Printf("Average Set latency: %v\n", m.totalSetLatency/time.Duration(m.setCalls))
	}
}
