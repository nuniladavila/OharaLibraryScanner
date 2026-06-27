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
	BasicBook
	Editor        string `json:"editor"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"published_date"`
	Edition       string `json:"edition"`
	Language      string `json:"language"`
	BookCover     string `json:"book_cover"`
}

type OharaBatchProperties struct {
	Category string `json:"category"`
	Location string `json:"location"`
}

func NewOharaBook(isbn string, batchProps OharaBatchProperties, GoogleBookInfo *GoogleBookInfo, read bool) *OharaBook {
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
		BasicBook: BasicBook{
			Title:         GoogleBookInfo.VolumeInfo.Title,
			Authors:       GoogleBookInfo.VolumeInfo.Authors,
			Category:      batchProps.Category,
			Subcategories: GoogleBookInfo.VolumeInfo.Categories,
			ShelfLocation: batchProps.Location,
			ISBN:          isbn,
			Read:          read,
			PageCount:     GoogleBookInfo.VolumeInfo.PageCount,
		},
		Editor:        "",
		Publisher:     GoogleBookInfo.VolumeInfo.Publisher,
		PublishedDate: GoogleBookInfo.VolumeInfo.PublishedDate,
		Edition:       "",
		Language:      lan,
		BookCover:     GoogleBookInfo.VolumeInfo.ImageLinks.Thumbnail,
	}
}

func BuildOharaBook(isbn string, batchProps OharaBatchProperties, googleBookInfo *GoogleBookInfo, manualEntryBook *BasicBook) *OharaBook {
	if googleBookInfo == nil {
		if manualEntryBook == nil {
			return nil
		}
		return &OharaBook{
			BasicBook: *manualEntryBook,
		}
	}

	var lan string
	switch googleBookInfo.VolumeInfo.Language {
	case "en":
		lan = "English"
	case "es":
		lan = "Spanish"
	default:
		lan = googleBookInfo.VolumeInfo.Language
	}

	return &OharaBook{
		BasicBook: BasicBook{
			Title:         googleBookInfo.VolumeInfo.Title,
			Authors:       googleBookInfo.VolumeInfo.Authors,
			Category:      batchProps.Category,
			Subcategories: googleBookInfo.VolumeInfo.Categories,
			ShelfLocation: batchProps.Location,
			ISBN:          isbn,
			Read:          false,
			PageCount:     googleBookInfo.VolumeInfo.PageCount,
		},
		Editor:        "",
		Publisher:     googleBookInfo.VolumeInfo.Publisher,
		PublishedDate: googleBookInfo.VolumeInfo.PublishedDate,
		Edition:       "",
		Language:      lan,
		BookCover:     googleBookInfo.VolumeInfo.ImageLinks.Thumbnail,
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
