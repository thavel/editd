package etcd

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

const (
	baseUri = "http://%s:%d"
)

type Client struct {
	url    string
	config client.Config
	etcd   client.KeysAPI
}

func NewClient(conf *Config) *Client {
	cli := new(Client)
	cli.url = fmt.Sprintf(baseUri, conf.address, conf.port)

	// Building a configuration structure
	cli.config = client.Config {
		Endpoints: []string{cli.url},
		Transport: client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	// Allocate a new etcd client
	c, _ := client.New(cli.config)
	cli.etcd = client.NewKeysAPI(c)
	return cli
}

func (cli *Client) Push(to string, data string) error {
	_, err := cli.etcd.Set(context.Background(), to, data, nil)
	return err
}

func (cli *Client) Pop(from string) (string, error) {
	resp, err := cli.etcd.Get(context.Background(), from, nil)
	if err != nil {
		return "", err
	}
	return resp.Node.Value, err
}