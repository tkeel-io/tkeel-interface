# tkeel-tool

tkeel-tool is a development tool for tKeel developers, which facilitates the rapid generation of framework code.

## Getting Started
### Required
-[go](https://golang.org/dl/)
-[protoc](https://github.com/protocolbuffers/protobuf)
-[protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)


### Create a service
```
# Create project template
tkeel-tool new helloworld

cd helloworld

# Download necessary plug-ins
make init

# Pull project dependencies
go mod download

# Generate proto template
tkeel-tool proto add api/helloworld/helloworld.proto
# Generate proto source code
tkeel-tool proto client api/helloworld/helloworld.proto
# Generate service template
tkeel-tool proto server api/helloworld/helloworld.proto -t internal/service
# Generate server template (this output needs to be manually added to cmd/helloworld/main.go)
tkeel-tool proto server api/helloworld/helloworld.proto

# Generate all proto source code
go generate ./...

# Run the program
go run cmd/helloworld/main.go
```