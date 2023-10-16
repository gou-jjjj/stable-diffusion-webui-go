package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
)

// ImagesIsNull errors
var (
	ImagesIsNull = errors.New("images is null")
)

var (
	Lg = new(MyLog)

	defaultImg2ImgRequest = new(Img2ImgRequest)
	defaultTxt2ImgRequest = new(Txt2ImgRequest)
)

func init() {
	loadconfig("./img2img.default.json", defaultImg2ImgRequest)
	loadconfig("./txt2img.default.json", defaultTxt2ImgRequest)
}

func loadconfig(path string, module any) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}

	bytes := make([]byte, 0)
	read := bufio.NewReader(file)
	for {
		b, err1 := read.ReadByte()
		if err1 != nil {
			break
		}
		bytes = append(bytes, b)
	}

	err = json.Unmarshal(bytes, module)
	if err != nil {
		log.Panicln(err)
	}
}

func DefaultImg2ImgReq() *Img2ImgRequest {
	defaultImg2ImgRequest.Seed = Seed()
	return defaultImg2ImgRequest
}

func DefaultTxt2ImgReq() *Txt2ImgRequest {
	defaultTxt2ImgRequest.Seed = Seed()
	return defaultTxt2ImgRequest
}
