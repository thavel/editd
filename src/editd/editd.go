package main

import (
	"log"
	"github.com/thavel/editd/etcd"
)

var (
	etcdApi  = "v2"
	etcdAddr = "etcd"
	etcdPort = 4001
)

func main() {
	client := Client{etcdApi, etcdAddr, etcdPort}
	log.Println(client.etcdUrl("/hello/prout"))
}