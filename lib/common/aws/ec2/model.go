package ec2

type InstanceType struct {
	Name		string
	CPU		int
	RAM		float32
	Disk		string
	Network		string
}


type CreateInstancesOptions struct {
	ImageId          string
	MinCount         int
	MaxCount         int
	KeyName          string
	InstanceType     string
	AvailabilityZone string
}


type CreateInstancesResp struct {
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
