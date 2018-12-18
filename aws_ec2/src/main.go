// aws create an instance
package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
  "fmt"
)

func main() {
  createEc2Instance("ami-07c20c8dd8c3136bf", "t2.micro", 1, 1)
}

func createEc2Instance(instanceId string, instanceType string, minCount int, maxCount int) (result string, err error) {
  // Creates a new session
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-1")
  },
  if err != nil {
    return err
  }
  // Returns connection error, else connects to EC2 service
  svc := ec2.New(sess)
  // Specify instance details
  runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
    ImageId:       aws.String(imageId),
    InstanceType:  aws.String(instanceType),
    MinCount:      aws.Int64(minCount),
    MaxCount:      aws.Int64(maxCount), 
  })
  if err != nil {
    return err
  }
  // Returns runInstance error, else returns created instance ID
  return *runResult.Instances[0].InstanceId
}
