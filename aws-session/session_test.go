package session

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockEC2Client struct {
	ec2iface.EC2API
}

func TestInitialize(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"eu-central-1", ""},
		{"us-east-1", ""},
		{"eu-west-1", ""},
		{"wrong-region", ""},
		{"a!@4%^&*", ""},
	}

	fmt.Println(tests)

	// ec2SvcClient, _ := Initialize("eu-central-1")

	// 	if reflect.ValueOf(ec2SvcClient) {
	// 		fmt.Println("wrong type")
	// 	}
}
