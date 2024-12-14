package handler

// Handlers is map that binds the user input commands to an executable function.
var (
	Handlers = map[string]func(argc int, argv []string) error{
		"request": Request,
	}
)

// Request accepts a (sender, receiver, amount) and dials a request.
func Request(argc int, argv []string) error {
	return nil
}
