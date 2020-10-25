package ec2session

import (
	"fmt"
	"testing"
)

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
