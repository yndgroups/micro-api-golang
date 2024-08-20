package create

import (
	"html/template"
	"strings"
)

var funcMap template.FuncMap

func createFuncMap() template.FuncMap {
	funcMap := make(map[string]any, 0)
	funcMap["index"] = index
	funcMap["varchar"] = varchar
	return funcMap
}

// 扩展索引
func index(in int) (out int) {
	out = in + 1
	return out
}

// 字符串类型
func varchar(in string) (out bool) {
	return strings.Contains(in, "varchar")
}
