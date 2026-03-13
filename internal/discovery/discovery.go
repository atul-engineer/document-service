package discovery

import (
	"context"

	clientV3 "go.etcd.io/etcd/client/v3"
)

type Discovery struct {
	Client *clientV3.Client
}

func NewDiscovery(client *clientV3.Client) *Discovery {
	return &Discovery{
		Client: client,
	}
}

func (d *Discovery) GetServices(srv string) ([]string, error) {
	resp, err := d.Client.Get(
		context.Background(),
		"/services/" + srv + "/",
		clientV3.WithPrefix(),
	)
	if err != nil {
		return nil, err
	}
	var services []string
	for _, kv := range resp.Kvs {
		services = append(services, string(kv.Value))
	}
	return services, nil
}
