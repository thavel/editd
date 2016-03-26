package etcd

import (
	"time"
)

type Data struct {
	key   string
	value string
	ttl   time.Duration
}

func NewData(key string, value string, ttl int) *Data {
	data := new(Data)
	data.key = key
	data.value = value
	data.ttl = time.Millisecond * time.Duration(ttl)
	return data
}