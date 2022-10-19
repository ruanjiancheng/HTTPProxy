package balancer

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandom_Add(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	cases := []struct {
		name   string
		lb     Balancer
		args   string
		expect Balancer
	}{
		{
			"test-1",
			&Random{hosts: []string{"http://127.0.0.1:1011",
				"http://127.0.0.1:1012", "http://127.0.0.1:1013"}, rnd: rnd},
			"http://127.0.0.1:1013",
			&Random{hosts: []string{"http://127.0.0.1:1011",
				"http://127.0.0.1:1012", "http://127.0.0.1:1013"}, rnd: rnd},
		},
		{
			"test-2",
			&Random{hosts: []string{"http://127.0.0.1:1011",
				"http://127.0.0.1:1012"}, rnd: rnd},
			"http://127.0.0.1:1012",
			&Random{hosts: []string{"http://127.0.0.1:1011",
				"http://127.0.0.1:1012"}, rnd: rnd},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.lb.Add(c.args)
			assert.Equal(t, c.expect, c.lb)
		})
	}
}

func TestRandom_Balance(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	cases := []struct {
		name   string
		lb     Balancer
		args   string
		expect error
	}{
		{
			name: "test-1",
			lb: &Random{hosts: []string{"http://127.0.0.1:1011",
				"http://127.0.0.1:1012", "http://127.0.0.1:1013"}, rnd: rnd},
			args:   "198.168.0.1",
			expect: nil,
		},
		{
			name:   "test-2",
			lb:     &Random{hosts: []string{}, rnd: rnd},
			args:   "198.168.0.1",
			expect: NoHostError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := c.lb.Balance(c.args)
			assert.Equal(t, err, c.expect)
		})
	}
}
