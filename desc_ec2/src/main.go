// Describe instances ec2
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
//	fmt.Println("Hello World!")
	// Session
        sess, err := session.NewSession(&aws.Config{
                Region: aws.String("us-east-1")},
        )
        // Create EC2 service client
        svc := ec2.New(sess)

	//Input
	input := &ec2.DescribeInstancesInput{
 		InstanceIds: []*string{
        		aws.String("INSTANCE-ID"),
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

