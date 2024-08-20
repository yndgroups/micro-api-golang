package create

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// UderscoreToUpperCamelCase 下划线单词转为大写驼峰单词
func UderscoreToUpperCamelCase(target string) string {
	target = strings.Replace(target, "_", " ", -1)
	target = cases.Title(language.English).String(target)
	return strings.Replace(target, " ", "", -1)
}

// UderscoreToLowerCamelCase 下划线单词转为小写驼峰单词
func UderscoreToLowerCamelCase(target string) string {
	target = UderscoreToUpperCamelCase(target)
	return string(unicode.ToLower(rune(target[0]))) + target[1:]
}

// CamelCaseToUdnderscore 驼峰单词转下划线单词
func CamelCaseToUdnderscore(target string) string {
	var output []rune
	for i, r := range target {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}
