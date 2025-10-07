package auth

// Auth
type LoginInput struct {
	Phone    string
	Password string
}

type GenerateTokenInput struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}

type GenerateTokenResponse struct {
	Token string
}
