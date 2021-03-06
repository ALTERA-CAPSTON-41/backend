package services

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/app/medical_record/mocks"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo mocks.Repositories
	services medicalrecord.Services

	sampleDomain            medicalrecord.Domain
	sampleNewDomainInput    medicalrecord.Domain
	sampleMedicalRecordUUID uuid.UUID
	samplePatientUUID       uuid.UUID
	sampleDoctorUUID        uuid.UUID
	sampleDomainList        []medicalrecord.Domain
)

func TestMain(m *testing.M) {
	services = NewServices(&mockRepo)
	sampleMedicalRecordUUID = uuid.Must(uuid.NewRandom())
	sampleDoctorUUID = uuid.Must(uuid.NewRandom())
	samplePatientUUID = uuid.Must(uuid.NewRandom())

	sampleDomain = medicalrecord.Domain{
		ID:               sampleMedicalRecordUUID,
		PatientStatus:    types.OUTPATIENT,
		Symptoms:         "fever, sneeze",
		ICD10Code:        "A750",
		ICD10Description: "Epidemic louse-borne typhus fever due to Rickettsia prowazekii",
		Suggestions:      "total rest 10 days",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		Patient: medicalrecord.PatientReference{
			ID:        samplePatientUUID,
			Name:      "Patient Capstone",
			NIK:       "3515111111010002",
			Phone:     "081010101001",
			Address:   "5-chome-77-7 Urafune-cho, Minami-ku, Kanagawa-ken",
			DOB:       utils.ConvertStringToDate("2001-11-11"),
			Gender:    types.MALE,
			BloodType: "O+",
		},
		Doctor: medicalrecord.DoctorReference{
			ID:     sampleDoctorUUID,
			Name:   "dr. Capstone",
			NIP:    "2015 028968 022",
			SIP:    "SIP22.912.2015.01",
			Gender: types.MALE,
		},
		Polyclinic: medicalrecord.PolyclinicReference{
			ID:   1,
			Name: "General",
		},
	}
	sampleDomain.Patient.Age = utils.CountIntervalByYearRoundDown(sampleDomain.Patient.DOB, sampleDomain.CreatedAt)

	sampleNewDomainInput = medicalrecord.Domain{
		ID:            sampleMedicalRecordUUID,
		PatientStatus: types.OUTPATIENT,
		Symptoms:      "fever, sneeze",
		ICD10Code:     "A750",
		Suggestions:   "total rest 10 days",
		Patient: medicalrecord.PatientReference{
			ID: samplePatientUUID,
		},
		Doctor: medicalrecord.DoctorReference{
			ID: sampleDoctorUUID,
		},
		Polyclinic: medicalrecord.PolyclinicReference{
			ID: 1,
		},
	}

	sampleDomainList = []medicalrecord.Domain{sampleDomain}

	os.Exit(m.Run())
}

func TestCreateMedicalRecord(t *testing.T) {
	const icd10DescriptionOfA750 string = "Epidemic louse-borne typhus fever due to Rickettsia prowazekii"

	t.Run("should create medical record", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return(icd10DescriptionOfA750, nil).Once()
		sampleNewDomainInput.ICD10Description = icd10DescriptionOfA750
		mockRepo.On("InsertData", sampleNewDomainInput).
			Return(sampleMedicalRecordUUID.String(), nil).Once()
		id, err := services.CreateMedicalRecord(sampleNewDomainInput)

		assert.Nil(t, err)
		assert.Equal(t, sampleMedicalRecordUUID.String(), id)
	})

	t.Run("should get error on consume external API", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return("", errors.New("can't get data from server")).Once()
		id, err := services.CreateMedicalRecord(sampleNewDomainInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	})

	t.Run("should get database error", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return(icd10DescriptionOfA750, nil).Once()
		mockRepo.On("InsertData", sampleNewDomainInput).
			Return("", errors.New("can't connect to the database")).Once()
		id, err := services.CreateMedicalRecord(sampleNewDomainInput)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	})
}

func TestGetAllMedicalRecordByPatientID(t *testing.T) {
	t.Run("should found some records", func(t *testing.T) {
		mockRepo.On("SelectDataByPatientID", sampleDomain.Patient.ID.String()).
			Return(sampleDomainList, nil).Once()
		domains, err := services.FindMedicalRecordByPatientID(sampleDomain.Patient.ID.String())

		assert.Nil(t, err)
		assert.Greater(t, len(domains), 0)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByPatientID", sampleDomain.Patient.ID.String()).
			Return(nil, errors.New("can't connect to the database")).Once()
		domains, err := services.FindMedicalRecordByPatientID(sampleDomain.Patient.ID.String())

		assert.NotNil(t, err)
		assert.Nil(t, domains)
	})
}

