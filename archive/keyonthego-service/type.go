package main

type NetworkInterface struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Dsn  string `json:"dsn"`
}

type Issue struct {
	Number      string `json:"number" validate:"required"`
	Copy        int    `json:"copy" validate:"required"`
	Description string `json:"description"`
}

type Status string

const (
	StatusPending Status = "pending"
	StatusExpired Status = "expired"
	StatusFailed  Status = "failed"
	StatusSuccess Status = "success"
)

// ==========================
// Request Body
// ==========================

// CreateSignRequest represents the request payload for creating a signing request
type CreateSignRequest struct {
	RequestUser string  `json:"request_user" validate:"required"`
	HolderID    string  `json:"holder_id" validate:"required"`
	HolderName  string  `json:"holder_name" validate:"required"`
	Notes       string  `json:"notes"`
	Issue       []Issue `json:"issue" validate:"required"`
}

// SignSubmitRequest represents the request payload for submitting a signing
type SignSubmitRequest struct {
	Sign string `json:"sign,omitempty" validate:"required"`
}

// ==========================
// Response Body
// ==========================

// CreateSignResponse represents the response payload for signing operations
type CreateSignResponse struct {
	RequestID  string             `json:"request_id"`
	Token      string             `json:"token"`
	Interfaces []NetworkInterface `json:"interfaces"`
}

// SignResponse represents the response payload for signing operations
type SignResponse struct {
	CreateSignRequest
	SignSubmitRequest
	RequestID string `json:"request_id"`
	Status    Status `json:"status"`
}

// ==========================
// Data that saved to file
// ==========================
type SignDataBody struct {
	CreateSignRequest
	Sign string `json:"sign,omitempty" validate:"required"`
}
type SignData struct {
	Token string       `json:"token"`
	Body  SignDataBody `json:"body"`
}
