package icd10_repositories

import (
	"clinic-api/src/app/icd10"
	"encoding/json"
	"fmt"
	"net/http"
)

type repository struct{}

// SearchDataByCode implements icd10.Repositories
func (repo *repository) SearchDataByCode(code string) ([]icd10.Domain, error) {
	endpoint := fmt.Sprintf("http://icd10api.com/?s=%s&desc=short&r=json", code)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	var body Data
	json.NewDecoder(resp.Body).Decode(&body)

	return MapToBatchDomain(body.Search), nil
}

func NewMySQLRepository() icd10.Repositories {
	return &repository{}
}
