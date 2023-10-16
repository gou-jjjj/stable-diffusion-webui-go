package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSD_NewTxt2Img(t *testing.T) {
	// baseUrl like: http://127.0.0.1:7861
	sd := NewSD("http://127.0.0.1:7861", 20*time.Minute, true)

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

func TestLog(t *testing.T) {
	log := NewMyLog(true)

	log.Info("qwewqe", "das", "Dasdas")
	log.Warn("qwewqe")

	log = NewMyLog(false)

	log.Info("qwewqe", "das", "Dasdas")
	log.Warn("qwewqe")
}
