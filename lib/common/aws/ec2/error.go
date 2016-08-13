package ec2

import (
	"fmt"
	"net/http"
	"encoding/xml"
)


// Error encapsulates an error returned by EC2.
type Error struct {
	StatusCode int
	Code       string // EC2 error code ("UnsupportedOperation", ...)
	Message    string
	RequestId  string `xml:"RequestID"`
}

func buildError(r *http.Response) error {
	errors := xmlErrors{}
	xml.NewDecoder(r.Body).Decode(&errors)
	var err Error
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

func (err *Error) Error() string {
	if err.Code == "" {
		return err.Message
	}

	return fmt.Sprintf("%s (%s)", err.Message, err.Code)
}

type xmlErrors struct {
	RequestId string  `xml:"RequestID"`
	Errors    []Error `xml:"Errors>Error"`
}

