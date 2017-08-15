package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::NetworkInterface.InstanceIpv6Address AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html
type AWSEC2NetworkInterface_InstanceIpv6Address struct {

	// Ipv6Address AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html#cfn-ec2-networkinterface-instanceipv6address-ipv6address

	Ipv6Address string `json:"Ipv6Address"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkInterface_InstanceIpv6Address) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterface.InstanceIpv6Address"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkInterface_InstanceIpv6Address) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkInterface_InstanceIpv6AddressResources retrieves all AWSEC2NetworkInterface_InstanceIpv6Address items from a CloudFormation template
func GetAllAWSEC2NetworkInterface_InstanceIpv6Address(template *Template) map[string]*AWSEC2NetworkInterface_InstanceIpv6Address {

	results := map[string]*AWSEC2NetworkInterface_InstanceIpv6Address{}
	for name, resource := range template.Resources {
		result := &AWSEC2NetworkInterface_InstanceIpv6Address{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkInterface_InstanceIpv6AddressWithName retrieves all AWSEC2NetworkInterface_InstanceIpv6Address items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2NetworkInterface_InstanceIpv6Address(name string, template *Template) (*AWSEC2NetworkInterface_InstanceIpv6Address, error) {

	result := &AWSEC2NetworkInterface_InstanceIpv6Address{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkInterface_InstanceIpv6Address{}, errors.New("resource not found")

}