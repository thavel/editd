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
	ttl      = flag.Int("ttl", 10000, "TTL duration for keys")
	nottl    = flag.Bool("nottl", false, "disable TTL duration for keys")
)

func main() {
	flag.Parse()

	config, err := etcd.NewConfig(*node)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := etcd.NewClient(config)

	// Compute pushing limit (onetime = 1 pushing)
	limit := -1
	if *onetime {
		limit = 1
	}

	// Compute TTL value (no TTL = no duration)
	ttlValue := *ttl
	if *nottl {
		ttlValue = 0
	}

	pusher := tasks.NewSync(client, *interval)
	pusher.Set(*key, *value, ttlValue)
	task := pusher.Start(limit)

	task.Wait()
	os.Exit(0)
}