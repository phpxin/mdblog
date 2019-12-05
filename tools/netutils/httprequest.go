package netutils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HttpRequest struct {
	url string
	method string
	body io.Reader
	contentType string
}

func NewHttpRequest(url string) (h *HttpRequest) {
	h = &HttpRequest{
		url:url,
		method:"GET",
		body:nil,
		contentType:"",
	}
	return h
}

func (h *HttpRequest) SetMethod(m string) error {
	m = strings.ToUpper(m)
	if m!="GET" && m!="POST" {
		return fmt.Errorf("only support GET or POST")
	}
	h.method = m
	return nil
}

func (h *HttpRequest) SetBodyStr(b string, contentType string) {
	h.body = strings.NewReader(b)
	h.contentType = contentType
}

func (h *HttpRequest) SetBodyBytes(b []byte, contentType string) {
	h.body = bytes.NewReader(b)
	h.contentType = contentType
}

func (h *HttpRequest) SetBodyFields(fields map[string]string) {
	h.contentType = "application/x-www-form-urlencoded"
	buf := bytes.NewBufferString("")
	for k,v := range fields {
		if buf.Len()>0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
	}
	h.body = buf

}

func (h *HttpRequest) Exec() (result []byte, err error) {
	client := http.Client{}
	req, _ := http.NewRequest(h.method, h.url, h.body)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	if h.method=="POST" && len(h.contentType)>0 {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)

	if err!=nil {
		return nil, err
	}

	defer func (){
		err := resp.Body.Close()
		if err!=nil {
			panic(err)
		}
	}()

	content,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		return nil, err
	}

	if "gzip" == resp.Header.Get("Content-Encoding") {
		r,err := gzip.NewReader(bytes.NewReader(content))
		if err!=nil {
			return nil, err
		}

		rs,err := ioutil.ReadAll(r)
		if err!=nil {
			return nil, err
		}
		result = rs
	}else{
		result = content
	}

	return result,nil
}