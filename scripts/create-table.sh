#!/bin/sh

aws dynamodb create-table --endpoint-url http://db:4566 --cli-input-json file:///scripts/dynamodb.json --region ca-central-1

aws dynamodb batch-write-item --endpoint-url http://db:4566  --request-items file:///scripts/data.json