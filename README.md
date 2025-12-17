## Tools

### Buf

Brew: `brew install bufbuild/buf/buf`

Alternatively (npm): `npm install @bufbuild/buf`

I needed to run: `buf dep update` to install deps

To generate stuff from protobuf: `buf generate`

### OpenAPI

To generate TypeScript web client from openapi: `bunx openapi-generator-cli generate -i ./gen/openapiv2/product.openapi.yaml -g typescript-fetch -o ./gen/ts-client --skip-validate-spec`

`--skip-validate-spec` is required because of missing version field - may need to look into it.

## Running locally

To start go server: `go run ./server/go/`

To run typescript client (against OpenAPI): `bun run client/typescript/call-simple-get.ts`

To run go client (against gRPC): `go run ./client/go/`
