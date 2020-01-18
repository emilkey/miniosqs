package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
    sqsClient *sqs.SQS
    queueUrl string
}

func main() {
    addr := flag.String("addr" , ":4000", "HTTP network address")
    sqsEndpointUrl := flag.String("sqsendpoint", "http://127.0.0.1:9324", "SQS endpoint")
    queueName := flag.String("qname", "test", "SQS queue name")
    createQueue := flag.Bool("createq", false, "Create SQS queue if it does not exist")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

    awsSess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    var sqsClient *sqs.SQS
    if *sqsEndpointUrl == "" {
        sqsClient = sqs.New(awsSess)
    } else {
        sqsClient = sqs.New(awsSess, aws.NewConfig().WithEndpoint(*sqsEndpointUrl))
    }

    queueUrl, err := getQueueUrl(queueName, sqsClient, *createQueue, infoLog)
    if err != nil {
        errorLog.Fatal("Failed to connect to queue: ", err)
    }
    infoLog.Printf("Using queue URL: %s", queueUrl)

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        sqsClient: sqsClient,
        queueUrl: queueUrl,
    }

    srv := &http.Server {
        Addr: *addr,
        ErrorLog: errorLog,
        Handler: app.routes(),
    }

    infoLog.Printf("Starting server on: %s", *addr)
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}

func getQueueUrl(queueName *string, sqsClient *sqs.SQS, create bool, infoLog *log.Logger) (string, error) {
    queueUrlResponse, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: queueName})
    if err != nil {
        if create && err.(awserr.Error).Code() == "AWS.SimpleQueueService.NonExistentQueue" {
            infoLog.Printf("Queue named %s does not exist; creating it now", *queueName)
            createQueueResponse, err2 := sqsClient.CreateQueue(&sqs.CreateQueueInput{QueueName: queueName})
            if err2 != nil {
                return "", fmt.Errorf("Failed to create queue: %s", err)
            }
            return *createQueueResponse.QueueUrl, nil
        } else {
            return "", fmt.Errorf("Failed to get queue URL: %s", err)
        }
    } else {
        return *queueUrlResponse.QueueUrl, nil    
    }
}
