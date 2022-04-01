package mytype

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/luebken/api-logistics/pkg/model"
)

// APILogisticsClient Some comment for the linter
type APILogisticsClient struct {
}

// Get does something for the linter
func (c *APILogisticsClient) Get(id uint64) model.Vessel {

	url := fmt.Sprintf("http://localhost:8090/vessels?id=%d", 1001)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	vessel := model.Vessel{}
	jsonErr := json.Unmarshal(body, &vessel)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vessel
}
