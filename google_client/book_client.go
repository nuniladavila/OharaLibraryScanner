package googleclient

import (
	inputmanagement "OharaLibraryScanner/input_management"
	"OharaLibraryScanner/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apiKey = "" // Set your Google Books API key here

func init() {
}

func GetBook(isbn string) *models.GoogleBookInfo {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s", isbn, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching book data:", err)
		return nil
	}

	//process response
	return ProcessGoogleBook(response)
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
	if result.TotalItems == 1 {
		fmt.Println("Book found!", result.Items[0].VolumeInfo.Title)
		return &result.Items[0]

	} else if result.TotalItems > 1 {
		return inputmanagement.ChooseBook(result.Items)

	} else {
		return nil
	}
}
