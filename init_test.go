package main

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Printf("%+v\n", DefaultTxt2ImgReq())

	fmt.Printf("%+v\n", DefaultImg2ImgReq())
}

func TestNewMyLog(t *testing.T) {
	log := NewMyLog(true)

	log.Info("hello world")
	log.Panic("hello world")
	log.Warn("hello world")
}
