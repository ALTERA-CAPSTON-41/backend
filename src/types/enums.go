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

type UserRoleEnum string

const (
	DOCTOR UserRoleEnum = "DOCTOR"
	NURSE  UserRoleEnum = "NURSE"
	ADMIN  UserRoleEnum = "ADMIN"
)
