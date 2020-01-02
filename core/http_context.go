package core

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type HttpContext struct {
	RawReq    *http.Request
	SessionId string
}

func (ctx *HttpContext) get(key string) (string, error) {
	qStr := ctx.RawReq.URL.Query()

	rawResult := qStr.Get(key)

	if rawResult == "" {
		b := ctx.RawReq.Body
		bodybytes := make([]byte, 1024)
		rl, err := b.Read(bodybytes)
		if err != nil {
			return "", err
		}

		if rl <= 0 {
			return "", nil
		}

		// todo application/json, x-http-request-urlencoded
		qStr, err = url.ParseQuery(string(bodybytes))
		if err != nil {
			return "", err
		}
		rawResult = qStr.Get(key)
	}

	return rawResult, nil
}

func (ctx *HttpContext) GetString(key string, def string) (string, error) {
	src, err := ctx.get(key)
	if err != nil {
		return def, err
	}

	hackStrings := []string{
		"'",
		"script",
		"frame",
	}

	srcLowercase := strings.ToLower(src)

	for _, v := range hackStrings {
		// src = strings.Replace(src, v, "", -1)
		if strings.Index(srcLowercase, v) > -1 {
			return "", fmt.Errorf("the paramter has invalid words")
		}
	}

	return src, nil
}

func (ctx *HttpContext) GetInt32(key string, def int32) (int32, error) {

	src, err := ctx.get(key)
	if err != nil {
		return def, err
	}

	if src == "" {
		return def, fmt.Errorf("not found")
	}

	r, err := strconv.ParseInt(src, 10, 32)
	if err != nil {
		return def, err
	}

	return int32(r), nil
}

func (ctx *HttpContext) GetInt64(key string, def int64) (int64, error) {

	src, err := ctx.get(key)
	if err != nil {
		return def, err
	}

	if src == "" {
		return def, fmt.Errorf("not found")
	}

	r, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return def, err
	}

	return r, nil
}

func (ctx *HttpContext) GetFloat64(key string, def float64) (float64, error) {

	src, err := ctx.get(key)
	if err != nil {
		return def, err
	}

	if src == "" {
		return def, fmt.Errorf("not found")
	}

	r, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return def, err
	}

	return r, nil
}
