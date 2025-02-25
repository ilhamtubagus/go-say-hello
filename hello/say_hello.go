package go_say_hello

import "fmt"

func SayHello(message string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}
	return "Hello, " + message, nil
}
