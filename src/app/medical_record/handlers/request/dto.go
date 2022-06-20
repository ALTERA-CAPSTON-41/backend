package request

type Request struct {
	Symptoms    string `json:"symptoms"`
	ICD10Code   string `json:"icd10_code"`
	Suggestions string `json:"suggestions"`
	PatientID   string `json:"patient_id"`
	DoctorID    string `json:"doctor_id"`
	Polyclinic  string `json:"polyclinic_id"`
}
