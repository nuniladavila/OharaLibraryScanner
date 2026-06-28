package dbclient

import (
	"OharaLibraryScanner/models"
	"fmt"
	"strings"
)

type DbClienter interface {
	AddBook(oharaBook *models.OharaBook)
}

type ExcelClient struct{}
type NotionClient struct{}
type SqlServerClient struct{}
type SqliteClient struct{}

func NewDBClient(name string) (DbClienter, error) {
	switch strings.ToLower(name) {
	case "excel":
		return &ExcelClient{}, nil
	case "notion":
		return &NotionClient{}, nil
	case "sqlserver", "sql_server", "mssql":
		return &SqlServerClient{}, nil
	case "sqlite":
		return &SqliteClient{}, nil
	default:
		return nil, fmt.Errorf("unsupported db client %q", name)
	}
}
