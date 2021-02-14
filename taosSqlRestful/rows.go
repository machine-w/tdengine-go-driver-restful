package taossqlrestful

import (
	"database/sql/driver"
	"io"
)

// RowS  RowS implemmet for driver.Rows
type RowS struct {
	Size int64
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *RowS) Columns() []string {
	return []string{
		"name",
		"age",
		"version",
	}
}

// Close closes the rows iterator.
func (r *RowS) Close() error {
	return nil
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
//
// The dest should not be written to outside of Next. Care
// should be taken when closing Rows not to modify
// a buffer held in dest.
func (r *RowS) Next(dest []driver.Value) error {
	if r.Size == 0 {
		return io.EOF
	}
	name := "dalong"
	age := 333
	version := "v1"
	dest[0] = name
	dest[1] = age
	dest[2] = version
	r.Size--
	return nil
}
