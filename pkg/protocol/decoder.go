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
	case CounterRegisterCmdType:
		return &payload.CounterRegisterCmd{}, nil
	case CounterIncCmdType:
		return &payload.CounterIncCmd{}, nil
	case CounterAddCMDType:
		return &payload.CounterAddCmd{}, nil
	case HistogramRegisterCmdType:
		return &payload.HistogramRegisterCmd{}, nil
	case HistogramObserveCmdType:
		return &payload.HistogramObserveCmd{}, nil
	case GaugeRegisterCmdType:
		return &payload.GaugeRegisterCmd{}, nil
	case GaugeIncCmdType:
		return &payload.GaugeIncCmd{}, nil
	case GaugeDecCmdType:
		return &payload.GaugeDecCmd{}, nil
	case GaugeSetCmdType:
		return &payload.GaugeSetCmd{}, nil
	case GaugeAddCmdType:
		return &payload.GaugeAddCmd{}, nil
	case GaugeSubCmdType:
		return &payload.GaugeSubCmd{}, nil
	case GaugeSetToCurrentTimeCmdType:
		return &payload.GaugeSetToCurrentTimeCmd{}, nil
	default:
		return nil, fmt.Errorf("unknown packet type %d", t)
	}
}
