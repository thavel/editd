package tasks


import (
	"time"
	"sync"

	"etcd"
)

type Pusher struct {
	ticker *time.Ticker
	period time.Duration
	limit  int
	client *etcd.Client
	key    string
	value  string
}

func NewSync(cli *etcd.Client, period int) *Pusher {
	pusher := new(Pusher)
	pusher.client = cli
	pusher.period = time.Duration(period)
	return pusher
}

func (push *Pusher) Set(endpoint string, data string) {
	push.key = endpoint
	push.value = data
}

func (push *Pusher) Start(limit int) *sync.WaitGroup {
	push.limit = limit
	push.ticker = time.NewTicker(time.Millisecond * push.period)

	task := new(sync.WaitGroup)
	task.Add(1)
	go func() {
		defer task.Done()
		// Ticker loop
		count := 0
		for range push.ticker.C {
			push.client.Push(push.key, push.value)
			// Use counter only if there is a limit to avoid memory overflow
			if (limit > 0) {
				count += 1
				if (count >= limit) {
					// Stop the goroutine
					return
				}
			}
		}
	}()
	return task
}

func (push *Pusher) Stop() bool {
	if push.ticker != nil {
		push.ticker.Stop()
		return true
	}
	return false
}