package misc

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Store struct {
	dynamoClient *dynamodb.DynamoDB
}

type ServiceProxyServicesEntry struct {
	ServiceName string `json:"servicename"`
}

func NewStore() *Store {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
		//		LogLevel: aws.LogLevel(aws.LogDebug),
	})
	if err != nil {
		panic(err)
	}

	return &Store{
		dynamoClient: dynamodb.New(sess),
	}
}

func (s *Store) PutService(entry ServiceProxyServicesEntry) {
	dynamoEntry, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		panic(err)
	}

	input := &dynamodb.PutItemInput{
		Item: dynamoEntry,
		TableName: aws.String("serviceproxy-services"),
	}

	_, err = s.dynamoClient.PutItem(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Stored service ", entry.ServiceName, " to table ", "serviceproxy-services")
}

func (s *Store) RemoveService(entry ServiceProxyServicesEntry) {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"servicename": {
				S: aws.String(entry.ServiceName),
			},
		},
		TableName: aws.String("serviceproxy-services"),
	}

	_, err := s.dynamoClient.DeleteItem(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Removed service ", entry.ServiceName, " from table ", "serviceproxy-services")
}
