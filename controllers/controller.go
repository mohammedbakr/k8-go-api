package controllers

import (
	"encoding/json"
	"log"

	"github.com/k8-proxy/k8-go-api/models"
)

func parseContentManagementFlagJSON(c []byte) (models.ContentManagementFlags, error) {

	var d models.ContentManagementFlags
	err := json.Unmarshal(c, &d)
	if err != nil {
		log.Println("unmarshall", err)
		return d, err

	}
	return d, nil
}
