package golic

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/nenavizhuleto/golic/driver"
)

var (
	connectionRequestQueueSize = 4
)

var nowFunc = time.Now

var (
	driversMu sync.Mutex
	drivers   = make(map[string]driver.Driver)
)

func Register(name string, driver driver.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("golic: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("golic: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

func OpenService(c driver.Connector) *Service {
	ctx, cancel := context.WithCancel(context.Background())
	service := &Service{
		connector: c,
		stop:      cancel,

		openerCh: make(chan struct{}, connectionRequestQueueSize),
	}

	service.openNewConnection(ctx)

	return service
}

func Open(driverName, dataSourceName string) (*Service, error) {
	driversMu.Lock()
	driveri, ok := drivers[driverName]
	driversMu.Unlock()
	if !ok {
		return nil, fmt.Errorf("golic: unknown driver %q", driverName)
	}

	if driverCtx, ok := driveri.(driver.DriverContext); ok {
		connector, err := driverCtx.OpenConnector(dataSourceName)
		if err != nil {
			return nil, err
		}
		return OpenService(connector), nil
	}

	return OpenService(dsnConnector{dsn: dataSourceName, driver: driveri}), nil
}
