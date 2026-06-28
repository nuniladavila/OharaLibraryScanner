package inputmanagement

import (
	"OharaLibraryScanner/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func ReadBookISBNInput() string {
	fmt.Print("Enter isbn or: [e] to exit, [c] to change batch properties:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}

func BuildBatchProperties() models.OharaBatchProperties {
	return models.OharaBatchProperties{
		Category: GetCategoryBatchProperty(),
		Location: GetLocationBatchProperty(),
	}
}

func GetCategoryBatchProperty() string {
	fmt.Print("Will this batch of books be fiction or non-fiction? (f/n): ")

	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	switch category {
	case "n":
		return "Non-Fiction"
	default:
		return "Fiction"
	}
}

func GetReadSingleProp(bookTitle string) bool {
	fmt.Printf("Did you read %s? (y/n): ", bookTitle)

	read, _ := reader.ReadString('\n')
	read = strings.TrimSpace(read)

	switch read {
	case "n":
		return false
	default:
		return true
	}
}

func GetLocationBatchProperty() string {
	fmt.Print("What's the location of these books?:\n")
	for i := 0; i < len(models.BOOKSHELF_LOCATIONS); i++ {
		fmt.Printf("[%d] %s\n", i, models.BOOKSHELF_LOCATIONS[i])
	}
	fmt.Printf("Choose the number from %d to %d belonging to the desired location: ", 0, len(models.BOOKSHELF_LOCATIONS)-1)

	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	num, err := strconv.Atoi(location)

	if err != nil {
		fmt.Println("Input a valid number.")
		return ""
	}

	if num >= 0 && num < len(models.BOOKSHELF_LOCATIONS) {
		return models.BOOKSHELF_LOCATIONS[num]
	} else {
		fmt.Println("Invalid selection. Please choose one of the numbers provided.")
		return ""
	}
}

func BuildRequiredBookDetailsManually(isbn string, batch models.OharaBatchProperties) *models.BasicBook {
	fmt.Println("The book isn't available in the Google API. Please add some required details manually.")

	fmt.Print("Title of the book?: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Author of the book?: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("SubCategory of the book?: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Page count of the book?: ")
	page, _ := reader.ReadString('\n')
	page = strings.TrimSpace(page)
	intPage, err := strconv.ParseInt(page, 10, 0)
	if err != nil {
		intPage = 0
	}

	return &models.BasicBook{
		Title:         title,
		Authors:       strings.Split(author, ","),
		Category:      batch.Category,
		Subcategories: strings.Split(category, ","),
		ShelfLocation: batch.Location,
		ISBN:          isbn,
		PageCount:     int(intPage),
	}

}

func ChooseBook(items []models.GoogleBookInfo) *models.GoogleBookInfo {
	fmt.Println("Multiple books found. Please confirm which book to add:")

	for i, item := range items {
		fmt.Printf("[%d] %s\n", i+1, item.VolumeInfo.Title)
	}

	reader := bufio.NewReader(os.Stdin)
	index, _ := reader.ReadString('\n')
	index = strings.TrimSpace(index)

	num, err := strconv.Atoi(index)

	if err != nil {
		fmt.Println("Input a valid number.")
		ChooseBook(items)
	}

	if num >= 0 && num < len(items) {
		return &items[num]
	} else {
		fmt.Println("Invalid selection. Please choose one of the numbers provided.")
		ChooseBook(items)
	}
	return nil
}
