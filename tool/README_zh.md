# tkeel-tool

tkeel-tool 是面向 tKeel 开发者的开发工具，方便快速生成框架代码。

## Getting Started
### Required
- [go](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)


### Create a service
```
# 创建项目模板
tkeel-tool new helloworld

cd helloworld
# 拉取项目依赖
go mod download

# 生成proto模板
tkeel-tool proto add api/helloworld/helloworld.proto
# 生成proto源码
tkeel-tool proto client api/helloworld/helloworld.proto
# 生成server模板
tkeel-tool proto server api/helloworld/helloworld.proto -t internal/service

# 生成所有proto源码、wire等等
go generate ./...

# 运行程序
go run cmd/helloworld/main.go
```
