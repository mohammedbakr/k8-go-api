package models

// Base64 model
type Base64 struct {
	Request struct {
		FileName               string                 `json:"FileName"`
		Base64                 string                 `json:"Base64"`
		ContentManagementFlags ContentManagementFlags `json:"ContentManagementFlags"`
	} `json:"request"`
}

type PdfContentManagement struct {
	Metadata           int `json:"Metadata"`
	InternalHyperlinks int `json:"InternalHyperlinks"`
	ExternalHyperlinks int `json:"ExternalHyperlinks"`
	EmbeddedFiles      int `json:"EmbeddedFiles"`
	EmbeddedImages     int `json:"EmbeddedImages"`
	Javascript         int `json:"Javascript"`
	Acroform           int `json:"Acroform"`
	ActionsAll         int `json:"ActionsAll"`
}
type OfficeContentManagement struct {
	Metadata           int `json:"Metadata"`
	InternalHyperlinks int `json:"InternalHyperlinks"`
	ExternalHyperlinks int `json:"ExternalHyperlinks"`
	EmbeddedFiles      int `json:"EmbeddedFiles"`
	EmbeddedImages     int `json:"EmbeddedImages"`
	Macros             int `json:"Macros"`
	ReviewComments     int `json:"ReviewComments"`
}

type ContentManagementFlags struct {
	PdfContentManagement        PdfContentManagement        `json:"PdfContentManagement"`
	ExcelContentManagement      OfficeContentManagement     `json:"ExcelContentManagement"`
	PowerPointContentManagement PowerPointContentManagement `json:"PowerPointContentManagement"`
	WordContentManagement       OfficeContentManagement     `json:"WordContentManagement"`
}

type PowerPointContentManagement struct {
	Metadata           int `json:"Metadata"`
	InternalHyperlinks int `json:"InternalHyperlinks"`
	ExternalHyperlinks int `json:"ExternalHyperlinks"`
	EmbeddedFiles      int `json:"EmbeddedFiles"`
	EmbeddedImages     int `json:"EmbeddedImages"`
	Macros             int `json:"Macros"`
	ReviewComments     int `json:"ReviewComments"`
}
