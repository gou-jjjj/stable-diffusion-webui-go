package main

import (
	"bufio"
	"os"
)

type Image struct {
	Images []string `json:"images"`
}

// NewImage You can use string(:must be base64), []byte, *os.File to create an *Image
func NewImage(image any) *Image {
	switch image.(type) {
	case string:
		return &Image{Images: []string{image.(string)}}

	case []byte:
		base64 := Base64(image.([]byte))
		return &Image{Images: []string{base64}}

	case *os.File:
		file := image.(*os.File)
		reader := bufio.NewReader(file)
		b := make([]byte, 0)
		for {
			readByte, err := reader.ReadByte()
			if err != nil {
				break
			}
			b = append(b, readByte)
		}
		base64 := Base64(b)
		return &Image{Images: []string{base64}}

	default:
		return &Image{Images: make([]string, 0)}
	}
}

func (r *Image) ToImage(path string, name string) error {
	if r.isNull() {
		return ImagesIsNull
	}
	return ImageStore(r.Images[0], path, name)
}

func (r *Image) ToSlice() ([]byte, error) {
	if r.isNull() {
		return nil, ImagesIsNull
	}
	img := r.Images[0]
	return []byte(img), nil
}

func (r *Image) ToString() (string, error) {
	if r.isNull() {
		return "", ImagesIsNull
	}
	return r.Images[0], nil
}

func (r *Image) isNull() bool {
	if r == nil || len(r.Images) <= 0 {
		return true
	}
	return false
}
