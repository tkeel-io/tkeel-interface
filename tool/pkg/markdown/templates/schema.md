| 字段名 | 类型 | 描述 |
| ---- | ---- | ----------- |  
{{- range $code, $resp := .Definition.Properties}}  
    {{- if eq $resp.Type  "array" }}   
        {{- if eq $resp.Items.Ref  "" }} 
| {{$code}} | Array[ {{FilterSchema $resp.Items.Type}} ] | {{$resp.Description}} | 
        {{- else}}  
| {{$code}} | Array[{{FilterSchema $resp.Items.Ref}}] | {{$resp.Description}} {{template "schema_description.md" $resp.Items.Ref}} | 
        {{- end}}  
    {{- else if $resp.Ref}}
| {{$code}} | Object | {{$resp.Description}} {{template "schema_description.md" $resp.Ref}}  |  
    {{- else}} 
| {{$code}} | {{$resp.Type}} | {{$resp.Description}} |  
    {{- end}} 
{{- end}}

{{if .Definition.Example }}
```jsx title="Example"
{{.Definition.Example}}
```
{{- end}}

{{if .Definition.ExternalDocs.Description }}
```jsx title="ExternalDocs.Description"
{{.Definition.ExternalDocs.Description}}
```
{{- end}}

{{$definitions := .Definitions}}
{{- range $code, $resp := .Definition.Properties -}}  
    {{- if eq $resp.Type  "array" -}}   
        {{- if ne $resp.Items.Ref  "" -}}
            {{- $nextRefName := (FilterSchema $resp.Items.Ref) -}}
            {{- if ne $nextRefName $.TopRef -}}
### {{$nextRefName}}
{{template "schema.md" CollectSchema $definitions  $resp.Items.Ref}}
            {{- end -}}
        {{- end -}}  
    {{- else -}}
        {{- if ne $resp.Ref  ""  -}}
            {{- $nextRefName := (FilterSchema $resp.Ref) -}}
            {{- if ne $nextRefName $.TopRef -}}
### {{$nextRefName}}
{{template "schema.md" CollectSchema $definitions  $resp.Ref}}
            {{- end -}}
        {{- end -}}  
    {{- end -}} 
 {{- end -}}

 
 

