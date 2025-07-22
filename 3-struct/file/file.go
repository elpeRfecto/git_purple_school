package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
	}
	fmt.Printf("Содержимое файла: %s", file)
}

func IsJSON(filename string) bool {
	if strings.Contains(filename, ".json") {
		return true
	}
	return false
}
