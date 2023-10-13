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

	NewImage, err := os.Create(path + name)
	if err != nil {
		return err
	}

	_, err = NewImage.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("\u001B[1;32mThe image is saved successfully\u001B[0m")
	return nil
}

func Seed() int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return int(r.Int31())
}
