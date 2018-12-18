// aws create an instance
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"fmt"
)

func main() {
	// Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	// Create EC2 service client
	svc := ec2.New(sess)

	// Specify the details of the instance that you want to create
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
	// An Amazon Linux AMI ID for t2.micro instances
		ImageId: 	aws.String("ami-07c20c8dd8c3136bf"),
		InstanceType:	aws.String("t2.micro"),
		MinCount:	aws.Int64(1),
		MaxCount:	aws.Int64(1),
	})

	if err != nil {
		fmt.Println("Could not create instance :(", err)
		return
	}

	fmt.Println("Created instance!", *runResult.Instances[0].InstanceId)
}
