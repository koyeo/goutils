package _http

import "github.com/bitly/go-simplejson"

func NewResponse() *Response {
	return &Response{}
}

type Response struct {
	bytes []byte
}

func (p *Response) Bytes() []byte {
	return p.bytes
}

func (p *Response) String() string {
	return string(p.bytes)
}

func (p *Response) Parse(parser Parser) (data interface{}, err error) {
	return parser(p.bytes)
}

func (p *Response) SimpleJson() (*simplejson.Json, error) {
	return simplejson.NewJson(p.bytes)
}
