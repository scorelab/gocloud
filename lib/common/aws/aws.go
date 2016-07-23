package aws

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Defines the valid signers
const (
	V2Signature = iota
)

// Defines the service endpoint and correct Signer implementation to use to sign requests for this endpoint
type ServiceInfo struct {
	Endpoint string
	Signer   uint
}

// Create a base set of params for an action
func MakeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}

// Create a new AWS server to handle making requests
func NewService(auth Auth, service ServiceInfo) (s *Service, err error) {
	var signer Signer
	switch service.Signer {
	case V2Signature:
		signer, err = NewV2Signer(auth, service)
	default:
		err = fmt.Errorf("Unsupported signer for service")
	}
	if err != nil {
		return
	}
	s = &Service{service: service, signer: signer}
	return
}



// Implements a Server Query/Post API to easily query AWS services and build
// errors when desired
type Service struct {
	service ServiceInfo
	signer  Signer
}

func (s *Service) Query(method, path string, params map[string]string) (resp *http.Response, err error) {
	params["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	u, err := url.Parse(s.service.Endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path
	s.signer.Sign(method, path, params)
	if method == "GET" {
		u.RawQuery = multimap(params).Encode()
		resp, err = http.Get(u.String())
	} else if method == "POST" {
		resp, err = http.PostForm(u.String(), multimap(params))
	}
	return
}




func (s *Service) BuildError(r *http.Response) error {
	errors := ErrorResponse{}
	xml.NewDecoder(r.Body).Decode(&errors)
	var err Error
	err = errors.Errors
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}





type ErrorResponse struct {
	Errors    Error  `xml:"Error"`
	RequestId string // A unique ID for tracking the request
}

type Error struct {
	StatusCode int
	Type       string
	Code       string
	Message    string
	RequestId  string
}

func (err *Error) Error() string {
	return fmt.Sprintf("Type: %s, Code: %s, Message: %s",
		err.Type, err.Code, err.Message,
	)
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

