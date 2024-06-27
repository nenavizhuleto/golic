package driver

import "context"

type Driver interface {
	Open(name string) (Conn, error)
}

type DriverContext interface {
	OpenConnector(name string) (Connector, error)
}

type Conn interface {
	Verify(context.Context) (bool, error)
	Scan(context.Context, interface{}) error
	Close() error
}

type Connector interface {
	Connect(context.Context) (Conn, error)
	Driver() Driver
}
