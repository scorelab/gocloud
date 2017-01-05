package aws

// Region defines the URLs where AWS services may be accessed.
type Region struct {
	Name        string
	EC2Endpoint string
	S3Endpoint  string
	IAMEndpoint string
	ELBEndpoint string
}

var Regions = map[string]Region{
	APNortheast.Name:  APNortheast,
	APSoutheast.Name:  APSoutheast,
	APSoutheast2.Name: APSoutheast2,
	EUCentral.Name:    EUCentral,
	EUWest.Name:       EUWest,
	USEast.Name:       USEast,
	USWest.Name:       USWest,
	USWest2.Name:      USWest2,
	USGovWest.Name:    USGovWest,
	SAEast.Name:       SAEast,
	CNNorth.Name:      CNNorth,
}

var USGovWest = Region{
	"us-gov-west-1",
	"https://ec2.us-gov-west-1.amazonaws.com",
	"https://s3-fips-us-gov-west-1.amazonaws.com",
	"https://iam.us-gov.amazonaws.com",
	"https://elasticloadbalancing.us-gov-west-1.amazonaws.com",
}

var USEast = Region{
	"us-east-1",
	"https://ec2.us-east-1.amazonaws.com",
	"https://s3.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-east-1.amazonaws.com",
}

var USWest = Region{
	"us-west-1",
	"https://ec2.us-west-1.amazonaws.com",
	"https://s3-us-west-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-west-1.amazonaws.com",
}

var USWest2 = Region{
	"us-west-2",
	"https://ec2.us-west-2.amazonaws.com",
	"https://s3-us-west-2.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-west-2.amazonaws.com",
}

var EUWest = Region{
	"eu-west-1",
	"https://ec2.eu-west-1.amazonaws.com",
	"https://s3-eu-west-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.eu-west-1.amazonaws.com",
}

var EUCentral = Region{
	"eu-central-1",
	"https://ec2.eu-central-1.amazonaws.com",
	"https://s3-eu-central-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.eu-central-1.amazonaws.com",
}

var APSoutheast = Region{
	"ap-southeast-1",
	"https://ec2.ap-southeast-1.amazonaws.com",
	"https://s3-ap-southeast-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-southeast-1.amazonaws.com",
}

var APSoutheast2 = Region{
	"ap-southeast-2",
	"https://ec2.ap-southeast-2.amazonaws.com",
	"https://s3-ap-southeast-2.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-southeast-2.amazonaws.com",
}

var APNortheast = Region{
	"ap-northeast-1",
	"https://ec2.ap-northeast-1.amazonaws.com",
	"https://s3-ap-northeast-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-northeast-1.amazonaws.com",
}

var SAEast = Region{
	"sa-east-1",
	"https://ec2.sa-east-1.amazonaws.com",
	"https://s3-sa-east-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.sa-east-1.amazonaws.com",
}

var CNNorth = Region{
	"cn-north-1",
	"https://ec2.cn-north-1.amazonaws.com.cn",
	"https://s3.cn-north-1.amazonaws.com.cn",
	"https://iam.cn-north-1.amazonaws.com.cn",
	"https://elasticloadbalancing.cn-north-1.amazonaws.com.cn",
}
