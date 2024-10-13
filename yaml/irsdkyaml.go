/*
  This file contains the Go struct for the YAML data of the Irsdk API.
  generated at 2023-07-22T21:27:06+02:00
*/

//nolint:tagliatelle,lll //by design
package yaml

type AeroBalanceCalc struct {
	FrontDownforce string `yaml:"FrontDownforce" json:"FrontDownforce"`
	FrontRhAtSpeed string `yaml:"FrontRhAtSpeed" json:"FrontRhAtSpeed"`
	RearRhAtSpeed  string `yaml:"RearRhAtSpeed" json:"RearRhAtSpeed"`
	WingSetting    string `yaml:"WingSetting" json:"WingSetting"`
}

type CameraInfo struct {
	Groups []Groups `yaml:"Groups" json:"Groups"`
}

type Cameras struct {
	CameraName string `yaml:"CameraName" json:"CameraName"`
	CameraNum  int    `yaml:"CameraNum" json:"CameraNum"`
}

type CarSetup struct {
	Chassis     Chassis   `yaml:"Chassis" json:"Chassis"`
	TiresAero   TiresAero `yaml:"TiresAero" json:"TiresAero"`
	UpdateCount int       `yaml:"UpdateCount" json:"UpdateCount"`
}

type Chassis struct {
	Front      Front      `yaml:"Front" json:"Front"`
	InCarDials InCarDials `yaml:"InCarDials" json:"InCarDials"`
	LeftFront  LeftFront  `yaml:"LeftFront" json:"LeftFront"`
	LeftRear   LeftRear   `yaml:"LeftRear" json:"LeftRear"`
	Rear       Rear       `yaml:"Rear" json:"Rear"`
	RightFront RightFront `yaml:"RightFront" json:"RightFront"`
	RightRear  RightRear  `yaml:"RightRear" json:"RightRear"`
}

type DriverInfo struct {
	DriverCarEngCylinderCount int       `yaml:"DriverCarEngCylinderCount" json:"DriverCarEngCylinderCount"`
	DriverCarEstLapTime       float64   `yaml:"DriverCarEstLapTime" json:"DriverCarEstLapTime"`
	DriverCarFuelKgPerLtr     float64   `yaml:"DriverCarFuelKgPerLtr" json:"DriverCarFuelKgPerLtr"`
	DriverCarFuelMaxLtr       float64   `yaml:"DriverCarFuelMaxLtr" json:"DriverCarFuelMaxLtr"`
	DriverCarGearNeutral      int       `yaml:"DriverCarGearNeutral" json:"DriverCarGearNeutral"`
	DriverCarGearNumForward   int       `yaml:"DriverCarGearNumForward" json:"DriverCarGearNumForward"`
	DriverCarGearReverse      int       `yaml:"DriverCarGearReverse" json:"DriverCarGearReverse"`
	DriverCarIdleRPM          float64   `yaml:"DriverCarIdleRPM" json:"DriverCarIdleRPM"`
	DriverCarIdx              int       `yaml:"DriverCarIdx" json:"DriverCarIdx"`
	DriverCarIsElectric       int       `yaml:"DriverCarIsElectric" json:"DriverCarIsElectric"`
	DriverCarMaxFuelPct       float64   `yaml:"DriverCarMaxFuelPct" json:"DriverCarMaxFuelPct"`
	DriverCarRedLine          float64   `yaml:"DriverCarRedLine" json:"DriverCarRedLine"`
	DriverCarSLBlinkRPM       float64   `yaml:"DriverCarSLBlinkRPM" json:"DriverCarSLBlinkRPM"`
	DriverCarSLFirstRPM       float64   `yaml:"DriverCarSLFirstRPM" json:"DriverCarSLFirstRPM"`
	DriverCarSLLastRPM        float64   `yaml:"DriverCarSLLastRPM" json:"DriverCarSLLastRPM"`
	DriverCarSLShiftRPM       float64   `yaml:"DriverCarSLShiftRPM" json:"DriverCarSLShiftRPM"`
	DriverCarVersion          string    `yaml:"DriverCarVersion" json:"DriverCarVersion"`
	DriverHeadPosX            float64   `yaml:"DriverHeadPosX" json:"DriverHeadPosX"`
	DriverHeadPosY            float64   `yaml:"DriverHeadPosY" json:"DriverHeadPosY"`
	DriverHeadPosZ            float64   `yaml:"DriverHeadPosZ" json:"DriverHeadPosZ"`
	DriverIncidentCount       int       `yaml:"DriverIncidentCount" json:"DriverIncidentCount"`
	DriverPitTrkPct           float64   `yaml:"DriverPitTrkPct" json:"DriverPitTrkPct"`
	DriverSetupIsModified     int       `yaml:"DriverSetupIsModified" json:"DriverSetupIsModified"`
	DriverSetupLoadTypeName   string    `yaml:"DriverSetupLoadTypeName" json:"DriverSetupLoadTypeName"`
	DriverSetupName           string    `yaml:"DriverSetupName" json:"DriverSetupName"`
	DriverSetupPassedTech     int       `yaml:"DriverSetupPassedTech" json:"DriverSetupPassedTech"`
	DriverUserID              int       `yaml:"DriverUserID" json:"DriverUserID"`
	Drivers                   []Drivers `yaml:"Drivers" json:"Drivers"`
	PaceCarIdx                int       `yaml:"PaceCarIdx" json:"PaceCarIdx"`
}

