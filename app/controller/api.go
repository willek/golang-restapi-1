package controller

import (
	"../../function"
)

// SetupEndpoint API
func SetupEndpoint() *function.Response {
	r := function.Response{
		Status:  "success",
		Message: "API Endpoint settled",
	}

	return &r
}
