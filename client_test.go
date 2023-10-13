package main

import (
	"fmt"
	"testing"
)

func TestSD_NewTxt2Img(t *testing.T) {
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
