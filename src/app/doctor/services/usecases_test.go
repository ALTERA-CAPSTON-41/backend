package services

import (
	"clinic-api/src/app/doctor"
	"clinic-api/src/app/doctor/mocks"
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
	mockRepo                       mocks.Repositories
	services                       doctor.Services
	sampleDomainList               []doctor.Domain
	sampleDomainIkoUwais           doctor.Domain
	sampleDomainIkoUwaisWithNoUUID doctor.Domain
	sampleUUIDIkoUwais             uuid.UUID
	sampleUUIDTheressaSiahaan      uuid.UUID
	sampleEmailIkoUwais            string
	samplePassword                 string
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUIDIkoUwais = uuid.Must(uuid.NewRandom())
	sampleUUIDTheressaSiahaan = uuid.Must(uuid.NewRandom())
	sampleEmailIkoUwais = "dr.ikouwais@example.com"
	samplePassword = "thestrongestpassword"

	sampleDomainIkoUwais = doctor.Domain{
		Name:    "dr. Iko Uwais",
		NIP:     "2022 089281 32",
		SIP:     "SIP.8198.31.2022",
		Address: "North Way 61 CA, Antartica",
		DOB:     utils.ConvertStringToDate("1997-08-19"),
		Gender:  types.MALE,
		Polyclinic: doctor.PolyclinicReference{
			ID:   1,
			Name: "General",
		},
		User: doctor.UserReference{
			ID:        sampleUUIDIkoUwais,
			Email:     "dr.ikouwais@example.com",
			Password:  samplePassword,
			Role:      types.DOCTOR,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	sampleDomainIkoUwaisWithNoUUID = doctor.Domain{
		Name:    "dr. Iko Uwais",
		NIP:     "2022 089281 32",
		SIP:     "SIP.8198.31.2022",
		Address: "North Way 61 CA, Antartica",
		DOB:     utils.ConvertStringToDate("1997-08-19"),
		Gender:  types.MALE,
		Polyclinic: doctor.PolyclinicReference{
			ID:   1,
			Name: "General",
		},
		User: doctor.UserReference{
			ID:        uuid.Nil,
			Email:     "dr.ikouwais@example.com",
			Password:  samplePassword,
			Role:      types.DOCTOR,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	sampleDomainList = []doctor.Domain{
		sampleDomainIkoUwais,
		{
			Name:    "dr. Theressa Siahaan",
			NIP:     "2022 089284 84",
			SIP:     "SIP.8908.74.2022",
			Address: "South Way 61 FR, Arctic",
			DOB:     utils.ConvertStringToDate("1998-01-18"),
			Gender:  types.FEMALE,
			Polyclinic: doctor.PolyclinicReference{
				ID:   1,
				Name: "General",
			},
			User: doctor.UserReference{
				ID:        sampleUUIDTheressaSiahaan,
				Email:     "dr.theressasiahaan@example.com",
				Password:  samplePassword,
				Role:      types.DOCTOR,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	os.Exit(m.Run())
}

func TestGetAllDoctors(t *testing.T) {
	t.Run("should got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData", 0, 0).Return(sampleDomainList, nil).Once()
		result, err := services.GetAllDoctors(0, 1)

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should got server error", func(t *testing.T) {
		mockRepo.On("SelectAllData", 0, 0).
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetAllDoctors(0, 1)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestGetDoctorByID(t *testing.T) {
	t.Run("should got data by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).
			Return(&sampleDomainIkoUwais, nil).Once()
		result, err := services.GetDoctorByID(sampleUUIDIkoUwais.String())

		assert.Nil(t, err)
		assert.Equal(t, sampleDomainIkoUwais.Name, result.Name)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).
			Return(nil, errors.New("can't connect to the database")).Once()
		result, err := services.GetDoctorByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUIDIkoUwais.String()).
			Return(nil, errors.New("record not found")).Once()
		result, err := services.GetDoctorByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestCreateDoctor(t *testing.T) {
	t.Run("should created a data", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).
			Return("", nil).Once()
		mockRepo.On("InsertData", sampleDomainIkoUwaisWithNoUUID).
			Return(sampleUUIDIkoUwais.String(), nil).Once()
		result, err := services.CreateDoctor(sampleDomainIkoUwaisWithNoUUID)

		assert.Nil(t, err)
		assert.Equal(t, sampleUUIDIkoUwais.String(), result)
	})

	t.Run("should got already used email error", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).
			Return(sampleEmailIkoUwais, nil).Once()
		result, err := services.CreateDoctor(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "already used")
		assert.Zero(t, result)
	})

	t.Run("should got database error while lookup data by email", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).
			Return("", errors.New("can't connect to the database")).Once()
		result, err := services.CreateDoctor(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Zero(t, result)
	})

	t.Run("should got an error while insert data", func(t *testing.T) {
		mockRepo.On("LookupDataByEmail", sampleEmailIkoUwais).
			Return("", nil).Once()
		mockRepo.On("InsertData", sampleDomainIkoUwaisWithNoUUID).
			Return(uuid.Nil.String(), errors.New("whoa there's some error")).Once()
		result, err := services.CreateDoctor(sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
		assert.Equal(t, uuid.Nil.String(), result)
	})
}

func TestUpdateDoctorByID(t *testing.T) {
	t.Run("should update data by id", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).
			Return(nil).Once()
		err := services.AmendDoctorByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).
			Return(errors.New("can't connect to the database")).Once()
		err := services.AmendDoctorByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID).
			Return(errors.New("record not found")).Once()
		err := services.AmendDoctorByID(sampleUUIDIkoUwais.String(), sampleDomainIkoUwaisWithNoUUID)

		assert.NotNil(t, err)
	})
}

func TestDeleteDoctorByID(t *testing.T) {
	t.Run("should delete data by id", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		mockRepo.On("DeleteUserByID", sampleUUIDIkoUwais.String()).Return(nil).Once()
		err := services.RemoveDoctorByID(sampleUUIDIkoUwais.String())

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).
			Return(errors.New("can't connect to the database")).Once()
		err := services.RemoveDoctorByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleUUIDIkoUwais.String()).
			Return(errors.New("record not found")).Once()
		err := services.RemoveDoctorByID(sampleUUIDIkoUwais.String())

		assert.NotNil(t, err)
	})
}
