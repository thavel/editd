package main

import (
	"log"
	"fmt"
	"strings"
)

var (
	etcdApi  = "v2"
	etcdAddr = "etcd"
	etcdPort = 4001
)

const (
	endpoint = "http://%s:%d/%s/keys/%s"
)

func etcdUrl(keyPath string) string {
	if strings.HasSuffix(keyPath, "/") {
		panic(fmt.Sprintf("'%s' is not a valid key!", keyPath))
	}

	keyPath = strings.TrimPrefix(keyPath, "/")
	return fmt.Sprintf(endpoint, etcdAddr, etcdPort, etcdApi, keyPath)
}

func main() {
	log.Println(etcdUrl("/hello/world"))
}