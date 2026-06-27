package models

// GoogleBooksResponse represents the top-level response from the Google Books API
type GoogleBooksResponse struct {
	Kind       string           `json:"kind"`
	TotalItems int              `json:"totalItems"`
	Items      []GoogleBookInfo `json:"items"`
}

// GoogleBook represents a single book item
type GoogleBookInfo struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo AccessInfo `json:"accessInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
}

type VolumeInfo struct {
	Title               string               `json:"title"`
	Authors             []string             `json:"authors"`
	Publisher           string               `json:"publisher"`
	PublishedDate       string               `json:"publishedDate"`
	Description         string               `json:"description"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
	ReadingModes        ReadingModes         `json:"readingModes"`
	PageCount           int                  `json:"pageCount"`
	PrintType           string               `json:"printType"`
	Categories          []string             `json:"categories"`
	MaturityRating      string               `json:"maturityRating"`
	AllowAnonLogging    bool                 `json:"allowAnonLogging"`
	ContentVersion      string               `json:"contentVersion"`
	ImageLinks          ImageLinks           `json:"imageLinks"`
	Language            string               `json:"language"`
	PreviewLink         string               `json:"previewLink"`
	InfoLink            string               `json:"infoLink"`
	CanonicalVolumeLink string               `json:"canonicalVolumeLink"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type ReadingModes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type SaleInfo struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"isEbook"`
}

type AccessInfo struct {
	Country                string `json:"country"`
	Viewability            string `json:"viewability"`
	Embeddable             bool   `json:"embeddable"`
	PublicDomain           bool   `json:"publicDomain"`
	TextToSpeechPermission string `json:"textToSpeechPermission"`
	WebReaderLink          string `json:"webReaderLink"`
	AccessViewStatus       string `json:"accessViewStatus"`
	QuoteSharingAllowed    bool   `json:"quoteSharingAllowed"`
}

type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}
