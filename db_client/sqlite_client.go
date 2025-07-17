package dbclient

import (
	"OharaLibraryScanner/models"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite", "./OharaLibrary_SQLite.db")
	if err != nil {
		fmt.Println("Error opening SQLite database:\n", err)
		return
	}
	fmt.Println("Connected to the SQLite database successfully.")
}

func CreateBooksTable() {
	query := `CREATE TABLE IF NOT EXISTS Books(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		BookTitle TEXT,
		Author TEXT,
		Editor TEXT,
		Category TEXT,
		SubCategory TEXT,
		Publisher TEXT,
		PublishedDate TEXT,
		Edition TEXT,
		Language TEXT,
		ShelfLocation TEXT,
		ISBN TEXT,
		Notes TEXT,
		Read INTEGER,
		DateAdded TEXT,
		DateAcquired TEXT
	)`

	_, err := db.Exec(query)

	fmt.Println("Books table created successfully.")
	if err != nil {
		fmt.Println("Error creating Books table:", err)
		return
	}
}

func AddBookToDatabase(oharaBook *models.OharaBook) {
	query := `INSERT INTO Books (
		BookTitle, Author, Editor, Category, SubCategory, Publisher, PublishedDate,
		Edition, Language, ShelfLocation, ISBN, Notes, Read, DateAdded, DateAcquired
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(
		query,
		oharaBook.Title,
		strings.Join(oharaBook.Authors, ", "),
		oharaBook.Editor,
		oharaBook.Category,
		strings.Join(oharaBook.Subcategories, ", "),
		oharaBook.Publisher,
		oharaBook.PublishedDate,
		oharaBook.Edition,
		oharaBook.Language,
		oharaBook.ShelfLocation,
		oharaBook.ISBN,
		"",                              //Notes
		0,                               // Read: 0 for false
		time.Now().Format(time.RFC3339), //Date Added
		time.Now().Format(time.RFC3339), //Date Acquired
	)
	if err != nil {
		fmt.Println("Error inserting book:", err)
		return
	}
	fmt.Println("Book added to SQLite database!")
}