func TestGetAllMedicalRecordByPatientNIK(t *testing.T) {
	t.Run("should found some records", func(t *testing.T) {
		mockRepo.On("SelectPatientIDByNIK", sampleDomain.Patient.NIK).
			Return(sampleDomain.Patient.ID.String(), nil).Once()
		mockRepo.On("SelectDataByPatientID", sampleDomain.Patient.ID.String()).
			Return(sampleDomainList, nil).Once()
		domains, err := services.FindMedicalRecordByPatientNIK(sampleDomain.Patient.NIK)

		assert.Nil(t, err)
		assert.Greater(t, len(domains), 0)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectPatientIDByNIK", sampleDomain.Patient.NIK).
			Return("", errors.New("can't connect to the database")).Once()
		domains, err := services.FindMedicalRecordByPatientNIK(sampleDomain.Patient.NIK)

		assert.NotNil(t, err)
		assert.Zero(t, domains)
	})

	t.Run("should got an error while select data by patient id", func(t *testing.T) {
		mockRepo.On("SelectPatientIDByNIK", sampleDomain.Patient.NIK).
			Return(sampleDomain.Patient.ID.String(), nil).Once()
		mockRepo.On("SelectDataByPatientID", sampleDomain.Patient.ID.String()).
			Return(nil, errors.New("can't connect to the database")).Once()
		domains, err := services.FindMedicalRecordByPatientNIK(sampleDomain.Patient.NIK)

		assert.NotNil(t, err)
		assert.Nil(t, domains)
	})

	t.Run("should get no data", func(t *testing.T) {
		mockRepo.On("SelectPatientIDByNIK", sampleDomain.Patient.NIK).Return(sampleDomain.Patient.ID.String(), nil).Once()
		mockRepo.On("SelectDataByPatientID", sampleDomain.Patient.ID.String()).
			Return(nil, errors.New("record not found"))
		domains, err := services.FindMedicalRecordByPatientNIK(sampleDomain.Patient.NIK)

		assert.NotNil(t, err)
		assert.Nil(t, domains)
	})
}

func TestGetMedicalRecordByID(t *testing.T) {
	t.Run("should found record by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleMedicalRecordUUID.String()).
			Return(&sampleDomain, nil).Once()
		domain, err := services.FindMedicalRecordByID(sampleMedicalRecordUUID.String())

		assert.Nil(t, err)
		assert.NotNil(t, domain)
	})

	t.Run("should got error not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleMedicalRecordUUID.String()).
			Return(nil, errors.New("record not found")).Once()
		domain, err := services.FindMedicalRecordByID(sampleMedicalRecordUUID.String())

		assert.NotNil(t, err)
		assert.Nil(t, domain)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleMedicalRecordUUID.String()).
			Return(nil, errors.New("can't connect to the database")).Once()
		domain, err := services.FindMedicalRecordByID(sampleMedicalRecordUUID.String())

		assert.NotNil(t, err)
		assert.Nil(t, domain)
	})
}

func TestAmendMedicalRecordByID(t *testing.T) {
	const icd10DescriptionOfA750 string = "Epidemic louse-borne typhus fever due to Rickettsia prowazekii"

	t.Run("should update record", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return(icd10DescriptionOfA750, nil).Once()
		mockRepo.On("UpdateByID", sampleNewDomainInput, samplePatientUUID.String()).
			Return(nil).Once()
		err := services.AmendMedicalRecordByID(sampleNewDomainInput, samplePatientUUID.String())

		assert.Nil(t, err)
	})

	t.Run("should return error fetch external", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return("", errors.New("can't get data from server")).Once()
		err := services.AmendMedicalRecordByID(sampleNewDomainInput, samplePatientUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should return error not found", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return(icd10DescriptionOfA750, nil).Once()
		mockRepo.On("UpdateByID", sampleNewDomainInput, samplePatientUUID.String()).
			Return(errors.New("record not found")).Once()
		err := services.AmendMedicalRecordByID(sampleNewDomainInput, samplePatientUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should return database error", func(t *testing.T) {
		mockRepo.On("LookupICD10Data", sampleNewDomainInput.ICD10Code).
			Return(icd10DescriptionOfA750, nil).Once()
		mockRepo.On("UpdateByID", sampleNewDomainInput, samplePatientUUID.String()).
			Return(errors.New("can't connect to the database")).Once()
		err := services.AmendMedicalRecordByID(sampleNewDomainInput, samplePatientUUID.String())

		assert.NotNil(t, err)
	})
}

func TestRemoveMedicalRecordByID(t *testing.T) {
	t.Run("should success remove record", func(t *testing.T) {
		mockRepo.On("DeleteByID", samplePatientUUID.String()).
			Return(nil).Once()
		err := services.RemoveMedicalRecordByID(samplePatientUUID.String())

		assert.Nil(t, err)
	})

	t.Run("should return not found error", func(t *testing.T) {
		mockRepo.On("DeleteByID", samplePatientUUID.String()).
			Return(errors.New("record not found")).Once()
		err := services.RemoveMedicalRecordByID(samplePatientUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should success remove record", func(t *testing.T) {
		mockRepo.On("DeleteByID", samplePatientUUID.String()).
			Return(errors.New("cn't connect to the database")).Once()
		err := services.RemoveMedicalRecordByID(samplePatientUUID.String())

		assert.NotNil(t, err)
	})
}