type Drivers struct {
	AbbrevName              string  `yaml:"AbbrevName" json:"AbbrevName"`
	BodyType                int     `yaml:"BodyType" json:"BodyType"`
	CarClassColor           int     `yaml:"CarClassColor" json:"CarClassColor"`
	CarClassDryTireSetLimit string  `yaml:"CarClassDryTireSetLimit" json:"CarClassDryTireSetLimit"`
	CarClassEstLapTime      float64 `yaml:"CarClassEstLapTime" json:"CarClassEstLapTime"`
	CarClassID              int     `yaml:"CarClassID" json:"CarClassID"`
	CarClassLicenseLevel    int     `yaml:"CarClassLicenseLevel" json:"CarClassLicenseLevel"`
	CarClassMaxFuelPct      string  `yaml:"CarClassMaxFuelPct" json:"CarClassMaxFuelPct"`
	CarClassPowerAdjust     string  `yaml:"CarClassPowerAdjust" json:"CarClassPowerAdjust"`
	CarClassRelSpeed        int     `yaml:"CarClassRelSpeed" json:"CarClassRelSpeed"`
	CarClassShortName       string  `yaml:"CarClassShortName" json:"CarClassShortName"`
	CarClassWeightPenalty   string  `yaml:"CarClassWeightPenalty" json:"CarClassWeightPenalty"`
	CarDesignStr            string  `yaml:"CarDesignStr" json:"CarDesignStr"`
	CarID                   int     `yaml:"CarID" json:"CarID"`
	CarIdx                  int     `yaml:"CarIdx" json:"CarIdx"`
	CarIsAI                 int     `yaml:"CarIsAI" json:"CarIsAI"`
	CarIsElectric           int     `yaml:"CarIsElectric" json:"CarIsElectric"`
	CarIsPaceCar            int     `yaml:"CarIsPaceCar" json:"CarIsPaceCar"`
	CarNumber               string  `yaml:"CarNumber" json:"CarNumber"`
	CarNumberDesignStr      string  `yaml:"CarNumberDesignStr" json:"CarNumberDesignStr"`
	CarNumberRaw            int     `yaml:"CarNumberRaw" json:"CarNumberRaw"`
	CarPath                 string  `yaml:"CarPath" json:"CarPath"`
	CarScreenName           string  `yaml:"CarScreenName" json:"CarScreenName"`
	CarScreenNameShort      string  `yaml:"CarScreenNameShort" json:"CarScreenNameShort"`
	CarSponsor_1            int     `yaml:"CarSponsor_1" json:"CarSponsor_1"`
	CarSponsor_2            int     `yaml:"CarSponsor_2" json:"CarSponsor_2"`
	ClubID                  int     `yaml:"ClubID" json:"ClubID"`
	ClubName                string  `yaml:"ClubName" json:"ClubName"`
	CurDriverIncidentCount  int     `yaml:"CurDriverIncidentCount" json:"CurDriverIncidentCount"`
	DivisionID              int     `yaml:"DivisionID" json:"DivisionID"`
	DivisionName            string  `yaml:"DivisionName" json:"DivisionName"`
	FaceType                int     `yaml:"FaceType" json:"FaceType"`
	HelmetDesignStr         string  `yaml:"HelmetDesignStr" json:"HelmetDesignStr"`
	HelmetType              int     `yaml:"HelmetType" json:"HelmetType"`
	Initials                string  `yaml:"Initials" json:"Initials"`
	IRating                 int     `yaml:"IRating" json:"IRating"`
	IsSpectator             int     `yaml:"IsSpectator" json:"IsSpectator"`
	LicColor                any     `yaml:"LicColor" json:"LicColor"`
	LicLevel                int     `yaml:"LicLevel" json:"LicLevel"`
	LicString               string  `yaml:"LicString" json:"LicString"`
	LicSubLevel             int     `yaml:"LicSubLevel" json:"LicSubLevel"`
	SuitDesignStr           string  `yaml:"SuitDesignStr" json:"SuitDesignStr"`
	TeamID                  int     `yaml:"TeamID" json:"TeamID"`
	TeamIncidentCount       int     `yaml:"TeamIncidentCount" json:"TeamIncidentCount"`
	TeamName                string  `yaml:"TeamName" json:"TeamName"`
	UserID                  int     `yaml:"UserID" json:"UserID"`
	UserName                string  `yaml:"UserName" json:"UserName"`
}

