package etcd

import (
	"fmt"
	"strings"
)

const (
	endpoint = "http://%s:%d/%s/keys/%s"
)

type Client struct {
	etcdApi  string
	etcdAddr string
	etcdPort int
}

func (cli Client) etcdUrl(keyPath string) string {
	if strings.HasSuffix(keyPath, "/") {
		panic(fmt.Sprintf("'%s' is not a valid key!", keyPath))
	}

	keyPath = strings.TrimPrefix(keyPath, "/")
	return fmt.Sprintf(
		endpoint,
		cli.etcdAddr, cli.etcdPort, cli.etcdApi, keyPath,
	)
}
