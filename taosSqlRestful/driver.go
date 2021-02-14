package taossqlrestful

import (
	"database/sql"
	"database/sql/driver"
	"log"
)

// Driver mydb driver for implement database/sql/driver
type restfulDriver struct {
}

func init() {
	log.Println("register taossqlrestful driver")
	sql.Register("taossqlrestful", &restfulDriver{})
}

// Open for implement driver interface
func (driver *restfulDriver) Open(name string) (driver.Conn, error) {
	log.Println("exec open driver")
	return &Conn{}, nil
}
