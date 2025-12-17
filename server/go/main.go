package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// Import your generated code
	pb "github.com/youruser/yourrepo/gen/go/product/v1"
)

// 1. Implement the gRPC Server
type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	// Logic to fetch product...
	return &pb.Product{
		Id:    req.ProductId,
		Name:  "Super Widget",
		Price: 19.99,
	}, nil
}

// 2. Middleware to add Cache-Control headers for Browser
func cachingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If it's a GET request, tell the browser to cache it for 1 hour
		if r.Method == "GET" {
			w.Header().Set("Cache-Control", "public, max-age=3600")
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	// --- Start gRPC Server (Internal) ---
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	go func() {
		log.Println("Serving gRPC on 0.0.0.0:9090")
		log.Fatalln(s.Serve(lis))
	}()

	// --- Start HTTP Gateway (External) ---
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9090",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterProductServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Wrap gateway with caching middleware
	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: cachingMiddleware(gwmux),
	}

	log.Println("Serving HTTP on 0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}