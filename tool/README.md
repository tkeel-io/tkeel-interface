# tkeel-tool

tkeel-tool is a development tool for tKeel developers, which facilitates the rapid generation of framework code.

## Getting Started
### Required
- [go](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)


### Quick Start

```
# Create project template
tkeel-tool new github.com/tkeel-io/helloworld

cd helloworld

# Download necessary plug-ins
make init

# Generate proto template
tkeel-tool proto add api/helloworld/helloworld.proto

# Generate proto source code
make api

# Generate service template
tkeel-tool proto service api/helloworld/helloworld.proto -t internal/service

# Generate server template (this output needs to be manually added to cmd/helloworld/main.go)
tkeel-tool proto server api/helloworld/helloworld.proto

# Run the program
go run cmd/helloworld/main.go
```



