# artisan (tkeel-tool)

The artisan is a development tool for tKeel developers, which facilitates the rapid generation of framework code.

## Getting Started
### Required
- [go](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)


### Quick Start

```bash
# Install
$ go get -u github.com/tkeel-io/tkeel-interface/tool/cmd/artisan

# Create project template
$ artisan new github.com/tkeel-io/helloworld

$ cd helloworld

# Download necessary plug-ins
$ make init

# Generate proto template
$ artisan proto add api/helloworld/v1/helloworld.proto

# Generate error proto template
$ artisan proto add api/helloworld/v1/error.proto

# Generate proto source code
$ make api

# Generate service template
$ artisan proto service api/helloworld/v1/helloworld.proto -t pkg/service

# Generate server template (this output needs to be manually added to cmd/helloworld/main.go)
$ artisan proto server api/helloworld/v1/helloworld.proto

# Generate API's makedown
$ artisan markdown -f api/apidocs.swagger.json  -t third_party/markdown-templates/ -o ./docs/API/Greeter -m all

# Run the program
$ go run cmd/helloworld/main.go
```



### Markdown

```bash
# Generate API list based on apidocs.swagger.json
$ artisan markdown -f apidocs.swagger.json -m tag -o .

# Generate API function details based on apidocs.swagger.json
$ artisan markdown -f apidocs.swagger.json -m method -o .

# Generate API list and API function details based on apidocs.swagger.json
$ artisan markdown -f apidocs.swagger.json -m all -o .

# Generated based on apidocs.swagger.json and excludes functions with tags 'Private' and 'Internal'
$ artisan markdown -f apidocs.swagger.json -m all -o . --exclude_tag 'Private' --exclude_tag 'Internal'

# Generate based on apidocs.swagger.json and specify the template directory
$ artisan markdown -f apidocs.swagger.json -m all -o . -t pkg/markdown/templates
```


