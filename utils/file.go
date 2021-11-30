package utils

import (
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)


func ReadFile(delimiter string) []string{
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	// parse directory with pathFromCaller (which could be relative to Directory)
	absolutePath := path.Join(path.Dir(filename), "input.txt")
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}

	fileContent := string(content)
	slicedContent := strings.Split(fileContent, delimiter)
	if delimiter == "\n" {
		return slicedContent[:len(slicedContent) - 1]
	} else {
		lastElement := slicedContent[len(slicedContent) - 1]
		slicedContent[len(slicedContent) - 1] = lastElement[:len(lastElement) - 1]

		return slicedContent
	}
}