package utils

import (
	"os"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

func GetUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func IsExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsExist(err) {
			return false
		}
		return false
	}
	return false
}
