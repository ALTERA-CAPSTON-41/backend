package request

import "clinic-api/src/app/admin"

type NewRequest struct {
	Name      string `json:"name"`
	NIPNumber string `json:"nip"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateRequest struct {
	Name      string `json:"name"`
	NIPNumber string `json:"nip"`
}

func (req *NewRequest) MapToDomain() admin.Domain {
	return admin.Domain{
		Name: req.Name,
		NIP:  req.NIPNumber,
		User: admin.UserReference{
			Email:    req.Email,
			Password: req.Password,
		},
	}
}

func (req *UpdateRequest) MapToDomain() admin.Domain {
	return admin.Domain{
		Name: req.Name,
		NIP:  req.NIPNumber,
	}
}
