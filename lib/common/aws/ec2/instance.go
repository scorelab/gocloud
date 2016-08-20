package ec2

import (
	"strconv"
)

func (ec2 *EC2) CreateInstances(options *CreateInstancesOptions) (resp *CreateInstancesResp, err error) {
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
	resp = &CreateInstancesResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}

func (ec2 *EC2) RebootInstance(ids ... string) (resp *SimpleResp, err error) {
	params := makeParams("RebootInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &SimpleResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}






