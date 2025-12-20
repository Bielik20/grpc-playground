## Tools

### Buf

Brew: `brew install bufbuild/buf/buf`

Alternatively (npm): `npm install @bufbuild/buf`

To generate stuff from protobuf: `buf generate`

## Running locally

To start go server: `go run ./server/go/`

<!--To run typescript client (against OpenAPI): `bun run client/typescript/call-simple-get.ts`-->

To run go client (against gRPC): `go run ./client/go/`

To call using JSON:
```
curl \
    --header "Content-Type: application/json" \
    --data '{"productId": "123"}' \
    http://localhost:8080/product.ProductService/GetProduct
```
