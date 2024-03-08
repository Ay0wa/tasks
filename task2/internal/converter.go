package internal

import (
	"encoding/json"
	"fmt"
	"task2/internal/models"

	"gopkg.in/yaml.v3"
)

func ConvertJSONtoYAML(jsonFile []byte) ([]byte, error) {
	var emp models.Employee
	err := json.Unmarshal(jsonFile, &emp)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	err = emp.Validate()
	if err != nil {
		return nil, err
	}

	yamlData, err := yaml.Marshal(&emp)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}
