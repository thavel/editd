package main

import (
	"log"

	"etcd"
)

var (
	etcdAddr = "192.168.9.128"
	etcdPort = 4001
)

func main() {
	client := etcd.NewClient(etcdAddr, etcdPort)
	err := client.Push("/test", `{"hello": "world"}`)
	log.Println(err)
	value, _ := client.Pop("/test")
	log.Println(value)
}