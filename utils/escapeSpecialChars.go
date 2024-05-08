package utils

import (
	"strings"
)

func EscapeSpecialChars(s string) string {
	// Заменяем кавычки и точки с запятой на безопасные эквиваленты
	s = strings.ReplaceAll(s, "'", "''")
	s = strings.ReplaceAll(s, ";", "")
	return s
}
