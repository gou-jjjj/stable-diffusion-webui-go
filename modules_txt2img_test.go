package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"stable-diffusion-webui-go/images"
	"testing"
)

func TestImage_Store(t *testing.T) {
	fmt.Printf("%v\n", ImageStore("", "iamges1", images.Img2_test))
	fmt.Printf("%v\n", ImageStore("./", "iamges1", ""))
	fmt.Printf("%v\n", ImageStore("./", "", images.Img2_test))

	fmt.Printf("%v\n", ImageStore("./images", "images2", images.Img2_test))
	fmt.Printf("%v\n", ImageStore("./images/", "images3", images.Img2_test))
	fmt.Printf("%v\n", ImageStore("./images/new", "images4", images.Img2_test))
	fmt.Printf("%v\n", ImageStore("./image", "images5", images.Img2_test))

	fmt.Printf("%v\n", ImageStore("./image", "images5", images.Img1_test))
}

func TestInit(t *testing.T) {
	file, err := os.Open("./txt2img.default.json")
	if err != nil {
		log.Fatalln(err)
	}

	bytes := make([]byte, 0)
	t1 := Txt2ImgRequest{}

	read := bufio.NewReader(file)

	for {
		b, err1 := read.ReadByte()
		if err1 != nil {
			break
		}
		bytes = append(bytes, b)
	}

	err = json.Unmarshal(bytes, &t1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v\n", t1)
}
