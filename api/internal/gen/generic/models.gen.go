// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version 2.3.0 DO NOT EDIT.
package api

// Response defines model for Response.
type Response struct {
	// Data Response payload that varies based on the endpoint
	Data    *map[string]interface{} `json:"data"`
	Message *string                 `json:"message,omitempty"`
	Status  *string                 `json:"status,omitempty"`
}
