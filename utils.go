package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func ImageStore(img string, path string, name string) error {
	if img == "" {
		return ImagesIsNull
	}

	if name == "" || path == "" {
		return errors.New("parameters cannot be empty")
	}

	data, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		return err
	}

	if _, err = os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0750)
		if err != nil {
			return err
		}
	}

	if path[len(path)-1] != '/' {
		path += "/"
	}

	ext := filepath.Ext(name)
	if ext != "png" && ext != "jpg" {
		name += ".png"
	}

	newImage, err := os.Create(path + name)
	if err != nil {
		return err
	}

	_, err = newImage.Write(data)
	if err != nil {
		return err
	}

	Lg.Info(fmt.Sprintf("The image is saved to [%s]", path+name))
	return nil
}

func Seed() int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return int(r.Int31())
}

func Base64(b []byte) string {
	// 将 []byte 转换为 Base64 编码的字符串
	encoded := base64.StdEncoding.EncodeToString(b)
	return encoded
}
