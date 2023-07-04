package model

import (
	// "fmt"

	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	awsv1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var dbclient *dynamo.DB

func Init() {
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
	dbclient = dynamo.New(sess)
	if dbclient == nil {
		panic("connect db wrong")
	}
	// err := TableClientCreate()
	// if err != nil {
	// 	panic(err)
	// }
	tableCheck()
}

func tableCheck() {
	list, err := dbclient.ListTables().All()
	if err != nil {
		panic(err)
	}
	hash := make(map[string]bool)
	for _, v := range list {
		hash[v] = true
	}

	if _, ok := hash[TableClient]; !ok {
		err = TableClientCreate()
		if err != nil {
			panic(err)
		}
	}

	if _, ok := hash[TableAdviser]; !ok {
		err = TableAdviserCreate()
		if err != nil {
			panic(err)
		}
	}
}