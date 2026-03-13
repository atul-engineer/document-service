package registry

import (
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"
)

type Registry struct {
	Client *clientV3.Client
}

func NewRegistry(endpoints []string) (*Registry, error) {
	cfg := clientV3.Config{
		Endpoints: endpoints,
		DialTimeout: 5 * time.Second,
	}
	client, err := clientV3.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Registry{Client: client}, nil
}


func (r *Registry) Register(service, instance string, addr string) error {
	key := "/services/" + service + "/" + instance
	_, err := r.Client.Put(r.Client.Ctx(), key, addr)
	return err
}
