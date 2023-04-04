package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type handler struct {
	Greeting string
}

func newHandler(greeting string) *handler {
	return &handler{
		Greeting: greeting,
	}
}

func (h *handler) handleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(h.Greeting)

	return events.APIGatewayProxyResponse{
		Body:       h.Greeting,
		StatusCode: http.StatusOK,
	}, nil
}
