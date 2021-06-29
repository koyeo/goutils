package _status

import (
	"fmt"
)

type Status struct {
	clone   bool
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (p *Status) fork() *Status {
	var r *Status
	if p.clone {
		r = p
	} else {
		r = &Status{Code: p.Code, Message: p.Message, clone: true}
	}
	return r
}

func (p *Status) Detailf(format string, args ...interface{}) *Status {
	r := p.fork()
	r.Detail = fmt.Sprintf(format, args)
	return r
}

func (p *Status) Messagef(args ...interface{}) *Status {
	r := p.fork()
	r.Message = fmt.Sprintf(p.Message, args)
	return r
}

func (p *Status) With(detail interface{}) *Status {
	r := p.fork()
	r.Detail = fmt.Sprintf("%v", detail)
	return r
}
