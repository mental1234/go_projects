package main

import (
	"github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"fmt"
	"os"
)

func main() {
	// Input
	region :=  "us-east-1"

	// Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	// Cloud trail client
	svc := cloudtrail.New(sess)

	resp, err := svc.DescribeTrails(&cloudtrail.DescribeTrailsInput{TrailNameList: nil})
	if err != nil {
		fmt.Println("Error getting cloudtrail")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Found",len(resp.TrailList),"trail(s) in", "us-west-1")
	
	for _, trail := range resp.TrailList {
    		fmt.Println("Trail name:  " + *trail.Name)
    		fmt.Println("Bucket name: " + *trail.S3BucketName)
		fmt.Println("")
	}
}
