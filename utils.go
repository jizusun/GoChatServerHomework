package main

import (
	"time"
)

type UtilitiesInterface interface {
	GetTimestamp() int64
}

type Utilities struct{}

func (ex Utilities) GetTimestamp() int64 {
	return time.Now().Unix()
}
