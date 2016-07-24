package ec2

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
