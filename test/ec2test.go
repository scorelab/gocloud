package main

import (
	"github.com/scorelab/gocloud/lib/common/aws"
	"github.com/scorelab/gocloud/lib/common/aws/ec2"
)

func main() {
	auth, err := aws.SharedAuth();
	if err != nil {
		panic(err.Error())
	}
	e := ec2.New(auth, aws.USEast)
	options := ec2.RunInstancesOptions{
		ImageId:      "ami-ccf405a5", // Ubuntu Maverick, i386, EBS store
		InstanceType: "t1.micro",
	}
	resp, err := e.RunInstances(&options)
	if err != nil {
		panic(err.Error())
	}

	for _, instance := range resp.Instances {
		println("Now running", instance.InstanceId)
	}
}

func createInstance() {

	//return resp.Instances
}