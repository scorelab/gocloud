package main

import (
	"github.com/scorelab/gocloud/lib/common/aws"
	"github.com/scorelab/gocloud/lib/common/aws/ec2"
	"fmt"
)

func main() {
	service := authorize();
	createInstance(service)
}


func authorize()(e *ec2.EC2){
	auth, err := aws.EnvAuth();
	if err != nil {
		panic(err.Error())
	}
	e = ec2.New(auth, aws.USEast)
	return
}

func createInstance(e *ec2.EC2){
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


func rebootInstance(e *ec2.EC2){
	resp, err := e.RebootInstance("i-b0906521")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp);
}

func startInstance(e * ec2.EC2){
	resp, err := e.StartInstances("i-16b58d18")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp);
}