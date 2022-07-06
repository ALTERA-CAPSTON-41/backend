package polyclinic

type Domain struct {
	ID          int
	Name        string
	TotalDoctor int
	TotalNurse  int
}

type Services interface {
	CreatePolyclinic(Domain) (int, error)
	GetAllPolyclinics() ([]Domain, error)
	GetPolyclinicByID(id int) (*Domain, error)
	GetAllPolyclinicsWithStats() ([]Domain, error)
	AmendPolyclinicByID(id int, polyclinic Domain) error
	RemovePolyclinicByID(id int) error
}

type Repositories interface {
	InsertData(data Domain) (int, error)
	SelectAllData() ([]Domain, error)
	SelectDataByID(id int) (*Domain, error)
	UpdateByID(id int, data Domain) error
	DeleteByID(id int) error
	CountDoctorByPolyclinic(polyclinic int) (int, error)
	CountNurseByPolyclinic(polyclinic int) (int, error)
}
