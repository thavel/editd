package main

import (
	"os"
	"flag"
	"fmt"

	"etcd"
	"tasks"
)

var (
	node     = flag.String("node", "", "etcd node")
	interval = flag.Int("interval", 5000, "synchronization interval")
	onetime  = flag.Bool("onetime", false, "run once and exit")
	safe     = flag.Bool("safe", false, "exit upon errors")
	key      = flag.String("key", "", "etcd key path")
	value    = flag.String("value", "", "specified key's value")
	ttl      = flag.Int("ttl", 10000, "key time to live")
)

func main() {
	flag.Parse()

	config, err := etcd.NewConfig(*node)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := etcd.NewClient(config)

	limit := -1
	if *onetime {
		limit = 1
	}

	pusher := tasks.NewSync(client, *interval)
	pusher.Set(*key, *value, *ttl)
	task := pusher.Start(limit)

	task.Wait()
	os.Exit(0)
}