package hugo

import (
	"time"
)

type Article struct {
	ID          string        `json:"Id"`
	ContentType string        `json:"ContentType"`
	CreatedDate time.Time     `json:"CreatedDate"`
	Description string        `json:"Description"`
	Editor      string        `json:"Editor"`
	Files       []File        `json:"Files"`
	Path        string        `json:"Path"`
	RelatedNews []interface{} `json:"RelatedNews"`
	StartDate   time.Time     `json:"StartDate"`
	Tags        []interface{} `json:"Tags"`
	Text        string        `json:"Text"`
	Title       string        `json:"Title"`
	URL         string        `json:"Url"`
	Writers     []interface{} `json:"Writers"`
}

type Metadata struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type File struct {
	FileURL  string   `json:"FileUrl"`
	Metadata Metadata `json:"Metadata"`
}
