package response

import (
	"clinic-api/src/types"

	"github.com/google/uuid"
)

type Response struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	NIK       string           `json:"nik"`
	Phone     string           `json:"phone"`
	Address   string           `json:"address"`
	DOB       string           `json:"dob"`
	Gender    types.GenderEnum `  json:"gender"`
	BloodType string           `json:"blood_type"`
}
