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
)

func main() {
    fileTransfer := FileTransferToS3{
        AccessKeyId:     "AKIAJQSYZINCSJO5YKBA",
        SecretAccessKey: "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",
        Region:          "ap-northeast-1",
        BucketName:      "data-ex2",
    }

    fileTransfer.PutToS3("./", "s3.go")
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
    region := "ap-northeast-1"
    creds  := credentials.NewStaticCredentials(f.AccessKeyId, f.SecretAccessKey, "")
    //config := aws.Config{Credentials: creds, Region: aws.String("ap-northeast-1")}
    config := aws.Config{Credentials: creds, Region: region}
    cli := s3.New(&config)

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

