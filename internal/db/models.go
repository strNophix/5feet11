package db

import "github.com/scylladb/gocqlx/v2/table"

var UrlTable = table.New(table.Metadata{
	Name:    "fivefeeteleven.urls",
	Columns: []string{"id", "redirect_url", "secret"},
	PartKey: []string{"id"},
	SortKey: []string{},
})

type UrlModel struct {
	Id          string
	RedirectUrl string
	Secret      *string
}