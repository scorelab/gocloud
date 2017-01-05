package aws

import (
	"math"
	"net"
	"net/http"
	"time"
	"github.com/scorelab/gocloud/lib/common/common"
)

var retryingTransport = &common.ResilientTransport{
	Deadline: func() time.Time {
		return time.Now().Add(5 * time.Second)
	},
	DialTimeout: 10 * time.Second,
	MaxTries:    3,
	ShouldRetry: awsRetry,
	Wait:        ExpBackoff,
}

// Exported default client
var RetryingClient = common.NewClient(retryingTransport)

// Decide if we should retry a request.
func awsRetry(req *http.Request, res *http.Response, err error) bool {
	retry := false
	if neterr, ok := err.(net.Error); ok {
		if neterr.Temporary() {
			retry = true
		}
	}
	if res != nil {
		if res.StatusCode >= 500 && res.StatusCode < 600 {
			retry = true
		}
	}

	return retry
}

func ExpBackoff(try int) {
	time.Sleep(100 * time.Millisecond *
	time.Duration(math.Exp2(float64(try))))
}
