package _robot

import (
	"sync"
	"time"
)

func NewBucket(duration time.Duration) *Bucket {
	b := new(Bucket)
	b.mu = new(sync.Mutex)
	b.duration = duration
	b.first = true
	b.change = false
	b.setNext()
	return b
}

type Bucket struct {
	duration time.Duration
	messages []interface{}
	mu       *sync.Mutex
	next     int64
	change   bool
	first    bool
}

func (p *Bucket) now() int64 {
	return time.Now().UnixNano()
}

func (p *Bucket) setNext() {
	p.next = p.now() + int64(p.duration)
}

func (p *Bucket) call(callback func(messages []interface{})) {
	p.mu.Lock()
	messages := make([]interface{}, len(p.messages))
	copy(messages, p.messages)
	p.mu.Unlock()
	
	callback(messages)
	p.messages = make([]interface{}, 0)
	p.setNext()
}

func (p *Bucket) Push(message interface{}) {
	p.mu.Lock()
	if len(p.messages) == 0 {
		p.change = true
	} else {
		p.change = false
	}
	p.messages = append(p.messages, message)
	p.mu.Unlock()
}

// PushWait 推入消息同时增加下次释放间隔
//func (p *Bucket) PushWait(message interface{}, wait time.Duration) {
//	p.mu.Lock()
//	p.next = p.now() + int64(wait)
//	if len(p.messages) == 0 {
//		p.change = true
//	} else {
//		p.change = false
//	}
//	p.messages = append(p.messages, message)
//	p.mu.Unlock()
//}

// Pop 所有的消息等待时间间隔结束后推出
func (p *Bucket) Pop(callback func(messages []interface{})) {
	for {
		if p.now() >= p.next {
			p.call(callback)
		}
	}
}

// Instant 一个时间间隔里的进入的第一条消息瞬时推出， 其余消息等待时间间隔结束后推出
func (p *Bucket) Instant(callback func(messages []interface{})) {
	for {
		if p.change && p.first {
			p.first = false
			p.call(callback)
		} else if p.now() >= p.next {
			p.first = true
			p.call(callback)
		}
	}
}
