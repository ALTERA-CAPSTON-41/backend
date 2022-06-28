package icd10_repositories

import "clinic-api/src/app/icd10"

type Data struct {
	Search []ICD10 `json:"Search"`
}

type ICD10 struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func MapToDomain(record ICD10) icd10.Domain {
	return icd10.Domain{
		Name:        record.Name,
		Description: record.Description,
	}
}

func MapToBatchDomain(records []ICD10) (domains []icd10.Domain) {
	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return
}
