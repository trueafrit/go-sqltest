package recorder

import (
	"database/sql/driver"
	"io"

	"github.com/DATA-DOG/go-sqlmock"
)

// Driver wraps the database driver.
type Driver struct {
	orig driver.Driver
	*output
}

// newDriver is the wrap-recorder of original driver.
func newDriver(orig driver.Driver, imports ImportList, code io.Writer, mock sqlmock.Sqlmock) driver.Driver {
	return &Driver{
		orig:   orig,
		output: newOutput(imports, code, mock),
	}
}

// Open opens a new connection to the database. name is a connection string.
func (d *Driver) Open(name string) (driver.Conn, error) {
	connection, err := d.orig.Open(name)
	if err != nil {
		return nil, err
	}

	return d.newConn(connection), nil
}
