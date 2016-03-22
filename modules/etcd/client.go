package etcd

import (
	"fmt"
	"bytes"
	"strings"
	"net/http"
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

func (cli *Client) Push(to string, data string) bool {
	// data should be json formatted: `{"key": "value"}`
	json := []byte(data)
	url := cli.GetUrl(to)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.Status == "200"
}
