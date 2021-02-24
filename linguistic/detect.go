package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-enry/go-enry/v2"
)

func main() {
	var files []string

	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	var result = map[string]int{}
	for _, file := range files {
		lang := enry.GetLanguage(file, []byte("<cpp-code>"))
		if len(lang) > 0 {
			result[lang]++

		}
	}
	fmt.Println(result)
}
