//
// Code generated by go-jet DO NOT EDIT.
// Generated at Wednesday, 17-Jul-19 13:11:01 CEST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Language struct {
	LanguageID int32 `sql:"primary_key"`
	Name       string
	LastUpdate time.Time
}
