package main

import (
	dbclient "OharaLibraryScanner/db_client"
	googleclient "OharaLibraryScanner/google_client"
	inputmanagement "OharaLibraryScanner/input_management"
	"OharaLibraryScanner/models"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Ohara Library Scanner!")

	AddBookProgram()
}

func AddBookProgram() {
	//Ask user initial questions
	batchProperties := inputmanagement.BuildBatchProperties()

	//Ask ISBN Input in loop
	for {
		isbn := inputmanagement.ReadBookISBNInput()

		if isbn == "" {
			fmt.Println("Exiting...")
			return
		}

		fmt.Println("Adding book with ISBN:", isbn)

		//send isbn api req
		googleBook := googleclient.GetBook(isbn)
		var manualEntryBook *models.BasicBook

		if googleBook == nil {
			fmt.Println("No book found with the provided ISBN :'(.")
			manualEntryBook = inputmanagement.BuildRequiredBookDetailsManually(isbn, batchProperties)
		}

		//Build base class book with google or manual entry
		oharaBook := models.BuildOharaBook(isbn, batchProperties, googleBook, manualEntryBook)

		//Ask if read
		read := inputmanagement.GetReadSingleProp(oharaBook.Title)
		oharaBook.Read = read

		AddBook(oharaBook)
	}
}

func AddBook(book *models.OharaBook) {
	AppendToInventory(book)
}

func AppendToInventory(book *models.OharaBook) {
	if book == nil {
		fmt.Println("No book to add.")
		return
	}

	// excelclient.AddToExcel(book)
	// dbclient.AddToBookDatabase(book)
	// dbclient.AddBookToDatabase(book)
	dbclient.AddBookToNotion(book)
}
