package model

import "time"

type Command struct {
	Name       string
	Key        string
	Value      interface{}
	ExpireTime time.Duration
	args       interface{}
}
