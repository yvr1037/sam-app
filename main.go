package main

import (
	// "context"
	"context"
	"errors"
	"hello-world/internal/handle"
	"hello-world/internal/model"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(context.TODO(), request)
}

func main() {
	g := gin.Default()
	err := handle.InitHandelApi(g)
	if err != nil {
		log.Println(err)
		return
	}
	model.Init()
	ginLambda = ginadapter.New(g)
	lambda.Start(handler)
}
