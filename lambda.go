package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/ec2"
	ec2session "github.com/burizz/whitelist-external-public-ip-aws/aws-session"
	updatesecuritygroup "github.com/burizz/whitelist-external-public-ip-aws/aws-sg"
)

var awsRegion string
var securityGroupIDs []string
var domainNames []string
var ec2SvcClient *ec2.EC2

// Init Vars
func init() {
	domainNames = strings.Split(os.Getenv("domainNames"), ",")
	securityGroupIDs = strings.Split(os.Getenv("securityGroupIDs"), ",")
	awsRegion = os.Getenv("awsRegion")
}

// Init Sessions
func init() {
	var err error
	// Initialized AWS EC2 Client Session
	ec2SvcClient, err = ec2session.Initialize(awsRegion)
	if err != nil {
		fmt.Println("Init Err: ", err)
	}
}

func main() {
	lambda.Start(LambdaHandler)
}

// LambdaHandler - entrypoint
func LambdaHandler() (msg string, err error) {
	var ipAddrList []string
	var errMsg string = "Error"

	// Get domain IP ranges
	if len(domainNames) == 0 || len(securityGroupIDs) == 0 || awsRegion == "" {
		// fmt.Println("Err: Need to set all environment variables - domainNames, securityGroupIDs, awsRegion")
		return errMsg, fmt.Errorf("Err: Missing Env Var; need to set all environment variables - domainNames, securityGroupIDs, awsRegion")
	}

	for _, domain := range domainNames {
		// Check a few times in case IPs change
		for count := 0; count <= 10; count++ {
			ipList, err := net.LookupHost(domain)
			if err != nil {
				// fmt.Println("LookupHost Err: ", err)
				return errMsg, fmt.Errorf("LookupHost Err: %v", err)
			}

			for _, ip := range ipList {
				ipAddrList = appendIfMissing(ipAddrList, ip)
			}
		}
	}

	// Whitelist IP ranges in SG Egress Rules - port 443
	for _, securityGroup := range securityGroupIDs {
		for _, ipAddr := range ipAddrList {
			ipAddr = ipAddr + "/32"
			if err := updatesecuritygroup.Egress(ec2SvcClient, ipAddr, securityGroup); err != nil {
				// fmt.Println("updateEgressErr: ", err)
				return errMsg, fmt.Errorf("updateEgressErr: %v", err)
			}
		}
	}
	fmt.Println("Lambda exeuction completed successfully")
	return fmt.Sprintf("Success"), nil
}

// Check if IP is in list, skip if it is
func appendIfMissing(ipList []string, i string) []string {
	for _, ip := range ipList {
		if ip == i {
			return ipList
		}
	}
	return append(ipList, i)
}
