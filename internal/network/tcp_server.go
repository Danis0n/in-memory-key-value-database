package network

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"io"
	"net"
	"sync"
	"time"
)

type Handler = func(context.Context, []byte) []byte

type TCPServer struct {
	logger         *zap.Logger
	address        string
	maxConnections int
	maxMessageSize int
	idleTimeout    time.Duration
}

func NewTCPServer(
	logger *zap.Logger,
	address string,
	maxConnections int,
	maxMessageSize int,
	idleTimeout time.Duration,
) (*TCPServer, error) {

	return &TCPServer{
		logger:         logger,
		address:        address,
		maxConnections: maxConnections,
		maxMessageSize: maxMessageSize,
		idleTimeout:    idleTimeout,
	}, nil
}

func (t *TCPServer) Start(ctx context.Context, handler Handler) error {

	listen, err := net.Listen("tcp", t.address)
	if err != nil {
		return err
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			t.logger.Error("error closing listener", zap.Error(err))
		}
	}(listen)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		err := func() error {
			for {
				connection, err := listen.Accept()
				if err != nil {

					t.logger.Error("error accepting connection", zap.Error(err))
					return errors.New("error accepting connection")
				}

				wg.Add(1)
				go func(connection net.Conn) {

					defer wg.Done()

					t.handleConnection(ctx, connection, handler)
				}(connection)
			}
		}()
		if err != nil {
			t.logger.Error("error accepting connection", zap.Error(err))
		}
	}()

	wg.Wait()

	return nil
}

func (t *TCPServer) handleConnection(
	ctx context.Context,
	connection net.Conn,
	handler Handler,
) {
	request := make([]byte, t.maxMessageSize)

	for {
		if err := connection.SetDeadline(time.Now().Add(t.idleTimeout)); err != nil {
			t.logger.Warn("failed to set read deadline", zap.Error(err))
			break
		}

		count, err := connection.Read(request)
		if err != nil {
			if err != io.EOF {
				t.logger.Warn("failed to read", zap.Error(err))
			}
			break
		}

		t.logger.Debug(
			"received request",
			zap.Any("request", request),
		)

		response := handler(ctx, request[:count])
		if _, err := connection.Write(response); err != nil {
			t.logger.Warn("failed to write", zap.Error(err))
			break
		}
	}
}
