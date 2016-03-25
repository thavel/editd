package main

import (
	"os"
	"flag"
	"fmt"

	"etcd"
)

var (
	node     = flag.String("node", "", "etcd node")
	interval = flag.Int("interval", 600, "synchronization interval")
	onetime  = flag.Bool("onetime", true, "run once and exit")
	safe     = flag.Bool("safe", false, "exit upon errors")
)

func main() {
	flag.Parse()

	config, err := etcd.NewConfig(*node)
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