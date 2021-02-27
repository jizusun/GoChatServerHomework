package main

import (
	"time"
)

type ExternalInterface interface {
	GetTimestamp() int64
}

type External struct{}

func (ex External) GetTimestamp() int64 {
	return time.Now().Unix()
}
