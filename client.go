package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SD struct {
	hc      *http.Client
	log     *MyLog
	baseUrl string
}

func NewSD(baseUrl string, timeout time.Duration, uselog bool) *SD {
	parse, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	return &SD{
		hc:      &http.Client{Timeout: timeout},
		baseUrl: parse.String(),
		log:     NewMyLog(uselog),
	}
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

	s.log.Info("Please wait, the image generation can take a long time")

	body := bytes.NewReader(reat2i)
	t2iReq, err = s.hc.Post(fmt.Sprintf("%s/sdapi/v1/txt2img", s.baseUrl), "application/json", body)
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

	s.log.Info(fmt.Sprintf("Use for %f Seconds", time.Now().Sub(begin).Seconds()))
	return rimgs, nil
}
