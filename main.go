package main

import (
	dbclient "OharaLibraryScanner/dbclient"
	googleclient "OharaLibraryScanner/googleclient"
	inputmanagement "OharaLibraryScanner/inputmanagement"
	"OharaLibraryScanner/models"
	"fmt"
)

var client dbclient.DbClienter

func main() {
	fmt.Println("Welcome to the Ohara Library Scanner!")

	client, _ = dbclient.NewDBClient("notion")

	AddBookProgram()
}

func AddBookProgram() {

	//Ask user initial questions
	batchProperties := inputmanagement.BuildBatchProperties()

	//Ask ISBN Input in loop
	for {
		isbn := inputmanagement.ReadBookISBNInput()

		switch isbn {
		case "", "e": //Exit
			fmt.Println("Exiting...")
			return
		case "c": //Change Shelf location
			batchProperties = inputmanagement.BuildBatchProperties()
			isbn = inputmanagement.ReadBookISBNInput()
		}

		fmt.Println("Checking if book is already added...")
		bookFound := client.FindBook(isbn)
		if bookFound != "" {
			fmt.Println("Book already exists! You already added", bookFound)
			continue
		}
		fmt.Println("New book confirmed. Adding book:", isbn)

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

	client.AddBook(book)
}
