package models

type GenderType string

const (
	MALE   GenderType = "MALE"
	FEMALE GenderType = "FEMALE"
)

type PatientStatus string

const (
	OUTPATIENT PatientStatus = "OUTPATIENT"
	REFERRED   PatientStatus = "REFERRED"
)
