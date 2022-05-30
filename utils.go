package main

import "os"

func readCode(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	code := string(content)
	return code, nil
}
