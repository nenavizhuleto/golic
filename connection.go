package golic

import (
	"context"
	"github.com/nenavizhuleto/golic/driver"
	"time"
)

type driverConn struct {
	s          *Service
	createdAt  time.Time
	returnedAt time.Time
	ci         driver.Conn
}

func (dc *driverConn) Close() error {
	return dc.ci.Close()
}

type dsnConnector struct {
	dsn    string
	driver driver.Driver
}

func (t dsnConnector) Connect(_ context.Context) (driver.Conn, error) {
	return t.driver.Open(t.dsn)
}

func (t dsnConnector) Driver() driver.Driver {
	return t.driver
}
