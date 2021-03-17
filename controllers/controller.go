package controllers

import (
	"encoding/json"

	"github.com/k8-proxy/k8-go-api/models"
)

func parseContentManagementFlagJSON(c string) (models.ContentManagementFlags, error) {
	var d models.ContentManagementFlags
	err := json.Unmarshal([]byte(c), &d)
	if err != nil {
		return d, err
	}

	return d, nil
}
