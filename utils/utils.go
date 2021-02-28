package utils

import (
	"encoding/json"
	"net/http"

	"github.com/k8-proxy/k8-go-api/models"
)

const (
	GwMetricFormFileRead = "gw-metric-formfileread"
	GwMetricFileSize     = "gw-metric-filesize"
	GwVersion            = "gw-version"
	GwMetricDetect       = "gw-metric-detect"
	GwMetricRebuild      = "gw-metric-rebuild"
)

// ResponseWithError to handle errors with JSON
func ResponseWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-type", "application/json")
	var error models.Error
	error.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

// AddGWHeader to add headers in repsponse
func AddGWHeader(w http.ResponseWriter, v models.GWCustomHeader) {
	w.Header().Set(GwMetricFormFileRead, v.MetricFormFileread)
	w.Header().Set(GwMetricFileSize, v.MetricFileSize)
	w.Header().Set(GwVersion, v.Version)
	w.Header().Set(GwMetricDetect, v.MetricDetect)
	w.Header().Set(GwMetricRebuild, v.MetricRebuild)

}
