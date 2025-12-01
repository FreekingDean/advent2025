package utils

import "os"

func LoadInput(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	return string(data), err
}
