package markdown

import (
	"strings"
	"testing"
)

func TestFilterSchema(t *testing.T) {
	schema := "/helloworld2/{name}"
	schema = strings.Replace(schema, "/", "_", -1)
	schema = strings.Replace(schema, "{", "_", -1)
	schema = strings.Replace(schema, "}", "_", -1)
	t.Log(schema)
}
