package go_say_hello

func SayHello(message string) (string, error) {
	if message == "" {
		return nil, fmt.Errorf("message cannot be empty")
	}
	return "Hello, " + message
}
