package ec2

import "encoding/xml"

type InstanceType struct {
	Name    string
	CPU     int
	RAM     float32
	Disk    string
	Network string
}

type instanceTypes struct {
	T2_NANO     InstanceType
	T2_MICRO    InstanceType
	T2_SMALL    InstanceType
	T2_MEDIUM   InstanceType
	T2_LARGE    InstanceType
	M4_LARGE    InstanceType
	M4_XLARGE   InstanceType
	M4_2XLARGE  InstanceType
	M4_4XLARGE  InstanceType
	M4_10XLARGE InstanceType
	M3_MEDIUM   InstanceType
	M3_LARGE    InstanceType
	M3_XLARGE   InstanceType
	M3_2XLARGE  InstanceType
	C4_LARGE    InstanceType
	C4_XLARGE   InstanceType
	C4_2XLARGE  InstanceType
	C4_4XLARGE  InstanceType
	C4_8XLARGE  InstanceType
	C3_LARGE    InstanceType
	C3_XLARGE   InstanceType
	C3_2XLARGE  InstanceType
	C3_4XLARGE  InstanceType
	C3_8XLARGE  InstanceType
	G2_2XLARGE  InstanceType
	G2_8XLARGE  InstanceType
	X1_32XLARGE InstanceType
	R3_LARGE    InstanceType
	R3_XLARGE   InstanceType
	R3_2XLARGE  InstanceType
	R3_4XLARGE  InstanceType
	R3_8XLARGE  InstanceType
	I2_XLARGE   InstanceType
	I2_2XLARGE  InstanceType
	I2_4XLARGE  InstanceType
	I2_8XLARGE  InstanceType
	D2_XLARGE   InstanceType
	D2_2XLARGE  InstanceType
	D2_4XLARGE  InstanceType
	D2_8XLARGE  InstanceType
}

var InstanceTypes = instanceTypes{
	InstanceType{"t2.nano", 1, 0.5, "EBS Only", "Low"},
	InstanceType{"t2.micro", 1, 1, "EBS Only", "Low to Moderate"},
	InstanceType{"t2.small", 1, 2, "EBS Only", "Low to Moderate"},
	InstanceType{"t2.medium", 2, 4, "EBS Only", "Low to Moderate"},
	InstanceType{"t2.large", 2, 8, "EBS Only", "Low to Moderate"},
	InstanceType{"m4.large", 2, 8, "EBS Only", "Moderate"},
	InstanceType{"m4.xlarge", 4, 16, "EBS Only", "High"},
	InstanceType{"m4.2xlarge", 8, 32, "EBS Only", "High"},
	InstanceType{"m4.4xlarge", 16, 64, "EBS Only", "High"},
	InstanceType{"m4.10xlarge", 40, 160, "EBS Only", "10 Gigabit"},
	InstanceType{"m3.medium", 1, 3.75, "1 x 4 SSD", "Moderate"},
	InstanceType{"m3.large", 2, 7.5, "1 x 32 SSD", "Moderate"},
	InstanceType{"m3.xlarge", 4, 15, "2 x 40 SSD", "High"},
	InstanceType{"m3.2xlarge", 8, 30, "2 x 80 SSD", "High"},
	InstanceType{"c4.large", 2, 3.75, "EBS Only", "Moderate"},
	InstanceType{"c4.xlarge", 4, 7.5, "EBS Only", "High"},
	InstanceType{"c4.2xlarge", 8, 15, "EBS Only", "High"},
	InstanceType{"c4.4xlarge", 16, 30, "EBS Only", "High"},
	InstanceType{"c4.8xlarge", 36, 60, "EBS Only", "10 Gigabit"},
	InstanceType{"c3.large", 2, 3.75, "2 x 16 SSD", "Moderate"},
	InstanceType{"c3.xlarge", 4, 7.5, "2 x 40 SSD", "Moderate"},
	InstanceType{"c3.2xlarge", 8, 15, "2 x 80 SSD", "High"},
	InstanceType{"c3.4xlarge", 16, 30, "2 x 160 SSD", "High"},
	InstanceType{"c3.8xlarge", 32, 60, "2 x 320 SSD", "10 Gigabit"},
	InstanceType{"g2.2xlarge", 8, 15, "1 x 60 SSD", "High"},
	InstanceType{"g2.8xlarge", 32, 60, "2 x 120 SSD", "10 Gigabit"},
	InstanceType{"x1.32xlarge", 128, 1952, "2 x 1,920 SSD", "20 Gigabit"},
	InstanceType{"r3.large", 2, 15.25, "1 x 32 SSD", "Moderate"},
	InstanceType{"r3.xlarge", 4, 30.5, "1 x 80 SSD", "Moderate"},
	InstanceType{"r3.2xlarge", 8, 61, "1 x 160 SSD", "High"},
	InstanceType{"r3.4xlarge", 16, 122, "1 x 320 SSD", "High"},
	InstanceType{"r3.8xlarge", 32, 244, "2 x 320 SSD", "10 Gigabit"},
	InstanceType{"i2.xlarge", 4, 30.5, "1 x 800 SSD", "Moderate"},
	InstanceType{"i2.2xlarge", 8, 61, "2 x 800 SSD", "High"},
	InstanceType{"i2.4xlarge", 16, 122, "4 x 800 SSD", "High"},
	InstanceType{"i2.8xlarge", 32, 244, "8 x 800 SSD", "10 Gigabit"},
	InstanceType{"d2.xlarge", 4, 30.5, "3 x 2000", "Moderate"},
	InstanceType{"d2.2xlarge", 8, 61, "6 x 2000", "High"},
	InstanceType{"d2.4xlarge", 16, 122, "12 x 2000", "High"},
	InstanceType{"d2.8xlarge", 36, 244, "24 x 2000", "10 Gigabit"},
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

type Image struct {
	Id                 string               `xml:"imageId"`
	Name               string               `xml:"name"`
	Description        string               `xml:"description"`
	Type               string               `xml:"imageType"`
	State              string               `xml:"imageState"`
	Location           string               `xml:"imageLocation"`
	Public             bool                 `xml:"isPublic"`
	Architecture       string               `xml:"architecture"`
	Platform           string               `xml:"platform"`
	ProductCodes       []string             `xml:"productCode>item>productCode"`
	KernelId           string               `xml:"kernelId"`
	RamdiskId          string               `xml:"ramdiskId"`
	StateReason        string               `xml:"stateReason"`
	OwnerId            string               `xml:"imageOwnerId"`
	OwnerAlias         string               `xml:"imageOwnerAlias"`
	RootDeviceType     string               `xml:"rootDeviceType"`
	RootDeviceName     string               `xml:"rootDeviceName"`
	VirtualizationType string               `xml:"virtualizationType"`
	Hypervisor         string               `xml:"hypervisor"`
}

type ImagesResp struct {
	RequestId string  `xml:"requestId"`
	Images    []Image `xml:"imagesSet>item"`
}

type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type SimpleResp struct {
	XMLName   xml.Name
	RequestId string `xml:"requestId"`
}


type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

type InstanceStateChange struct {
	InstanceId    string        `xml:"instanceId"`
	CurrentState  InstanceState `xml:"currentState"`
	PreviousState InstanceState `xml:"previousState"`
}

type InstanceStateReason struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

type StartInstanceResp struct {
	RequestId    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}