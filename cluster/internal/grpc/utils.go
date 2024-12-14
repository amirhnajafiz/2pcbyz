package grpc

import (
	"fmt"
	"strings"
)

// parseFullMethod accepts a gRPC full-method and returns the service name and the method.
func parseFullMethod(fullMethod string) (service, method string, err error) {
	// ensure the method starts with a slash
	if !strings.HasPrefix(fullMethod, "/") {
		return "", "", fmt.Errorf("invalid full method format: %s", fullMethod)
	}

	// split by the last slash
	parts := strings.SplitN(fullMethod[1:], "/", 2) // remove the leading slash before splitting
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid full method format: %s", fullMethod)
	}

	// remove the dot from the service name (e.g., api.Service -> apiService)
	service = strings.ReplaceAll(parts[0], ".", "")

	// convert the method name to lowercase
	method = strings.ToLower(parts[1])

	return service, method, nil
}
