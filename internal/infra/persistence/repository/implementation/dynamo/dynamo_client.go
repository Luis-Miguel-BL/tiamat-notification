package dynamo

import (
	"context"
	"fmt"
	"strings"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoClient struct {
	cfg  config.DBConfig
	log  logger.Logger
	conn *dynamodb.Client
}

func NewDynamoClient(ctx context.Context, cfg config.DBConfig, log logger.Logger) (*DynamoClient, error) {
	dynamoConfig, err := aws_config.LoadDefaultConfig(
		ctx,
		aws_config.WithRegion(cfg.DynamoRegion),
		aws_config.WithRetryMaxAttempts(3),
		aws_config.WithRetryMode(aws.RetryModeStandard),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, " + err.Error())
	}

	conn := dynamodb.NewFromConfig(dynamoConfig)

	return &DynamoClient{
		cfg:  cfg,
		log:  log,
		conn: conn,
	}, nil
}

func (dc *DynamoClient) BatchWrite(ctx context.Context, items []map[string]types.AttributeValue, tableName string) (err error) {
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
			return DynamoThottleErr(err)
		}
		return err
	}
	return nil
}

type DynamoThottleErr error

func isThottleErr(err error) bool {
	return strings.Contains(err.Error(), "rate limit token")
}
