package gen

import (
	"strings"
	"unicode"
)

// 处理名称
func handleName(name string) string {
	var result strings.Builder
	// 用非字母数字字符分割（例如下划线、连字符）
	words := strings.FieldsFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		if word == "" {
			continue
		}
		// 处理每个单词：首字母大写，其余小写
		if len(word) > 0 {
			first := strings.ToUpper(string(word[0]))
			rest := strings.ToLower(word[1:])
			result.WriteString(first)
			result.WriteString(rest)
		}
	}
	return result.String()
}
