package main

import (
	"fmt"
	"net"

	"github.com/aws/aws-sdk-go/service/ec2"
	ec2session "github.com/burizz/whitelist-external-public-ip-aws/aws-session"
	updatesecuritygroup "github.com/burizz/whitelist-external-public-ip-aws/aws-sg"
)

var awsRegion string
var securityGroupIDs []string
var domainName string
var ec2SvcClient *ec2.EC2

// Init Vars
func init() {
	//TODO: Add Env vars
	domainName = "hub.docker.com"
	securityGroupIDs = []string{"sg-00ffabccebd5efda2"}
	awsRegion = "eu-central-1"
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
	// Get domains IP ranges
	ipAddrList, err := net.LookupHost(domainName)
	if err != nil {
		fmt.Println("LookupHost error: %w", err)
	}

	// Check a few times in case IPs change
	for count := 0; count <= 10; count++ {
		ipList, err := net.LookupHost(domainName)
		if err != nil {
			fmt.Println("LookupHost error: %w", err)
		}

		for _, ip := range ipList {
			ipAddrList = appendIfMissing(ipAddrList, ip)
		}
	}

	fmt.Println(ipAddrList)

	// Whitelist IP ranges in SG Egress Rules - port 443
	for _, securityGroup := range securityGroupIDs {
		for _, ipAddr := range ipAddrList {
			if err := updatesecuritygroup.Egress(ec2SvcClient, ipAddr, securityGroup); err != nil {
				fmt.Println("updateEgressErr : %w", err)
			}
		}
	}
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
