package irsdk

const (
	MemMapFileSize = 1164 * 1024
	MemMapFile     = "Local\\IRSDKMemMapFileName"

	IrsdkMaxBufs   = 3
	IrsdkMaxString = 32
	IrsdkMaxDesc   = 64
)

type (
	VarType int32
	irValue interface {
		int32 | float32 | float64 | bool
	}
)

const (
	// 1 byte
	irsdkChar VarType = iota
	irsdkBool
	// 4 bytes
	irsdkInt
	irsdkBitField
	irsdkFloat
	// 8 bytes
	irsdkDouble
	irsdkETCount
)

type (
	MMap   []byte
	Header struct {
		Version           int32
		Status            int32
		TickRate          int32
		SessionInfoUpdate int32
		SessionInfoLen    int32
		SessionInfoOffset int32
		NumVars           int32
		VarOffsetHeader   int32
		NumBuf            int32
		BufLen            int32
		Pad               [2]int32
		VarBufs           [IrsdkMaxBufs]VarBuffer
	}
)

type VarHeader struct {
	Type        VarType
	Offset      int32
	Count       int32
	CountAsTime bool
	Pad         [3]byte
	Name        [IrsdkMaxString]byte
	Desc        [IrsdkMaxDesc]byte
	Unit        [IrsdkMaxString]byte
}
type VarBuffer struct {
	TickCount int32
	BufOffset int32
	Pad       [2]int32
}
