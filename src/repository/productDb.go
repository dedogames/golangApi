package repository

import (
	"fmt"

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

func (mp *ProductDb) Delete(product *entities.ProductBody) error {

	conditionExpr := aws.String("attribute_exists(#id)") // Check if item with that ID exists

	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(mp.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {N: aws.String(fmt.Sprintf("%d", product.Id))}, // Assuming "id" is your primary key
		},
		ConditionExpression: conditionExpr,
		ExpressionAttributeNames: map[string]*string{
			"#id": aws.String("id"), // Assuming "id" is your attribute name
		},
	}

	_, err := mp.DynamodbClient.DeleteItem(deleteItemInput)
	if err != nil {
		// Handle error, potentially checking if it's due to the condition not being met

		// Other error handling
		fmt.Println("Error deleting item:", err)

	}
	return nil
}
func (mp *ProductDb) Find(product *entities.ProductBody) ([]*entities.ProductBody, error) {

	input := &dynamodb.QueryInput{
		TableName:              aws.String(mp.TableName),
		IndexName:              aws.String("id-index"), // Replace with your GSI name
		KeyConditionExpression: aws.String("#pk = :pk"),
		ExpressionAttributeNames: map[string]*string{
			"#pk": aws.String("id"), // Assuming "Id" is your GSI partition key
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {N: aws.String(fmt.Sprintf("%d", product.Id))},
		},
	}
	result, err := mp.DynamodbClient.Query(input)
	if err != nil {
		return nil, err
	}

	var products []*entities.ProductBody
	for _, item := range result.Items {
		var product entities.ProductBody
		err := dynamodbattribute.UnmarshalMap(item, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
func (mp *ProductDb) SelectAll() ([]*entities.ProductBody, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(mp.TableName),
	}

	result, err := mp.DynamodbClient.Scan(input)
	if err != nil {
		return nil, err
	}

	var products []*entities.ProductBody
	for _, item := range result.Items {
		var product entities.ProductBody
		err := dynamodbattribute.UnmarshalMap(item, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
