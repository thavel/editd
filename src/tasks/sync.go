package tasks

import (
	"fmt"
	"time"
	"sync"

	"etcd"
)

type Pusher struct {
	ticker *time.Ticker
	period time.Duration
	limit  int
	client *etcd.Client
	data   *etcd.Data
	safe   bool
}

func NewSync(cli *etcd.Client, period int, safe bool) *Pusher {
	pusher := new(Pusher)
	pusher.client = cli
	pusher.period = time.Duration(period)
	pusher.safe = safe
	return pusher
}

func (push *Pusher) Set(endpoint string, data string, ttl int) {
	push.data = etcd.NewData(endpoint, data, ttl)
}

func (push *Pusher) Start(limit int) (*sync.WaitGroup, *int) {
	push.limit = limit
	push.ticker = time.NewTicker(time.Millisecond * push.period)

	task := new(sync.WaitGroup)
	task.Add(1)
	state := 0

	go func() {
		// Mark this task done when the goroutine exits
		defer task.Done()

		// Ticker loop
		count := 0
		for range push.ticker.C {
			err := push.client.Push(push.data)
			if err != nil {
				fmt.Println(err)
				if push.safe {
					// Exit if the safe mode is set
					state = 1
					return
				}
			}
			// Use counter only if there is a limit to avoid memory overflow
			if (limit > 0) {
				count += 1
				if (count >= limit) {
					// Stop the goroutine
					state = 0
					return
				}
			}
		}
	}()
	return task, &state
}

func (push *Pusher) Stop() bool {
	if push.ticker != nil {
		push.ticker.Stop()
		return true
	}
	return false
}