package main

import (
	// "context"
	"context"
	// "errors"
	"fmt"
	"hello-world/global"
	"hello-world/internal/handle"
	"hello-world/internal/model"
	"os"

	// "hello-world/internal/model"
	// "log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	awsv1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
)

var ginLambda *ginadapter.GinLambda

// var (
// 	// DefaultHTTPGetAddress Default Address
// 	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

// 	// ErrNoIP No IP found in response
// 	ErrNoIP = errors.New("No IP in HTTP response")

// 	// ErrNon200Response non 200 status code in response
// 	ErrNon200Response = errors.New("Non 200 Response found")
// )

// func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return ginLambda.ProxyWithContext(context.TODO(), request)
// }



// var dbclient *dynamo.DB

func initDynamoDbTables() {
	var expectedTables = map[string]interface{}{
		model.TableClient:  model.Client{},
		model.TableAdviser: model.Adviser{},
	}

	var tableHasCreate = map[string]bool{
		model.TableClient:  false,
		model.TableAdviser: false,
	}

	tableLists, err := global.DB.ListTables().All()
	if err != nil {
		// panic("get all table err")
		fmt.Println("get all table err:", err.Error())

	}
	for _, table := range tableLists {
		if _, ok := expectedTables[table]; ok {
			tableHasCreate[table] = true
		}
	}

	for k, v := range tableHasCreate {
		if !v {
			err := global.DB.CreateTable(k, expectedTables[k]).Run()
			if err != nil {
				fmt.Println("create ", k, "err ", err.Error())
			} else {
				fmt.Println("create table k sucess", k)
			}
		}
	}
	fmt.Println("initDynamoDbTables end.......")
}

func InitDynamoDB() {
	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")
	creds := credentials.NewEnvCredentials()
	creds.Get()
	sess := session.Must(session.NewSession(&awsv1.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://docker.for.mac.localhost:8000"),
		Credentials: creds,
	}))
	global.DB = dynamo.New(sess)
	if global.DB == nil {
		panic("connect db wrong")
	}

	// initDynamoDbTables()
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	InitDynamoDB()
	if ginLambda == nil {
		r := gin.Default()

		api := r.Group("/api")
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{"data": "health"})
		})

		client := api.Group("/client")
		{
			client.POST("/register", handle.RegisterClient)
			client.POST("/change", handle.ChangeClient)
			client.POST("/delete", handle.DeleteClient)

		}

		adviser := api.Group("/adviser")
		{
			adviser.POST("/register")
		}

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// g := gin.Default()
	// err := handle.InitHandelApi(g)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// model.Init()
	// ginLambda = ginadapter.New(g)
	lambda.Start(Handler)
}
