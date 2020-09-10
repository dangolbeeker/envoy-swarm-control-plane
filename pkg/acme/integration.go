package acme

import (
	"github.com/go-acme/lego/v4/lego"
	acme "github.com/nstapelbroek/envoy-swarm-control-plane/pkg/acme/storage"
	"github.com/nstapelbroek/envoy-swarm-control-plane/pkg/logger"
	tls "github.com/nstapelbroek/envoy-swarm-control-plane/pkg/provider/tls/storage"
)

type Integration struct {
	http01Port  uint
	lego        *lego.Client
	userStorage acme.User
	certStorage tls.Certificate
	logger      logger.Logger
}

func NewIntegration(port uint, userEmail string, userStorage acme.User, certStorage tls.Certificate, log logger.Logger) *Integration {
	c := &Integration{http01Port: port, userStorage: userStorage, certStorage: certStorage, logger: log}

	return c
}
