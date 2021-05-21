//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var FilmCategory = newFilmCategoryTable("dvds", "film_category", "")

type filmCategoryTable struct {
	postgres.Table

	//Columns
	FilmID     postgres.ColumnInteger
	CategoryID postgres.ColumnInteger
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FilmCategoryTable struct {
	filmCategoryTable

	EXCLUDED filmCategoryTable
}

// AS creates new FilmCategoryTable with assigned alias
func (a FilmCategoryTable) AS(alias string) *FilmCategoryTable {
	return newFilmCategoryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FilmCategoryTable with assigned schema name
func (a FilmCategoryTable) FromSchema(schemaName string) *FilmCategoryTable {
	return newFilmCategoryTable(schemaName, a.TableName(), a.Alias())
}

func newFilmCategoryTable(schemaName, tableName, alias string) *FilmCategoryTable {
	return &FilmCategoryTable{
		filmCategoryTable: newFilmCategoryTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newFilmCategoryTableImpl("", "excluded", ""),
	}
}

func newFilmCategoryTableImpl(schemaName, tableName, alias string) filmCategoryTable {
	var (
		FilmIDColumn     = postgres.IntegerColumn("film_id")
		CategoryIDColumn = postgres.IntegerColumn("category_id")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
		allColumns       = postgres.ColumnList{FilmIDColumn, CategoryIDColumn, LastUpdateColumn}
		mutableColumns   = postgres.ColumnList{LastUpdateColumn}
	)

	return filmCategoryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		FilmID:     FilmIDColumn,
		CategoryID: CategoryIDColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
