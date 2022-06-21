---
title: '{{.Summary}}'
description: "{{.OperationID}}"
---
{{- $definitions := .Definitions}}
## 接口说明
调用该接口{{.Summary}}。

{{- .Description}}

## URI

```
{{.Operation}} {{.Path }}
```

## 请求参数

| 名称 | 参数位置 | 类型 | 描述 |  是否必须 |
| ---- | ---------- | ----------- | ----------- | ----------- | 
{{- with $paths := FilterParameters .Parameters "path"}}    {{range $param := $paths}}
| {{$param.Name}} | path | {{$param.Type}} | {{$param.Description}} |  Required | {{end}}{{end}}
{{- with $queries := FilterParameters .Parameters "query"}} {{range $param := $queries}}
| {{$param.Name}} | query | {{$param.Type}} | {{$param.Description}} |  {{$param.Required}} |{{end}}{{end}}

{{- with $bodies := FilterParameters .Parameters "body"}}

### 请求Body

{{- range $resp := $bodies }}
{{- if eq $resp.Type  "array" }}   
| 描述 | 类型 |
| ----------- | ------ |
| {{$resp.Description}} | Array[[{{FilterSchema $resp.Schema.Items.Ref}}](#{{FilterSchema $resp.Items.Ref}})] |

#### {{FilterSchema $resp.Items.Ref}}

{{- template "schema.md" CollectSchema $definitions  $resp.Items.Ref}}
{{ else if $resp.Schema.Ref }}
| 描述 | 类型 |
| ----------- | ------ |
| {{$resp.Description}} | Object([{{FilterSchema $resp.Schema.Ref}}](#{{FilterSchema $resp.Schema.Ref}})) |

#### {{FilterSchema $resp.Schema.Ref}}

{{- template "schema.md" CollectSchema $definitions  $resp.Schema.Ref}}

{{ else }} 
| 描述 | 类型 |
| ----------- | ------ |
| {{$resp.Description}} | Object(<业务对象>) |

{{- end }}
{{- end }}
{{- end }}

## 响应

{{- range $code, $resp := .Responses}}
{{if ne $code "default"}}

### 响应<{{$code}}>

{{- if ne $resp.Schema.Items.Ref  ""}}   
| Code | 描述 | 类型 |
| ---- | ----------- | ------ |
| {{$code}} | {{$resp.Description}} | Array[{{FilterSchema $resp.Schema.Items.Ref}}](#{{FilterSchema $resp.Schema.Items.Ref}}) |

#### {{FilterSchema $resp.Schema.Items.Ref}}

{{ template "schema.md" CollectSchema $definitions  $resp.Schema.Items.Ref}}
{{- else if ne $resp.Schema.Ref  "" }}
| Code | 描述 | 类型 |
| ---- | ----------- | ------ | 
| {{$code}} | {{$resp.Description}} | Object([{{FilterSchema $resp.Schema.Ref}}](#{{FilterSchema $resp.Schema.Ref}})) |

#### {{FilterSchema $resp.Schema.Ref}}

{{ template "schema.md" CollectSchema $definitions  $resp.Schema.Ref}}
{{- else if eq $resp.Schema.Type  "" }}
| Code | 描述 | 类型 |
| ---- | ----------- | ------ | 
| {{$code}} | {{$resp.Description}} | - |

{{- else}}
| Code | 描述 | 类型 |
| ---- | ----------- | ------ |
| {{$code}} | {{$resp.Description}} | {{$resp.Schema}} |
{{- end}} 

{{- end}}
{{- end}}

