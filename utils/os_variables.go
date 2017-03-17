package utils

import (
	"os"
	"regexp"
	"strings"
)

func ReplaceOSVariables(expr string) string {
	b := `\${\w*}`
	r, _ := regexp.Compile(b)
	vars := r.FindAllString(expr, -1)
	for _, v := range vars {
		key := v[2 : len(v)-1]
		expr = strings.Replace(expr, v, os.Getenv(key), -1)
	}
	return expr
}
