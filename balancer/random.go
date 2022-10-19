package balancer

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	factories[RandomBalancer] = NewRandom
}

type Random struct {
	sync.RWMutex
	hosts []string
	rnd   *rand.Rand
}

func NewRandom(hosts []string) Balancer {
	return &Random{hosts: hosts,
		rnd: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (r *Random) Add(host string) {
	r.Lock()
	defer r.Unlock()
	for _, h := range r.hosts {
		if h == host {
			return
		}
	}
	r.hosts = append(r.hosts, host)
}

func (r *Random) Delete(host string) {
	r.Lock()
	defer r.Unlock()
	for idx, h := range r.hosts {
		if h == host {
			r.hosts = append(r.hosts[:idx], r.hosts[idx+1:]...)
		}
	}
}

func (r *Random) Balance(host string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", NoHostError
	}
	return r.hosts[r.rnd.Intn(len(r.hosts))], nil
}

func (r *Random) Inc(host string) {}

func (r *Random) Done(host string) {}
