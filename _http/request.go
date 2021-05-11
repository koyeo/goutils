package _http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

func NewRequest() *Request {
	return &Request{}
}

type Request struct {
	headers *Headers      // 请求头
	timeout time.Duration // 超时时间
	body    map[string]string
}

func (p *Request) Headers() *Headers {
	if p.headers == nil {
		p.headers = NewHeaders()
	}
	return p.headers
}

func (p *Request) SetHeader(key, value string) *Request {
	p.Headers().Set(key, value)
	return p
}

func (p *Request) SetContentType(contentType string) *Request {
	p.Headers().Set(CONTENT_TYPE, contentType)
	return p
}

func (p *Request) request(method, url string, data interface{}) (response *Response, err error) {
	
	contentType := p.Headers().Get(CONTENT_TYPE)
	if contentType == "" {
		p.SetContentType(APPLICATION_JSON)
	}
	
	body, err := p.formatBody(contentType, data)
	if err != nil {
		err = fmt.Errorf("format request data error: %s", err)
		return
	}
	
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("http.NewRequest error: %s", err)
		return
	}
	
	for key, val := range p.Headers().Headers() {
		req.Header.Add(key, val)
	}
	
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do error: %s", err)
		return
	}
	
	if resp.StatusCode != 200 {
		err = fmt.Errorf("response code error: %d", resp.StatusCode)
		return
	}
	
	response = NewResponse()
	response.bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("reade body error: %s", err)
		return
	}
	
	return
}

func (p *Request) formatBody(contentType string, data interface{}) (body io.Reader, err error) {
	switch contentType {
	case APPLICATION_X_WWW_FORM_URLENCODED:
	case APPLICATION_JSON:
		return p.formatJson(data)
	case TEXT_XML:
	case MULTIPART_FORM_DATA:
	
	}
	return
}

func (p *Request) formatJson(data interface{}) (body io.Reader, err error) {
	
	t := reflect.TypeOf(data)
	switch t.Kind().String() {
	case "slice":
		if t.Elem().String() == "uint8" {
			body = bytes.NewReader(data.([]byte))
			return
		}
	case "string":
		body = bytes.NewReader([]byte(data.(string)))
		return
	}
	
	// TODO 判断其它类型
	
	var bs []byte
	bs, err = json.Marshal(data)
	if err != nil {
		return
	}
	body = bytes.NewReader(bs)
	
	return
}

func (p *Request) Get(url string) (*Response, error) {
	return p.request(GET, url, nil)
}

func (p *Request) Post(url string, data interface{}) (*Response, error) {
	return p.request(POST, url, data)
}
