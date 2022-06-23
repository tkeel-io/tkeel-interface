package markdown

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

func DecodeJSON(data []byte) (*API, error) {
	var api API
	err := json.Unmarshal(data, &api)
	return &api, err
}

func DecodeYAML(data []byte) (*API, error) {
	var api API
	err := yaml.Unmarshal(data, &api)
	return &api, err
}

func FilterParameters(params []Parameter, t string) []Parameter {
	var results []Parameter
	for _, param := range params {
		if strings.ToLower(param.In) == t {
			results = append(results, param)
		}
	}
	return results
}

func FilterSchema(schema string) string {
	schema = strings.Replace(schema, "#/definitions/", "", 1)
	schema = strings.TrimPrefix(schema, "request.")
	return schema
}

func CollectSchema(definitions map[string]Definition, schema string) SchemaContext {
	schema = strings.Replace(schema, "#/definitions/", "", 1)
	return SchemaContext{
		schema,
		definitions,
		definitions[schema],
	}
}

func FormatAnchor(schema string) string {
	schema = strings.Replace(schema, "/:/g", "", 0)
	schemas := strings.Split(schema, " ")
	schema = strings.Join(schemas, "-")
	schema = strings.ToLower(schema)
	return schema
}

func FormatPath(schema string) string {
	//schema = strings.Replace(schema, "/()/g", "_", 0)
	schema = strings.Replace(schema, "/", "_", -1)
	schema = strings.Replace(schema, "{", "_", -1)
	schema = strings.Replace(schema, "}", "", -1)
	return "Path_" + schema
}

func StringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsExcludeTag(tags []string) bool {
	for _, tag := range tags {
		if tag == "Internal" {
			fmt.Println(1)
		}
		if ret := StringContains(ExcludeTags, tag); ret {
			return ret
		}
	}
	return false
}

type APIError map[string]string

type API struct {
	Swagger     string                `json,yaml:"swagger"`
	Info        Info                  `json,yaml:"info"`
	Host        string                `json,yaml:"host"`
	BasePath    string                `json,yaml:"basePath"`
	Schemes     []string              `json,yaml:"schemes"`
	Consumes    []string              `json,yaml:"consumes"`
	Produces    []string              `json,yaml:"produces"`
	Paths       map[string]Operations `json,yaml:"paths"`
	Definitions map[string]Definition `json,yaml:"definitions"`
}

type Operations map[string]*Operation

type Method struct {
	Index int
	*Operation
}

type Tag struct {
	TagIndex int
	Tag      string
	BasePath string
	Methods  map[int]*Method
}

type Operation struct {
	API         *API
	Operation   string
	Path        string
	Tags        []string              `json,yaml:"tags"`
	Description string                `json,yaml:"description"`
	OperationID string                `json,yaml:"operationId"`
	Summary     string                `json,yaml:"summary"`
	Parameters  []Parameter           `json,yaml:"parameters"`
	Responses   map[string]Response   `json,yaml:"responses"`
	Definitions map[string]Definition `json,yaml:"definitions"`
}

type Response struct {
	Description string       `json,yaml:"description"`
	Schema      SchemaObject `json,yaml:"schema"`
}

type SchemaObject struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Ref         string `json:"$ref"`
	Items       struct {
		Type string `json:"type"`
		Ref  string `json:"$ref"`
	} `json:"items"`
}

type Parameter struct {
	Name             string                 `json,yaml:"name"`
	In               string                 `json,yaml:"in"`
	Description      string                 `json,yaml:"description"`
	Required         bool                   `json,yaml:"required"`
	Type             string                 `json,yaml:"type"`
	CollectionFormat string                 `json,yaml:"collectionFormat"`
	Items            map[string]interface{} `json,yaml:"items"`
	Schema           SchemaObject           `json,yaml:"schema"`
}

type Info struct {
	Version        string  `json,yaml:"version"`
	Title          string  `json,yaml:"title"`
	Description    string  `json,yaml:"description"`
	TermsOfService string  `json,yaml:"termsOfService"`
	Contact        Contact `json,yaml:"contact"`
	License        License `json,yaml:"license"`
}

type Contact struct {
	Name  string `json,yaml:"name"`
	Email string `json,yaml:"email"`
	URL   string `json,yaml:"url"`
}

type License struct {
	Name string `json,yaml:"name"`
	URL  string `json,yaml:"url"`
}

type Definition struct {
	Type         string                  `json,yaml:"type"`
	Example      string                  `json,yaml:"example"`
	Properties   map[string]SchemaObject `json,yaml:"properties"`
	ExternalDocs ExternalDocumentation   `json,yaml:"external_docs"`
}

type ExternalDocumentation struct {
	Description string `json,yaml:"description"`
	URL         string `json,yaml:"url"`
}

func (u *Definition) UnmarshalJSON(data []byte) error {
	type (
		Alias Definition
	)
	aux := &struct {
		Example interface{} `json:"example"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Example != nil {
		if ret, err := json.MarshalIndent(aux.Example, "", "    "); err != nil {
			return err
		} else {
			u.Example = string(ret)
		}
	}
	return nil
}

type SchemaContext struct {
	TopRef      string
	Definitions map[string]Definition
	Definition  Definition
}