type Frequencies struct {
	CanScan       int    `yaml:"CanScan" json:"CanScan"`
	CanSquawk     int    `yaml:"CanSquawk" json:"CanSquawk"`
	CarIdx        int    `yaml:"CarIdx" json:"CarIdx"`
	ClubID        int    `yaml:"ClubID" json:"ClubID"`
	EntryIdx      int    `yaml:"EntryIdx" json:"EntryIdx"`
	FrequencyName string `yaml:"FrequencyName" json:"FrequencyName"`
	FrequencyNum  int    `yaml:"FrequencyNum" json:"FrequencyNum"`
	IsDeletable   int    `yaml:"IsDeletable" json:"IsDeletable"`
	IsMutable     int    `yaml:"IsMutable" json:"IsMutable"`
	Muted         int    `yaml:"Muted" json:"Muted"`
	Priority      int    `yaml:"Priority" json:"Priority"`
}

type Front struct {
	ArbBlades        string `yaml:"ArbBlades" json:"ArbBlades"`
	ArbOuterDiameter string `yaml:"ArbOuterDiameter" json:"ArbOuterDiameter"`
	BrakePads        string `yaml:"BrakePads" json:"BrakePads"`
	CrossWeight      string `yaml:"CrossWeight" json:"CrossWeight"`
	FrontMasterCyl   string `yaml:"FrontMasterCyl" json:"FrontMasterCyl"`
	RearMasterCyl    string `yaml:"RearMasterCyl" json:"RearMasterCyl"`
	ToeIn            string `yaml:"ToeIn" json:"ToeIn"`
}

type Groups struct {
	Cameras   []Cameras `yaml:"Cameras" json:"Cameras"`
	GroupName string    `yaml:"GroupName" json:"GroupName"`
	GroupNum  int       `yaml:"GroupNum" json:"GroupNum"`
}

type InCarDials struct {
	AbsSetting             string `yaml:"AbsSetting" json:"AbsSetting"`
	BrakePressureBias      string `yaml:"BrakePressureBias" json:"BrakePressureBias"`
	ThrottleShapeSetting   string `yaml:"ThrottleShapeSetting" json:"ThrottleShapeSetting"`
	TractionControlSetting string `yaml:"TractionControlSetting" json:"TractionControlSetting"`
}

type IrsdkYaml struct {
	CameraInfo         CameraInfo         `yaml:"CameraInfo" json:"CameraInfo"`
	CarSetup           CarSetup           `yaml:"CarSetup" json:"CarSetup"`
	DriverInfo         DriverInfo         `yaml:"DriverInfo" json:"DriverInfo"`
	QualifyResultsInfo QualifyResultsInfo `yaml:"QualifyResultsInfo" json:"QualifyResultsInfo"`
	RadioInfo          RadioInfo          `yaml:"RadioInfo" json:"RadioInfo"`
	SessionInfo        SessionInfo        `yaml:"SessionInfo" json:"SessionInfo"`
	SplitTimeInfo      SplitTimeInfo      `yaml:"SplitTimeInfo" json:"SplitTimeInfo"`
	WeekendInfo        WeekendInfo        `yaml:"WeekendInfo" json:"WeekendInfo"`
}

