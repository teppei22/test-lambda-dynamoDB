package lamf

import (
	"encoding/json"

	"github.com/k0kubun/pp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func UploadPostHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	method := request.HTTPMethod

	// DBと接続するセッションを作る→DB接続
	sess, err := session.NewSession()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	db := dynamodb.New(sess)

	// リクエストボディのjsonから、File構造体(DB用データの構造体)を作成
	reqBody := request.Body
	resBodyJSONBytes := ([]byte)(reqBody)
	item := File{
		UUID: UUID(),
	}
	if err := json.Unmarshal(resBodyJSONBytes, &item); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}
	pp.Print("item", item)

	// Item構造体から、inputするデータを用意
	inputAV, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}
	pp.Print("inputAV", inputAV)
	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      inputAV,
	}

	// insert実行
	_, err = db.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	// httpレスポンス作成
	res := Response{
		RequestMethod: method,
	}
	jsonBytes, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}
