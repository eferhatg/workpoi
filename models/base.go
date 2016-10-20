package models

import (
	// "database/sql"
	// "errors"
	// "fmt"
	"github.com/jmoiron/sqlx"
	//"strings"
)

type Base struct {
	db    *sqlx.DB
	table string
	hasID bool
}
