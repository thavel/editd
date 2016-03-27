package main

import (
	"os"
	"flag"
	"fmt"
	"io/ioutil"

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
	fvalue   = flag.String("fvalue", "", "file content value")
	ttl      = flag.Int("ttl", 10000, "TTL duration for keys")
	nottl    = flag.Bool("nottl", false, "disable TTL duration for keys")
)

func main() {
	// Parsing flags (aka command arguments)
	flag.Parse()

	// Instanciating etcd client and pusher
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

	// Compute key's value
	val := *value
	if len(*fvalue) > 0 {
		if len(*value) > 0 {
			fmt.Println("Can't use both value and fvalue!")
			os.Exit(1)
		}
		val, err = readfvalue(*fvalue)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Creating a new asynchronous task to push keys/values.
	pusher := tasks.NewSync(client, *interval, *safe)
	err = pusher.Set(*key, val, ttlValue)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	task, state := pusher.Start(limit)

	// Await for task to complete
	task.Wait()
	os.Exit(*state)
}

func readfvalue(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	return string(data), err
}
