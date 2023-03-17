package main

import (
	"aws.tappc.tap.vmare.com/pkg"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(cfn.LambdaWrap(pkg.CAGeneratorFunc))
}
