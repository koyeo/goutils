package _robot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/koyeo/goutils/_http"
	"log"
	"net/url"
	"strconv"
	"time"
)

func NewDingTalkRobot(config *DingTalkConfig) *DingTalkRobot {
	return &DingTalkRobot{config: config}
}

type DingTalkRobot struct {
	config   *DingTalkConfig
	messages []*DingTalkMessage
}

func (p *DingTalkRobot) Push(message ...*DingTalkMessage) {
	for _, v := range message {
		p.messages = append(p.messages, v)
	}
	p.request()
}

func (p *DingTalkRobot) sign() (timestamp int64, sign string) {
	timestamp = time.Now().UnixNano() / 1e6
	str := fmt.Sprintf("%d\n%s", timestamp, p.config.SignSecret)
	h := hmac.New(sha256.New, []byte(p.config.SignSecret))
	h.Write([]byte(str))
	sign = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

func (p *DingTalkRobot) request() {
	timestamp, sign := p.sign()
	
	req := _http.NewRequest()
	req.SetContentType(_http.APPLICATION_JSON)
	
	address, err := url.Parse(p.config.Webhook)
	if err != nil {
		log.Println(err)
		return
	}
	query := address.Query()
	query.Add("timestamp", strconv.FormatInt(timestamp, 10))
	query.Add("sign", sign)
	address.RawQuery = query.Encode()
	resp, err := req.Post(address.String(), p.messages[0])
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp.String())
}
