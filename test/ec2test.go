package main

import (
	"github.com/scorelab/gocloud/lib/common/aws"
	"github.com/scorelab/gocloud/lib/common/aws/ec2"
	"fmt"
)

func main() {
	service := authorize();
	//createInstance(service)
	//startInstance(service)
	//rebootInstance(service)
	//stopInstance(service)
	//terminateInstance(service)
	listImages(service);

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

func stopInstance(e * ec2.EC2){
	resp, err := e.StopInstances("i-0b53b39a")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp);
}

func terminateInstance(e * ec2.EC2){
	resp, err := e.TerminateInstances("i-b0906521")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp);
}

func listImages(e * ec2.EC2) {
	var idList =[]string{"ami-0022c769"};
	filter := ec2.NewFilter();
	//filter.Add("architecture","i386")

	resp, err := e.ListImages(idList,filter)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Number of instances ",len(resp.Images));
	for _, image := range resp.Images {
		fmt.Printf("image id : %v  , type : %v\n", image.Id,image.Type)
	}

}
