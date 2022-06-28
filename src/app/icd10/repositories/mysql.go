package icd10_repositories

import (
	"clinic-api/src/app/icd10"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// SearchDataByCode implements icd10.Repositories
func (repo *repository) SearchDataByCode(code string) ([]icd10.Domain, error) {
	endpoint := fmt.Sprintf("http://icd10api.com/?s=%s&desc=short&r=json", code)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var body Data
	json.Unmarshal(bodyBytes, &body)

	return MapToBatchDomain(body.Search), nil
}

func NewMySQLRepository(DB *gorm.DB) icd10.Repositories {
	return &repository{DB}
}
