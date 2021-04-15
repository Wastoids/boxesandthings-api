package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	tableName = "box_things"
	debug     = "DEBUG"
	address   = "http://localhost:4566"
	region    = "ca-central-1"
)

type dao struct {
	dynamoDB *dynamodb.Client
}

func newDao() (d dao, err error) {
	var cfg aws.Config
	if len(os.Getenv(debug)) > 0 {
		customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           address,
				SigningRegion: region,
			}, nil
		})
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
		if err != nil {
			log.Printf("error: %v", err)
			return dao{}, err
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			log.Printf("error: %v", err)
			return dao{}, err
		}
	}

	return dao{dynamoDB: dynamodb.NewFromConfig(cfg)}, nil
}

func (d dao) getTopLevelBoxesForUser(userName string) (boxes []model.Box, err error) {
	result, err := d.dynamoDB.Query(
		context.Background(),
		&dynamodb.QueryInput{
			TableName:              aws.String(tableName),
			KeyConditionExpression: aws.String("#pk = :userName AND begins_with(#sk, :topBoxes)"),
			ExpressionAttributeNames: map[string]string{
				"#pk": "pk",
				"#sk": "sk",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":userName": &types.AttributeValueMemberS{
					Value: userName,
				},
				":topBoxes": &types.AttributeValueMemberS{
					Value: "top#",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	for _, item := range result.Items {
		boxes = append(boxes, toBox(item))
	}

	fmt.Printf("%v", result)
	return boxes, nil
}

func toBox(attributeMap map[string]types.AttributeValue) model.Box {
	name := attributeMap["name"].(*types.AttributeValueMemberS)
	return model.Box{
		Name: name.Value,
	}
}
