package producers

import (
	"fmt"
	"sync"

	"github.com/zalando-incubator/mate/pkg"
	"github.com/zalando-incubator/mate/producers/fake"
	"github.com/zalando-incubator/mate/producers/kubernetes"
	"github.com/zalando-incubator/mate/producers/null"
)

type Producer interface {
	Endpoints() ([]*pkg.Endpoint, error)
	Monitor(chan *pkg.Endpoint, chan error, chan struct{}, *sync.WaitGroup)
}

func New(name string) (Producer, error) {
	switch name {
	case "kubernetes":
		return kubernetes.NewProducer()
	case "fake":
		return fake.NewFake()
	case "null":
		return null.NewNull()
	}
	return nil, fmt.Errorf("Unknown producer '%s'.", name)
}
