package irsdk

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/exp/slices"
	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	goyaml "gopkg.in/yaml.v3"

	"github.com/mpapenbr/goirsdk/yaml"
)

const (
	SimStatusUrl       = "http://127.0.0.1:32034/get_sim_status?object=simStatus"
	DataValidEventName = "Local\\IRSDKDataValidEvent"
	BroadcastMsg       = "IRSDK_BROADCASTMSG"
)

//nolint:lll // readabilty
var (
	ErrInvalidDataRequest = errors.New("Invalid data request")
	ErrNoMatchingDataType = errors.New("requested data type does not match iRacing data type")
	ErrBroadcastError     = errors.New("could not send broadcast message")
)

type Irsdk struct {
	SimIsRunning       bool
	sharedMem          MMap
	sharedMemHandle    windows.Handle
	dataEventHandle    windows.Handle
	hdr                Header                 // will never change during connection
	varHeaders         []VarHeader            // VarHeaders
	vHeaderLookup      map[string]VarHeader   // lookup for VarHeader by name
	latestVarBuffer    []byte                 // copy of the latest VarBuf
	latestYamlBuffer   []byte                 // copy of the latest yaml raw data
	lastValidTickCount int32                  // tickCount of the latest...Buffer
	lastValidTime      time.Time              // timestamp of the latest valid data
	previousYamlString string                 // yaml string from previous call
	dataDict           map[string]interface{} // holds all data from API
	irsdkYaml          yaml.IrsdkYaml         // parsed yaml data
	cfg                irsdkConfig            // config values
	user32             *windows.LazyDLL       // user32.dll used for broadcast messages
	broadcastId        uintptr                // handle for broadcast messages
}

type Option interface {
	apply(*irsdkConfig)
}
type optionFunc func(*irsdkConfig)

func (f optionFunc) apply(cfg *irsdkConfig) { f(cfg) }

func WithWaitForValidDataTimeout(t time.Duration) Option {
	return optionFunc(func(cfg *irsdkConfig) {
		cfg.waitForSingleObjectTimeout = t
	})
}

type irsdkConfig struct {
	waitForSingleObjectTimeout time.Duration
}
type ApiConfig struct {
	WaitForSimTimeout int
}

func NewIrsdk(opts ...Option) *Irsdk {
	cfg := &irsdkConfig{
		waitForSingleObjectTimeout: 32 * time.Millisecond,
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}

	mem, h := openSharedMem()
	dataEvent := openDataValidEventHandler()
	ret := Irsdk{
		sharedMem:       mem,
		dataEventHandle: dataEvent,
		sharedMemHandle: h,
		cfg:             *cfg,
	}
	ret.initialize()
	return &ret
}

func NewIrsdkWithFile(f *os.File) *Irsdk {
	mem := openTelemetryDump(f)

	ret := Irsdk{sharedMem: mem}
	ret.initialize()
	//nolint:lll  //by design
	copy(ret.latestYamlBuffer,
		ret.sharedMem[ret.hdr.SessionInfoOffset:ret.hdr.SessionInfoOffset+ret.hdr.SessionInfoLen])
	return &ret
}

// @deprecated use IsSimRunning instead
func CheckIfSimIsRunning() bool {
	// by default we use the default http client. Note: There is no timeout set
	client := http.DefaultClient
	resp, err := IsSimRunning(context.Background(), client)
	if err != nil {
		log.Printf("Could not connect to iRacing service: %v\n", err)
		return false
	}
	return resp
}

// checks if the iRacing simulation is running
func IsSimRunning(ctx context.Context, client *http.Client) (bool, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		SimStatusUrl,
		http.NoBody)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	return strings.Contains(string(body), "running:1"), nil
}

func (irsdk *Irsdk) Close() {
	if err := windows.CloseHandle(irsdk.sharedMemHandle); err != nil {
		log.Printf("Error closing shared mem handle: %v\n", err)
	}
	if err := windows.CloseHandle(irsdk.dataEventHandle); err != nil {
		log.Printf("Error closing dataEvent handle: %v\n", err)
	}
}

func (irsdk *Irsdk) WaitForValidData() bool {
	return irsdk.waitForValidData(irsdk.cfg.waitForSingleObjectTimeout)
}

func (irsdk *Irsdk) waitForValidData(t time.Duration) bool {
	h, err := windows.WaitForSingleObject(irsdk.dataEventHandle,
		uint32(t.Milliseconds()))
	if err != nil {
		log.Fatalf("waitForValidData: %v\n", err)
	}
	return h == 0
}

