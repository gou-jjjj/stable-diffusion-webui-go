package main

import (
	"fmt"
	"log"
)

type MyLog struct {
	use bool
}

func NewMyLog(use bool) *MyLog {
	return &MyLog{
		use: use,
	}
}

func (l MyLog) Info(msg ...string) {
	if l.use {
		m := fmt.Sprint(msg)
		log.Printf("\u001B[1;32m%s\u001B[0m", m)
	}
}

func (l MyLog) Warn(msg ...string) {
	if l.use {
		m := fmt.Sprintln(msg)
		log.Printf("\u001B[1;32m%s\u001B[0m", m)
	}
}
