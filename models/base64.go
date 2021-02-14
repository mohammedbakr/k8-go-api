package models

// Base64 struct
type Base64 struct {
	Request struct {
		FileName               string `json:"FileName"`
		Base64                 string `json:"Base64"`
		ContentManagementFlags string `json:"ContentManagementFlags"`
	} `json:"request"`
}
