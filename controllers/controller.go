package controllers

import (
	"encoding/json"
	"log"

	"github.com/mohammedbakr/k8-go-api/models"
)

func parsecontentManagementFlagJSON(c []byte) (models.ContentManagementFlagContent, error) {

	var d models.ContentManagementFlagContent
	err := json.Unmarshal(c, &d)
	if err != nil {
		log.Println("unmarshall", err)
		return d, err

	}
	return d, nil
}
