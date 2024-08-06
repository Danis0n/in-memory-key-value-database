package initialization

import (
	"errors"
	"go.uber.org/zap"
	"in-memory-key-value-database/internal/configuration"
	"in-memory-key-value-database/internal/network"
	"in-memory-key-value-database/internal/tools"
	"time"
)

const (
	defaultMessageSize    = 10000
	defaultMaxConnections = 2048
	defaultAddress        = "localhost:3223"
	defaultIdleTimeout    = time.Minute * 5
)

func CreateNetwork(cfg *configuration.NetworkConfiguration, logger *zap.Logger) (*network.TCPServer, error) {
	if cfg == nil {
		return nil, errors.New("network configuration invalid")
	}

	address := defaultAddress
	maxConnections := defaultMaxConnections
	maxMessageSize := defaultMessageSize
	idleTimeout := defaultIdleTimeout

	if cfg.Address != "" {
		address = cfg.Address
	}

	if cfg.MaxConnections != 0 {
		maxConnections = cfg.MaxConnections
	}

	if cfg.MaxMessageSize != "" {
		size, err := tools.ParseSize(cfg.MaxMessageSize)
		if err != nil {
			return nil, errors.New("incorrect max message size")
		}
		maxMessageSize = size
	}

	if cfg.IdleTimeout != 0 {
		idleTimeout = cfg.IdleTimeout
	}

	return network.NewTCPServer(logger, address, maxConnections, maxMessageSize, idleTimeout)
}
