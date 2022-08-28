package db

import (
	"time"

	"github.com/scylladb/gocqlx/v2/table"
)

var UrlTable = table.New(table.Metadata{
	Name:    "fivefeeteleven.urls",
	Columns: []string{"id", "long_url", "lifespan", "created_at"},
	PartKey: []string{"id"},
	SortKey: []string{},
})

type UrlModel struct {
	ID        string
	LongUrl   string
	Lifespan  int64
	CreatedAt time.Time
}
