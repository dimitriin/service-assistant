package protocol

import (
	"encoding/binary"
	"fmt"

	"github.com/dimitriin/service-assistant/pkg/protocol/payload"
	"github.com/golang/protobuf/proto"
)

type Decoder struct{}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) Decode(data []byte, p *Packet) error {
	packetType := binary.LittleEndian.Uint16(data[:2])

	packetValue, err := d.getProtoMessageByType(packetType)

	if err != nil {
		return err
	}

	err = proto.Unmarshal(data[2:], packetValue)

	if err != nil {
		return err
	}

	p.Type = packetType
	p.Value = packetValue

	return nil
}

func (d *Decoder) getProtoMessageByType(t uint16) (proto.Message, error) {
	switch t {
	case ReadyzBitType:
		return &payload.ReadyBit{}, nil
	case HealthzBitType:
		return &payload.HealthBit{}, nil
	case CounterIncCmdType:
		return &payload.CounterIncCmd{}, nil
	case CounterAddCMDType:
		return &payload.CounterAddCmd{}, nil
	default:
		return nil, fmt.Errorf("unknown packet type %d", t)
	}
}
