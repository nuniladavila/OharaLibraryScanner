package dbclient

import (
	"OharaLibraryScanner/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	godotenv.Load() // Loads .env file from project root (relative path)
}

func AddToBookDatabase(oharaBook *models.OharaBook) {
	var server = os.Getenv("DB_SERVER")
	var database = os.Getenv("DB_DATABASE")
	var port = 1433
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")

	// Build connection string with pwd since GO doesn't support Entra login
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected!")

	query := `INSERT INTO Books 
	(BookTitle, Author, Editor, Category, SubCategory, Publisher, PublishedDate, 
	Edition, Language, ShelfLocation, ISBN, Notes, [Read], DateAdded, DateAcquired)
	VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15)`

	_, err = db.ExecContext(ctx, query,
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
		"",
		false, // Assuming Read is false by default
		time.Now(),
		time.Now(), // Assuming DateAdded and DateAcquired are the current time
	)
	if err != nil {
		fmt.Println("Error inserting book:", err)
		return
	}
	fmt.Println("Book added to database!")
}
