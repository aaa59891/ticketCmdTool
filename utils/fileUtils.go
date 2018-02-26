package utils

import (
	"bytes"
	"io/ioutil"
)

func RemoveFirstLine(filepath string) error {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	index := bytes.IndexByte(b, '\n')
	var content []byte
	if index < 0 {
		content = []byte{}
	} else {
		content = b[index+1:]
	}
	return ioutil.WriteFile(filepath, content, 0664)
}
