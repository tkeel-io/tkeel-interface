package add

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type Proto interface {
	Generate() error
}

// APIProto is a proto generator.
type APIProto struct {
	Name        string
	Path        string
	Service     string
	Package     string
	GoPackage   string
	JavaPackage string
}

func (p *APIProto) execute() ([]byte, error) {
	funcMap := template.FuncMap{
		"toLower": strings.ToLower,
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Funcs(funcMap).Parse(strings.TrimSpace(protoTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Generate generate a proto template.
func (p *APIProto) Generate() error {
	body, err := p.execute()
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	to := path.Join(wd, p.Path)
	if _, err := os.Stat(to); os.IsNotExist(err) {
		if err := os.MkdirAll(to, 0o700); err != nil {
			return err
		}
	}
	name := path.Join(to, p.Name)
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", p.Name)
	}
	return ioutil.WriteFile(name, body, 0o644)
}

type ErrorProto struct {
	Name        string
	Path        string
	Package     string
	GoPackage   string
	JavaPackage string
}

func (p *ErrorProto) execute() ([]byte, error) {
	funcMap := template.FuncMap{
		"toLower": strings.ToLower,
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Funcs(funcMap).Parse(strings.TrimSpace(errTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *ErrorProto) Generate() error {
	body, err := p.execute()
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	to := path.Join(wd, p.Path)
	if _, err := os.Stat(to); os.IsNotExist(err) {
		if err := os.MkdirAll(to, 0o700); err != nil {
			return err
		}
	}
	name := path.Join(to, p.Name)
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", p.Name)
	}
	return ioutil.WriteFile(name, body, 0o644)
}
