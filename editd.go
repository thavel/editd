package main

import (
	"log"
	"github.com/thavel/editd/modules/etcd"
)

var (
	etcdAddr = "etcd"
	etcdPort = 4001
)

func main() {
	client := etcd.NewClient(etcdAddr, etcdPort)
	log.Println(client.GetUrl("/hello/world"))
}