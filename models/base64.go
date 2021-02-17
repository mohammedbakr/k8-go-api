package models

// Base64 model
type Base64 struct {
	Request struct {
		FileName               string                       `json:"FileName"`
		Base64                 string                       `json:"Base64"`
		ContentManagementFlags ContentManagementFlagContent `json:"ContentManagementFlags"`
	} `json:"request"`
}

// JSONMeta to mock JSON meta header
type JSONMeta struct {
	Metadata           int `json:"Metadata"`
	InternalHyperlinks int `json:"InternalHyperlinks"`
	ExternalHyperlinks int `json:"ExternalHyperlinks"`
	EmbeddedFiles      int `json:"EmbeddedFiles"`
	EmbeddedImages     int `json:"EmbeddedImages"`
	Javascript         int `json:"Javascript"`
	Acroform           int `json:"Acroform"`
	ActionsAll         int `json:"ActionsAll"`
}

// ContentManagementFlagContent to mock header
type ContentManagementFlagContent struct {
	PdfContentManagement        JSONMeta `json:"PdfContentManagement"`
	ExcelContentManagement      JSONMeta `json:"ExcelContentManagement"`
	PowerPointContentManagement JSONMeta `json:"PowerPointContentManagement"`
	WordContentManagement       JSONMeta `json:"WordContentManagement"`
}
