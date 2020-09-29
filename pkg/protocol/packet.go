package protocol

const (
	NullType = 0
	CmdCounterIncType = iota
	CmdCounterAddType
)

type Packet struct {
	Type uint16
	Value interface{}
}