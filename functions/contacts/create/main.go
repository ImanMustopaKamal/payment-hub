package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ImanMustopaKamal/payment-hub/config"
	awsconfig "github.com/ImanMustopaKamal/payment-hub/config/aws"
	"github.com/ImanMustopaKamal/payment-hub/internal/entities"
	. "github.com/ImanMustopaKamal/payment-hub/internal/types"
	. "github.com/ImanMustopaKamal/payment-hub/internal/utils"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func Handler(ctx context.Context, request Request) (Response, error) {
	var dto entities.ContactCreateDto
	if err := json.Unmarshal([]byte(request.Body), &dto); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return NewJSONResponse(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		fmt.Println("Validation errors:", err)
		return NewJSONResponse(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	if err := Action(&dto); err != nil {
		return NewJSONResponse(http.StatusInternalServerError, map[string]string{"message": "Failed to create contact"}) // Or a more specific error message based on the error details
	}

	return NewJSONResponse(http.StatusCreated, map[string]string{"message": "Contact created succesfully!"})
}

func Action(item *entities.ContactCreateDto) (err error) {
	svc := awsconfig.GetDynamoDBClient()

	selectedKeys := map[string]string{
		"account_id": fmt.Sprintf("account_id", item.AccountID),
	}
	key, err := attributevalue.MarshalMap(selectedKeys)
	result, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("contacts"),
		Key:       key,
	},
	)
	if result.Item == nil {
		fmt.Println("GetItem: Data not found, %v", nil)
	}

	// store
	// Generate a new UUID
	newUUID := uuid.New().String()

	contact := &entities.Contact{
		ID:         newUUID,
		Name:       item.Name,
		AccountID:  item.AccountID,
		CardNumber: item.CardNumber,
		BankName:   item.BankName,
	}

	data, err := attributevalue.MarshalMap(contact)
	if err != nil {
		log.Println("error MarshalMap, %v", err)
	}
	fmt.Println("MarshalMap, %v", data)

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("contacts"),
		Item:      data,
	})

	if err != nil {
		fmt.Println("error dynamodb, %v", err)
	}
	return nil
}

func main() {
	config.Init()
	lambda.Start(Handler)
}
