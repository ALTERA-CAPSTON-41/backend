package response

import "clinic-api/src/types"

type AuthResponse struct {
	Token string             `json:"token"`
	Role  types.UserRoleEnum `json:"role"`
}

type AuthErrorResponse struct {
	Reason string `json:"reason"`
}
