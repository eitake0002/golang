package main

import (
  "log"
  "os"
  "fmt"
  "reflect"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/awsutil"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/aws/session"
)

func main() {
    fileTransfer := FileTransferToS3{
        AccessKeyId:     "AKIAJQSYZINCSJO5YKBA",
        SecretAccessKey: "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",
        Region:          "ap-northeast-1",
        BucketName:      "data-ex2",
    }

    fileTransfer.PutToS3("./", "s3_aws_sdk.go")
}

type FileTransferToS3 struct {
    AccessKeyId     string
    SecretAccessKey string
    Region          string
    BucketName      string
}

func (f *FileTransferToS3) PutToS3(path string, filename string) {
    file, err := os.Open(fmt.Sprintf("%s%s", path, filename))
    if err != nil {
        log.Println(err.Error())
    }
    defer file.Close()

    log.Println(reflect.TypeOf(f.Region))
    creds  := credentials.NewStaticCredentials(f.AccessKeyId, f.SecretAccessKey, "")
    config := aws.Config{Credentials: creds, Region: &f.Region}
    cli := s3.New(session.New(), &config)
    resp, err := cli.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(f.BucketName),
        Key:    aws.String(filename),
        Body:   file,
    })
    if err != nil {
        log.Println(err.Error())
    }

    log.Println(awsutil.StringValue(resp))
}