type LeftFront struct {
	Camber            string `yaml:"Camber" json:"Camber"`
	Caster            string `yaml:"Caster" json:"Caster"`
	CompDamping       string `yaml:"CompDamping" json:"CompDamping"`
	CornerWeight      string `yaml:"CornerWeight" json:"CornerWeight"`
	RbdDamping        string `yaml:"RbdDamping" json:"RbdDamping"`
	RideHeight        string `yaml:"RideHeight" json:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset" json:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate" json:"SpringRate"`
	SpringSelected    int    `yaml:"SpringSelected" json:"SpringSelected"`
}

type LeftRear struct {
	Camber            string `yaml:"Camber" json:"Camber"`
	CompDamping       string `yaml:"CompDamping" json:"CompDamping"`
	CornerWeight      string `yaml:"CornerWeight" json:"CornerWeight"`
	RbdDamping        string `yaml:"RbdDamping" json:"RbdDamping"`
	RideHeight        string `yaml:"RideHeight" json:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset" json:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate" json:"SpringRate"`
	SpringSelected    int    `yaml:"SpringSelected" json:"SpringSelected"`
	ToeIn             string `yaml:"ToeIn" json:"ToeIn"`
}

type QualifyResultsInfo struct {
	Results []Results `yaml:"Results" json:"Results"`
}

type RadioInfo struct {
	Radios           []Radios `yaml:"Radios" json:"Radios"`
	SelectedRadioNum int      `yaml:"SelectedRadioNum" json:"SelectedRadioNum"`
}

type Radios struct {
	Frequencies         []Frequencies `yaml:"Frequencies" json:"Frequencies"`
	HopCount            int           `yaml:"HopCount" json:"HopCount"`
	NumFrequencies      int           `yaml:"NumFrequencies" json:"NumFrequencies"`
	RadioNum            int           `yaml:"RadioNum" json:"RadioNum"`
	ScanningIsOn        int           `yaml:"ScanningIsOn" json:"ScanningIsOn"`
	TunedToFrequencyNum int           `yaml:"TunedToFrequencyNum" json:"TunedToFrequencyNum"`
}

type Rear struct {
	ArbBlades        string `yaml:"ArbBlades" json:"ArbBlades"`
	ArbOuterDiameter string `yaml:"ArbOuterDiameter" json:"ArbOuterDiameter"`
	DiffPreload      string `yaml:"DiffPreload" json:"DiffPreload"`
	FrictionFaces    int    `yaml:"FrictionFaces" json:"FrictionFaces"`
	FuelLevel        string `yaml:"FuelLevel" json:"FuelLevel"`
	SixthGear        string `yaml:"SixthGear" json:"SixthGear"`
	WingSetting      string `yaml:"WingSetting" json:"WingSetting"`
}

type Results struct {
	CarIdx        int     `yaml:"CarIdx" json:"CarIdx"`
	ClassPosition int     `yaml:"ClassPosition" json:"ClassPosition"`
	FastestLap    int     `yaml:"FastestLap" json:"FastestLap"`
	FastestTime   float64 `yaml:"FastestTime" json:"FastestTime"`
	Position      int     `yaml:"Position" json:"Position"`
}

type ResultsFastestLap struct {
	CarIdx      int     `yaml:"CarIdx" json:"CarIdx"`
	FastestLap  int     `yaml:"FastestLap" json:"FastestLap"`
	FastestTime float64 `yaml:"FastestTime" json:"FastestTime"`
}