// returns true if new valid data is copied from iRacing telemetry to this Irdsk struct
func (irsdk *Irsdk) GetData() bool {
	return irsdk.getDataInternal(irsdk.cfg.waitForSingleObjectTimeout)
}

// returns true if new valid data is copied from iRacing telemetry to this Irdsk struct
// The call will wait up to dataReadyTimeout for new data to arrive
func (irsdk *Irsdk) GetDataWithDataReadyTimeout(dataReadyTimeout time.Duration) bool {
	return irsdk.getDataInternal(dataReadyTimeout)
}

//nolint:lll // better readability for copy
func (irsdk *Irsdk) getDataInternal(waitForData time.Duration) bool {
	ret := irsdk.waitForValidData(waitForData)
	if !ret {
		return false
	}

	latestBuf := irsdk.hdr.VarBufs[0]
	for _, v := range irsdk.hdr.VarBufs {
		if v.TickCount > 0 && v.TickCount > latestBuf.TickCount {
			latestBuf = v
		}
	}
	// corner case: sometimes we get just 0x00 content. In such cases just don't update
	if isZeroed(irsdk.sharedMem[latestBuf.BufOffset : latestBuf.BufOffset+100]) {
		return false
	}

	irsdk.lastValidTickCount = latestBuf.TickCount
	irsdk.lastValidTime = time.Now()

	copy(irsdk.latestVarBuffer, irsdk.sharedMem[latestBuf.BufOffset:latestBuf.BufOffset+irsdk.hdr.BufLen])
	copy(irsdk.latestYamlBuffer, irsdk.sharedMem[irsdk.hdr.SessionInfoOffset:irsdk.hdr.SessionInfoOffset+irsdk.hdr.SessionInfoLen])
	//nolint:errcheck // by design
	irsdk.GetYaml()
	return true
}

func (irsdk *Irsdk) GetYaml() (*yaml.IrsdkYaml, error) {
	yamlStr := irsdk.GetYamlString()
	// if the yaml string is the same as the previous one, we don't need to parse it again
	if yamlStr == irsdk.previousYamlString {
		return &irsdk.irsdkYaml, nil
	}
	var yamlData yaml.IrsdkYaml
	var err error

	err = goyaml.Unmarshal([]byte(yamlStr), &yamlData)
	if err != nil {
		// maybe the yaml is just not valid (see issue #6)
		// let's try to fix it and try again
		err = goyaml.Unmarshal([]byte(irsdk.RepairedYaml(yamlStr)), &yamlData)
		if err != nil {
			return nil, err
		}
	}
	if irsdk.validYamlData(&yamlData) {
		irsdk.previousYamlString = yamlStr
		irsdk.irsdkYaml = yamlData
	}

	return &irsdk.irsdkYaml, nil
}

func (irsdk *Irsdk) GetLatestYaml() *yaml.IrsdkYaml {
	return &irsdk.irsdkYaml
}

//nolint:errcheck // by design
func (irsdk *Irsdk) GetYamlString() string {
	var b bytes.Buffer
	wInUtf8 := transform.NewWriter(&b, charmap.Windows1252.NewDecoder())
	wInUtf8.Write([]byte(extractCStringFromSlice(irsdk.latestYamlBuffer)))
	wInUtf8.Close()

	return b.String()
}

// replaces the yaml team and user name with a quoted string
// these values are not quoted in the original yaml and most certainly cause issues
func (irsdk *Irsdk) RepairedYaml(s string) string {
	work := s
	for _, key := range []string{"TeamName", "UserName"} {
		re := regexp.MustCompile(fmt.Sprintf("%s: (.*)", key))
		work = re.ReplaceAllString(work, fmt.Sprintf("%s: \"$1\"", key))
	}
	return work
}

// the yaml is considered valid if it contains an entry for the pace car ;)
func (irsdk *Irsdk) validYamlData(yamlData *yaml.IrsdkYaml) bool {
	if yamlData == nil {
		return false
	}
	foundPaceCar := false
	for i := range yamlData.DriverInfo.Drivers {
		if yamlData.DriverInfo.Drivers[i].CarIsPaceCar > 0 {
			foundPaceCar = true
		}
	}

	return foundPaceCar
}

func (irsdk *Irsdk) GetValue(name string) (any, error) {
	return irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
}

