package main

import (
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/sqs"
)

var svc *sqs.SQS

const (
    AWS_REGION      = "ap-northeast-1"
    MAIN_QUEUE_NAME = "main-queue"
    DEAD_QUEUE_NAME = "dead-queue"
)

// エラーになったメッセージを溜め込むためのデッドキューを作る。
func CreateDeadQueue(queuename string) (string, error) {
    params := &sqs.CreateQueueInput{
        QueueName: aws.String(queuename),
    }

    queueurl, err := createQueue(params)
    if err != nil {
        return "", err
    }

    return queueurl, nil
}

// キューを作成する。
func CreateMainQueue(queuename, deadarn string) (string, error) {
    // 一定回数の処理に失敗したメッセージをデッドキューに入れるためのポリシー。
    // ここでは3回失敗したらデッドキューに移すようにしている。
    // 設定内容はjson文字列を組み立てる。わりと原始的。
    redrivePolicy := fmt.Sprintf(
        "{\"deadLetterTargetArn\":\"%s\",\"maxReceiveCount\":%d}",
        deadarn,
        3,
    )

    // キュー名を指定してキュー作成。
    // 設定内容はAttributesに足していく。
    params := &sqs.CreateQueueInput{
        QueueName: aws.String(queuename),
        Attributes: map[string]*string{
            // VisibilityTimeout:取得したメッセージは指定した秒数の間、他から見えなくする。
            "VisibilityTimeout": aws.String("30"),

            // ReceiveMessageWaitTimeSeconds: ロングポーリングの秒数。
            // キューが空だった場合、どれだけ待つか。
            // ここで指定しなくても、メッセージ取得時に指定は可能。
            "ReceiveMessageWaitTimeSeconds": aws.String("20"),

            // RedrivePolicy: デッドキュー用ポリシー。先に作っておいた値を設定。
            "RedrivePolicy": aws.String(redrivePolicy),
        },
    }

    queueurl, err := createQueue(params)
    if err != nil {
        return "", err
    }

    return queueurl, nil
}

// 指定キューの情報を取得する。
// デッドキュー指定時にARNが必要だが、キュー生成時にはQueueURLしか返してくれないので、
// 別途取得している。
func GetQueueAttributes(queueurl string) (map[string]*string, error) {
    params := &sqs.GetQueueAttributesInput{
        QueueUrl: aws.String(queueurl),
        AttributeNames: []*string{
            // ほしいパラメータ名を指定。Allにすると全部取得できる。
            aws.String("All"),
        },
    }
    resp, err := svc.GetQueueAttributes(params)

    if err != nil {
        return nil, err
    }

    return resp.Attributes, nil
}

// キュー生成用の共通関数
func createQueue(params *sqs.CreateQueueInput) (string, error) {
    resp, err := svc.CreateQueue(params)

    if err != nil {
        return "", err
    }

    return *resp.QueueUrl, nil
}

func main() {
    // クライアント生成
    svc = sqs.New(&aws.Config{Region: aws.String(AWS_REGION)})

    // デッドキューを先に作成
    deadurl, err := CreateDeadQueue(DEAD_QUEUE_NAME)
    if err != nil {
        log.Fatal(err)
    }

    // デッドキューのARNを取得
    attrs, err := GetQueueAttributes(deadurl)
    if err != nil {
        log.Fatal(err)
    }

    deadarn := *attrs["QueueArn"]

    // メイン処理用のキューを作成
    mainurl, err := CreateMainQueue(MAIN_QUEUE_NAME, deadarn)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Main queue URL:%s\nDead queue URL:%s\n", mainurl, deadurl)
}