type ResultsPositions struct {
	CarIdx            int     `yaml:"CarIdx" json:"CarIdx"`
	ClassPosition     int     `yaml:"ClassPosition" json:"ClassPosition"`
	FastestLap        int     `yaml:"FastestLap" json:"FastestLap"`
	FastestTime       float64 `yaml:"FastestTime" json:"FastestTime"`
	Incidents         int     `yaml:"Incidents" json:"Incidents"`
	JokerLapsComplete int     `yaml:"JokerLapsComplete" json:"JokerLapsComplete"`
	Lap               int     `yaml:"Lap" json:"Lap"`
	LapsComplete      int     `yaml:"LapsComplete" json:"LapsComplete"`
	LapsDriven        float64 `yaml:"LapsDriven" json:"LapsDriven"`
	LapsLed           int     `yaml:"LapsLed" json:"LapsLed"`
	LastTime          float64 `yaml:"LastTime" json:"LastTime"`
	Position          int     `yaml:"Position" json:"Position"`
	ReasonOutId       int     `yaml:"ReasonOutId" json:"ReasonOutId"`
	ReasonOutStr      string  `yaml:"ReasonOutStr" json:"ReasonOutStr"`
	Time              float64 `yaml:"Time" json:"Time"`
}

type RightFront struct {
	Camber            string `yaml:"Camber" json:"Camber"`
	Caster            string `yaml:"Caster" json:"Caster"`
	CompDamping       string `yaml:"CompDamping" json:"CompDamping"`
	CornerWeight      string `yaml:"CornerWeight" json:"CornerWeight"`
	RbdDamping        string `yaml:"RbdDamping" json:"RbdDamping"`
	RideHeight        string `yaml:"RideHeight" json:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset" json:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate" json:"SpringRate"`
	SpringSelected    int    `yaml:"SpringSelected" json:"SpringSelected"`
}

type RightRear struct {
	Camber            string `yaml:"Camber" json:"Camber"`
	CompDamping       string `yaml:"CompDamping" json:"CompDamping"`
	CornerWeight      string `yaml:"CornerWeight" json:"CornerWeight"`
	RbdDamping        string `yaml:"RbdDamping" json:"RbdDamping"`
	RideHeight        string `yaml:"RideHeight" json:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset" json:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate" json:"SpringRate"`
	SpringSelected    int    `yaml:"SpringSelected" json:"SpringSelected"`
	ToeIn             string `yaml:"ToeIn" json:"ToeIn"`
}

type Sectors struct {
	SectorNum      int     `yaml:"SectorNum" json:"SectorNum"`
	SectorStartPct float64 `yaml:"SectorStartPct" json:"SectorStartPct"`
}

type SessionInfo struct {
	Sessions []Sessions `yaml:"Sessions" json:"Sessions"`
}

type Sessions struct {
	// Note: QualifyPositions is only used in Heat races
	QualifyPositions                 []Results           `yaml:"QualifyPositions" json:"QualifyPositions"`
	ResultsAverageLapTime            float64             `yaml:"ResultsAverageLapTime" json:"ResultsAverageLapTime"`
	ResultsFastestLap                []ResultsFastestLap `yaml:"ResultsFastestLap" json:"ResultsFastestLap"`
	ResultsLapsComplete              int                 `yaml:"ResultsLapsComplete" json:"ResultsLapsComplete"`
	ResultsNumCautionFlags           int                 `yaml:"ResultsNumCautionFlags" json:"ResultsNumCautionFlags"`
	ResultsNumCautionLaps            int                 `yaml:"ResultsNumCautionLaps" json:"ResultsNumCautionLaps"`
	ResultsNumLeadChanges            int                 `yaml:"ResultsNumLeadChanges" json:"ResultsNumLeadChanges"`
	ResultsOfficial                  int                 `yaml:"ResultsOfficial" json:"ResultsOfficial"`
	ResultsPositions                 []ResultsPositions  `yaml:"ResultsPositions" json:"ResultsPositions"`
	SessionEnforceTireCompoundChange int                 `yaml:"SessionEnforceTireCompoundChange" json:"SessionEnforceTireCompoundChange"`
	SessionLaps                      string              `yaml:"SessionLaps" json:"SessionLaps"`
	SessionName                      string              `yaml:"SessionName" json:"SessionName"`
	SessionNum                       int                 `yaml:"SessionNum" json:"SessionNum"`
	SessionNumLapsToAvg              int                 `yaml:"SessionNumLapsToAvg" json:"SessionNumLapsToAvg"`
	SessionRunGroupsUsed             int                 `yaml:"SessionRunGroupsUsed" json:"SessionRunGroupsUsed"`
	SessionSkipped                   int                 `yaml:"SessionSkipped" json:"SessionSkipped"`
	SessionTime                      string              `yaml:"SessionTime" json:"SessionTime"`
	SessionTrackRubberState          string              `yaml:"SessionTrackRubberState" json:"SessionTrackRubberState"`
	SessionType                      string              `yaml:"SessionType" json:"SessionType"`
}

