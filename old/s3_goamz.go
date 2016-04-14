package main

import (
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
  "log"
  "fmt"
)

func main() {
  auth, err := aws.EnvAuth()
  if err != nil {
    log.Fatal(err)
  }
  client := s3.New(auth, aws.APNortheast)
  resp, err := client.ListBuckets()

  if err != nil {
    log.Fatal(err)
  }

  log.Print(fmt.Sprintf("%T %+v", resp.Buckets[0], resp.Buckets[0]))
}
