package token

import "time"

// Interface for managing tokens
type Maker interface {
	// Creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// Checks if token is valid
	VerifyToken(token string) (*Payload, error)
}