type SplitTimeInfo struct {
	Sectors []Sectors `yaml:"Sectors" json:"Sectors"`
}

type TelemetryOptions struct {
	TelemetryDiskFile string `yaml:"TelemetryDiskFile" json:"TelemetryDiskFile"`
}

type TiresAero struct {
	AeroBalanceCalc AeroBalanceCalc `yaml:"AeroBalanceCalc" json:"AeroBalanceCalc"`
}

type WeekendInfo struct {
	BuildTarget            string           `yaml:"BuildTarget" json:"BuildTarget"`
	BuildType              string           `yaml:"BuildType" json:"BuildType"`
	BuildVersion           string           `yaml:"BuildVersion" json:"BuildVersion"`
	Category               string           `yaml:"Category" json:"Category"`
	DCRuleSet              string           `yaml:"DCRuleSet" json:"DCRuleSet"`
	EventType              string           `yaml:"EventType" json:"EventType"`
	HeatRacing             int              `yaml:"HeatRacing" json:"HeatRacing"`
	LeagueID               int              `yaml:"LeagueID" json:"LeagueID"`
	MaxDrivers             int              `yaml:"MaxDrivers" json:"MaxDrivers"`
	MinDrivers             int              `yaml:"MinDrivers" json:"MinDrivers"`
	NumCarClasses          int              `yaml:"NumCarClasses" json:"NumCarClasses"`
	NumCarTypes            int              `yaml:"NumCarTypes" json:"NumCarTypes"`
	Official               int              `yaml:"Official" json:"Official"`
	QualifierMustStartRace int              `yaml:"QualifierMustStartRace" json:"QualifierMustStartRace"`
	RaceWeek               int              `yaml:"RaceWeek" json:"RaceWeek"`
	SeasonID               int              `yaml:"SeasonID" json:"SeasonID"`
	SeriesID               int              `yaml:"SeriesID" json:"SeriesID"`
	SessionID              int              `yaml:"SessionID" json:"SessionID"`
	SimMode                string           `yaml:"SimMode" json:"SimMode"`
	SubSessionID           int              `yaml:"SubSessionID" json:"SubSessionID"`
	TeamRacing             int              `yaml:"TeamRacing" json:"TeamRacing"`
	TelemetryOptions       TelemetryOptions `yaml:"TelemetryOptions" json:"TelemetryOptions"`
	TrackAirPressure       string           `yaml:"TrackAirPressure" json:"TrackAirPressure"`
	TrackAirTemp           string           `yaml:"TrackAirTemp" json:"TrackAirTemp"`
	TrackAltitude          string           `yaml:"TrackAltitude" json:"TrackAltitude"`
	TrackCity              string           `yaml:"TrackCity" json:"TrackCity"`
	TrackCleanup           int              `yaml:"TrackCleanup" json:"TrackCleanup"`
	TrackConfigName        string           `yaml:"TrackConfigName" json:"TrackConfigName"`
	TrackCountry           string           `yaml:"TrackCountry" json:"TrackCountry"`
	TrackDirection         string           `yaml:"TrackDirection" json:"TrackDirection"`
	TrackDisplayName       string           `yaml:"TrackDisplayName" json:"TrackDisplayName"`
	TrackDisplayShortName  string           `yaml:"TrackDisplayShortName" json:"TrackDisplayShortName"`
	TrackDynamicTrack      int              `yaml:"TrackDynamicTrack" json:"TrackDynamicTrack"`
	TrackFogLevel          string           `yaml:"TrackFogLevel" json:"TrackFogLevel"`
	TrackID                int              `yaml:"TrackID" json:"TrackID"`
	TrackLatitude          string           `yaml:"TrackLatitude" json:"TrackLatitude"`
	TrackLength            string           `yaml:"TrackLength" json:"TrackLength"`
	TrackLengthOfficial    string           `yaml:"TrackLengthOfficial" json:"TrackLengthOfficial"`
	TrackLongitude         string           `yaml:"TrackLongitude" json:"TrackLongitude"`
	TrackName              string           `yaml:"TrackName" json:"TrackName"`
	TrackNorthOffset       string           `yaml:"TrackNorthOffset" json:"TrackNorthOffset"`
	TrackNumTurns          int              `yaml:"TrackNumTurns" json:"TrackNumTurns"`
	TrackPitSpeedLimit     string           `yaml:"TrackPitSpeedLimit" json:"TrackPitSpeedLimit"`
	TrackRelativeHumidity  string           `yaml:"TrackRelativeHumidity" json:"TrackRelativeHumidity"`
	TrackSkies             string           `yaml:"TrackSkies" json:"TrackSkies"`
	TrackSurfaceTemp       string           `yaml:"TrackSurfaceTemp" json:"TrackSurfaceTemp"`
	TrackType              string           `yaml:"TrackType" json:"TrackType"`
	TrackVersion           string           `yaml:"TrackVersion" json:"TrackVersion"`
	TrackWeatherType       string           `yaml:"TrackWeatherType" json:"TrackWeatherType"`
	TrackWindDir           string           `yaml:"TrackWindDir" json:"TrackWindDir"`
	TrackWindVel           string           `yaml:"TrackWindVel" json:"TrackWindVel"`
	WeekendOptions         WeekendOptions   `yaml:"WeekendOptions" json:"WeekendOptions"`
}

