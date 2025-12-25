package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// Import the generated code
	productv1 "github.com/Bielik20/grpc-playground/gen/go"
	productv1connect "github.com/Bielik20/grpc-playground/gen/go/_goconnect"
)

// ProductServer implements the ProductServiceHandler interface.
type ProductServer struct{}

// GetProduct implements the logic for the GetProduct RPC.
func (s *ProductServer) GetProduct(
	ctx context.Context,
	req *connect.Request[productv1.GetProductRequest],
) (*connect.Response[productv1.Product], error) {

	log.Printf("Received request for product ID: %s ContentType: %s HttpMethod: %s", req.Msg.ProductId, req.Header().Get("Content-Type"), req.HTTPMethod())

	// Mock logic: In a real app, you would query a database here.
	if req.Msg.ProductId == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("product_id is required"))
	}

	product := &productv1.Product{
		Id:    req.Msg.ProductId,
		Name:  "Super Cool Gadget",
		Price: 99.99,
	}

	res := connect.NewResponse(product)
	res.Header().Set("Product-Version", "v1")

	return res, nil
}

func main() {
	// 1. Instantiate the implementation
	producter := &ProductServer{}

	// 2. Create the path and handler using the generated constructor.
	// NewProductServiceHandler returns the path prefix (e.g., "/product.v1.ProductService/")
	// and the http.Handler that handles the logic.
	path, handler := productv1connect.NewProductServiceHandler(producter)

	// 3. Mount the handler on a standard Go http.ServeMux
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	// 4. Start the server.
	// To support gRPC (which requires HTTP/2) over unencrypted HTTP (h2c) locally,
	// we wrap the mux in h2c.NewHandler.
	// If you are behind a load balancer doing TLS termination or using pure JSON,
	// standard http.ListenAndServe is often sufficient, but h2c is best for local gRPC testing.
	address := "localhost:8080"
	fmt.Printf("Server listening on http://%s\n", address)

	// Use h2c to support HTTP/2 without TLS (required for standard gRPC clients locally)
	err := http.ListenAndServe(
		address,
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
