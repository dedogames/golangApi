package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/crud/configuration"
	"github.com/crud/entities"
	"github.com/crud/lib"
)

type ProductDb struct {
	TableName      string
	RegionName     string
	DynamodbClient dynamodbiface.DynamoDBAPI
}

func NewProductDb() *ProductDb {

	TableName := configuration.Cfg.DynamoDBConfig.ProductTable
	RegionName := configuration.Cfg.DynamoDBConfig.Region

	config := &aws.Config{
		Region:     aws.String(RegionName),
		MaxRetries: aws.Int(configuration.Cfg.DynamoDBConfig.MaxAttempts),
	}
	//TODO: refector this with dependency injection help with tests and others aws services
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	lib.Logger.Info("ProductDb initialized with TableName: " + TableName)
	return &ProductDb{
		TableName:      TableName,
		RegionName:     RegionName,
		DynamodbClient: dynamodb.New(session, config),
	}
}

func (mp *ProductDb) Count() (int, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(mp.TableName),
	}
	result, err := mp.DynamodbClient.Scan(input)
	if err != nil {
		return 0, err
	}
	return len(result.Items), nil
}
func (mp *ProductDb) Save(product *entities.ProductBody) error {
	lastId, err := mp.Count()
	if err != nil {
		return err
	}
	product.Id = lastId + 1
	item, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(mp.TableName),
		Item:      item,
	}

	_, err = mp.DynamodbClient.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}
