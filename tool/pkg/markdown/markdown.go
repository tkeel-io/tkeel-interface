package markdown

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
)

var (
	ExcludeTags = []string{}
)

type Template struct {
	mode string
	tmpl *template.Template
}

func New(templatePath, mode string) *Template {
	SwagdownTemplates := filepath.Join(filepath.Dir(templatePath), "*")
	funcMap := template.FuncMap{
		"FilterParameters": FilterParameters,
		"FilterSchema":     FilterSchema,
		"CollectSchema":    CollectSchema,
		"FormatAnchor":     FormatAnchor,
		"FormatPath":       FormatPath,
		"StringContains":   StringContains,
		"IsExcludeTag":     IsExcludeTag,
	}
	for name, fn := range sprig.FuncMap() {
		funcMap[name] = fn
	}

	if templatePath == "" {
		return &Template{
			mode,
			template.Must(template.New(fmt.Sprintf("%s.md", mode)).Funcs(funcMap).ParseFS(f, "templates/*")),
		}
	} else {
		return &Template{
			mode,
			template.Must(template.New(fmt.Sprintf("%s.md", mode)).Funcs(funcMap).ParseGlob(SwagdownTemplates)),
		}
	}
}

func RenderFromJSON(w Writer, r io.Reader, tmpl *Template) error {
	return renderAPI(tmpl, w, r, DecodeJSON)
}

func RenderFromYAML(w Writer, r io.Reader, tmpl *Template) error {
	return renderAPI(tmpl, w, r, DecodeYAML)
}

func renderAPI(tmpl *Template, w Writer, r io.Reader, decode func(data []byte) (*API, error)) error {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}

	api, err := decode(buf.Bytes())
	if err != nil {
		return err
	}

	switch tmpl.mode {
	case "tag": // 文档目录
		return renderTag(tmpl, w, api)
	case "method": // 文档函数
		return renderMethod(tmpl, w, api)
	case "sdk.ts": // 文档函数
		return renderSDK(tmpl, w, api)
	}

	return errors.New("error mode")
}

func renderMethod(tmpl *Template, w Writer, api *API) error {
	tags := parseTags(api)
	for _, tag := range tags {
		//tagName := tag.Tag
		for _, method := range tag.Methods {
			if IsExcludeTag(method.Tags) {
				continue
			}
			filename := method.OperationID
			if filename == "" {
				continue
			}
			filename = fmt.Sprintf("%s_%s.md", "method", filename)
			if err := tmpl.tmpl.Execute(w.For(filename), method); err != nil {
				return err
			}
			//if err := FormatMarkdownFile(filename); err != nil {
			//	return err
			//}
		}
	}
	return nil
}

func renderTag(tmpl *Template, w Writer, api *API) error {
	tags := parseTags(api)
	filename := fmt.Sprintf("tag.md")
	if err := tmpl.tmpl.Execute(w.For(filename), tags); err != nil {
		return err
	}
	//if err := FormatMarkdownFile(filename); err != nil {
	//	return err
	//}
	return nil
}

func renderSDK(tmpl *Template, w Writer, api *API) error {
	tags := parseTags(api)
	for _, tag := range tags {
		// tagName := tag.Tag
		for _, method := range tag.Methods {
			filename := method.OperationID
			if filename == "" {
				continue
			}
			if err := tmpl.tmpl.Execute(w.For(fmt.Sprintf("%s_%s.ts", "method", filename)), method); err != nil {
				return err
			}
		}
	}
	return nil
}

func parseTags(api *API) map[string]*Tag {
	tagIndex := 0
	methodIndex := 0
	step := getMaxMethodNum(api)
	tags := make(map[string]*Tag)
	for path, operations := range api.Paths {
		for typ, openator := range operations {
			method := &Method{Operation: openator}
			if len(method.Tags) == 0 {
				fmt.Printf("skip:method(%v) without tag\n", method.Summary)
				continue
			}

			if IsExcludeTag(method.Tags) {
				fmt.Printf("skip:method(%v) with Tag(:%v) in ExcludeTag(%v)\n", method.Summary, method.Tags, ExcludeTags)
				continue
			}

			key := method.Tags[0]
			if key == "" {
				fmt.Printf("error:method(%v) tag empty\n", method.Summary)
				continue
			}

			if method.OperationID == "" {
				fmt.Printf("error:tag(%v)method(%v) OperationID empty\n", method.Tags, method.Summary)
				continue
			}

			tag, ok := tags[key]
			if !ok {
				tag = &Tag{TagIndex: tagIndex, Tag: key, Methods: make(map[int]*Method, 0)}
				tags[key] = tag
				tagIndex += 10
			}
			method.Definitions = api.Definitions
			method.Operation.Operation = typ
			method.Path = filepath.Join(api.BasePath, path)
			method.Index = tagIndex*step + methodIndex
			tag.Methods[method.Index] = method
			methodIndex++
		}
	}
	return tags
}

func getMaxMethodNum(api *API) int {
	maxnum := 0
	for _, methods := range api.Paths {
		if len(methods) > maxnum {
			maxnum = len(methods)
		}
	}
	return maxnum
}
