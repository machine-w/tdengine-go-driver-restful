package taossqlrestful

import (
	"database/sql/driver"
	"errors"
	"log"
)

// Stmt is stmt
type Stmt struct {
}

// Close  implement for stmt
func (stmt *Stmt) Close() error {
	return nil
}

// Query  implement for Query
func (stmt *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	log.Println("do query", args)
	myrows := RowS{
		Size: 3,
	}
	return &myrows, nil
}

// NumInput row numbers
func (stmt *Stmt) NumInput() int {
	// don't know how many row numbers
	return -1
}

// Exec exec  implement
func (stmt *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("some wrong")
}
