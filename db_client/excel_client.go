package dbclient

import (
	"OharaLibraryScanner/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

func init() {
	godotenv.Load() // Loads .env file from project root (relative path)
}

var SHEET_NAME string
var INVENTORY_FILE_PATH string

func AddToExcel(oharaBook *models.OharaBook) {
	SHEET_NAME = os.Getenv("SHEET_NAME")
	INVENTORY_FILE_PATH = os.Getenv("EXCEL_FILE_PATH")

	file, err := excelize.OpenFile(INVENTORY_FILE_PATH)
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return
	}

	rows, err := file.GetRows(SHEET_NAME)
	if err != nil {
		fmt.Println("Error reading rows:", err)
		return
	}

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

	fmt.Println("Book was successfully added!")
}
