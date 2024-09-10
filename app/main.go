package main

import (
	"cdk-example/app/api"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"net/http"
)

func main() {
	api.CreateApi(".")
	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
