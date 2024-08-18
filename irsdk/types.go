//nolint:lll // readability
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
	SessionState  int32
	Flags         int64
	TrackLocation int32
	TrackSurface  int32
)

const (
	// 1 byte
	IrsdkTypeChar VarType = iota
	IrsdkTypeBool
	// 4 bytes
	IrsdkTypeInt
	IrsdkTypeBitField
	IrsdkTypeFloat
	// 8 bytes
	IrsdkTypeDouble
	IrsdkTypeETCount
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

const (
	TrackWetnessUnknown = iota
	TrackWetnessDry
	TrackWetnessMostlyDry
	TrackWetnessVeryLightlyWet
	TrackWetnessLightlyWet
	TrackWetnessModeratelyWet
	TrackWetnessVeryWet
	TrackWetnessExtremeWet
)

const (
	TrackLocationNotInWorld TrackLocation = iota - 1
	TrackLocationOffTrack
	TrackLocationInPitStall
	TrackLocationAproachingPits
	TrackLocationOnTrack
)

const (
	TrackSurfaceNotInWorld TrackSurface = iota - 1
	TrackSurfaceUndefined
	TrackSurfaceAsphalt1
	TrackSurfaceAsphalt2
	TrackSurfaceAsphalt3
	TrackSurfaceAsphalt4
	TrackSurfaceConcrete1
	TrackSurfaceConcrete2
	TrackSurfaceRacingDirt1
	TrackSurfaceRacingDirt2
	TrackSurfacePaint1
	TrackSurfacePaint2
	TrackSurfaceRumble1
	TrackSurfaceRumble2
	TrackSurfaceRumble3
	TrackSurfaceRumble4
	TrackSurfaceGrass1
	TrackSurfaceGrass2
	TrackSurfaceGrass3
	TrackSurfaceGrass4
	TrackSurfaceDirt1
	TrackSurfaceDirt2
	TrackSurfaceDirt3
	TrackSurfaceDirt4
	TrackSurfaceSand
	TrackSurfaceGravel1
	TrackSurfaceGravel2
	TrackSurfaceGrasscrete
	TrackSurfaceAstroturf
)

type BroadcastCmd int32

const (
	BroadcastCamSwitchPos            BroadcastCmd = iota // car position, group, camera
	BroadcastCamSwitchNum                                // driver #, group, camera
	BroadcastCamSetState                                 // CameraState, unused, unused
	BroadcastReplaySetPlaySpeed                          // speed, slowMotion, unused
	BroadcastReplaySetPlayPosition                       // ReplayPosMode, Frame Number (high, low)
	BroadcastReplaySearch                                // ReplaySearchMode, unused, unused
	BroadcastReplaySetState                              // ReplayStateMode, unused, unused
	BroadcastReloadTextures                              // ReloadTexturesMode, carIdx, unused
	BroadcastChatComand                                  // ChatCommand, subcommand, unused
	BroadcastPitCommand                                  // PitCommand, parameter, unused
	BroadcastTelemetryCommand                            // TelemetryCommand, unused, unused
	BroadcastFFBCommand                                  // FFBCommand, value (float, high,low)
	BroadcastReplaySearchSessionTime                     //
	BroadcastVideoCapture
)

type CameraState uint32

const (
	CameraStateIsSessionScreen = 0x0001 // the camera tool can only be activated if viewing the session screen (out of car)
	CameraStateIsScenicActive  = 0x0002 // the scenic camera is active (no focus car)
	// these can be changed with broadcast messages
	CameraStateCamToolActive         = 0x0004
	CameraStateUiHidden              = 0x0008
	CameraStateUseAutoShotSelection  = 0x0010
	CameraStateUseTemporaryEdits     = 0x0020
	CameraStateUseKeyAcceleration    = 0x0040
	CameraStateUseKey10xAcceleration = 0x0080
	CameraStateUseMouseAimMode       = 0x0100
)

type ChatCommand uint32

const (
	ChatCommandMacro     ChatCommand = iota // pass in a number from 1-15 representing the chat macro to launch
	ChatCommandBeginChat                    // Open up a new chat window
	ChatCommandReply                        // Reply to last private chat
	ChatCommandCancel                       // Close chat window
)

type PitCommand uint32

const (
	PitCommandClear      PitCommand = iota // Clear all pit checkboxes
	PitCommandWS                           // clean the windshield, using one tear-off
	PitCommandFuel                         // add fuel, optionally specify the amount in liters
	PitCommandLF                           // change the left front tire, optionally specify the pressure in KPa
	PitCommandRF                           // change the right front tire, optionally specify the pressure in KPa
	PitCommandLR                           // change the left rear tire, optionally specify the pressure in KPa
	PitCommandRR                           // change the right rear tire, optionally specify the pressure in KPa
	PitCommandClearTires                   // clear tire pit checkboxes
	PitCommandFR                           // use fast repair
	PitCommandClearWS                      // uncheck the windshield checkbox
	PitCommandClearFR                      // clear fast repair checkbox
	PitCommandClearFuel                    // clear fuel checkbox
)

type TelemetryCommand uint32

const (
	TelemetryCommandStop TelemetryCommand = iota
	TelemetryCommandStart
	TelemetryCommandRestart
)

type FFBCommand uint32

const (
	FFBCommandMaxForce FFBCommand = iota
)

type ReplayStateMode uint32

const (
	ReplayStateModeEraseTape ReplayStateMode = iota
)

type ReplayPosMode uint32

const (
	ReplayPosModeBegin ReplayPosMode = iota
	ReplayPosModeCurrent
	ReplayPosModeEnd
)

type ReplaySearchMode uint32

const (
	ReplaySearchModeStart ReplaySearchMode = iota
	ReplaySearchModeEnd
	ReplaySearchModePrevSession
	ReplaySearchModeNextSession
	ReplaySearchModePrevLap
	ReplaySearchModeNextLap
	ReplaySearchModePrevFrame
	ReplaySearchModeNextFrame
	ReplaySearchModePrevIncident
	ReplaySearchModeNextIncident
)

type ReloadTexturesMode uint32

const (
	ReloadTexturesModeAll ReloadTexturesMode = iota
	ReloadTexturesModeCarIdx
)

type VideoCaptureMode uint32

const (
	VideoCaptureModeTriggerSreenshot VideoCaptureMode = iota
	VideoCaptureModeStartCapture
	VideoCaptureModeStopCapture
	VideoCaptureModeToggleVideoCapture
	VideoCaptureModeShowVideoTimer
	VideoCaptureModeHideVideoTimer
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
