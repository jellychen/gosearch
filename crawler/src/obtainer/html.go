package obtainer

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Html struct {
	client  *http.Client
	request *http.Request
	url     string
	timeout uint16
}

func NewHtml() *Html {
	return &Html{nil, nil, "", 20}
}

func (self *Html) LoadUrl(url string, timeout uint16) bool {
	self.url = url
	if timeout != self.timeout {
		self.timeout = timeout
		self.client = nil
	}

	if nil == self.client {
		duration := (time.Duration)(timeout)
		self.client = &http.Client{
			Timeout: duration * time.Second,
		}
	}
	return true
}

func (self *Html) SetProxy(addr string) bool {
	if nil == self.client {
		return false
	}

	urli := url.URL{}
	urlproxy, err := urli.Parse(addr)
	if nil != err {
		return false
	}
	self.client.Transport = &http.Transport{
		Proxy: http.ProxyURL(urlproxy),
	}
	return true
}

func (self *Html) Get(arguments string) bool {
	reader := strings.NewReader(arguments)
	request, err := http.NewRequest("GET", self.url, reader)
	self.request = request
	return nil == err
}

func (self *Html) HeaderSet(k, v string) bool {
	if nil == self.request {
		return false
	}
	self.request.Header.Set(k, v)
	return true
}

func (self *Html) SetUserAgent(v string) bool {
	return self.HeaderSet("User-Agent", v)
}

func (self *Html) WaitContent() (string, bool) {
	res, err := self.client.Do(self.request)
	if nil != err || nil == res {
		return "", false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if nil != err {
		return "", false
	}
	return (string)(body), true
}
