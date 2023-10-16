package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSD_NewTxt2Img(t *testing.T) {
	// baseUrl like: http://127.0.0.1:7861
	sd := NewSD("http://127.0.0.1:7861", 20*time.Minute, true)

	req := DefaultTxt2ImgReq()
	img, err := sd.Txt2Img(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	println(img.ToString())
}

func TestImage_ToImage(t *testing.T) {
	sd := NewSD("http://127.0.0.1:7861", 20*time.Minute, true)

	open, err := os.Open("./images/my.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	req := DefaultImg2ImgReq()
	req.SetImage(NewImage(open))

	img, err := sd.Img2Img(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(img.ToImage("./images", "ff"))
}

func TestLog(t *testing.T) {
	log := NewMyLog(true)

	log.Info("qwewqe", "das", "Dasdas")
	log.Warn("qwewqe")

	log = NewMyLog(false)

	log.Info("qwewqe", "das", "Dasdas")
	log.Warn("qwewqe")
}
