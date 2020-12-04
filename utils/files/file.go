package files

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string, delimiter string) []string {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

	return slicedContent
}
