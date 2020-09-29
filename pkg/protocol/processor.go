package protocol

import (
	"net"
	"sync"

	"go.uber.org/zap"
)

type HandlerInterface interface {
	Handle(value interface{}) error
}

type DecoderInterface interface {
	Decode(buf []byte, p *Packet) error
}

type PacketStreamProcessor struct {
	conn    net.PacketConn
	decoder DecoderInterface
	handler HandlerInterface
	logger  *zap.Logger
	wg      sync.WaitGroup
}

func NewPacketStreamProcessor(conn net.PacketConn, decoder DecoderInterface, handler HandlerInterface, logger *zap.Logger) *PacketStreamProcessor {
	return &PacketStreamProcessor{conn: conn, decoder: decoder, handler: handler, logger: logger}
}

func (u *PacketStreamProcessor) Process() error {
	for {
		buf := make([]byte, 1024)

		n, _, readErr := u.conn.ReadFrom(buf)

		if n > 0 {
			p := &Packet{}

			if err := u.decoder.Decode(buf[:n], p); err != nil {
				u.logger.Error("packet decode error", zap.Binary("buf", buf[:n]), zap.Error(err))

				continue
			}

			u.wg.Add(1)

			go func(p *Packet) {
				defer u.wg.Done()

				if err := u.handler.Handle(p.Value); err != nil {
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
