start-ddb:
		docker run -dp 4566:4566 localstack/localstack:0.12.9.1
		sleep 30

create-db: start-ddb
		aws --endpoint-url http://localhost:4566 dynamodb create-table --cli-input-json file://scripts/dynamodb.json
