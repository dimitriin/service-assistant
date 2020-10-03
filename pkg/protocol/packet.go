package protocol

const (
	NullType      = 0
	ReadyzBitType = iota
	HealthzBitType
	CounterRegisterCmdType
	CounterIncCmdType
	CounterAddCMDType
	HistogramRegisterCmdType
	HistogramObserveCmdType
	GaugeRegisterCmdType
	GaugeIncCmdType
	GaugeDecCmdType
	GaugeSetCmdType
	GaugeAddCmdType
	GaugeSubCmdType
	GaugeSetToCurrentTimeCmdType
)

type Packet struct {
	Type  uint16
	Value interface{}
}
