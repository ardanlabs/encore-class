// Package auth represent the encore application.
package auth

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"encore.dev"
	"github.com/ardanlabs/conf/v3"
	"github.com/ardanlabs/encore/app/sdk/auth"
	"github.com/ardanlabs/encore/foundation/keystore"
	"github.com/ardanlabs/encore/foundation/logger"
)

// Service represents the encore service application.
//
//encore:service
type Service struct {
	log  *logger.Logger
	auth *auth.Auth
}

// NewService is called to create a new encore Service.
func NewService(log *logger.Logger, ath *auth.Auth) (*Service, error) {
	s := Service{
		log:  log,
		auth: ath,
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

	auth, err := startup(log)
	if err != nil {
		return nil, err
	}

	return NewService(log, auth)
}

func startup(log *logger.Logger) (*auth.Auth, error) {
	ctx := context.Background()

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "initService", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Auth struct {
			ActiveKID string `conf:"default:54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"`
			Issuer    string `conf:"default:service project"`
		}
	}{
		Version: conf.Version{
			Build: encore.Meta().Environment.Name,
			Desc:  "Sales",
		},
	}

	const prefix = "AUTH"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil, err
		}
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.Info(ctx, "initService", "environment", encore.Meta().Environment.Name)

	out, err := conf.String(&cfg)
	if err != nil {
		return nil, fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "initService", "config", out)

	// -------------------------------------------------------------------------
	// Auth Support

	log.Info(ctx, "initService", "status", "initializing authentication support")

	// Load the private keys files from disk. We can assume some system like
	// Vault has created these files already. How that happens is not our
	// concern.

	ks := keystore.New()
	if err := ks.LoadKey(secrets.KeyID, secrets.KeyPEM); err != nil {
		return nil, fmt.Errorf("reading keys: %w", err)
	}

	authCfg := auth.Config{
		Log:       log,
		KeyLookup: ks,
		Issuer:    cfg.Auth.Issuer,
	}

	auth, err := auth.New(authCfg)
	if err != nil {
		return nil, fmt.Errorf("constructing auth: %w", err)
	}

	return auth, nil
}
