# MinioSQS

MinioSQS is a small utility that uses [MinIO](https://github.com/minio/minio) and [ElasticMQ](https://github.com/softwaremill/elasticmq) to enable local testing of systems that use [AWS S3 event notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).

## Basic Instructions

1. Download and run ElasticMQ
2. Download and run MinioSQS
3. Download and run MinIO
4. Create a bucket and congfigure it to send event notifications via webhook to the MinioSQS endpoint
5. Configure your application to use the MinIO endpoint for S3 operations, and the ElasticMQ endpoint for SQS operations
