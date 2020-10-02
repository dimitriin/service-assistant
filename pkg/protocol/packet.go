package protocol

const (
	NullType      = 0
	ReadyzBitType = iota
	HealthzBitType
	CounterIncCmdType
	CounterAddCMDType
)

type Packet struct {
	Type  uint16
	Value interface{}
}
