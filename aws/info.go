package main

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
)

func main() {
    // Load session from shared config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create new EC2 client
    ec2Svc := ec2.New(sess)

    // Call to get detailed information on each instance
    resp, err := ec2Svc.DescribeInstances(nil)
    if err != nil {
        fmt.Println("Error", err)
    } else {
            //fmt.Println(resp)

            for idx, res := range resp.Reservations {
                    for _, inst := range resp.Reservations[idx].Instances {
                            if *inst.State.Name == "running"{
                                fmt.Println("  > Number of instances: ", len(res.Instances))
                                fmt.Println("    - Instance ID: ", *inst.InstanceId)
                                fmt.Println("    - Public IP  : ", *inst.PublicIpAddress)
                                fmt.Println("    - PublicDnsName: ", *inst.PublicDnsName)
                                fmt.Println("    - Username: ec2-user" )

                            }
                    }
            }

    }
}