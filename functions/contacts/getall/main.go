package main

import (
	"context"
	"net/http"

	. "github.com/ImanMustopaKamal/payment-hub/internal/types"
	. "github.com/ImanMustopaKamal/payment-hub/internal/utils"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request Request) (Response, error) {
	return NewJSONResponse(http.StatusCreated, map[string]string{"message": "Contact found!"})
}

func main() {
	lambda.Start(Handler)
}
