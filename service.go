package golic

import (
	"context"
	"sync"

	"github.com/nenavizhuleto/golic/driver"
)

type Service struct {
	mu         sync.Mutex
	connector  driver.Connector
	connection *driverConn
	stop       context.CancelFunc

	openerCh chan struct{}
}

func (s *Service) VerifyContext(ctx context.Context) (bool, error) {
	return s.connection.ci.Verify(context.Background())
}

func (s *Service) Verify() (bool, error) {
	return s.VerifyContext(context.Background())
}

func (s *Service) ScanContext(ctx context.Context, result interface{}) error {
	return s.connection.ci.Scan(ctx, result)
}

func (s *Service) Scan(result interface{}) error {
	return s.ScanContext(context.Background(), result)
}

func (s *Service) Close() error {
	s.stop()
	return s.connection.Close()
}

func (s *Service) openNewConnection(ctx context.Context) {
	ci, err := s.connector.Connect(ctx)
	s.mu.Lock()
	defer s.mu.Unlock()

	if err != nil {
		return
	}

	dc := &driverConn{
		s:          s,
		createdAt:  nowFunc(),
		returnedAt: nowFunc(),
		ci:         ci,
	}

	s.connection = dc
}
