package lamf

import (
	"time"
)

const TableName string = "files"

type File struct {
	CompanyID  int64  `json:"-"`
	UUID       string `dynamodbav:"uuid" json:"uuid"`
	Name       string `dynamodbav:"name" json:"name"`
	Key        string
	URL        string `dynamodbav:"url" json:"url"`
	Extension  string `dynamodbav:"extension" json:"extension"`
	Size       int64
	RecordedAt *time.Time `dynamodbav:"recorded_at" json:"recorded_at"`
	CreatedAt  *time.Time `dynamodbav:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `dynamodbav:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

type Response struct {
	RequestMethod  string `json:"RequestMethod"`
	RequestBody    string `json:"RequestBody"`
	PathParameter  string `json:"PathParameter"`
	QueryParameter string `json:"QueryParameter"`
}
