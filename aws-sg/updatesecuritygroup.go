package updatesecuritygroup

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Egress - updates a security groups Egress rules to whitelist destination IP address on port 443
func Egress(ec2ClientSvc *ec2.EC2, ipForWhitelist string, securityGroup string) error {
	// Define Egress rule Input
	input := &ec2.AuthorizeSecurityGroupEgressInput{
		GroupId: aws.String(securityGroup),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(443),
				ToPort:     aws.Int64(443),
				IpProtocol: aws.String("tcp"),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String(ipForWhitelist),
					},
				},
			},
		},
	}

	// Update Security Group Egress rule from Input
	_, err := ec2ClientSvc.AuthorizeSecurityGroupEgress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "InvalidPermission.Duplicate":
				fmt.Printf("Skipping - [%v] already exists in [%v] Egress rules\n", ipForWhitelist, securityGroup)
			default:
				return fmt.Errorf("awsUpdateSg: %v", aerr)
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and Message from an error.
			fmt.Println(err.Error())
		}
	} else {
		fmt.Printf("IP Range [%v] added successfully\n", ipForWhitelist)
	}
	return nil
}
