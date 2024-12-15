package handler

// contextKey is used in order to store key-value pairs inside a context
type contextKey string

func (c contextKey) String() string {
	return string(c)
}
