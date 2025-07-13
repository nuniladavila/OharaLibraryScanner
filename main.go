package main

import (
	"OharaLibraryScanner/models"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

var SHEET_NAME string = "BookCollection"

func main() {
	// AddBook("9780134190440") // Example ISBN
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Will this batch of books be fiction or non-fiction? (f/n): ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	switch category {
	case "n":
		category = "Non-Fiction"
	default:
		category = "Fiction"
	}

	fmt.Print("What's the location of these books?: ")
	for i := 0; i < len(models.BOOKSHELF_LOCATIONS); i++ {
		fmt.Printf("[%d] %s\n", i, models.BOOKSHELF_LOCATIONS[i])
	}
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	num, err := strconv.Atoi(location)

	if err != nil {
		fmt.Println("Input a valid number.")
		return
	}

	if num >= 0 && num < len(models.BOOKSHELF_LOCATIONS) {
		location = models.BOOKSHELF_LOCATIONS[num]
	} else {
		fmt.Println("Invalid selection. Please choose one of the numbers provided.")
		return
	}

	batchProps := models.OharaBatchProperties{
		Category: category,
		Location: location,
	}

	//Ask ISBN Input in loop
	for {
		fmt.Print("Enter isbn (or 'e' to exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "e":
			fmt.Println("Exiting...")
			return
		default:
			AddBook(input, batchProps)
		}
	}
}

func AddBook(isbn string, batchProps models.OharaBatchProperties) {
	//init input
	fmt.Println("Adding book with ISBN:", isbn)

	//send isbn api req
	apiKey := "AIzaSyBq7vz6P8Xjhk2tI2LW4oGXBg7jRJuZvB0"
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s", isbn, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching book data:", err)
		return
	}

	//process response
	newBook := ProcessGoogleBook(response)

	//add book to excel
	if newBook != nil {
		AddBookToInventory(isbn, batchProps, newBook)
	} else {
		fmt.Println("No book found with the provided ISBN :'(.")
	}
}

func ProcessGoogleBook(response *http.Response) *models.GoogleBookInfo {
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var result models.GoogleBooksResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	//If only one, add w/o issues, if more than one, ask for confirmation
	fmt.Printf("Total items: %d\n", result.TotalItems)
	if result.TotalItems == 1 {
		fmt.Println("Book title:", result.Items[0].VolumeInfo.Title)
		return &result.Items[0]

	} else if result.TotalItems > 1 {
		return ChooseBook(result.Items)

	} else {
		fmt.Println("No books found for the given ISBN.")
		return nil
	}
}

func AddBookToInventory(isbn string, batchProps models.OharaBatchProperties, newBook *models.GoogleBookInfo) {
	if newBook == nil {
		fmt.Println("No book to add.")
		return
	}

	fmt.Println("Adding book to inventory:")
	oharaBook := models.NewOharaBook(isbn, batchProps, newBook)

	fmt.Println("My book:", oharaBook)

	AddToExcel(oharaBook)
}

func AddToExcel(oharaBook *models.OharaBook) {
	file, err := excelize.OpenFile("C:\\Users\\nunil\\OneDrive\\Documents\\Libros@BibliotecaOhara.xlsx")
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return
	}

	rows, err := file.GetRows(SHEET_NAME)
	if err != nil {
		fmt.Println("Error reading rows:", err)
		return
	}

	fmt.Println("Total rows", len(rows))

	lastRow := 0
	for i := len(rows) - 1; i >= 0; i-- {
		for _, cell := range rows[i] {
			if cell != "" {
				lastRow = i + 1 // Excel rows are 1-indexed
				break
			}
		}
		if lastRow != 0 {
			break
		}
	}

	WriteToRow(file, lastRow+1, oharaBook)
}

func WriteToRow(file *excelize.File, row int, book *models.OharaBook) {
	fmt.Println("Adding book to row:", row)

	for key, value := range models.BuildBookPropToExcelCellMap(*book) {
		cell := fmt.Sprintf("%s%d", key, row)
		if err := file.SetCellValue(SHEET_NAME, cell, value); err != nil {
			fmt.Println("Error setting cell value:", err)
			return
		}
	}

	if err := file.Save(); err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	fmt.Printf("Book title '%s' added to row %d!\n", book.Title, row)
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
