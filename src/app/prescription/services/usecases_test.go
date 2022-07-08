package services

import (
	"clinic-api/src/app/prescription"
	"clinic-api/src/app/prescription/mocks"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo          mocks.Repositories
	services          prescription.Services
	sampleDomainInput prescription.Domain
	sampleUUID        uuid.UUID
	sampleDomainList  []prescription.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUID = uuid.Must(uuid.NewRandom())

	sampleDomainList = []prescription.Domain{
		{
			Name:        "paracetamol 500 mg",
			Quantity:    10,
			Dosage:      "3x1",
			Preparatory: "tablet",
			Description: "before eating",
		},
	}

	os.Exit(m.Run())
}

func TestCreatePrescription(t *testing.T) {
	t.Run("should create prescription", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainInput).
			Return(sampleUUID.String(), nil).Once()
		id, err := services.CreatePrescription(sampleDomainInput)

		assert.Nil(t, err)
		assert.Equal(t, sampleUUID.String(), id)
	})

	t.Run("should get database error", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainInput).
			Return("", errors.New("can't connect to the database")).Once()
		id, err := services.CreatePrescription(sampleDomainInput)

		assert.NotNil(t, err)
		assert.Zero(t, id)
	})
}

func TestGetAllPrescriptionsByID(t *testing.T) {
	t.Run("should found some records", func(t *testing.T) {
		mockRepo.On("SelectAllDataByMedRecordID", sampleUUID.String()).
			Return(sampleDomainList, nil).Once()
		domains, err := services.FindPrescriptionsByID(sampleUUID.String())

		assert.Nil(t, err)
		assert.Greater(t, len(domains), 0)
	})

	t.Run("should get database error", func(t *testing.T) {
		mockRepo.On("SelectAllDataByMedRecordID", sampleUUID.String()).
			Return(nil, errors.New("can't connect to the database")).Once()
		domains, err := services.FindPrescriptionsByID(sampleUUID.String())

		assert.NotNil(t, err)
		assert.Nil(t, domains)
	})
}

func TestUpdatePrescriptionByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUID.String(), sampleDomainInput).Return(nil).Once()
		err := services.AmendPrescriptionByID(sampleUUID.String(), sampleDomainInput)

		assert.Nil(t, err)
	})

	t.Run("should get an error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUID.String(), sampleDomainInput).
			Return(errors.New("can't connect to the database")).Once()
		err := services.AmendPrescriptionByID(sampleUUID.String(), sampleDomainInput)

		assert.NotNil(t, err)
	})
}

func TestDeletePrescriptionByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUID.String()).Return(nil).Once()
		err := services.RemovePrescriptionByID(sampleUUID.String())

		assert.Nil(t, err)
	})

	t.Run("should get an error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUID.String()).
			Return(errors.New("can't connect to the database")).Once()
		err := services.RemovePrescriptionByID(sampleUUID.String())

		assert.NotNil(t, err)
	})
}
