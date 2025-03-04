package hello

import "fmt"

func SayHellow(message string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}
	return "Hello, " + message, nil
}
