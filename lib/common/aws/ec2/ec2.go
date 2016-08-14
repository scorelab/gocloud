package ec2

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
	"github.com/scorelab/gocloud/lib/common/aws"
)

const debug = false

type EC2 struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
	private    byte
	aws.Service
}


func New(auth aws.Auth, region aws.Region) *EC2 {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, client *http.Client) *EC2 {
	s,_ :=aws.NewService(auth,aws.ServiceInfo{region.EC2Endpoint,0})
	return &EC2{auth, region, client, 0, *s} //check 0
}

var timeNow = time.Now

func (ec2 *EC2) query(params map[string]string, resp interface{}) error {
	params["Version"] = "2014-02-01"
	params["Timestamp"] = timeNow().In(time.UTC).Format(time.RFC3339)
	endpoint, err := url.Parse(ec2.Region.EC2Endpoint)
	if err != nil {
		return err
	}
	if endpoint.Path == "" {
		endpoint.Path = "/"
	}
	endpoint.RawQuery = multimap(params).Encode()
	if debug {
		log.Printf("get { %v } -> {\n", endpoint.String())
	}
	r, err := ec2.httpClient.Get(endpoint.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n}\n", string(dump))
	}
	if r.StatusCode != 200 {
		return buildError(r)
	}
	err = xml.NewDecoder(r.Body).Decode(resp)
	return err
}


