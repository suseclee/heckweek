package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

// Usage:
// go run main.go <instance id>
func main() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

    // Create EC2 service client
    svc := ec2.New(sess)

    input := &ec2.TerminateInstancesInput{
        InstanceIds: []*string{
            aws.String(os.Args[1]),
        },
        DryRun: aws.Bool(true),
    }
    result, err := svc.TerminateInstances(input)
    awsErr, ok := err.(awserr.Error)

    // If the error code is `DryRunOperation` it means we have the necessary
    // permissions to Start this instance
    if ok && awsErr.Code() == "DryRunOperation" {
        // Let's now set dry run to be false. This will allow us to reboot the instances
        input.DryRun = aws.Bool(false)
        result, err = svc.TerminateInstances(input)
        if err != nil {
            fmt.Println("Error", err)
        } else {
            fmt.Println("Success", result)
        }
    } else { // This could be due to a lack of permissions
        fmt.Println("Error", err)
    }

}