package main

import (
	"fmt"
	"testing"
)

func TestSD_NewTxt2Img(t *testing.T) {
	// If your sd service is not in 127.0.0.1:7861, you must use NewSD() to initialize the client
	sd := DefaultSD()

	img, err := sd.Txt2Img(DefaultTxt2ImgRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = img.ToImage("./images", "myimage")
	if err != nil {
		fmt.Println(err)
		return
	}
}
