package main

import (
	"os"
	"flag"
	"fmt"

	"etcd"
)

var (
	node     string
	interval int
	onetime  bool
)

func init() {
	flag.StringVar(&node, "node", "", "etcd node")
	flag.IntVar(&interval, "interval", 600, "synchronization interval")
	flag.BoolVar(&onetime, "onetime", true, "run once and exit")
}

func main() {
	flag.Parse()

	config, err := etcd.NewConfig(node)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := etcd.NewClient(config)

	err = client.Push("/test", `{"hello": "world"}`)
	fmt.Println(err)
	value, _ := client.Pop("/test")
	fmt.Println(value)
}