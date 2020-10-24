package main

import (
	"fmt"
	"net"
	"reflect"

	"github.com/aws/aws-sdk-go/service/ec2"
	session "github.com/burizz/whitelist-external-public-ip-aws/aws-session"
)

var awsRegion string
var securityGroupIDs []string
var domainName string
var ec2ClienSess *ec2.EC2

// Init Vars
func init() {
	//TODO: Add Env vars
	awsRegion = "eu-central-1"
	securityGroupIDs = []string{"sg-00ffabccebd5efda2"}
	domainName = "hub.docker.com"
}

// Init Sessions
func init() {
	var err error
	// Initialized AWS EC2 Client Session
	ec2ClienSess, err = session.Initialize(awsRegion)
	if err != nil {
		fmt.Println("Init Err: ", err)
	}
}

func main() {
	// TODO: Check multiple times to get a full list of IPs in case they change
	addr, lookUperr := net.LookupHost(domainName)
	if lookUperr != nil {
		fmt.Println("LookupHost error: %w", lookUperr)
	}

	fmt.Println(addr)
	fmt.Println(reflect.TypeOf(ec2ClienSess))
	// Whitelist IP into SG egress rules
}
