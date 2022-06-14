package services

import (
	"clinic-api/src/app/queue"
	"clinic-api/src/app/queue/mocks"
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
	mockRepo            mocks.Repositories
	services            queue.Services
	sampleDomainList    []queue.Domain
	sampleDomainEko     queue.Domain
	sampleRequestDomain queue.Domain
	sampleIDEko         uuid.UUID
	sampleIDKoko        uuid.UUID
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleIDEko = uuid.Must(uuid.NewRandom())
	sampleIDKoko = uuid.Must(uuid.NewRandom())

	sampleDomainEko = queue.Domain{
		ID: sampleIDEko,
		Patient: queue.PatientReference{
			ID:     sampleIDEko,
			Name:   "Eko",
			Gender: types.MALE,
		},
		Polyclinic: queue.PolyclinicReference{
			ID: 1, Name: "UMUM",
		},
		PatientStatus:    types.OUTPATIENT,
		DailyQueueNumber: 1,
		DailyQueueDate:   utils.ConvertStringToDate("2022-01-11"),
		ServiceDoneAt:    time.Now(),
	}

	sampleDomainList = []queue.Domain{
		sampleDomainEko,
		{
			ID: sampleIDEko,
			Patient: queue.PatientReference{
				ID:     sampleIDKoko,
				Name:   "Koko",
				Gender: types.MALE,
			},
			Polyclinic: queue.PolyclinicReference{
				ID: 1, Name: "UMUM",
			},
			PatientStatus:    types.OUTPATIENT,
			DailyQueueNumber: 2,
			DailyQueueDate:   utils.ConvertStringToDate("2022-01-11"),
			ServiceDoneAt:    time.Now(),
		},
	}

	sampleRequestDomain = queue.Domain{
		PatientID:      sampleIDEko,
		PolyclinicID:   1,
		PatientStatus:  types.OUTPATIENT,
		DailyQueueDate: utils.ConvertStringToDate("2022-01-11"),
		ServiceDoneAt:  time.Now(),
	}

	os.Exit(m.Run())
}

func TestCreateQueue(t *testing.T) {
	t.Run("should created a data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleRequestDomain).Return(sampleIDEko.String(), nil).Once()
		result, err := services.CreateQueue(sampleRequestDomain)

		assert.Nil(t, err)
		assert.Equal(t, sampleIDEko.String(), result)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("InsertData", sampleRequestDomain).
			Return(uuid.Nil.String(), errors.New("can't connect to the database")).Once()
		result, err := services.CreateQueue(sampleRequestDomain)

		assert.NotNil(t, err)
		assert.Equal(t, uuid.Nil.String(), result)
	})
}

func TestGetAllQueues(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData", "", "").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllQueues("", "")

		assert.Nil(t, err)
		assert.Greater(t, len(result), 1)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectAllData", "", "").
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAllQueues("", "")

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdateQueueByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleIDEko.String(), sampleRequestDomain).Return(nil).Once()
		err := services.AmendQueueByID(sampleIDEko.String(), sampleRequestDomain)

		assert.Nil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleIDEko.String(), sampleRequestDomain).
			Return(errors.New("record not found")).Once()
		err := services.AmendQueueByID(sampleIDEko.String(), sampleRequestDomain)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleIDEko.String(), sampleRequestDomain).
			Return(errors.New("can't connect to the database")).Once()
		err := services.AmendQueueByID(sampleIDEko.String(), sampleRequestDomain)

		assert.NotNil(t, err)
	})
}

func TestRemoveQueueByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleIDEko.String()).Return(nil).Once()
		err := services.RemoveQueueByID(sampleIDEko.String())

		assert.Nil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleIDEko.String()).
			Return(errors.New("record not found")).Once()
		err := services.RemoveQueueByID(sampleIDEko.String())

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleIDEko.String()).
			Return(errors.New("can't connect to the database")).Once()
		err := services.RemoveQueueByID(sampleIDEko.String())

		assert.NotNil(t, err)
	})
}
