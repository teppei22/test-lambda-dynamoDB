package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	lamf "github.com/teppei22/test-lambda-dynamoDB/lambdaFunc"
)

func main() {
	lambda.Start(lamf.UploadTest)
}
