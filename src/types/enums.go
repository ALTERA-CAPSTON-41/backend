package types

type PatientStatusEnum string

const (
	OUTPATIENT PatientStatusEnum = "OUTPATIENT"
	REFERRED   PatientStatusEnum = "REFERRED"
)

type GenderTypeEnum string

const (
	MALE   GenderTypeEnum = "MALE"
	FEMALE GenderTypeEnum = "FEMALE"
)
