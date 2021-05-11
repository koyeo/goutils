package _http

func NewHeaders() *Headers {
	return &Headers{}
}

type Headers struct {
	headers map[string]string
}

func (p *Headers) Headers() map[string]string {
	if p.headers == nil {
		p.headers = map[string]string{}
	}
	return p.headers
}

func (p *Headers) Set(name, val string) {
	p.Headers()[name] = val
}

func (p *Headers) Get(name string) string {
	v, _ := p.Headers()[name]
	return v
}
