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
	sampleNurseUUID  uuid.UUID

	sampleDoctorAccountDomain account.Domain
	sampleAdminAccountDomain  account.Domain
	sampleNurseAccountDomain  account.Domain

	sampleDoctorUserDataDomain account.UserDataDomain
	sampleAdminUserDataDomain  account.UserDataDomain
	sampleNurseUserDataDomain  account.UserDataDomain

	samplePassword string

	sampleNIP string
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)

	sampleDoctorUUID = uuid.Must(uuid.NewRandom())
	sampleAdminUUID = uuid.Must(uuid.NewRandom())
	sampleNurseUUID = uuid.Must(uuid.NewRandom())

	samplePassword = "thestrongestpassword"
	hashedSamplePassword, _ := utils.CreateHash(samplePassword)

	sampleNIP = "2015 031033 32"

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
	sampleNurseAccountDomain = account.Domain{
		ID:       sampleNurseUUID,
		Email:    "nurse.capstone@example.com",
		Password: hashedSamplePassword,
		Role:     types.NURSE,
	}

	sampleDoctorUserDataDomain = account.UserDataDomain{
		ID:   sampleDoctorUUID,
		Name: "dr. Capstone",
		NIP:  sampleNIP,
		Role: types.ADMIN,
	}
	sampleAdminUserDataDomain = account.UserDataDomain{
		ID:   sampleAdminUUID,
		Name: "Admin Capstone",
		NIP:  sampleNIP,
		Role: types.ADMIN,
	}
	sampleNurseUserDataDomain = account.UserDataDomain{
		ID:   sampleNurseUUID,
		Name: "Nurse Capstone",
		NIP:  sampleNIP,
		Role: types.NURSE,
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
		token, role, err := services.AttemptLogin(loginData)

		assert.Nil(t, err)
		assert.NotEqual(t, "", token)
		assert.Equal(t, types.DOCTOR, role)
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
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})

	t.Run("should got wrong passphrase doctor account error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleDoctorAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleDoctorAccountDomain, nil).Once()
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
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
		token, role, err := services.AttemptLogin(loginData)

		assert.Nil(t, err)
		assert.NotEqual(t, "", token)
		assert.Equal(t, types.ADMIN, role)
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
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})

	t.Run("should got wrong passphrase admin account error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleAdminAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleAdminAccountDomain, nil).Once()
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})

	// nurse tests
	t.Run("should got nurse login information", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleNurseAccountDomain.Email,
			Password: samplePassword,
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleNurseAccountDomain, nil).Once()
		mockRepo.On("LookupNurseByUserID", sampleNurseAccountDomain.ID.String()).
			Return(&sampleNurseUserDataDomain, nil).Once()
		token, role, err := services.AttemptLogin(loginData)

		assert.Nil(t, err)
		assert.NotEqual(t, "", token)
		assert.Equal(t, types.NURSE, role)
	})

	t.Run("should got database error while querying nurse data", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleNurseAccountDomain.Email,
			Password: samplePassword,
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleNurseAccountDomain, nil).Once()
		mockRepo.On("LookupNurseByUserID", sampleNurseAccountDomain.ID.String()).
			Return(nil, errors.New("something error with database")).Once()
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})

	t.Run("should got wrong passphrase nurse account error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleNurseAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(&sampleNurseAccountDomain, nil).Once()
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})

	// another errors
	t.Run("should got email unregistered", func(t *testing.T) {
		loginData := account.Domain{
			Email:    "some.email@example.com",
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(nil, errors.New("record not found"))
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})

	t.Run("should got server error", func(t *testing.T) {
		loginData := account.Domain{
			Email:    sampleDoctorAccountDomain.Email,
			Password: "anotherpassword",
		}

		mockRepo.On("LookupAccountByEmail", loginData.Email).
			Return(nil, errors.New("can't connect to the database")).Once()
		token, role, err := services.AttemptLogin(loginData)

		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, types.UserRoleEnum(""), role)
	})
}
