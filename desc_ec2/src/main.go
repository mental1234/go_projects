// Describe instances ec2
package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	instance_aws := os.Args[1]
	// Session
        sess, err := session.NewSession(&aws.Config{
                Region: aws.String("us-east-1")},
        )
        // Create EC2 service client
        svc := ec2.New(sess)

	//Input
	input := &ec2.DescribeInstancesInput{
 		InstanceIds: []*string{
        		aws.String(instance_aws),
    		},
	}

	// Describe instances
	result, err := svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Sucess", result)
	}
	fmt.Println(result)
}

