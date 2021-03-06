package vpn

import (
	"context"
	"sync"
	"time"

	"github.com/zerops-io/zcli/src/daemonStorage"
	"github.com/zerops-io/zcli/src/dnsServer"
	"github.com/zerops-io/zcli/src/utils/logger"
)

const (
	wireguardPort  = "51820"
	vpnApiGrpcPort = ":64510"
)

type Config struct {
	VpnCheckInterval   time.Duration
	VpnCheckRetryCount int
	VpnCheckTimeout    time.Duration
}

type Handler struct {
	config    Config
	logger    logger.Logger
	storage   *daemonStorage.Handler
	dnsServer *dnsServer.Handler

	lock sync.RWMutex
}

func New(
	config Config,
	logger logger.Logger,
	daemonStorage *daemonStorage.Handler,
	dnsServer *dnsServer.Handler,
) *Handler {
	return &Handler{
		config:    config,
		logger:    logger,
		storage:   daemonStorage,
		dnsServer: dnsServer,
	}
}

func (h *Handler) Run(ctx context.Context) error {
	t := time.NewTicker(h.config.VpnCheckInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			h.vpnStatusCheck(ctx)
		}
	}
}
