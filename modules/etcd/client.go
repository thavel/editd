package etcd

import (
	"fmt"
	"strings"
)

const (
	baseUri = "http://%s:%d/%s/keys"
	defaultVersion = "v2"
)

type Client struct {
	url string
}

func NewClient(addr string, port int) *Client {
	cli := new(Client)
	cli.url = fmt.Sprintf(baseUri, addr, port, defaultVersion)
	return cli
}

func (cli *Client) GetUrl(keyPath string) string {
	if strings.HasSuffix(keyPath, "/") {
		panic(fmt.Sprintf("'%s' is not a valid key!", keyPath))
	}

	keyPath = strings.TrimPrefix(keyPath, "/")
	return fmt.Sprintf("%s/%s", cli.url, keyPath)
}
