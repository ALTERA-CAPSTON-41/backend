package polyclinic

type Domain struct {
	ID   int
	Name string
}

type Services interface {
	CreatePolyclinic(Domain) (int, error)
	GetAllPolyclinics() ([]Domain, error)
	GetPolyclinicByID(id int) (*Domain, error)
	AmendPolyclinicByID(id string, polyclinic Domain) error
	RemovePolyclinicByID(id int) error
}

type Repositories interface {
	InsertData(data Domain) (int, error)
	SelectAllData() ([]Domain, error)
	SelectDataByID(id int) (*Domain, error)
	UpdateByID(id int, data Domain) error
	DeleteByID(id int) error
}
