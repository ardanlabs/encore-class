// Package auth represent the encore application.
package auth

import (
	"context"
	"runtime"

	"github.com/ardanlabs/encore/foundation/logger"
)

// Service represents the encore service application.
//
//encore:service
type Service struct {
	log *logger.Logger
}

// NewService is called to create a new encore Service.
func NewService(log *logger.Logger) (*Service, error) {
	s := Service{
		log: log,
	}

	return &s, nil
}

// Shutdown implements a function that will be called by encore when the service
// is signaled to shutdown.
func (s *Service) Shutdown(force context.Context) {
	ctx := context.Background()

	defer s.log.Info(ctx, "shutdown", "status", "shutdown complete")
	s.log.Info(ctx, "shutdown", "status", "stopping database support")
}

// =============================================================================

// initService is called by Encore to initialize the service.
//
//lint:ignore U1000 "called by encore"
func initService() (*Service, error) {
	log := logger.New("auth")

	if err := startup(log); err != nil {
		return nil, err
	}

	return NewService(log)
}

func startup(log *logger.Logger) error {
	ctx := context.Background()

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "initService", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	return nil
}