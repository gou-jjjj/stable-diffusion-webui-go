package main

import (
	"os"
	"testing"
)

func TestNewImage(t *testing.T) {
	open, err := os.Open("./images/my.png")
	if err != nil {
		panic(err)
		return
	}
	image := NewImage(open)

	println(image.ToString())
}
