package types

const ()

type MinioFile struct {
	EncryptedURL   string         `json:"encryptedURL"`
	AdaptationFile AdaptationFile `json:"adaptationFile"`
}

type AdaptationFile struct {
	FileID              string `json:"file-id"`
	SourceFileLocation  string `json:"source-file-location"`
	RebuiltFileLocation string `json:"rebuilt-file-location"`
	GenerateReport      string `json:"generate-report"`
}
