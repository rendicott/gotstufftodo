package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func readDynamoConfig() (tableName string, err error) {
	sess := session.New()
	dsvc := dynamodb.New(sess)
	configTable := os.Getenv("CONFIG_TABLE")
	var attributesToGet []*string
	setting := "setting"
	value := "value"
	attributesToGet = append(attributesToGet, &setting)
	attributesToGet = append(attributesToGet, &value)
	scanInput := dynamodb.ScanInput{AttributesToGet: attributesToGet,
		TableName: &configTable}
	results, err := dsvc.Scan(&scanInput)
	if err != nil {
		return tableName, err
	}
	for item := range results.Items {
		setting := results.Items[item]["setting"].S
		log.Println(*setting)
		if *setting == "table-name" {
			tn := results.Items[item]["value"].S
			log.Println(tn)
			tableName = *tn
		}
	}
	return tableName, err
}

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	tableName, err := readDynamoConfig()
	if err != nil {
		msg := fmt.Sprintf("Got error: '%s'", err.Error())
		return Response{
			Message: msg,
		}, err
	}
	msg := fmt.Sprintf("Got config '%s'", tableName)
	return Response{
		Message: msg,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
