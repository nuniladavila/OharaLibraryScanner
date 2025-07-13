package models

import "strings"

var BOOKSHELF_LOCATIONS = []string{
	"Spanish Non-Fiction",
	"Spanish Fiction",
	"New TBR",
	"Comics",
	"English Non-Fiction",
	"English General Fiction",
	"English Classics",
	"English Speculative",
}

type OharaBook struct {
	Title         string   `json:"title"`
	Authors       []string `json:"author"`
	Editor        string   `json:"editor"`
	Category      string   `json:"category"`
	Subcategories []string `json:"subcategory"`
	Publisher     string   `json:"publisher"`
	PublishedDate string   `json:"published_date"`
	Edition       string   `json:"edition"`
	Language      string   `json:"language"`
	ShelfLocation string   `json:"shelf_location"`
	ISBN          string   `json:"isbn"`
}

type OharaBatchProperties struct {
	Category string `json:"category"`
	Location string `json:"location"`
}

func NewOharaBook(isbn string, batchProps OharaBatchProperties, GoogleBookInfo *GoogleBookInfo) *OharaBook {
	if GoogleBookInfo == nil {
		return nil
	}

	var lan string
	switch GoogleBookInfo.VolumeInfo.Language {
	case "en":
		lan = "English"
	case "es":
		lan = "Spanish"
	default:
		lan = GoogleBookInfo.VolumeInfo.Language
	}

	return &OharaBook{
		Title:         GoogleBookInfo.VolumeInfo.Title,
		Authors:       GoogleBookInfo.VolumeInfo.Authors,
		Editor:        "",                                   // Assuming editor is not available in GoogleBookInfo
		Category:      batchProps.Category,                  // Assuming category is not available in GoogleBookInfo
		Subcategories: GoogleBookInfo.VolumeInfo.Categories, // Assuming subcategory is not available in GoogleBookInfo
		Publisher:     GoogleBookInfo.VolumeInfo.Publisher,
		PublishedDate: GoogleBookInfo.VolumeInfo.PublishedDate,
		Edition:       "", // Assuming edition is not available in GoogleBookInfo
		Language:      lan,
		ShelfLocation: batchProps.Location, // Assuming shelf location is not available in GoogleBookInfo
		ISBN:          isbn,                // Assuming ISBN is not available in GoogleBookInfo
	}
}

func BuildBookPropToExcelCellMap(book OharaBook) map[string]string {
	return map[string]string{
		"B": book.Title,
		"C": strings.Join(book.Authors, ", "),
		"D": book.Editor,
		"E": book.Category,
		"F": strings.Join(book.Subcategories, ", "),
		"G": book.Publisher,
		"H": book.PublishedDate,
		"I": book.Edition,
		"J": book.Language,
		"K": book.ShelfLocation,
		"L": book.ISBN,
	}
}
