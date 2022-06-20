package services

import (
	"clinic-api/src/app/nurse"
	"clinic-api/src/app/nurse/mocks"
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
	mockRepo                  mocks.Repositories
	services                  nurse.Services
	sampleDomainList          []nurse.Domain
	sampleDomainEko           nurse.Domain
	sampleDomainEkoWithNoUUID nurse.Domain
	sampleUUIDEko             uuid.UUID
	sampleUUIDSusan           uuid.UUID
	samplePassword            string
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUIDEko = uuid.Must(uuid.NewRandom())
	sampleUUIDSusan = uuid.Must(uuid.NewRandom())
	samplePassword = "strongpassword"

	sampleDomainEko = nurse.Domain{
		Name:    "Eko",
		NIP:     "2022 0211021 43",
		SIP:     "SIP.2132.12.2322",
		Address: "Sehat Selalu",
		DOB:     utils.ConvertStringToDate("1981-08-01"),
		Gender:  types.MALE,
		Polyclinic: nurse.PolyclinicReference{
			ID:   1,
			Name: "General",
		},
		User: nurse.UserReference{
			ID:        sampleUUIDEko,
			Email:     "eko@example.com",
			Password:  samplePassword,
			Role:      types.NURSE,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	sampleDomainEkoWithNoUUID = nurse.Domain{
		Name:    "Eko",
		NIP:     "2022 0211021 43",
		SIP:     "SIP.2132.12.2322",
		Address: "Sehat Selalu",
		DOB:     utils.ConvertStringToDate("1981-08-01"),
		Gender:  types.MALE,
		Polyclinic: nurse.PolyclinicReference{
			ID:   1,
			Name: "General",
		},
		User: nurse.UserReference{
			Email:     "eko@example.com",
			Password:  samplePassword,
			Role:      types.NURSE,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	sampleDomainList = []nurse.Domain{
		sampleDomainEko,
		{
			Name:    "Susan",
			NIP:     "2013 02131 13",
			Address: "New Greeen",
			DOB:     utils.ConvertStringToDate("2001-03-20"),
			Gender:  types.FEMALE,
			Polyclinic: nurse.PolyclinicReference{
				ID:   1,
				Name: "General",
			},
			User: nurse.UserReference{
				ID:        sampleUUIDSusan,
				Email:     "susan@example.com",
				Password:  samplePassword,
				Role:      types.NURSE,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	os.Exit(m.Run())
}

func TestCreateNurse(t *testing.T) {
	t.Run("should created a data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainEkoWithNoUUID).
			Return(sampleUUIDEko.String(), nil).Once()
		result, err := services.CreateNurse(sampleDomainEkoWithNoUUID)

		assert.Nil(t, err)
		assert.Equal(t, sampleUUIDEko.String(), result)
	})

	t.Run("should got an error", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainEkoWithNoUUID).
			Return(uuid.Nil.String(), errors.New("can't connect to the database")).Once()
		result, err := services.CreateNurse(sampleDomainEkoWithNoUUID)

		assert.NotNil(t, err)
		assert.Equal(t, uuid.Nil.String(), result)
	})
}

func TestGetAllNurses(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllNurses()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 1)
	})

	t.Run("should got an error", func(t *testing.T) {
		mockRepo.On("SelectAllData").
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAllNurses()

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetNurseByID(t *testing.T) {
	t.Run("should got data by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDEko.String()).
			Return(&sampleDomainEko, nil).Once()
		result, err := services.GetNurseByID(sampleUUIDEko.String())

		assert.Nil(t, err)
		assert.Equal(t, sampleDomainEko.Name, result.Name)
	})

	t.Run("should got an error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDEko.String()).
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetNurseByID(sampleUUIDEko.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdateNurseByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDEko.String(), sampleDomainEkoWithNoUUID).
			Return(nil).Once()
		err := services.AmendNurseByID(sampleUUIDEko.String(), sampleDomainEkoWithNoUUID)

		assert.Nil(t, err)
	})

	t.Run("should got an error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDEko.String(), sampleDomainEkoWithNoUUID).
			Return(errors.New("can't connect to the database")).Once()
		err := services.AmendNurseByID(sampleUUIDEko.String(), sampleDomainEkoWithNoUUID)

		assert.NotNil(t, err)
	})
}

func TestDeleteDoctorByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDEko.String()).Return(nil).Once()
		mockRepo.On("DeleteUserByID", sampleUUIDEko.String()).Return(nil).Once()
		err := services.RemoveNurseByID(sampleUUIDEko.String())

		assert.Nil(t, err)
	})

	t.Run("should got an error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDEko.String()).
			Return(errors.New("can't connect to the database")).Once()
		err := services.RemoveNurseByID(sampleUUIDEko.String())

		assert.NotNil(t, err)
	})
}
