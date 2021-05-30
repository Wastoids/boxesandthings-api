package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/Wastoids/boxesandthings-api/service"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	tableName    = "box_things"
	debug        = "DEBUG"
	address      = "http://localhost:4566"
	region       = "ca-central-1"
	boxEntity    = "box#"
	details      = "details"
	thingEntity  = "thing#"
	topBoxEntity = "top#"
	box          = "box"
	thing        = "thing"
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
					Value: topBoxEntity,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		boxes = append(boxes, toBox(item))
	}

	return boxes, nil
}

func (d dao) saveBox(b model.Box) error {
	_, err := d.dynamoDB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: fmt.Sprintf("%v%v", boxEntity, b.ID),
			},
			"sk": &types.AttributeValueMemberS{
				Value: details,
			},
			"id": &types.AttributeValueMemberS{
				Value: b.ID,
			},
			"name": &types.AttributeValueMemberS{
				Value: b.Name,
			},
			"type": &types.AttributeValueMemberS{
				Value: box,
			},
		},
		TableName: aws.String(tableName),
	})

	return err
}

func (d dao) saveTopBox(userName string, b model.Box) error {
	_, err := d.dynamoDB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: userName,
			},
			"sk": &types.AttributeValueMemberS{
				Value: fmt.Sprintf("%v%v", topBoxEntity, b.ID),
			},
			"id": &types.AttributeValueMemberS{
				Value: b.ID,
			},
			"name": &types.AttributeValueMemberS{
				Value: b.Name,
			},
			"type": &types.AttributeValueMemberS{
				Value: box,
			},
		},
		TableName: aws.String(tableName),
	})
	return err
}

func (d dao) saveThing(t model.Thing, boxID string) error {
	result, err := d.dynamoDB.PutItem(context.Background(), &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"pk":          &types.AttributeValueMemberS{Value: fmt.Sprintf("%v%v", boxEntity, boxID)},
			"sk":          &types.AttributeValueMemberS{Value: fmt.Sprintf("%v%v", thingEntity, t.ID)},
			"id":          &types.AttributeValueMemberS{Value: t.ID},
			"name":        &types.AttributeValueMemberS{Value: t.Name},
			"description": &types.AttributeValueMemberS{Value: t.Description},
			"type":        &types.AttributeValueMemberS{Value: thing},
		},
		TableName: aws.String(tableName),
	})
	if err != nil {
		return err
	}
	fmt.Printf("got this: %v", result)
	return nil
}

func (d dao) getBoxContent(boxID string) (service.BoxContentResult, error) {
	output, err := d.dynamoDB.Query(context.Background(),
		&dynamodb.QueryInput{
			TableName:              aws.String(tableName),
			KeyConditionExpression: aws.String("#pk = :boxID"),
			ExpressionAttributeNames: map[string]string{
				"#pk": "pk",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":boxID": &types.AttributeValueMemberS{Value: fmt.Sprintf("%v%v", boxEntity, boxID)},
			},
		})
	if err != nil {
		return service.BoxContentResult{}, err
	}
	return toBoxContentResult(output.Items), nil

}

func toBox(attributeMap map[string]types.AttributeValue) model.Box {
	return model.Box{
		ID:   attributeMap["id"].(*types.AttributeValueMemberS).Value,
		Name: attributeMap["name"].(*types.AttributeValueMemberS).Value,
	}
}

func toThing(attributeMap map[string]types.AttributeValue) model.Thing {
	return model.Thing{
		ID:          attributeMap["id"].(*types.AttributeValueMemberS).Value,
		Name:        attributeMap["name"].(*types.AttributeValueMemberS).Value,
		Description: attributeMap["description"].(*types.AttributeValueMemberS).Value}
}

func toBoxContentResult(attributeMap []map[string]types.AttributeValue) service.BoxContentResult {
	var boxes []model.Box
	var things []model.Thing
	var result service.BoxContentResult

	for _, entity := range attributeMap {
		if entity["type"].(*types.AttributeValueMemberS).Value == box {
			boxes = append(boxes, toBox(entity))
		} else {
			things = append(things, toThing(entity))
		}
	}

	result.Boxes = boxes
	result.Things = things
	return result
}
