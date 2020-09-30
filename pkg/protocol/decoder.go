package protocol

import (
	"encoding/binary"
	"fmt"

	"github.com/dimitriin/service-assistant/pkg/protocol/cmd"
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
	case CmdCounterIncType:
		return &cmd.CounterIncCmd{}, nil
	case CmdCounterAddType:
		return &cmd.CounterAddCmd{}, nil
	default:
		return nil, fmt.Errorf("unknown packet type %d", t)
	}
}
