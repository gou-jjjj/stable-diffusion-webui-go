package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SD struct {
	hc *http.Client

	ip      string
	port    int
	timeout time.Duration
}

func NewSD(ip string, port int, timeout time.Duration) *SD {
	return &SD{
		hc:      &http.Client{Timeout: 20 * time.Minute},
		ip:      ip,
		port:    port,
		timeout: timeout,
	}
}

func DefaultSD() *SD {
	return NewSD("127.0.0.1", 7861, 60*time.Minute)
}

func (s SD) Txt2Img(t2i Txt2ImgRequest) (*Txt2imgResponse, error) {
	var (
		err    error
		imgRes []byte
		reat2i []byte
		t2iReq *http.Response
		begin  = time.Now()

		rimgs = new(Txt2imgResponse)
	)

	reat2i, err = json.Marshal(t2i)
	if err != nil {
		return nil, err
	}
	fmt.Println("\033[1;32mPlease wait, the image generation can take a long time\033[0m")
	body := bytes.NewReader(reat2i)
	t2iReq, err = s.hc.Post(fmt.Sprintf("http://%s:%d/sdapi/v1/txt2img", s.ip, s.port), "application/json", body)
	if err != nil {
		return nil, err
	}
	defer t2iReq.Body.Close()

	imgRes, err = io.ReadAll(t2iReq.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(imgRes, rimgs)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\u001B[1;32mUse for %f Seconds\n\u001B[0m", time.Now().Sub(begin).Seconds())
	return rimgs, nil
}
