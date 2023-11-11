package main

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
	SessionState int32
	Flags        int64
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

const (
	FlagCheckered Flags = 1 << iota
	FlagWhite
	FlagGreen
	FlagYello
	FlagRed
	FlagBlue
	FlagDebris
	FlagCrossed
	FlagYellowWaving
	FlagOneLapToGreen
	FlagGreenHeld
	FlagTenToGo
	FlagFiveToGo
	FlagRandomWaving
	FlagCaution
	FlagCautionWaving

	// driver black flacks
	FlagBlack      Flags = 0x010000
	FlagDisqualify Flags = 0x020000
	FlagServicible Flags = 0x040000
	FlagFurled     Flags = 0x080000
	FlagRepair     Flags = 0x100000

	// start lights
	FlagStartHidden Flags = 0x10000000
	FlagStartReady  Flags = 0x20000000
	FlagStartSet    Flags = 0x40000000
	FlagStartGo     Flags = 0x80000000
)

const (
	StateInvalid SessionState = iota
	StateGetInCar
	StateWarmup
	StateParadeLaps
	StateRacing
	StateCheckered
	StateCoolDown
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
