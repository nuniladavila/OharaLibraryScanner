package main

import (
	dbclient "OharaLibraryScanner/dbclient"
	googleclient "OharaLibraryScanner/googleclient"
	inputmanagement "OharaLibraryScanner/inputmanagement"
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

		switch isbn {
		case "", "e":
			fmt.Println("Exiting...")
			return
		case "c":
			batchProperties = inputmanagement.BuildBatchProperties()
			isbn = inputmanagement.ReadBookISBNInput()
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

		AppendToInventory(oharaBook)
	}
}

func AppendToInventory(book *models.OharaBook) {
	if book == nil {
		fmt.Println("No book to add.")
		return
	}

	client, err := dbclient.NewDBClient("notion")

	if err != nil {
		fmt.Println("Error creating client.", err)
		return
	}

	client.AddBook(book)
}
