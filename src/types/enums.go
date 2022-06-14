package types

type PatientStatusEnum string

const (
	OUTPATIENT PatientStatusEnum = "OUTPATIENT"
	REFERRED   PatientStatusEnum = "REFERRED"
)

type GenderEnum string

const (
	MALE   GenderEnum = "MALE"
	FEMALE GenderEnum = "FEMALE"
)
