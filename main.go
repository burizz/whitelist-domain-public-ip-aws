package main

import (
	"fmt"
)

func init() {
	// AWS Sessions
	// Env vars
	fmt.Println("Add init logic")
}

func main() {
	// DNS Check domain for its external IPs
	// Check it at least 10 times to get a full list of all IPs
	// Parse IP range into a /24
	// Whitelist IP into SG egress rules
	fmt.Println("Main")
}
