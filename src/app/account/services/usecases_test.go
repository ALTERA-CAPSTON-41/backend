package services

import (
	"clinic-api/src/app/account"
	"clinic-api/src/app/account/mocks"
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
	services account.Services

	sampleDoctorUUID uuid.UUID
	sampleAdminUUID  uuid.UUID

	sampleDoctorAccountDomain account.Domain
	sampleAdminAccountDomain  account.Domain

	sampleDoctorUserDataDomain account.UserDataDomain
	sampleAdminUserDataDomain  account.UserDataDomain

	samplePassword string
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)

	sampleDoctorUUID = uuid.Must(uuid.NewRandom())
	sampleAdminUUID = uuid.Must(uuid.NewRandom())

	samplePassword = "thestrongestpassword"
	hashedSamplePassword, _ := utils.CreateHash(samplePassword)

	sampleDoctorAccountDomain = account.Domain{
		ID:       sampleDoctorUUID,
		Email:    "dr.capstone@example.com",
		Password: hashedSamplePassword,
		Role:     types.DOCTOR,
	}
	sampleAdminAccountDomain = account.Domain{
		ID:       sampleAdminUUID,
		Email:    "admin.capstone@example.com",
		Password: hashedSamplePassword,
		Role:     types.ADMIN,
	}

	sampleDoctorUserDataDomain = account.UserDataDomain{
		ID:   sampleDoctorUUID,
		Name: "dr. Capstone",
		NIP:  "2015 031033 32",
		Role: types.ADMIN,
	}
	sampleAdminUserDataDomain = account.UserDataDomain{
		ID:   sampleAdminUUID,
		Name: "Admin Capstone",
		NIP:  "2015 031033 32",
		Role: types.ADMIN,
	}

	os.Exit(m.Run())
}

func TestAttemptLogin(t *testing.T) {
	// doctor tests
	t.Run("should got doctor login information", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleDoctorAccountDomain.Email,
			Password: samplePassword,
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleDoctorAccountDomain, nil).Once()
		mockRepo.On("LookupDoctorByUserID", sampleDoctorAccountDomain.ID.String()).
			Return(&sampleDoctorUserDataDomain, nil).Once()
		result, err := services.AttemptLogin(loginData)

		assert.Nil(t, err)
		assert.Equal(t, sampleDoctorUserDataDomain.Name, result.Name)
		assert.NotEqual(t, uuid.Nil, result.ID)
	})

	t.Run("should got database error while querying doctor data", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleDoctorAccountDomain.Email,
			Password: samplePassword,
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleDoctorAccountDomain, nil).Once()
		mockRepo.On("LookupDoctorByUserID", sampleDoctorAccountDomain.ID.String()).
			Return(nil, errors.New("something error with database")).Once()
		result, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got wrong passphrase doctor account error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleDoctorAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleDoctorAccountDomain, nil).Once()
		result, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	// admin tests
	t.Run("should got admin login information", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleAdminAccountDomain.Email,
			Password: samplePassword,
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleAdminAccountDomain, nil).Once()
		mockRepo.On("LookupAdminByUserID", sampleAdminAccountDomain.ID.String()).
			Return(&sampleAdminUserDataDomain, nil).Once()
		result, err := services.AttemptLogin(loginData)

		assert.Nil(t, err)
		assert.Equal(t, sampleAdminUserDataDomain.Name, result.Name)
		assert.NotEqual(t, uuid.Nil, result.ID)
	})

	t.Run("should got database error while querying admin data", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleAdminAccountDomain.Email,
			Password: samplePassword,
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleAdminAccountDomain, nil).Once()
		mockRepo.On("LookupAdminByUserID", sampleAdminAccountDomain.ID.String()).
			Return(nil, errors.New("something error with database")).Once()
		result, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got wrong passphrase admin account error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleAdminAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleAdminAccountDomain, nil).Once()
		result, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	// another errors
	t.Run("should got email unregistered", func(t *testing.T) {
		loginData := account.Domain{
			Email:    "some.email@example.com",
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(nil, errors.New("record not found"))
		result, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got server error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleDoctorAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
