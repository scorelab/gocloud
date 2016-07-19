package ec2

import (
	"strconv"
	"encoding/hex"
	"crypto/rand"
)

type RunInstancesOptions struct {
	ImageId          string
	MinCount         int
	MaxCount         int
	KeyName          string
	InstanceType     string
	AvailabilityZone string
}

// Response to a RunInstances request.
type RunInstancesResp struct {
	RequestId      string          `xml:"requestId"`
	ReservationId  string          `xml:"reservationId"`
	OwnerId        string          `xml:"ownerId"`
	SecurityGroups []SecurityGroup `xml:"groupSet>item"`
	Instances      []Instance      `xml:"instancesSet>item"`
}

type SecurityGroup struct {
	Id          string `xml:"groupId"`
	Name        string `xml:"groupName"`
	Description string `xml:"groupDescription"`
	VpcId       string `xml:"vpcId"`
}
type Instance struct {
	InstanceId       string              `xml:"instanceId"`
	InstanceType     string              `xml:"instanceType"`
	AvailabilityZone string              `xml:"placement>availabilityZone"`
	Reason           string              `xml:"reason"`
	ImageId          string              `xml:"imageId"`
	KeyName          string              `xml:"keyName"`
	Monitoring       string              `xml:"monitoring>state"`
	LaunchTime       string              `xml:"launchTime"`
}

// RunInstances starts new instances in EC2.
func (ec2 *EC2) RunInstances(options *RunInstancesOptions) (resp *RunInstancesResp, err error) {
	params := makeParams("RunInstances")
	params["ImageId"] = options.ImageId
	params["InstanceType"] = options.InstanceType
	var min, max int
	if options.MinCount == 0 && options.MaxCount == 0 {
		min = 1
		max = 1
	} else if options.MaxCount == 0 {
		min = options.MinCount
		max = min
	} else {
		min = options.MinCount
		max = options.MaxCount
	}
	params["MinCount"] = strconv.Itoa(min)
	params["MaxCount"] = strconv.Itoa(max)
	token, err := clientToken()
	if err != nil {
		return nil, err
	}
	params["ClientToken"] = token
	if options.KeyName != "" {
		params["KeyName"] = options.KeyName
	}
	resp = &RunInstancesResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}

func clientToken() (string, error) {
	// Maximum EC2 client token size is 64 bytes.
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}



