package controllers

import (
	"testing"
)

const (
	contentManagementFlagJSON = `{
		"PdfContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "Javascript": 1,
		  "Acroform": 1,
		  "ActionsAll": 1
		},
		"ExcelContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "DynamicDataExchange": 1,
		  "Macros": 1,
		  "ReviewComments": 1
		},
		"PowerPointContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "Macros": 1,
		  "ReviewComments": 1
		},
		"WordContentManagement": {
		  "Metadata": 1,
		  "InternalHyperlinks": 1,
		  "ExternalHyperlinks": 1,
		  "EmbeddedFiles": 1,
		  "EmbeddedImages": 1,
		  "DynamicDataExchange": 1,
		  "Macros": 1,
		  "ReviewComments": 1
		}
	}`
)

const (
	SampleDataPath = "../sampledata/"
	PdfFileName    = "file.pdf"
	ZipFileName    = "file.zip"
)

func TestParsecontentManagementFlagJSON(t *testing.T) {
	var tests = []struct {
		content string
		err     bool
	}{
		{contentManagementFlagJSON, false},
		{"", true},
	}

	for _, test := range tests {
		if _, err := parseContentManagementFlagJSON(test.content); checkError(err) != test.err {
			t.Errorf("Test Failed: Error {%s} ", err)
		}
	}
}

func checkError(i error) bool {
	if i != nil {
		return true
	}
	return false
}
