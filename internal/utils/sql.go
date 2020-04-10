package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Script returnes sequence  of sql commands
func Script(dirName, fileName string) ([]string, error) {
	dirName, err := filepath.Abs(dirName)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dirName, fileName)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sdata := strings.TrimRight(string(data), ";\n")
	return strings.Split(sdata, ";"), nil
}
