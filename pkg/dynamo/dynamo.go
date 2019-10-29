package dynamo

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Store struct {
	dynamo *dynamodb.DynamoDB
	name   string
	table  string
}

type Item struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewStore(name string, table string, cfgs ...*aws.Config) *Store {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess, cfgs...)

	return &Store{
		dynamo: svc,
		name:   name,
		table:  table,
	}
}

func (s *Store) Name() string {
	return s.name
}

func (s *Store) Put(key string, value []byte) error {
	item := &Item{
		Key:       key,
		Value:     string(value),
		UpdatedAt: time.Now(),
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = s.dynamo.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: &s.table,
	})

	return err
}

func (s *Store) Get(key string) (value []byte, err error) {
	result, err := s.dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: &s.table,
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var item Item
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	return []byte(item.Value), err
}

func (s *Store) Delete(key string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
		},
		TableName: &s.table,
	}

	_, err := s.dynamo.DeleteItem(input)
	return err
}
