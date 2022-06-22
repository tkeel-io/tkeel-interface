# artisan (tkeel-tool)

The artisan 是面向 tKeel 开发者的开发工具，方便快速生成框架代码。

## Getting Started
### Required
- [go](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

### Quick Start
```bash
# 安装
$ go get -u github.com/tkeel-io/tkeel-interface/tool/cmd/artisan

# 创建项目模板
$ artisan new github.com/tkeel-io/helloworld

$ cd helloworld

# 下载必须的插件
$ make init

# 生成 proto 模板
$ artisan proto add api/helloworld/v1/helloworld.proto

# 生成 error proto 模板
$ artisan proto add api/helloworld/v1/error.proto

# 下载必须的插件
$ make api

# 生成 service 模板
$ artisan proto service api/helloworld/v1/helloworld.proto -t pkg/service

# 生成 server 模板(此输出需要手工加入 cmd/helloworld/main.go 中)
$ artisan proto server api/helloworld/v1/helloworld.proto

# 生成 api 的 makedown 文件
$ artisan markdown -f api/apidocs.swagger.json  -t third_party/markdown-templates/ -o ./docs/API/Greeter -m all

# 运行程序
$ go run cmd/helloworld/main.go
```




### Markdown

```bash
# 基于 apidocs.swagger.json 生成 API 列表
$ artisan markdown -f apidocs.swagger.json -m tag -o .

# 基于 apidocs.swagger.json 生成 API 函数详情
$ artisan markdown -f apidocs.swagger.json -m method -o .

# 基于 apidocs.swagger.json 生成 API 列表 以及 API 函数详情
$ artisan markdown -f apidocs.swagger.json -m all -o .

# 基于 apidocs.swagger.json 生成，且排除 tag 为 'Private' 以及 'Internal' 的函数
$ artisan markdown -f apidocs.swagger.json -m all -o . --exclude_tag 'Private' --exclude_tag 'Internal'

# 基于 apidocs.swagger.json 生成，且指定模板目录
$ artisan markdown -f apidocs.swagger.json -m all -o . -t pkg/markdown/templates
```

