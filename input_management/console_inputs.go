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
	fmt.Print("Enter isbn (or 'e' to exit): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "e":
		return ""
	default:
		return input
	}
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
