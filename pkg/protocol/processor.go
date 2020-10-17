package protocol

import (
	"net"
	"sync"

	"github.com/golang/protobuf/proto"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"

	"go.uber.org/zap"
)

type PacketStreamProcessor struct {
	conn    net.PacketConn
	handler *PacketHandler
	logger  *zap.Logger
	wg      sync.WaitGroup
}

func NewPacketStreamProcessor(conn net.PacketConn, handler *PacketHandler, logger *zap.Logger) *PacketStreamProcessor {
	return &PacketStreamProcessor{conn: conn, handler: handler, logger: logger}
}

func (u *PacketStreamProcessor) Process() error {
	for {
		buf := make([]byte, 1024)

		n, _, readErr := u.conn.ReadFrom(buf)

		if n > 0 {
			p := &payload.Packet{}

			if err := proto.Unmarshal(buf[:n], p); err != nil {
				u.logger.Error("packet decode error", zap.Binary("buf", buf[:n]), zap.Error(err))

				continue
			}

			u.logger.Info("incoming packet", zap.Any("packet", p))

			u.wg.Add(1)

			go func(p *payload.Packet) {
				defer u.wg.Done()

				if err := u.handler.Handle(p); err != nil {
					u.logger.Error("packet handle error", zap.Any("packet", p), zap.Error(err))
				}
			}(p)
		}

		if readErr != nil {
			u.wg.Wait()

			return readErr
		}
	}
}