//nolint:exhaustive // by design
func (irsdk *Irsdk) GetIntValue(name string) (int32, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return 0, err
	}
	switch irsdk.vHeaderLookup[name].Type {
	case IrsdkTypeBitField, IrsdkTypeInt:
		return v.(int32), nil
	}
	return 0, ErrNoMatchingDataType
}

//nolint:exhaustive // by design
func (irsdk *Irsdk) GetIntValues(name string) ([]int32, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return nil, err
	}
	if irsdk.vHeaderLookup[name].Count == 1 {
		return nil, ErrNoMatchingDataType
	}
	switch irsdk.vHeaderLookup[name].Type {
	case IrsdkTypeBitField, IrsdkTypeInt:
		return v.([]int32), nil
	}
	return nil, ErrNoMatchingDataType
}

//nolint:exhaustive,gocritic // by design
func (irsdk *Irsdk) GetFloatValue(name string) (float32, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return 0, err
	}
	switch irsdk.vHeaderLookup[name].Type {
	case IrsdkTypeFloat:
		return v.(float32), nil
	}
	return 0, ErrNoMatchingDataType
}

//nolint:exhaustive,gocritic // by design
func (irsdk *Irsdk) GetDoubleValue(name string) (float64, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return 0, err
	}
	switch irsdk.vHeaderLookup[name].Type {
	case IrsdkTypeDouble:
		return v.(float64), nil
	}
	return 0, ErrNoMatchingDataType
}

func (irsdk *Irsdk) GetDoubleValues(name string) ([]float64, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return nil, err
	}
	if irsdk.vHeaderLookup[name].Type == IrsdkTypeDouble &&
		irsdk.vHeaderLookup[name].Count > 1 {

		return v.([]float64), nil
	}
	return nil, ErrNoMatchingDataType
}

func (irsdk *Irsdk) GetFloatValues(name string) ([]float32, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return nil, err
	}
	if irsdk.vHeaderLookup[name].Type == IrsdkTypeFloat &&
		irsdk.vHeaderLookup[name].Count > 1 {

		return v.([]float32), nil
	}
	return nil, ErrNoMatchingDataType
}

//nolint:exhaustive,gocritic // by design
func (irsdk *Irsdk) GetBoolValue(name string) (bool, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return false, err
	}
	switch irsdk.vHeaderLookup[name].Type {
	case IrsdkTypeBool:
		return v.(bool), nil
	}
	return false, ErrNoMatchingDataType
}

func (irsdk *Irsdk) GetBoolValues(name string) ([]bool, error) {
	v, err := irsdk.getValueFromBuf(name, irsdk.latestVarBuffer)
	if err != nil {
		return nil, err
	}
	if irsdk.vHeaderLookup[name].Type == IrsdkTypeBool && irsdk.vHeaderLookup[name].Count > 1 {
		return v.([]bool), nil
	}
	return nil, ErrNoMatchingDataType
}

