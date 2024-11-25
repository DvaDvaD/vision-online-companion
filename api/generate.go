package main

//go:generate -command sqlc sqlc generate
//go:generate -command api ./scripts/generate-api.sh

//go:generate sqlc
//go:generate api

// Example usage:
// go generate -run sqlc
// go generate -run api

// Notes:
// Make sure you already install sqlc and oapi-codegen. Check readme.