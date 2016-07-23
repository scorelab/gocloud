package common

import (
	"net"
	"net/http"
	"time"
	"math"
)

type RetryableFunc func(*http.Request, *http.Response, error) bool
type WaitFunc func(try int)
type DeadlineFunc func() time.Time

type ResilientTransport struct {
	DialTimeout time.Duration
	MaxTries    int
	Deadline    DeadlineFunc
	ShouldRetry RetryableFunc
	Wait        WaitFunc
	transport   *http.Transport
}

//method for creating an http client
func NewClient(rt *ResilientTransport) *http.Client {
	rt.transport = &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, rt.DialTimeout)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(rt.Deadline())
			return c, nil
		},
		Proxy: http.ProxyFromEnvironment,
	}
	return &http.Client{
		Transport: rt,
	}
}

var retryingTransport = &ResilientTransport{
	Deadline: func() time.Time {
		return time.Now().Add(5 * time.Second)
	},
	DialTimeout: 10 * time.Second,
	MaxTries:    3,
	ShouldRetry: func(*http.Request, *http.Response, error) bool { return true},
	Wait:        ExpBackoff,
}

// Exported default client
var RetryingClient = NewClient(retryingTransport)

func (t *ResilientTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.tries(req)
}

func ExpBackoff(try int) {
	time.Sleep(100 * time.Millisecond *
	time.Duration(math.Exp2(float64(try))))
}


func (t *ResilientTransport) tries(req *http.Request) (res *http.Response, err error) {
	for try := 0; try < t.MaxTries; try += 1 {
		res, err = t.transport.RoundTrip(req)

		if !t.ShouldRetry(req, res, err) {
			break
		}
		if res != nil {
			res.Body.Close()
		}
		if t.Wait != nil {
			t.Wait(try)
		}
	}

	return
}
