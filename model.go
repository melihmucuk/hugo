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

// Query is a struct that contains query properties for endpoints.
type Query struct {

	// Filter oData filtreleme yapısını kullanır. örnek: "Path eq '/teknoloji/'"
	Filter string

	// Select istenilen kolonları sonuç setine dahil eder. örnek: ["Id", "Description"]
	Select []string

	// Top sonuç setini limitlemek için kullanılır. örnek: 5 (sadece 5 adet sonuç döner)
	Top int

	// Skip sonuç setini limitlemek için kullanılır. örnek: 5 (5 adet sonucu es geçer)
	Skip int

	// S yeniden eskiye ya da eskiden yeniye sıralama için kullanılır. örnek: "-1" (yeniden eskiye)
	S string
}

type Article struct {

	// ID Haber id'sini temsil eder.
	ID string `json:"Id"`

	// ContentType Haber içerik tipini temsil eder. Olası değerler: Article, Column, NewsPhotoGallery, Page, Folder
	ContentType string `json:"ContentType"`

	//CreatedDate Haberin oluşturulma tarihini temsil eder.
	CreatedDate HurriyetTime `json:"CreatedDate"`

	//Description Haberin spot metnini temsil eder.
	Description string `json:"Description"`

	// Editor Haberin editörünü temsil eder.
	Editor string `json:"Editor"`

	// Files Haberin görsellerini temsil eder.
	Files []File `json:"Files"`

	// Path Haberin bağlı bulunduğu dizini temsil eder. /{dizin}/ şeklindeki ifade biçimi, ağaç yapı olarak devam etmektedir. Örneğin; /spor/futbol/
	Path string `json:"Path"`

	// RelatedNews Haberle ilişkilendirilmiş haberleri temsil eder.
	RelatedNews []Article    `json:"RelatedNews"`
	StartDate   HurriyetTime `json:"StartDate"`

	// Tags Haberle ilişkilendirilmiş etiketleri temsil eder.
	Tags []string `json:"Tags"`

	// Text Haberin orijinal metnini temsil eder.
	Text string `json:"Text"`

	// Title Haberin başlığını temsil eder.
	Title string `json:"Title"`

	// URL Haberin gerçek adresini temsil eder.
	URL string `json:"Url"`

	// Haber eğer köşe yazısı tipinde ise köşe yazarının bilgisini temsil eder.
	Writers []Writer `json:"Writers"`
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

	// ID Yazar id'sini temsil eder.
	ID string `json:"Id"`

	// Fullname Yazarın tam adını temsil eder.
	Fullname string `json:"Fullname"`

	// ContentType Yazar içerik tipini temsil eder. Sabit değer: PersonContainer
	ContentType string `json:"ContentType"`

	// CreatedDate Yazarın sistemde oluşturulma tarihini temsil eder.
	CreatedDate HurriyetTime `json:"CreatedDate"`

	// Files Yazarın görsellerini temsil eder.
	Files []File `json:"Files"`

	// Path Yazarın sistemde bağlı olduğu dizini temsil eder.
	Path string `json:"Path"`

	// URL Yazarın gerçek adresini temsil eder.
	URL string `json:"Url"`
}

type Search struct {
	// Count arama sonucunda bulunan eleman sayısı
	Count   int64     `json:"Count"`
	Results []Article `json:"List"`
}

type Path struct {

	// ID Dizin id'sini temsil eder.
	ID string `json:"Id"`

	// Path Dizini temsil eder.
	Path string `json:"Path"`

	// Title Dizin başlığını temsil eder.
	Title string `json:"Title"`
}

type Page struct {

	// ID Sayfa id'sini temsil eder.
	ID string `json:"Id"`

	// CreatedDate Sayfanın oluşturulma tarihini temsil eder.
	CreatedDate HurriyetTime `json:"CreatedDate"`

	RelatedNews []Article `json:"RelatedNews"`

	// PageNews Sayfaya atanmış haberleri temsil eder.
	PageNews []Article `json:"PageNews"`

	// Title Sayfa başlığını temsil eder.
	Title string `json:"Title"`

	// URL Sayfanın gerçek adresini temsil eder.
	URL string `json:"Url"`
}

type Column struct {

	// ID Köşe yazısının id'sini temsil eder.
	ID string `json:"Id"`

	// Fullname Köşe yazarının tam adını temsil eder.
	Fullname string `json:"Fullname"`

	// ContentType Köşe yazısı içerik tipini temsil eder. Sabit değer: Column
	ContentType string `json:"ContentType"`

	// CreatedDate Köşe yazısının oluşturulma tarihini temsil eder.
	CreatedDate HurriyetTime `json:"CreatedDate"`

	// Description Köşe yazısının spot metnini temsil eder.
	Description string `json:"Description"`

	// Files Köşe yazısının görsellerini temsil eder.
	Files []File `json:"Files"`

	// Path Köşe yazısının bağlı bulunduğu dizini temsil eder. /{dizin}/ şeklindeki ifade biçimi, ağaç yapı olarak devam etmektedir. Örneğin; /spor/futbol/
	Path      string       `json:"Path"`
	StartDate HurriyetTime `json:"StartDate"`

	// Title Köşe yazısının başlığını temsil eder.
	Title string `json:"Title"`

	// URL Köşe yazısının gerçek adresini temsil eder.
	URL string `json:"Url"`

	// WriterID Köşe yazarının id'sini temsil eder. Bu id üzerinden tekil sorgu alınabilir.
	WriterID string `json:"WriterId"`
}
