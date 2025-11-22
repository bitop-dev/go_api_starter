package auth

// JWTProvider defines the expected behavior for JWT-based auth implementations.
type JWTProvider interface {
	Generate(subject string, claims map[string]interface{}) (string, error)
	Validate(token string) (map[string]interface{}, error)
}
