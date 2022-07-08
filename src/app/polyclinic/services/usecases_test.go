package services

import (
	"clinic-api/src/app/polyclinic"
	"clinic-api/src/app/polyclinic/mocks"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockRepo                  mocks.Repositories
	services                  polyclinic.Services
	sampleDomainList          []polyclinic.Domain
	sampleDomainUmum          polyclinic.Domain
	sampleDomainListWithStats []polyclinic.Domain
	sampleDomainUmumWithNoID  polyclinic.Domain
	sampleIDUmum              int
	sampleIDMata              int
	sampleTotalDoctor         int
	sampleTotalNurse          int
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleIDUmum = 1
	sampleIDMata = 2
	sampleTotalDoctor = 10
	sampleTotalNurse = 15

	sampleDomainUmum = polyclinic.Domain{
		ID:   sampleIDUmum,
		Name: "umum",
	}

	sampleDomainList = []polyclinic.Domain{
		sampleDomainUmum,
		{
			ID:   sampleIDMata,
			Name: "mata",
		},
	}

	sampleDomainListWithStats = []polyclinic.Domain{
		{
			ID:          sampleIDUmum,
			Name:        "umum",
			TotalDoctor: sampleTotalDoctor,
			TotalNurse:  sampleTotalNurse,
		},
		{
			ID:          sampleIDMata,
			Name:        "mata",
			TotalDoctor: sampleTotalDoctor,
			TotalNurse:  sampleTotalNurse,
		},
	}

	sampleDomainUmumWithNoID = polyclinic.Domain{Name: "umum"}

	os.Exit(m.Run())
}

func TestCreatePolyclinic(t *testing.T) {
	t.Run("should created a data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainUmumWithNoID).Return(sampleIDUmum, nil).Once()
		result, err := services.CreatePolyclinic(sampleDomainUmumWithNoID)

		assert.Nil(t, err)
		assert.Equal(t, sampleIDUmum, result)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomainUmumWithNoID).
			Return(0, errors.New("can't connect to the database")).Once()
		result, err := services.CreatePolyclinic(sampleDomainUmumWithNoID)

		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestGetAllPolyclinics(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllPolyclinics()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 1)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectAllData").
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAllPolyclinics()

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetAllPolyclinicsWithStats(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		for _, domain := range sampleDomainList {
			mockRepo.On("CountDoctorByPolyclinic", domain.ID).Return(sampleTotalDoctor, nil).Once()
			mockRepo.On("CountNurseByPolyclinic", domain.ID).Return(sampleTotalNurse, nil).Once()
		}
		result, err := services.GetAllPolyclinicsWithStats()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 1)

	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAllPolyclinicsWithStats()

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got an error while counting doctor", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		mockRepo.On("CountDoctorByPolyclinic", sampleIDUmum).Return(0, errors.New("whoa there is an error")).Once()
		result, err := services.GetAllPolyclinicsWithStats()

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got an error while counting nurse", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		mockRepo.On("CountDoctorByPolyclinic", sampleIDUmum).Return(sampleTotalDoctor, nil).Once()
		mockRepo.On("CountNurseByPolyclinic", sampleIDUmum).Return(0, errors.New("whoa there is an error")).Once()
		result, err := services.GetAllPolyclinicsWithStats()

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetPolyclinicByID(t *testing.T) {
	t.Run("should got data by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleIDUmum).Return(&sampleDomainUmum, nil).Once()
		result, err := services.GetPolyclinicByID(sampleIDUmum)

		assert.Nil(t, err)
		assert.Equal(t, sampleIDUmum, result.ID)
		assert.Equal(t, sampleDomainUmum.Name, result.Name)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleIDUmum).Return(nil, errors.New("record not found"))
		result, err := services.GetPolyclinicByID(sampleIDUmum)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
		assert.Nil(t, result)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleIDUmum).
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetPolyclinicByID(sampleIDUmum)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdatePolycllinicByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleIDUmum, sampleDomainUmumWithNoID).Return(nil).Once()
		err := services.AmendPolyclinicByID(sampleIDUmum, sampleDomainUmumWithNoID)

		assert.Nil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleIDUmum, sampleDomainUmumWithNoID).
			Return(errors.New("record not found")).Once()
		err := services.AmendPolyclinicByID(sampleIDUmum, sampleDomainUmumWithNoID)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleIDUmum, sampleDomainUmumWithNoID).
			Return(errors.New("can't connect to the database")).Once()
		err := services.AmendPolyclinicByID(sampleIDUmum, sampleDomainUmumWithNoID)

		assert.NotNil(t, err)
	})
}

func TestRemovePolyclinicByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleIDUmum).Return(nil).Once()
		err := services.RemovePolyclinicByID(sampleIDUmum)

		assert.Nil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleIDUmum).Return(errors.New("record not found")).Once()
		err := services.RemovePolyclinicByID(sampleIDUmum)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleIDUmum).
			Return(errors.New("can't connect to the database")).Once()
		err := services.RemovePolyclinicByID(sampleIDUmum)

		assert.NotNil(t, err)
	})
}
