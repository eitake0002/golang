package main

import (
  "fmt"
  "log"
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
)

func main(){
  //start := time.Now()
  fmt.Println("Start")

  //access_key := "AKIAJQSYZINCSJO5YKBA"
  //secret_key := "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"
  //region     := "ap-northeast-1"
  //path     := "/home/ec2-user/go_program"
  //filename := "test.jpg"

  auth, _ := aws.EnvAuth()
  client := s3.New(auth, aws.APNortheast)
  //fmt.Println(client)
  //resp := client.Bucket("data-ex2")
  resp := client.Bucket("aws-logs-462629360846-ap-northeast-1")
  fmt.Println(resp)
  fmt.Println(resp.Get("/output/"))
  log.Print("End")
}
