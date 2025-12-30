# gRPC Playground

## Tools

### Buf

Install Buff: `bun install -D @bufbuild/buf`

Install TypeScript plugin: `bun install -D @bufbuild/protoc-gen-es`

Install Go plugins:
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.1`
- `go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.19.1`

To generate stuff from protobuf: `PATH="$(go env GOPATH)/bin:$(go env GOBIN):${PATH}" bunx buf generate`

Or as a shorthand: `make generate`

## Running locally

To start ts server: `bun run ./server/typescript/main.ts`

To start go server: `go run ./server/go/`

To run typescript client (as JSON): `bun run client/typescript/call-json.ts`

To run typescript client (as gRPC): `bun run client/typescript/call-grpc.ts`

To run go client (against gRPC): `go run ./client/go/`

To call using JSON:

```
curl \
    --header "Content-Type: application/json" \
    --data '{"productId": "123"}' \
    http://localhost:8080/product.ProductService/GetProduct
```

## Links

### Buf

- https://buf.build/docs/cli
- https://buf.build/plugins/other

### Connect

- https://connectrpc.com/
- TypeScript: https://www.npmjs.com/package/@connectrpc/connect
  - https://connectrpc.com/docs/web/getting-started/
  - https://connectrpc.com/docs/node/getting-started/
- Go: https://connectrpc.com/docs/go/getting-started
