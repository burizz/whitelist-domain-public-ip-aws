// Package ec2session initializes an AWS session and returns an ec2 service client with that session
package ec2session

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Initialize returns an EC2 svc client with AWS session - default credentials and region (set in ENV vars)
func Initialize(awsRegion string) (*ec2.EC2, error) {
	if isNotValid(awsRegion, ValidRegions) {
		return nil, fmt.Errorf("Invalid AWS Region %v", awsRegion)
	}

	// Create AWS session with default credentials and region (in ENV vars)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return nil, fmt.Errorf("Initialize: Cannot create AWS config sessions: %w", err)
	}

	// Create an AWS EC2 service client
	svc := ec2.New(sess)
	return svc, nil
}

func isNotValid(awsRegion string, validRegions []string) bool {
	for _, validRegion := range validRegions {
		if validRegion == awsRegion {
			return false
		}
	}
	return true
}
