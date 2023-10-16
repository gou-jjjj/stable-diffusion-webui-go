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
	baseUrl string
}

func NewSD(baseUrl string, timeout time.Duration, uselog bool) *SD {
	parse, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	Lg = NewMyLog(uselog)
	return &SD{
		hc:      &http.Client{Timeout: timeout},
		baseUrl: parse.String(),
	}
}

func (s *SD) Txt2Img(t2i *Txt2ImgRequest) (*Image, error) {
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

	Lg.Info("Please wait, the image generation can take a long time")

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

	Lg.Info(fmt.Sprintf("Use for %f Seconds", time.Now().Sub(begin).Seconds()))
	return &Image{Images: rimgs.Images}, nil
}

func (s *SD) Img2Img(i2i *Img2ImgRequest) (*Image, error) {
	var (
		err    error
		imgRes []byte
		reat2i []byte
		i2iReq *http.Response
		begin  = time.Now()

		rimgs = new(Img2ImgResponse)
	)

	reat2i, err = json.Marshal(i2i)
	if err != nil {
		return nil, err
	}

	Lg.Info("Please wait, the image generation can take a long time")

	body := bytes.NewReader(reat2i)
	i2iReq, err = s.hc.Post(fmt.Sprintf("%s/sdapi/v1/img2img", s.baseUrl), "application/json", body)
	if err != nil {
		return nil, err
	}
	defer i2iReq.Body.Close()

	imgRes, err = io.ReadAll(i2iReq.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(imgRes, rimgs)
	if err != nil {
		return nil, err
	}

	Lg.Info(fmt.Sprintf("Use for %f Seconds", time.Now().Sub(begin).Seconds()))
	return &Image{Images: rimgs.Images}, nil
}
