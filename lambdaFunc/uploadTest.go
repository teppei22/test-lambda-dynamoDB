package lambdaFunc

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func uploadTestHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// httpリクエストの情報を取得
	method := request.HTTPMethod
	body := request.Body
	pathParam := request.PathParameters["pathparam"]
	queryParam := request.QueryStringParameters["queryparam"]

	// レスポンスとして返すjson文字列を作る
	res := Response{
		RequestMethod:  method,
		RequestBody:    body,
		PathParameter:  pathParam,
		QueryParameter: queryParam,
	}
	jsonBytes, _ := json.Marshal(res)

	// 返り値としてレスポンスを返す
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}
