package main

import (
	dbclient "OharaLibraryScanner/db_client"
	googleclient "OharaLibraryScanner/google_client"
	inputmanagement "OharaLibraryScanner/input_management"
	"OharaLibraryScanner/models"
	"fmt"
)

func main() {
	//Ask user initial questions
	batchProperties := inputmanagement.BuildBatchProperties()

	//Ask ISBN Input in loop
	for {
		isbn := inputmanagement.ReadBookISBNInput()

		if isbn == "" {
			fmt.Println("Exiting...")
			return
		}

		AddBook(isbn, batchProperties)
	}
}

func AddBook(isbn string, batchProps models.OharaBatchProperties) {
	fmt.Println("Adding book with ISBN:", isbn)

	//send isbn api req
	googleBook := googleclient.GetBook(isbn)

	if googleBook == nil {
		fmt.Println("No book found with the provided ISBN :'(.")
		return
	}

	oharaBook := models.NewOharaBook(isbn, batchProps, googleBook)

	//add book to excel
	AppendToInventory(oharaBook)
}

func AppendToInventory(book *models.OharaBook) {
	if book == nil {
		fmt.Println("No book to add.")
		return
	}

	// excelclient.AddToExcel(book)
	dbclient.AddToBookDatabase(book)
}
