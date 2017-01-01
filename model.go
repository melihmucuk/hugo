package hugo

import (
	"strings"
	"time"
)

type HurriyetTime struct {
	time.Time
}

func (self *HurriyetTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		t, err = time.Parse(time.RFC3339, s)
	}
	self.Time = t
	return
}

type Query struct {
	Filter []string
	Select []string
	Top    string
	Skip   string
	S      string
}

type Article struct {
	ID          string        `json:"Id"`
	ContentType string        `json:"ContentType"`
	CreatedDate HurriyetTime  `json:"CreatedDate"`
	Description string        `json:"Description"`
	Editor      string        `json:"Editor"`
	Files       []File        `json:"Files"`
	Path        string        `json:"Path"`
	RelatedNews []interface{} `json:"RelatedNews"`
	StartDate   HurriyetTime  `json:"StartDate"`
	Tags        []interface{} `json:"Tags"`
	Text        string        `json:"Text"`
	Title       string        `json:"Title"`
	URL         string        `json:"Url"`
	Writers     []Writer      `json:"Writers"`
}

type Metadata struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type File struct {
	FileURL  string   `json:"FileUrl"`
	Metadata Metadata `json:"Metadata"`
}

type Writer struct {
	ID          string       `json:"Id"`
	Fullname    string       `json:"Fullname"`
	ContentType string       `json:"ContentType"`
	CreatedDate HurriyetTime `json:"CreatedDate"`
	Files       []File       `json:"Files"`
	Path        string       `json:"Path"`
	URL         string       `json:"Url"`
}

type Search struct {
	Count   int64     `json:"Count"`
	Results []Article `json:"List"`
}

type Path struct {
	ID    string `json:"Id"`
	Path  string `json:"Path"`
	Title string `json:"Title"`
}
