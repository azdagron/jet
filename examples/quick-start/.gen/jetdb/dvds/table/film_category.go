//
// Code generated by go-jet DO NOT EDIT.
// Generated at Wednesday, 17-Jul-19 13:11:01 CEST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet"
)

var FilmCategory = newFilmCategoryTable()

type FilmCategoryTable struct {
	jet.Table

	//Columns
	FilmID     jet.ColumnInteger
	CategoryID jet.ColumnInteger
	LastUpdate jet.ColumnTimestamp

	AllColumns     jet.ColumnList
	MutableColumns jet.ColumnList
}

// creates new FilmCategoryTable with assigned alias
func (a *FilmCategoryTable) AS(alias string) *FilmCategoryTable {
	aliasTable := newFilmCategoryTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newFilmCategoryTable() *FilmCategoryTable {
	var (
		FilmIDColumn     = jet.IntegerColumn("film_id")
		CategoryIDColumn = jet.IntegerColumn("category_id")
		LastUpdateColumn = jet.TimestampColumn("last_update")
	)

	return &FilmCategoryTable{
		Table: jet.NewTable("dvds", "film_category", FilmIDColumn, CategoryIDColumn, LastUpdateColumn),

		//Columns
		FilmID:     FilmIDColumn,
		CategoryID: CategoryIDColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     jet.ColumnList{FilmIDColumn, CategoryIDColumn, LastUpdateColumn},
		MutableColumns: jet.ColumnList{LastUpdateColumn},
	}
}