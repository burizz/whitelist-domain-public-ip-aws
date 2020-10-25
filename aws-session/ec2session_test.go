package ec2session

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

var ec2SvcClient *ec2.EC2

func TestInitialize(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"eu-central-1"},
		{"us-east-1"},
		{"eu-west-1"},
	}

	var err error

	for _, test := range tests {
		fmt.Println(test.input)
		ec2SvcClient, err = Initialize(test.input)
		if err != nil {
			t.Error("Test failed: {} input, received: {}", test.input, err)
		}
	}
}
