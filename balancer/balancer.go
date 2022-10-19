package balancer

import "errors"

var (
	NoHostError             = errors.New("no host")
	AlgorithmNoSupportError = errors.New("algorithm not support")
)

type Balancer interface {
	Add(string)
	Delete(string)
	Balance(string) (string, error)
	Inc(string)
	Done(string)
}

type Factory func([]string) Balancer

var factories = make(map[string]Factory)

func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, AlgorithmNoSupportError
	}
	return factory(hosts), nil
}
