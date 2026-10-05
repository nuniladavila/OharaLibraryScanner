package dbclient

import (
	"OharaLibraryScanner/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var notionEndpoint = os.Getenv("NOTION_BASE_URL")
var datasourceId = os.Getenv("BOOKS_DATASOURCE_ID")
var colorList = []string{"pink", "blue", "green", "orange", "red", "purple", "brown", "yellow"}
var shelfLocationColorMapping = map[string]string{
	"Spanish Non-Fiction":     "purple",
	"Spanish Fiction":         "red",
	"New TBR":                 "pink",
	"Comics":                  "yellow",
	"English Non-Fiction":     "brown",
	"English General Fiction": "green",
	"English Classics":        "orange",
	"English Speculative":     "blue",
	"English Horror/Suspense": "red",
}
var categoryColorMapping = map[string]string{
	"Non-Fiction": "brown",
	"Fiction":     "red",
}
var langColorMapping = map[string]string{
	"Spanish": "red",
	"English": "blue",
}

// sanitizeSelectName removes commas and trims whitespace for SelectItem names
func sanitizeSelectName(s string) string {
	s = strings.ReplaceAll(s, ",", "")
	return strings.TrimSpace(s)
}

func (c *NotionClient) FindBook(isbn string) string {
	url := fmt.Sprintf("%sdata_sources/%s/query", notionEndpoint, datasourceId)

	payloadJSON, err := BuildFindBookPayload(isbn)
	if err != nil {
		log.Println("error building Notion query payload:", err)
		return ""
	}

	payload := strings.NewReader(payloadJSON)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println("error creating Notion request:", err)
		return ""
	}
	req.Header.Add("Notion-Version", "2026-03-11")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		log.Println("Error querying Notion:", err, res.Status)
		return ""
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading Notion response body:", err)
		return ""
	}

	var response models.NotionQueryResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("error decoding Notion response:", err)
		return ""
	}

	if len(response.Results) == 1 {
		log.Println("One book found!")
		return response.Results[0].Properties.Title.Title[0].Text.Content
	}

	if len(response.Results) == 0 {
		return ""
	} else {
		log.Println("There was more than one book with the same ISBN. Please clean your db.")
		for i, book := range response.Results {
			log.Println(i, book)
		}
	}

	return ""
}

func (c *NotionClient) AddBook(oharaBook *models.OharaBook) {
	url := notionEndpoint + "pages"
	payloadJSON, err := GeneratePayload(oharaBook)

	if err != nil {
		log.Println("error generating payload:", err)
		return
	}

	payload := strings.NewReader(payloadJSON)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Notion-Version", "2026-03-11")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		log.Println("Error happened when trying to add book to the database", err, res.Status)
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if body != nil {
		log.Printf("\nSuccessfully added %s to the Notion Library!\n", oharaBook.Title)
	}

}

func BuildFindBookPayload(isbn string) (string, error) {
	queryPayload := map[string]interface{}{
		"filter": map[string]interface{}{
			"rich_text": map[string]string{"equals": isbn},
			"property":  "ISBN",
		},
		"start_cursor": "",
		"page_size":    5,
		"result_type":  "page",
	}

	payloadBytes, err := json.Marshal(queryPayload)
	if err != nil {
		return "", err
	}

	return string(payloadBytes), nil
}

// GeneratePayload builds a JSON payload for Notion from an OharaBook.
// It returns the JSON string or an error.
func GeneratePayload(oharaBook *models.OharaBook) (string, error) {
	if oharaBook == nil {
		return "", fmt.Errorf("OharaBook is null")
	}
	// build multi-select items from subcategories
	multiItems := make([]models.SelectItem, 0, len(oharaBook.Subcategories))
	for i, sc := range oharaBook.Subcategories {
		color := colorList[i%len(colorList)]
		multiItems = append(multiItems, models.SelectItem{Name: sanitizeSelectName(sc), Color: color})
	}

	// pick first author and sanitize
	authorName := ""
	if len(oharaBook.Authors) > 0 {
		authorName = sanitizeSelectName(oharaBook.Authors[0])
	}

	bookCoverOpenLib := fmt.Sprintf("https://covers.openlibrary.org/b/isbn/%s-M.jpg", oharaBook.ISBN)

	// Build properties map and only include non-null / non-empty values
	props := make(map[string]interface{})

	// Always present
	props["Title"] = models.TitleProp{Title: []models.TitleText{{Text: models.TextContent{Content: oharaBook.Title}}}}
	props["ISBN"] = models.RichTextProp{RichText: []models.RichTextItem{{Text: models.TextContent{Content: oharaBook.ISBN}}}}
	props["Read"] = models.CheckboxProp{Checkbox: oharaBook.Read}
	props["Date Added"] = models.DateProp{Date: models.DateStart{Start: time.Now().Format("2006-01-02")}}
	props["Language"] = models.SelectProp{Select: models.SelectItem{Name: sanitizeSelectName(oharaBook.Language), Color: langColorMapping[oharaBook.Language]}}
	props["Shelf Location"] = models.SelectProp{Select: models.SelectItem{Name: sanitizeSelectName(oharaBook.ShelfLocation), Color: shelfLocationColorMapping[oharaBook.ShelfLocation]}}
	props["Category"] = models.SelectProp{Select: models.SelectItem{Name: sanitizeSelectName(oharaBook.Category), Color: categoryColorMapping[oharaBook.Category]}}

	if authorName != "" {
		props["Author"] = models.SelectProp{Select: models.SelectItem{Name: authorName, Color: "default"}}
	}
	if oharaBook.Editor != "" {
		props["Editor"] = models.SelectProp{Select: models.SelectItem{Name: sanitizeSelectName(oharaBook.Editor), Color: "default"}}
	}
	if oharaBook.Publisher != "" {
		props["Publisher"] = models.SelectProp{Select: models.SelectItem{Name: sanitizeSelectName(oharaBook.Publisher), Color: "default"}}
	}
	if len(multiItems) > 0 {
		props["SubCategory"] = models.MultiSelectProp{MultiSelect: multiItems}
	}
	if oharaBook.PublishedDate != "" {
		props["Published Date"] = models.DateProp{Date: models.DateStart{Start: oharaBook.PublishedDate}}
	}
	if oharaBook.Edition != "" {
		props["Edition"] = models.RichTextProp{RichText: []models.RichTextItem{{Text: models.TextContent{Content: oharaBook.Edition}}}}
	}
	if oharaBook.PageCount > 0 {
		props["Page Count"] = models.NumberProp{Number: float64(oharaBook.PageCount)}
	}
	if oharaBook.BookCover != "" {
		props["Book Cover"] = models.FilesProp{Files: []models.FileItem{{External: models.FileExternal{URL: bookCoverOpenLib}, Name: oharaBook.Title}}}
	}

	payload := map[string]interface{}{
		"parent":     models.ParentProp{DataSourceID: datasourceId},
		"properties": props,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
