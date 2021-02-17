package models

var (
	// Temp => temporary values for headers
	Temp = GWCustomHeader{"0.01", "5 mb", "1.39", "0.02", "0.03"}
)

// GWCustomHeader to mock responce from ICAP server
type GWCustomHeader struct {
	MetricFormFileread string
	MetricFileSize     string
	Version            string
	MetricDetect       string
	MetricRebuild      string
}
