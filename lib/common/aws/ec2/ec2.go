package ec2

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/http/httputil"
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

func (ec2 *EC2) query(params map[string]string, resp interface{}) error {
	r, err := ec2.Service.Query("GET","/",params);
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
