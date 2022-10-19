package proxy

import (
	"log"
	"time"
)

func (h *HTTPProxy) SetAlive(host string, alive bool) {
	h.Lock()
	defer h.Unlock()
	h.alive[host] = alive
}

func (h *HTTPProxy) GetAlive(host string) bool {
	h.RLock()
	defer h.RUnlock()
	return h.alive[host]
}

func (h *HTTPProxy) healthCheck(host string, interval uint) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	log.Printf("Checking site %s health", host)
	for range ticker.C {
		if !IsBackendAlive(host) && h.GetAlive(host) {
			log.Printf("Site unreachable, remove %s from load balancer.", host)

			h.SetAlive(host, false)
			h.lb.Delete(host)
		} else if IsBackendAlive(host) && !h.GetAlive(host) {
			log.Printf("Site reachable, add %s to load balancer.", host)

			h.SetAlive(host, true)
			h.lb.Add(host)
		}
	}
}

func (h *HTTPProxy) HealthCheck(interval uint) {
	for host := range h.hostMap {
		go h.healthCheck(host, interval)
	}
}