type WeekendOptions struct {
	CommercialMode             string `yaml:"CommercialMode" json:"CommercialMode"`
	CourseCautions             string `yaml:"CourseCautions" json:"CourseCautions"`
	EarthRotationSpeedupFactor int    `yaml:"EarthRotationSpeedupFactor" json:"EarthRotationSpeedupFactor"`
	FastRepairsLimit           string `yaml:"FastRepairsLimit" json:"FastRepairsLimit"`
	FogLevel                   string `yaml:"FogLevel" json:"FogLevel"`
	GreenWhiteCheckeredLimit   int    `yaml:"GreenWhiteCheckeredLimit" json:"GreenWhiteCheckeredLimit"`
	HardcoreLevel              int    `yaml:"HardcoreLevel" json:"HardcoreLevel"`
	HasOpenRegistration        int    `yaml:"HasOpenRegistration" json:"HasOpenRegistration"`
	IncidentLimit              string `yaml:"IncidentLimit" json:"IncidentLimit"`
	IsFixedSetup               int    `yaml:"IsFixedSetup" json:"IsFixedSetup"`
	NightMode                  string `yaml:"NightMode" json:"NightMode"`
	NumJokerLaps               int    `yaml:"NumJokerLaps" json:"NumJokerLaps"`
	NumStarters                int    `yaml:"NumStarters" json:"NumStarters"`
	QualifyScoring             string `yaml:"QualifyScoring" json:"QualifyScoring"`
	RelativeHumidity           string `yaml:"RelativeHumidity" json:"RelativeHumidity"`
	Restarts                   string `yaml:"Restarts" json:"Restarts"`
	ShortParadeLap             int    `yaml:"ShortParadeLap" json:"ShortParadeLap"`
	Skies                      string `yaml:"Skies" json:"Skies"`
	StandingStart              int    `yaml:"StandingStart" json:"StandingStart"`
	StartingGrid               string `yaml:"StartingGrid" json:"StartingGrid"`
	StrictLapsChecking         string `yaml:"StrictLapsChecking" json:"StrictLapsChecking"`
	TimeOfDay                  string `yaml:"TimeOfDay" json:"TimeOfDay"`
	Unofficial                 int    `yaml:"Unofficial" json:"Unofficial"`
	WeatherTemp                string `yaml:"WeatherTemp" json:"WeatherTemp"`
	WeatherType                string `yaml:"WeatherType" json:"WeatherType"`
	WindDirection              string `yaml:"WindDirection" json:"WindDirection"`
	WindSpeed                  string `yaml:"WindSpeed" json:"WindSpeed"`
}
