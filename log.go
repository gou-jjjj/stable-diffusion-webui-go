package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type MyLog struct {
	use   bool
	info  *log.Logger
	warn  *log.Logger
	panic *log.Logger
}

func NewMyLog(use bool) *MyLog {
	info := log.New(os.Stdout, fmt.Sprintf("[Info ][%s]", time.Now().Format("2006-01-02 15:04:05")), log.Lmsgprefix)
	pic := log.New(os.Stdout, fmt.Sprintf("[Panic][%s]", time.Now().Format("2006-01-02 15:04:05")), log.Lmsgprefix)
	warn := log.New(os.Stdout, fmt.Sprintf("[Warn ][%s]", time.Now().Format("2006-01-02 15:04:05")), log.Lmsgprefix)

	return &MyLog{
		use:   use,
		info:  info,
		panic: pic,
		warn:  warn,
	}
}

func (l MyLog) Info(msg ...string) {
	if l.use {
		m := fmt.Sprint(msg)
		l.info.Printf("\u001B[32m%s\u001B[0m", m)
	}
}

func (l MyLog) Warn(msg ...string) {
	if l.use {
		m := fmt.Sprint(msg)
		l.warn.Printf("\u001B[31m%s\u001B[0m", m)
	}
}

func (l MyLog) Panic(msg ...string) {
	m := fmt.Sprint(msg)
	if l.use {
		l.panic.Printf("\u001B[33m%s\u001B[0m", m)
	} else {
		panic(m)
	}
}
