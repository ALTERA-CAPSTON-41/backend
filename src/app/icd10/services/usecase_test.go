package services

import (
	"clinic-api/src/app/icd10"
	"clinic-api/src/app/icd10/mocks"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockRepo         mocks.Repositories
	services         icd10.Services
	sampleDomain     icd10.Domain
	sampleCode       string
	sampleDomainList []icd10.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleCode = "A7"
	sampleDomain = icd10.Domain{
		Name:        "A750",
		Description: "Epidemic louse-borne typhus fever due to Rickettsia prowazekii",
	}
	sampleDomainList = []icd10.Domain{
		sampleDomain,
		{
			Name:        "A70",
			Description: "Chlamydia psittaci infections",
		},
	}

	os.Exit(m.Run())
}

func TestSearchICD10ByCode(t *testing.T) {
	t.Run("should found some records", func(t *testing.T) {
		mockRepo.On("SearchDataByCode", sampleCode).Return(sampleDomainList, nil).Once()
		result, err := services.FindICD10ByCode(sampleCode)

		assert.Nil(t, err)
		assert.Greater(t, len(result), 1)
	})

	t.Run("should get database error", func(t *testing.T) {
		mockRepo.On("SearchDataByCode", sampleCode).Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.FindICD10ByCode(sampleCode)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
