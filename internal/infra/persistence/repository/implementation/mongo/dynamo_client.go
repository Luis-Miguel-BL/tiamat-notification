package dynamo

import (
	"context"
	"fmt"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type MongoClient struct {
	cfg  config.DBConfig
	log  logger.Logger
	conn *dynamodb.Client
}

func NewMongoClient(ctx context.Context, cfg config.DBConfig, log logger.Logger) (*MongoClient, error) {
	dynamoConfig, err := aws_config.LoadDefaultConfig(
		ctx,
		aws_config.WithRegion(cfg.MongoRegion),
		aws_config.WithRetryMaxAttempts(3),
		aws_config.WithRetryMode(aws.RetryModeStandard),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, " + err.Error())
	}

	conn := dynamodb.NewFromConfig(dynamoConfig)

	return &MongoClient{
		cfg:  cfg,
		log:  log,
		conn: conn,
	}, nil
}

func (dc *MongoClient) BatchWrite(ctx context.Context, items []map[string]types.AttributeValue, tableName string) (err error) {
	var writeItems []types.WriteRequest

	for _, item := range items {
		writeReq := types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		}
		writeItems = append(writeItems, writeReq)
	}

	batchInput := dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			tableName: writeItems,
		},
	}

	_, err = dc.conn.BatchWriteItem(ctx, &batchInput)
	if err != nil {
		if isThottleErr(err) {
			return MongoThottleErr(err)
		}
		return err
	}
	return nil
}

func (dc *MongoClient) QueryByPK(ctx context.Context, tableName string, pk string) (items []map[string]types.AttributeValue, count int32, err error) {
	result, err := dc.conn.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    aws.String("#pk = :pk"),
		ExpressionAttributeNames:  map[string]string{"#pk": "PK"},
		ExpressionAttributeValues: map[string]types.AttributeValue{":pk": &types.AttributeValueMemberS{Value: pk}},
	})

	if err != nil {
		if isThottleErr(err) {
			return items, count, MongoThottleErr(err)
		}
		return items, count, err
	}

	return result.Items, result.Count, nil
}

func (dc *MongoClient) QueryByIndex(ctx context.Context, tableName string, index string, attrName string, attrValue string) (items []map[string]types.AttributeValue, count int32, err error) {
	result, err := dc.conn.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(index),
		KeyConditionExpression: aws.String("#an = :av"),
		ExpressionAttributeNames: map[string]string{
			"#an": attrName,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":av": &types.AttributeValueMemberS{Value: attrValue},
		},
	})

	if err != nil {
		if isThottleErr(err) {
			return items, count, MongoThottleErr(err)
		}
		return items, count, err
	}

	return result.Items, result.Count, nil
}

type MongoThottleErr error

func isThottleErr(err error) bool {
	return strings.Contains(err.Error(), "rate limit token")
}
