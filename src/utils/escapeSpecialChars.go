package utils

import (
	"fmt"
	"strings"
)

func EscapeSpecialChars(s string) string {
	// Заменяем кавычки и точки с запятой на безопасные эквиваленты
	s = strings.ReplaceAll(s, "'", "''")
	s = strings.ReplaceAll(s, ";", "")
	fmt.Println("pasrsed \n", s)
	return s
}
