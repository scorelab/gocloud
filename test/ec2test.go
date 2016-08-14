package main

import (
	"github.com/scorelab/gocloud/lib/common/aws"
	"github.com/scorelab/gocloud/lib/common/aws/ec2"
)

func main() {

	auth, err := aws.EnvAuth();
	if err != nil {
		panic(err.Error())
	}
	e := ec2.New(auth, aws.USEast)
	options := ec2.CreateInstancesOptions{
		ImageId:      "ami-ccf405a5",
		InstanceType: "t1.micro",
	}
	resp, err := e.CreateInstances(&options)
	if err != nil {
		panic(err.Error())
	}

	for _, instance := range resp.Instances {
		println("Now running", instance.InstanceId)
	}
}

