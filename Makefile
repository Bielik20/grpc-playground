# Define versions here for easy updates
VER_GO_PROTO    := v1.36.11
VER_CONNECT     := v1.19.1

.PHONY: install-tools generate

install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(VER_GO_PROTO)
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@$(VER_CONNECT)

generate: install-tools
  # `shell` is required here in Makefile, I don't know why :/
	PATH="$(shell go env GOPATH)/bin:$(shell go env GOBIN):${PATH}" bunx buf generate
