package response

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthErrorResponse struct {
	Reason string `json:"reason"`
}
