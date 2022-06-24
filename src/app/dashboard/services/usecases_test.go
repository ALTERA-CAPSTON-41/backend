package services

import (
	"clinic-api/src/app/dashboard"
	"clinic-api/src/app/dashboard/mocks"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockRepo           mocks.Repositories
	services           dashboard.Services
	sampleInput        string
	sampleInputQueue   string
	sampleInvalidInput string
	sampleOutput       int
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleInput = "patient"
	sampleInputQueue = "queue"
	sampleInvalidInput = "something"
	sampleOutput = 10
	os.Exit(m.Run())
}

func getTableName(input string) string {
	return fmt.Sprintf("%ss", input)
}

func TestGetTotal(t *testing.T) {
	t.Run("should got total data", func(t *testing.T) {
		mockRepo.On("CountData", getTableName(sampleInput)).Return(sampleOutput, nil).Once()
		result, err := services.GetTotal(sampleInput)

		assert.Nil(t, err)
		assert.Equal(t, sampleOutput, result)
	})

	t.Run("should got total queue data", func(t *testing.T) {
		mockRepo.On("CountQueueData").Return(sampleOutput, nil).Once()
		result, err := services.GetTotal(sampleInputQueue)

		assert.Nil(t, err)
		assert.Equal(t, sampleOutput, result)
	})

	t.Run("should got an error while count data", func(t *testing.T) {
		mockRepo.On("CountData", getTableName(sampleInput)).Return(0, errors.New("error")).Once()
		result, err := services.GetTotal(sampleInput)

		assert.NotNil(t, err)
		assert.Zero(t, result)
	})

	t.Run("should got an error while count queue data", func(t *testing.T) {
		mockRepo.On("CountQueueData").Return(0, errors.New("error")).Once()
		result, err := services.GetTotal(sampleInputQueue)

		assert.NotNil(t, err)
		assert.Zero(t, result)
	})

	t.Run("should got an error while accessing unallowed table", func(t *testing.T) {
		result, err := services.GetTotal(sampleInvalidInput)

		assert.NotNil(t, err)
		assert.Zero(t, result)
	})
}
