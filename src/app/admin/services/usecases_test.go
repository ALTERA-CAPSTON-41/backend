package services

import (
	"clinic-api/src/app/admin"
	"clinic-api/src/app/admin/mocks"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo                       mocks.Repositories
	services                       admin.Services
	sampleDomainList               []admin.Domain
	sampleDomainIkoUwais           admin.Domain
	sampleDomainIkoUwaisWithNoUUID admin.Domain
	sampleUUIDIkoUwais             uuid.UUID
	sampleUUIDTheressaSiahaan      uuid.UUID
	sampleEmailIkoUwais            string
	samplePassword                 string
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUIDIkoUwais = uuid.Must(uuid.NewRandom())
	sampleUUIDTheressaSiahaan = uuid.Must(uuid.NewRandom())
	sampleEmailIkoUwais = "ikouwais@example.com"
	samplePassword = "thestrongestpassword"

	sampleDomainList = []admin.Domain{
		{
			Name: "Iko Uwais",
			NIP:  "2022 083917 36",
			User: admin.UserReference{
				ID:        sampleUUIDIkoUwais,
				Email:     "ikouwais@example.com",
				Password:  samplePassword,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			Name: "Therresa Siahaan",
			NIP:  "2022 738232 43",
			User: admin.UserReference{
				ID:        sampleUUIDTheressaSiahaan,
				Email:     "theressa.siahaan@example.com",
				Password:  samplePassword,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	sampleDomainIkoUwais = admin.Domain{
		Name: "Iko Uwais",
		NIP:  "2022 083917 36",
		User: admin.UserReference{
			ID:        sampleUUIDIkoUwais,
			Email:     "ikouwais@example.com",
			Password:  samplePassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	sampleDomainIkoUwaisWithNoUUID = admin.Domain{
		Name: "Iko Uwais",
		NIP:  "2022 083917 36",
		User: admin.UserReference{
			Email:     "ikouwais@example.com",
			Password:  samplePassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	os.Exit(m.Run())
}

func TestGetAllAdmin(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllAdmins()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should got server error", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAllAdmins()

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetAdminByID(t *testing.T) {
	t.Run("should got data by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).Return(&sampleDomainIkoUwais, nil).Once()
		result, err := services.GetAdminByID(sampleUUIDIkoUwais.String())

		assert.Nil(t, err)
		assert.Equal(t, sampleDomainIkoUwais.Name, result.Name)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAdminByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).Return(nil, errors.New("record not found"))
		result, err := services.GetAdminByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestCreateAdmin(t *testing.T) {
	t.Run("should created a data", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).Return("", nil).Once()
		mockRepo.On("InsertData", sampleDomainIkoUwaisWithNoUUID).
			Return(sampleUUIDIkoUwais.String(), nil).Once()
		result, err := services.CreateAdmin(sampleDomainIkoUwaisWithNoUUID)

		assert.Nil(t, err)
		assert.Equal(t, sampleUUIDIkoUwais.String(), result)
	})

	t.Run("should got already used email error", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).Return(sampleEmailIkoUwais, nil).Once()
		result, err := services.CreateAdmin(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "already used")
		assert.Zero(t, result)
	})

	t.Run("should got an error while lookup data by email", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).Return("", errors.New("can't connect to the database")).Once()
		result, err := services.CreateAdmin(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Zero(t, result)
	})

	t.Run("should got an error while insert data", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).Return("", nil).Once()
		mockRepo.On("InsertData", sampleDomainIkoUwaisWithNoUUID).Return(uuid.Nil.String(), errors.New("can't connect to the database")).Once()
		result, err := services.CreateAdmin(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Equal(t, uuid.Nil.String(), result)
	})
}

func TestUpdateAdminByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).Return(nil).Once()
		err := services.AmendAdminByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).Return(errors.New("can't connect to the database")).Once()
		err := services.AmendAdminByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).Return(errors.New("record not found")).Once()
		err := services.AmendAdminByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
	})
}

func TestRemoveAdminByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		mockRepo.On("DeleteUserByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		err := services.RemoveAdminByID(sampleUUIDIkoUwais.String())

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(errors.New("can't connect to the database")).Once()
		mockRepo.On("DeleteUserByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		err := services.RemoveAdminByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(errors.New("record not found")).Once()
		mockRepo.On("DeleteUserByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		err := services.RemoveAdminByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
	})
}
