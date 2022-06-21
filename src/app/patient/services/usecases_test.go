package services

import (
	"clinic-api/src/app/patient"
	"clinic-api/src/app/patient/mocks"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo mocks.Repositories
	services patient.Services

	sampleDomainList []patient.Domain

	sampleUUIDIkoUwais        uuid.UUID
	sampleUUIDTheressaSiahaan uuid.UUID

	sampleDomainIkoUwais           patient.Domain
	sampleDomainIkoUwaisWithNoUUID patient.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUIDIkoUwais = uuid.Must(uuid.NewRandom())
	sampleUUIDTheressaSiahaan = uuid.Must(uuid.NewRandom())

	sampleDomainIkoUwaisWithNoUUID = patient.Domain{
		Name:      "Iko Uwais",
		NIK:       "1122330201060002",
		Phone:     "+62 825 1111 1111",
		Address:   "Jalan Sama Dia 21, North Way, Antarctica",
		DOB:       utils.ConvertStringToDate("2006-01-02"),
		Gender:    types.MALE,
		BloodType: "O+",
	}

	sampleDomainIkoUwais = sampleDomainIkoUwaisWithNoUUID
	sampleDomainIkoUwais.ID = sampleUUIDIkoUwais

	sampleDomainList = []patient.Domain{
		sampleDomainIkoUwais,
		{
			ID:        sampleUUIDTheressaSiahaan,
			Name:      "Theressa Siahaan",
			NIK:       "2211334201060002",
			Phone:     "+62 825 1111 1111",
			Address:   "Jalan Sama Dia 21, North Way, Antarctica",
			DOB:       utils.ConvertStringToDate("2006-01-02"),
			Gender:    types.MALE,
			BloodType: "A+",
		},
	}

	os.Exit(m.Run())
}

func TestHuntPatientByNameOrNIKOrAll(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData", 0).Return(sampleDomainList, nil).Once()
		result, err := services.HuntPatientByNameOrNIKOrAll(patient.Domain{}, 1)

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should found data by querying name", func(t *testing.T) {
		var nameOnlyDomain patient.Domain
		nameOnlyDomain.Name = sampleDomainIkoUwais.Name

		mockRepo.On("SearchDataByNameParam", nameOnlyDomain.Name, 0).
			Return([]patient.Domain{sampleDomainIkoUwais}, nil).Once()
		result, err := services.HuntPatientByNameOrNIKOrAll(nameOnlyDomain, 1)

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should found data by querying NIK", func(t *testing.T) {
		var nikOnlyDomain patient.Domain
		nikOnlyDomain.NIK = sampleDomainIkoUwais.NIK

		mockRepo.On("SearchDataByNIKParam", nikOnlyDomain.NIK).
			Return([]patient.Domain{sampleDomainIkoUwais}, nil).Once()
		result, err := services.HuntPatientByNameOrNIKOrAll(nikOnlyDomain, 0)

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should got empty array on search data by name not found", func(t *testing.T) {
		var nameOnlyDomain patient.Domain
		nameOnlyDomain.Name = sampleDomainIkoUwais.Name

		mockRepo.On("SearchDataByNameParam", nameOnlyDomain.Name, 0).
			Return(nil, nil).Once()
		result, err := services.HuntPatientByNameOrNIKOrAll(nameOnlyDomain, 1)

		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	})

	t.Run("should got empty array on search data by NIK not found", func(t *testing.T) {
		var nikOnlyDomain patient.Domain
		nikOnlyDomain.NIK = sampleDomainIkoUwais.NIK

		mockRepo.On("SearchDataByNIKParam", nikOnlyDomain.NIK).
			Return(nil, nil).Once()
		result, err := services.HuntPatientByNameOrNIKOrAll(nikOnlyDomain, 0)

		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	})

	t.Run("should got server error", func(t *testing.T) {
		mockRepo.On("SelectAllData", 0).Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.HuntPatientByNameOrNIKOrAll(patient.Domain{}, 1)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetPatientByID(t *testing.T) {
	t.Run("should got data by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).Return(&sampleDomainIkoUwais, nil).Once()
		result, err := services.GetPatientByID(sampleUUIDIkoUwais.String())

		assert.Nil(t, err)
		assert.Equal(t, sampleDomainIkoUwais.Name, result.Name)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetPatientByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).Return(nil, errors.New("record not found"))
		result, err := services.GetPatientByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestCreatePatient(t *testing.T) {
	t.Run("should created a data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainIkoUwaisWithNoUUID).Return(sampleUUIDIkoUwais.String(), nil).Once()
		result, err := services.CreatePatient(sampleDomainIkoUwaisWithNoUUID)

		assert.Nil(t, err)
		assert.Equal(t, sampleUUIDIkoUwais.String(), result)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainIkoUwaisWithNoUUID).Return(uuid.Nil.String(), errors.New("can't connect to the database")).Once()
		result, err := services.CreatePatient(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Equal(t, uuid.Nil.String(), result)
	})
}

func TestUpdatePatientByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).Return(nil).Once()
		err := services.AmendPatientByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).Return(errors.New("can't connect to the database")).Once()
		err := services.AmendPatientByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).Return(errors.New("record not found")).Once()
		err := services.AmendPatientByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
	})
}

func TestRemovePatientByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		err := services.RemovePatientByID(sampleUUIDIkoUwais.String())

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(errors.New("can't connect to the database")).Once()
		err := services.RemovePatientByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(errors.New("record not found")).Once()
		err := services.RemovePatientByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
	})
}
