package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	productv1 "github.com/Bielik20/grpc-playground/gen/go"
	productv1connect "github.com/Bielik20/grpc-playground/gen/go/_goconnect"
)

func main() {
	client := productv1connect.NewProductServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	
	res, err := client.GetProduct(context.Background(), connect.NewRequest(&productv1.GetProductRequest{
		ProductId: "123",
	}))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg)
}