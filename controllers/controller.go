package controllers

import (
	"encoding/json"
	"k8-go-api/models"
	"log"
)

func parsecontentManagementFlagJson(c []byte) (models.ContentManagementFlagContent, error) {

	var d models.ContentManagementFlagContent
	err := json.Unmarshal(c, &d)
	if err != nil {
		log.Println("unlarshall", err)
		return d, err

	}
	return d, nil
}
