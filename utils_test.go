package main

import (
	"fmt"
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

func TestName(t *testing.T) {
	fmt.Println(fmt.Sprintf("%s", fmt.Sprintf("The image is saved to [%s]", "hello world")))
	fmt.Println(fmt.Sprintf("The image is saved to [%s]", "hello world"))
}
