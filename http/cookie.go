package http

import "sync"

type Cookie struct {
	sync.Mutex
	cookie string
}

func NewCookie(s string) *Cookie {
	return &Cookie{
		cookie: s,
	}
}

func (c *Cookie) SetCookie(s string) {
	c.Lock()
	defer c.Unlock()
	c.cookie = s
}

func (c *Cookie) GetCookie() string {
	c.Lock()
	defer c.Unlock()
	return c.cookie
}
