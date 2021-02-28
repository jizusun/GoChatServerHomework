// Author: Jizu Sun
package main

import (
	"time"
)

type utilitiesInterface interface {
	GetTimestamp() int64
}

type utilities struct{}

func (ex utilities) GetTimestamp() int64 {
	return time.Now().Unix()
}
