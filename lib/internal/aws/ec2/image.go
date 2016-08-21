package ec2

import "strconv"

func (ec2 *EC2) ListImages(ids []string, filter *Filter) (resp *ImagesResp, err error) {
	params := makeParams("DescribeImages")
	for i, id := range ids {
		params["ImageId."+strconv.Itoa(i+1)] = id
	}
	filter.addParams(params)
	resp = &ImagesResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}
