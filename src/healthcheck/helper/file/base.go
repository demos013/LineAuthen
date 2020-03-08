package file

import (
	"io/ioutil"
	"os"
)

func WriteFile(path, text string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = f.WriteString(text)
	if err != nil {
		f.Close()
		return err
	}
	return nil
}

func ReadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	text := string(content)
	return text, nil

}
