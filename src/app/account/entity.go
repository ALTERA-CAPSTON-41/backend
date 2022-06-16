package account

import (
	"clinic-api/src/types"

	"github.com/google/uuid"
)

type Domain struct {
	ID       uuid.UUID
	Email    string
	Password string
	Role     types.UserRoleEnum
}
