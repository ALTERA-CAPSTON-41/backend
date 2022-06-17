package request

type Request struct {
	Name      string `json:"name"`
	NIK       string `json:"nik"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	BloodType string `json:"blood_type"`
}
