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
	"English Horror/Suspense",
}

type OharaBook struct {
	BasicBook
	Editor        string `json:"editor"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"published_date"`
	Edition       string `json:"edition"`
	BookCover     string `json:"book_cover"`
}

type OharaBatchProperties struct {
	Category string `json:"category"`
	Location string `json:"location"`
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

	var lan = NormalizeLanguage(googleBookInfo.VolumeInfo.Language)

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
			Language:      lan,
		},
		Editor:        "",
		Publisher:     googleBookInfo.VolumeInfo.Publisher,
		PublishedDate: googleBookInfo.VolumeInfo.PublishedDate,
		Edition:       "",
		BookCover:     googleBookInfo.VolumeInfo.ImageLinks.Thumbnail,
	}
}

func NormalizeLanguage(inputLan string) string {
	switch inputLan {
	case "en":
		return "English"
	case "es":
		return "Spanish"
	default:
		return inputLan
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
