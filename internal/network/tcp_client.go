package network

import (
	"go.uber.org/zap"
	"net"
)

type TCPClient struct {
	logger         *zap.Logger
	maxMessageSize int
	connection     net.Conn
}
