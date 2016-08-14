package common


import (
	"net"
	"net/http"
	"time"
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

func (t *ResilientTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.tries(req)
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
