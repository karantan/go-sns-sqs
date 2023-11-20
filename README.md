# go-sns-sqs

## Introduction

`go-sns-sqs` is a sandbox project demonstrating the use of AWS services to create a
serverless architecture in Go. This project is centered around a Lambda function named
`calculate`, which is triggered by an SNS topic called `dispatch`. Depending on the
outcome of the calculation, the result is either sent to a success SQS queue or a
failure SQS queue. This setup serves as a foundation for further development and
integration with end applications.

## Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.20 or higher)
- [Serverless Framework](https://www.serverless.com/)

Additionally, you should have:
- An AWS account with appropriate permissions to create and manage Lambda functions,
SNS topics, and SQS queues.
- Basic understanding of AWS services and the Go programming language.

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/go-sns-sqs.git
   cd go-sns-sqs
   ```

2. **Install Dependencies:**

   Install the required Go dependencies (if any).

   ```bash
   go mod tidy
   ```

3. **Deploy with Serverless Framework:**

   Use the Serverless Framework to deploy the application to AWS.

   ```bash
   make deploy
   ```

## Usage

After deployment, the `calculate` Lambda function will be triggered by messages published
to the `dispatch` SNS topic. The function will process the message and send the result
to either the success or failure SQS queue.

- **To Trigger the Lambda Function:**

  Publish a message to the `dispatch` SNS topic. You can do this via e.g. the AWS
  Management Console, the AWS CLI etc.

- **Monitoring Results:**

  Check the success and failure SQS queues to see the results of the processed messages.
  You can configure additional applications or services to process these messages further.

## Further Configuration

This project serves as a basic example. You can extend or modify the `serverless.yml`
and the Go code in `calculate` function to suit your specific requirements.


## Additional resources:

- [Go AWS examples](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/).
- [Configuring a queue to use with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-queueconfig)
- [Lambda Destinations](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html?icmpid=docs_lambda_help#invocation-async-destinations)
- [Dead-Letter Queue](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html?icmpid=docs_lambda_help#invocation-dlq)
- [SNS FAQs](https://aws.amazon.com/sns/faqs/)
