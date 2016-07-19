package ec2

import "fmt"


// Error encapsulates an error returned by EC2.
type Error struct {
	StatusCode int
	Code       string // EC2 error code ("UnsupportedOperation", ...)
	Message    string
	RequestId  string `xml:"RequestID"`
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