func (irsdk *Irsdk) GetValueKeys() []string {
	keys := make([]string, 0, len(irsdk.vHeaderLookup))
	for k := range irsdk.vHeaderLookup {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func (irsdk *Irsdk) GetVarHeaders() []VarHeader {
	ret := make([]VarHeader, 0, len(irsdk.vHeaderLookup))
	for _, v := range irsdk.vHeaderLookup {
		ret = append(ret, v)
	}
	slices.SortFunc(ret, func(i, j VarHeader) int {
		return strings.Compare(string(i.Name[:]), string(j.Name[:]))
	})
	return ret
}

func (irsdk *Irsdk) WriteDump(w io.Writer) (int, error) {
	return w.Write(irsdk.sharedMem)
}

//nolint:gocritic // by design
func (irsdk *Irsdk) DumpHeaders() {
	fmt.Printf("Header\n%+v\n", irsdk.hdr)
	checkForNullbytes := func(buf []byte) bool {
		for i, b := range buf {
			if b != 0 {
				fmt.Printf("i: %d, b: %v", i, b)
				return false
			}
		}
		return true
	}

	for i, v := range irsdk.hdr.VarBufs {
		fmt.Printf("VarBuf: %d TickCount: %d allZero: %v\n%s\n",
			i, v.TickCount,
			checkForNullbytes(irsdk.sharedMem[v.BufOffset:v.BufOffset+irsdk.hdr.BufLen]),
			// checkForNullbytes(irsdk.sharedMem[v.BufOffset:v.BufOffset+100]),
			hex.Dump(irsdk.sharedMem[v.BufOffset:v.BufOffset+100]))
		for j, h := range irsdk.varHeaders {
			name := extractCStringFromSlice(h.Name[:])
			val, _ := irsdk.getValueFromBuf(name,
				irsdk.sharedMem[v.BufOffset:v.BufOffset+irsdk.hdr.BufLen])
			fmt.Printf("%3d:%-32s: %v\n", j, name, val)
		}
	}
}

// convenience function to for BroadcastCamSwitchPos
func (irsdk *Irsdk) CamSwitchPos(position, group, camera int32) error {
	return irsdk.BroadcastMsg(BroadcastCamSwitchPos, position, group, camera)
}

// convenience function to for BroadcastCamSwitchNum
//
//nolint:lll // better readability
func (irsdk *Irsdk) CamSwitchNum(carNum string, group, camera int32) error {
	return irsdk.BroadcastMsg(BroadcastCamSwitchNum, irsdk.padCarNum(carNum), group, camera)
}

// convenience function to for BroadcastCamSetState
func (irsdk *Irsdk) CamSetState(mode CameraState) error {
	return irsdk.BroadcastMsg(BroadcastCamSetState, int32(mode), 0, 0)
}

// convenience function to for BroadcastReplaySetPlaySpeed
func (irsdk *Irsdk) ReplaySetPlaySpeed(speed int32, slowMotion bool) error {
	b := 0
	if slowMotion {
		b = 1
	}
	return irsdk.BroadcastMsg(BroadcastReplaySetPlaySpeed, speed, int32(b), 0)
}

// convenience function to for BroadcastReplaySetPlayPosition
func (irsdk *Irsdk) ReplaySetPlayPosition(mode ReplayPosMode, frameNum int32) error {
	return irsdk.BroadcastMsg(BroadcastReplaySetPlayPosition, int32(mode), frameNum, 0)
}

// convenience function to for BroadcaseReplaySearch
func (irsdk *Irsdk) ReplaySearch(mode ReplaySearchMode) error {
	return irsdk.BroadcastMsg(BroadcastReplaySearch, int32(mode), 0, 0)
}

// convenience function to for BroadcaseReplaySetState
func (irsdk *Irsdk) ReplaySetState(mode ReplayStateMode) error {
	return irsdk.BroadcastMsg(BroadcastReplaySetState, int32(mode), 0, 0)
}

// convenience function to for BroadcastReloadTextures
func (irsdk *Irsdk) ReloadAllTextures() error {
	return irsdk.BroadcastMsg(BroadcastReloadTextures, int32(ReloadTexturesModeAll), 0, 0)
}

// convenience function to for BroadcastReloadTextures
//
//nolint:lll // better readability
func (irsdk *Irsdk) ReloadTexture(carIdx int32) error {
	return irsdk.BroadcastMsg(BroadcastReloadTextures, int32(ReloadTexturesModeCarIdx), carIdx, 0)
}

// convenience function to for BroadcastChatComand
func (irsdk *Irsdk) ChatCommand() error {
	return irsdk.BroadcastMsg(BroadcastChatComand, int32(ChatCommandBeginChat), 0, 0)
}

// convenience function to for BroadcastChatComand
func (irsdk *Irsdk) ChatCommandMacro(macroNum int32) error {
	return irsdk.BroadcastMsg(BroadcastChatComand, int32(ChatCommandMacro), macroNum, 0)
}

// convenience function to for BroadcastPitComand
func (irsdk *Irsdk) PitCommand(cmd PitCommand, val int32) error {
	return irsdk.BroadcastMsg(BroadcastPitCommand, int32(cmd), val, 0)
}

// convenience function to for BroadcastTelemetryCommand
func (irsdk *Irsdk) TelemetryCommand(cmd TelemetryCommand) error {
	return irsdk.BroadcastMsg(BroadcastTelemetryCommand, int32(cmd), 0, 0)
}

// convenience function to for BroadcastReplaySearchSessionTime
//
//nolint:lll // better readability
func (irsdk *Irsdk) ReplaySearchSearchSessionTime(sessionNum, sessionTimeMs int32) error {
	return irsdk.BroadcastMsg(BroadcastReplaySearchSessionTime, sessionNum, sessionTimeMs, 0)
}

// convenience function to for BroadcastVideoCapture
func (irsdk *Irsdk) VideoCapture(cmd VideoCaptureMode) error {
	return irsdk.BroadcastMsg(BroadcastVideoCapture, int32(cmd), 0, 0)
}

func (irsdk *Irsdk) BroadcastMsg(cmd BroadcastCmd, var1, var2, var3 int32) error {
	if id := irsdk.getBroadcastId(); id != 0 {
		sendMsg := irsdk.user32.NewProc("SendNotifyMessageW")
		ret, _, err := sendMsg.Call(0xffff,
			id,
			uintptr(cmd)|uintptr(var1)<<16,
			uintptr(var2)|uintptr(var3)<<16,
		)
		if ret == 1 {
			return nil
		} else {
			return err
		}
	}
	return ErrBroadcastError
}

func (irsdk *Irsdk) getBroadcastId() uintptr {
	if irsdk.broadcastId == 0 {
		irsdk.broadcastId, _ = irsdk.setupBroadcastChannel()
	}
	return irsdk.broadcastId
}

func (irsdk *Irsdk) setupBroadcastChannel() (uintptr, error) {
	irsdk.user32 = windows.NewLazyDLL("user32.dll")
	registerWindowMessageW := irsdk.user32.NewProc("RegisterWindowMessageW")

	msgPtr2, _ := windows.UTF16PtrFromString(BroadcastMsg)
	ret, _, err := registerWindowMessageW.Call(uintptr(unsafe.Pointer(msgPtr2)))

	if ret == 0 {
		return 0, err
	}
	return ret, nil
}

// special handling for carNums starting with '0', foe example '01', '007',...
func (irsdk *Irsdk) padCarNum(carNum string) int32 {
	num, _ := strconv.Atoi(carNum)
	// count number of '0' prefixes in carNum (for carNums like '01', '007',...)
	zeros := 0
	for _, c := range carNum {
		if c == '0' {
			zeros++
		} else {
			break
		}
	}
	retVal := num
	numPlace := 1
	if num > 99 {
		numPlace = 3
	} else if num > 9 {
		numPlace = 2
	}
	if zeros > 0 {
		numPlace += zeros
		retVal = num + 1000*numPlace
	}

	return int32(retVal)
}

//nolint:lll,errcheck,gocritic // by design
func (irsdk *Irsdk) initialize() {
	buf := bytes.NewReader(irsdk.sharedMem[0:unsafe.Sizeof(irsdk.hdr)])
	err := binary.Read(buf, binary.LittleEndian, &irsdk.hdr)
	if err != nil {
		log.Fatalf("Could not read basic static header: %v\n", err)
	}

	irsdk.varHeaders = make([]VarHeader, irsdk.hdr.NumVars)
	buf = bytes.NewReader(irsdk.sharedMem[irsdk.hdr.VarOffsetHeader : irsdk.hdr.VarOffsetHeader+irsdk.hdr.NumVars*int32(unsafe.Sizeof(VarHeader{}))])
	binary.Read(buf, binary.LittleEndian, &irsdk.varHeaders)
	irsdk.vHeaderLookup = make(map[string]VarHeader, irsdk.hdr.NumVars)
	for _, v := range irsdk.varHeaders {
		irsdk.vHeaderLookup[extractCStringFromSlice(v.Name[:])] = v
	}

	irsdk.latestVarBuffer = make([]byte, irsdk.hdr.BufLen)
	irsdk.latestYamlBuffer = make([]byte, irsdk.hdr.SessionInfoLen)

	irsdk.dataDict = make(map[string]interface{})
	irsdk.readYaml()
}

//nolint:nestif,exhaustive,gocognit,cyclop // by design
func (irsdk *Irsdk) getValueFromBuf(name string, buf []byte) (any, error) {
	if v, ok := irsdk.vHeaderLookup[name]; ok {
		switch v.Type {
		case IrsdkTypeBool:
			if v.Count == 1 {
				return extractGen[bool](buf[v.Offset : v.Offset+1]), nil
			} else {
				return extractGenArray[bool](buf[v.Offset:v.Offset+v.Count], int(v.Count)), nil
			}

		case IrsdkTypeInt, IrsdkTypeBitField:
			if v.Count == 1 {
				return extractGen[int32](buf[v.Offset : v.Offset+4]), nil
			} else {
				return extractGenArray[int32](buf[v.Offset:v.Offset+4*v.Count], int(v.Count)), nil
			}
		case IrsdkTypeFloat:
			if v.Count == 1 {
				return extractGen[float32](buf[v.Offset : v.Offset+4]), nil
			} else {
				return extractGenArray[float32](buf[v.Offset:v.Offset+4*v.Count], int(v.Count)), nil
			}
		case IrsdkTypeDouble:
			if v.Count == 1 {
				return extractGen[float64](buf[v.Offset : v.Offset+8]), nil
			} else {
				return extractGenArray[float64](buf[v.Offset:v.Offset+8*v.Count], int(v.Count)), nil
			}
		default:
			return 0, ErrInvalidDataRequest
		}
	}
	return 0, ErrInvalidDataRequest
}

//nolint:lll,errcheck // by design
func (irsdk *Irsdk) readYaml() {
	yamlBuf := irsdk.sharedMem[irsdk.hdr.SessionInfoOffset : irsdk.hdr.SessionInfoOffset+irsdk.hdr.SessionInfoLen]
	var b bytes.Buffer
	wInUtf8 := transform.NewWriter(&b, charmap.Windows1252.NewDecoder())
	wInUtf8.Write([]byte(extractCStringFromSlice(yamlBuf)))
	wInUtf8.Close()

	err := goyaml.Unmarshal(b.Bytes(), &irsdk.dataDict)
	if err != nil {
		log.Printf("Error parsing yaml file: %v", err)
		return
	}
}

func openSharedMem() (MMap, windows.Handle) {
	fnPtr, _ := syscall.UTF16PtrFromString(MemMapFile)

	flProtect := uint32(windows.PAGE_READONLY)

	h, errno := windows.CreateFileMapping(
		windows.InvalidHandle,
		nil,
		flProtect,
		0,
		MemMapFileSize,
		fnPtr)
	if h == 0 {
		log.Fatal("could not open memmap file: ", errno)
	}

	addr, errno := windows.MapViewOfFile(h,
		windows.FILE_MAP_READ,
		0,
		0,
		uintptr(MemMapFileSize))
	if addr == 0 {
		log.Printf("error in MapViewOfFile: %v", errno)
	}
	m := MMap{}
	dh := m.header()
	dh.Data = addr
	dh.Len = MemMapFileSize
	dh.Cap = dh.Len
	return m, h
}

func openTelemetryDump(f *os.File) MMap {
	flProtect := uint32(windows.PAGE_READONLY)

	h, errno := windows.CreateFileMapping(
		windows.Handle(f.Fd()),
		nil,
		flProtect,
		0,
		MemMapFileSize,
		nil)
	if h == 0 {
		log.Fatal("could not open memmap file: ", errno)
	}

	addr, errno := windows.MapViewOfFile(h,
		windows.FILE_MAP_READ,
		0,
		0,
		uintptr(MemMapFileSize))
	if addr == 0 {
		log.Printf("error in MapViewOfFile: %v", errno)
	}
	m := MMap{}
	dh := m.header()
	dh.Data = addr
	dh.Len = MemMapFileSize
	dh.Cap = dh.Len
	return m
}

func openDataValidEventHandler() windows.Handle {
	strPtr, _ := syscall.UTF16PtrFromString(DataValidEventName)
	h, err := windows.OpenEvent(0x00100000, false, strPtr)
	if err != nil {
		log.Fatalf("Could not create: %v\n", err)
	}
	return h
}

func (m *MMap) header() *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(m))
}

//nolint:unused // by design
func (m *MMap) addrLen() (uintptr, uintptr) {
	header := m.header()
	return header.Data, uintptr(header.Len)
}

//nolint:errcheck //by design
func extractGen[N irValue](buf []byte) N {
	var ret N
	reader := bytes.NewReader(buf)
	binary.Read(reader, binary.LittleEndian, &ret)
	return ret
}

//nolint:errcheck //by design
func extractGenArray[N irValue](buf []byte, num int) []N {
	ret := make([]N, num)
	reader := bytes.NewReader(buf)
	binary.Read(reader, binary.LittleEndian, &ret)
	return ret
}

func extractCStringFromSlice(buf []byte) string {
	for i, v := range buf {
		if v == 0 {
			return string((buf)[:i])
		}
	}
	return ""
}

// returns true if slice contains just null-bytes
func isZeroed(buf []byte) bool {
	for _, b := range buf {
		if b != 0 {
			return false
		}
	}
	return true
}
