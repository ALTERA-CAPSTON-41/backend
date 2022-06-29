package icd10

type Domain struct {
	Name        string
	Description string
}

type Services interface {
	FindICD10ByCode(code string) ([]Domain, error)
}

type Repositories interface {
	SearchDataByCode(code string) ([]Domain, error)
}
