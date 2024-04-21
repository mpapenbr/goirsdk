# Variables

| Name                            | Desc                                                            | Unit                 | Type         |
| ------------------------------- | --------------------------------------------------------------- | -------------------- | ------------ |
| AirDensity                      | Density of air at start/finish line                             | kg/m^3               | float        |
| AirPressure                     | Pressure of air at start/finish line                            | Pa                   | float        |
| AirTemp                         | Temperature of air at start/finish line                         | C                    | float        |
| Brake                           | 0=brake released to 1=max pedal force                           | %                    | float        |
| BrakeABSactive                  | true if abs is currently reducing brake force pressure          |                      | bool         |
| BrakeRaw                        | Raw brake input 0=brake released to 1=max pedal force           | %                    | float        |
| CFshockDefl                     | CF shock deflection                                             | m                    | float        |
| CFshockDefl_ST                  | CF shock deflection at 360 Hz                                   | m                    | float[6]     |
| CFshockVel                      | CF shock velocity                                               | m/s                  | float        |
| CFshockVel_ST                   | CF shock velocity at 360 Hz                                     | m/s                  | float[6]     |
| CRshockDefl                     | CR shock deflection                                             | m                    | float        |
| CRshockDefl_ST                  | CR shock deflection at 360 Hz                                   | m                    | float[6]     |
| CRshockVel                      | CR shock velocity                                               | m/s                  | float        |
| CRshockVel_ST                   | CR shock velocity at 360 Hz                                     | m/s                  | float[6]     |
| CamCameraNumber                 | Active camera number                                            |                      | int          |
| CamCameraState                  | State of camera system                                          | irsdk_CameraState    | bitfield     |
| CamCarIdx                       | Active camera's focus car index                                 |                      | int          |
| CamGroupNumber                  | Active camera group number                                      |                      | int          |
| CarIdxBestLapNum                | Cars best lap number                                            |                      | int[64]      |
| CarIdxBestLapTime               | Cars best lap time                                              | s                    | float[64]    |
| CarIdxClass                     | Cars class id by car index                                      |                      | int[64]      |
| CarIdxClassPosition             | Cars class position in race by car index                        |                      | int[64]      |
| CarIdxEstTime                   | Estimated time to reach current location on track               | s                    | float[64]    |
| CarIdxF2Time                    | Race time behind leader or fastest lap time otherwise           | s                    | float[64]    |
| CarIdxFastRepairsUsed           | How many fast repairs each car has used                         |                      | int[64]      |
| CarIdxGear                      | -1=reverse 0=neutral 1..n=current gear by car index             |                      | int[64]      |
| CarIdxLap                       | Laps started by car index                                       |                      | int[64]      |
| CarIdxLapCompleted              | Laps completed by car index                                     |                      | int[64]      |
| CarIdxLapDistPct                | Percentage distance around lap by car index                     | %                    | float[64]    |
| CarIdxLastLapTime               | Cars last lap time                                              | s                    | float[64]    |
| CarIdxOnPitRoad                 | On pit road between the cones by car index                      |                      | bool[64]     |
| CarIdxP2P_Count                 | Push2Pass count of usage (or remaining in Race)                 |                      | int[64]      |
| CarIdxP2P_Status                | Push2Pass active or not                                         |                      | bool[64]     |
| CarIdxPaceFlags                 | Pacing status flags for each car                                | irsdk_PaceFlags      | bitfield[64] |
| CarIdxPaceLine                  | What line cars are pacing in or -1 if not pacing                |                      | int[64]      |
| CarIdxPaceRow                   | What row cars are pacing in or -1 if not pacing                 |                      | int[64]      |
| CarIdxPosition                  | Cars position in race by car index                              |                      | int[64]      |
| CarIdxQualTireCompound          | Cars Qual tire compound                                         |                      | int[64]      |
| CarIdxQualTireCompoundLocked    | Cars Qual tire compound is locked-in                            |                      | bool[64]     |
| CarIdxRPM                       | Engine rpm by car index                                         | revs/min             | float[64]    |
| CarIdxSessionFlags              | Session flags for each player                                   | irsdk_Flags          | bitfield[64] |
| CarIdxSteer                     | Steering wheel angle by car index                               | rad                  | float[64]    |
| CarIdxTireCompound              | Cars current tire compound                                      |                      | int[64]      |
| CarIdxTrackSurface              | Track surface type by car index                                 | irsdk_TrkLoc         | int[64]      |
| CarIdxTrackSurfaceMaterial      | Track surface material type by car index                        | irsdk_TrkSurf        | int[64]      |
| CarLeftRight                    | Notify if car is to the left or right of driver                 | irsdk_CarLeftRight   | int          |
| ChanAvgLatency                  | Communications average latency                                  | s                    | float        |
| ChanClockSkew                   | Communications server clock skew                                | s                    | float        |
| ChanLatency                     | Communications latency                                          | s                    | float        |
| ChanPartnerQuality              | Partner communications quality                                  | %                    | float        |
| ChanQuality                     | Communications quality                                          | %                    | float        |
| Clutch                          | 0=disengaged to 1=fully engaged                                 | %                    | float        |
| ClutchRaw                       | Raw clutch input 0=disengaged to 1=fully engaged                | %                    | float        |
| CpuUsageBG                      | Percent of available tim bg thread took with a 1 sec avg        | %                    | float        |
| CpuUsageFG                      | Percent of available tim fg thread took with a 1 sec avg        | %                    | float        |
| DCDriversSoFar                  | Number of team drivers who have run a stint                     |                      | int          |
| DCLapStatus                     | Status of driver change lap requirements                        |                      | int          |
| DisplayUnits                    | Default units for the user interface 0 = english 1 = metric     |                      | int          |
| DriverMarker                    | Driver activated flag                                           |                      | bool         |
| Engine0_RPM                     | Engine0Engine rpm                                               | revs/min             | float        |
| EngineWarnings                  | Bitfield for warning lights                                     | irsdk_EngineWarnings | bitfield     |
| EnterExitReset                  | Indicate action the reset key will take 0 enter 1 exit 2 reset  |                      | int          |
| FastRepairAvailable             | How many fast repairs left 255 is unlimited                     |                      | int          |
| FastRepairUsed                  | How many fast repairs used so far                               |                      | int          |
| FogLevel                        | Fog level at start/finish line                                  | %                    | float        |
| FrameRate                       | Average frames per second                                       | fps                  | float        |
| FrontTireSetsAvailable          | How many front tire sets are remaining 255 is unlimited         |                      | int          |
| FrontTireSetsUsed               | How many front tire sets used so far                            |                      | int          |
| FuelLevel                       | Liters of fuel remaining                                        | l                    | float        |
| FuelLevelPct                    | Percent fuel remaining                                          | %                    | float        |
| FuelPress                       | Engine fuel pressure                                            | bar                  | float        |
| FuelUsePerHour                  | Engine fuel used instantaneous                                  | kg/h                 | float        |
| Gear                            | -1=reverse 0=neutral 1..n=current gear                          |                      | int          |
| GpuUsage                        | Percent of available tim gpu took with a 1 sec avg              | %                    | float        |
| HandbrakeRaw                    | Raw handbrake input 0=handbrake released to 1=max force         | %                    | float        |
| IsDiskLoggingActive             | 0=disk based telemetry file not being written 1=being written   |                      | bool         |
| IsDiskLoggingEnabled            | 0=disk based telemetry turned off 1=turned on                   |                      | bool         |
| IsGarageVisible                 | 1=Garage screen is visible                                      |                      | bool         |
| IsInGarage                      | 1=Car in garage physics running                                 |                      | bool         |
| IsOnTrack                       | 1=Car on track physics running with player in car               |                      | bool         |
| IsOnTrackCar                    | 1=Car on track physics running                                  |                      | bool         |
| IsReplayPlaying                 | 0=replay not playing 1=replay playing                           |                      | bool         |
| LFTiresAvailable                | How many left front tires are remaining 255 is unlimited        |                      | int          |
| LFTiresUsed                     | How many left front tires used so far                           |                      | int          |
| LFbrakeLinePress                | LF brake line pressure                                          | bar                  | float        |
| LFcoldPressure                  | LF tire cold pressure as set in the garage                      | kPa                  | float        |
| LFshockDefl                     | LF shock deflection                                             | m                    | float        |
| LFshockDefl_ST                  | LF shock deflection at 360 Hz                                   | m                    | float[6]     |
| LFshockVel                      | LF shock velocity                                               | m/s                  | float        |
| LFshockVel_ST                   | LF shock velocity at 360 Hz                                     | m/s                  | float[6]     |
| LFtempCL                        | LF tire left carcass temperature                                | C                    | float        |
| LFtempCM                        | LF tire middle carcass temperature                              | C                    | float        |
| LFtempCR                        | LF tire right carcass temperature                               | C                    | float        |
| LFwearL                         | LF tire left percent tread remaining                            | %                    | float        |
| LFwearM                         | LF tire middle percent tread remaining                          | %                    | float        |
| LFwearR                         | LF tire right percent tread remaining                           | %                    | float        |
| LRTiresAvailable                | How many left rear tires are remaining 255 is unlimited         |                      | int          |
| LRTiresUsed                     | How many left rear tires used so far                            |                      | int          |
| LRbrakeLinePress                | LR brake line pressure                                          | bar                  | float        |
| LRcoldPressure                  | LR tire cold pressure as set in the garage                      | kPa                  | float        |
| LRshockDefl                     | LR shock deflection                                             | m                    | float        |
| LRshockDefl_ST                  | LR shock deflection at 360 Hz                                   | m                    | float[6]     |
| LRshockVel                      | LR shock velocity                                               | m/s                  | float        |
| LRshockVel_ST                   | LR shock velocity at 360 Hz                                     | m/s                  | float[6]     |
| LRtempCL                        | LR tire left carcass temperature                                | C                    | float        |
| LRtempCM                        | LR tire middle carcass temperature                              | C                    | float        |
| LRtempCR                        | LR tire right carcass temperature                               | C                    | float        |
| LRwearL                         | LR tire left percent tread remaining                            | %                    | float        |
| LRwearM                         | LR tire middle percent tread remaining                          | %                    | float        |
| LRwearR                         | LR tire right percent tread remaining                           | %                    | float        |
| Lap                             | Laps started count                                              |                      | int          |
| LapBestLap                      | Players best lap number                                         |                      | int          |
| LapBestLapTime                  | Players best lap time                                           | s                    | float        |
| LapBestNLapLap                  | Player last lap in best N average lap time                      |                      | int          |
| LapBestNLapTime                 | Player best N average lap time                                  | s                    | float        |
| LapCompleted                    | Laps completed count                                            |                      | int          |
| LapCurrentLapTime               | Estimate of players current lap time as shown in F3 box         | s                    | float        |
| LapDeltaToBestLap               | Delta time for best lap                                         | s                    | float        |
| LapDeltaToBestLap_DD            | Rate of change of delta time for best lap                       | s/s                  | float        |
| LapDeltaToBestLap_OK            | Delta time for best lap is valid                                |                      | bool         |
| LapDeltaToOptimalLap            | Delta time for optimal lap                                      | s                    | float        |
| LapDeltaToOptimalLap_DD         | Rate of change of delta time for optimal lap                    | s/s                  | float        |
| LapDeltaToOptimalLap_OK         | Delta time for optimal lap is valid                             |                      | bool         |
| LapDeltaToSessionBestLap        | Delta time for session best lap                                 | s                    | float        |
| LapDeltaToSessionBestLap_DD     | Rate of change of delta time for session best lap               | s/s                  | float        |
| LapDeltaToSessionBestLap_OK     | Delta time for session best lap is valid                        |                      | bool         |
| LapDeltaToSessionLastlLap       | Delta time for session last lap                                 | s                    | float        |
| LapDeltaToSessionLastlLap_DD    | Rate of change of delta time for session last lap               | s/s                  | float        |
| LapDeltaToSessionLastlLap_OK    | Delta time for session last lap is valid                        |                      | bool         |
| LapDeltaToSessionOptimalLap     | Delta time for session optimal lap                              | s                    | float        |
| LapDeltaToSessionOptimalLap_DD  | Rate of change of delta time for session optimal lap            | s/s                  | float        |
| LapDeltaToSessionOptimalLap_OK  | Delta time for session optimal lap is valid                     |                      | bool         |
| LapDist                         | Meters traveled from S/F this lap                               | m                    | float        |
| LapDistPct                      | Percentage distance around lap                                  | %                    | float        |
| LapLasNLapSeq                   | Player num consecutive clean laps completed for N average       |                      | int          |
| LapLastLapTime                  | Players last lap time                                           | s                    | float        |
| LapLastNLapTime                 | Player last N average lap time                                  | s                    | float        |
| LatAccel                        | Lateral acceleration (including gravity)                        | m/s^2                | float        |
| LatAccel_ST                     | Lateral acceleration (including gravity) at 360 Hz              | m/s^2                | float[6]     |
| LeftTireSetsAvailable           | How many left tire sets are remaining 255 is unlimited          |                      | int          |
| LeftTireSetsUsed                | How many left tire sets used so far                             |                      | int          |
| LoadNumTextures                 | True if the car_num texture will be loaded                      |                      | bool         |
| LongAccel                       | Longitudinal acceleration (including gravity)                   | m/s^2                | float        |
| LongAccel_ST                    | Longitudinal acceleration (including gravity) at 360 Hz         | m/s^2                | float[6]     |
| ManifoldPress                   | Engine manifold pressure                                        | bar                  | float        |
| ManualBoost                     | Hybrid manual boost state                                       |                      | bool         |
| ManualNoBoost                   | Hybrid manual no boost state                                    |                      | bool         |
| MemPageFaultSec                 | Memory page faults per second                                   |                      | float        |
| MemSoftPageFaultSec             | Memory soft page faults per second                              |                      | float        |
| OilLevel                        | Engine oil level                                                | l                    | float        |
| OilPress                        | Engine oil pressure                                             | bar                  | float        |
| OilTemp                         | Engine oil temperature                                          | C                    | float        |
| OkToReloadTextures              | True if it is ok to reload car textures at this time            |                      | bool         |
| OnPitRoad                       | Is the player car on pit road between the cones                 |                      | bool         |
| PaceMode                        | Are we pacing or not                                            | irsdk_PaceMode       | int          |
| PitOptRepairLeft                | Time left for optional repairs if repairs are active            | s                    | float        |
| PitRepairLeft                   | Time left for mandatory pit repairs if repairs are active       | s                    | float        |
| PitSvFlags                      | Bitfield of pit service checkboxes                              | irsdk_PitSvFlags     | bitfield     |
| PitSvFuel                       | Pit service fuel add amount                                     | l or kWh             | float        |
| PitSvLFP                        | Pit service left front tire pressure                            | kPa                  | float        |
| PitSvLRP                        | Pit service left rear tire pressure                             | kPa                  | float        |
| PitSvRFP                        | Pit service right front tire pressure                           | kPa                  | float        |
| PitSvRRP                        | Pit service right rear tire pressure                            | kPa                  | float        |
| PitSvTireCompound               | Pit service pending tire compound                               |                      | int          |
| Pitch                           | Pitch orientation                                               | rad                  | float        |
| PitchRate                       | Pitch rate                                                      | rad/s                | float        |
| PitchRate_ST                    | Pitch rate at 360 Hz                                            | rad/s                | float[6]     |
| PitsOpen                        | True if pit stop is allowed for the current player              |                      | bool         |
| PitstopActive                   | Is the player getting pit stop service                          |                      | bool         |
| PlayerCarClass                  | Player car class id                                             |                      | int          |
| PlayerCarClassPosition          | Players class position in race                                  |                      | int          |
| PlayerCarDriverIncidentCount    | Teams current drivers incident count for this session           |                      | int          |
| PlayerCarDryTireSetLimit        | Players dry tire set limit                                      |                      | int          |
| PlayerCarIdx                    | Players carIdx                                                  |                      | int          |
| PlayerCarInPitStall             | Players car is properly in their pitstall                       |                      | bool         |
| PlayerCarMyIncidentCount        | Players own incident count for this session                     |                      | int          |
| PlayerCarPitSvStatus            | Players car pit service status bits                             | irsdk_PitSvStatus    | int          |
| PlayerCarPosition               | Players position in race                                        |                      | int          |
| PlayerCarPowerAdjust            | Players power adjust                                            | %                    | float        |
| PlayerCarSLBlinkRPM             | Shift light blink rpm                                           | revs/min             | float        |
| PlayerCarSLFirstRPM             | Shift light first light rpm                                     | revs/min             | float        |
| PlayerCarSLLastRPM              | Shift light last light rpm                                      | revs/min             | float        |
| PlayerCarSLShiftRPM             | Shift light shift rpm                                           | revs/min             | float        |
| PlayerCarTeamIncidentCount      | Players team incident count for this session                    |                      | int          |
| PlayerCarTowTime                | Players car is being towed if time is greater than zero         | s                    | float        |
| PlayerCarWeightPenalty          | Players weight penalty                                          | kg                   | float        |
| PlayerFastRepairsUsed           | Players car number of fast repairs used                         |                      | int          |
| PlayerTireCompound              | Players car current tire compound                               |                      | int          |
| PlayerTrackSurface              | Players car track surface type                                  | irsdk_TrkLoc         | int          |
| PlayerTrackSurfaceMaterial      | Players car track surface material type                         | irsdk_TrkSurf        | int          |
| Precipitation                   | Precipitation at start/finish line                              | %                    | float        |
| PushToPass                      | Push to pass button state                                       |                      | bool         |
| PushToTalk                      | Push to talk button state                                       |                      | bool         |
| RFTiresAvailable                | How many right front tires are remaining 255 is unlimited       |                      | int          |
| RFTiresUsed                     | How many right front tires used so far                          |                      | int          |
| RFbrakeLinePress                | RF brake line pressure                                          | bar                  | float        |
| RFcoldPressure                  | RF tire cold pressure as set in the garage                      | kPa                  | float        |
| RFshockDefl                     | RF shock deflection                                             | m                    | float        |
| RFshockDefl_ST                  | RF shock deflection at 360 Hz                                   | m                    | float[6]     |
| RFshockVel                      | RF shock velocity                                               | m/s                  | float        |
| RFshockVel_ST                   | RF shock velocity at 360 Hz                                     | m/s                  | float[6]     |
| RFtempCL                        | RF tire left carcass temperature                                | C                    | float        |
| RFtempCM                        | RF tire middle carcass temperature                              | C                    | float        |
| RFtempCR                        | RF tire right carcass temperature                               | C                    | float        |
| RFwearL                         | RF tire left percent tread remaining                            | %                    | float        |
| RFwearM                         | RF tire middle percent tread remaining                          | %                    | float        |
| RFwearR                         | RF tire right percent tread remaining                           | %                    | float        |
| RPM                             | Engine rpm                                                      | revs/min             | float        |
| RRTiresAvailable                | How many right rear tires are remaining 255 is unlimited        |                      | int          |
| RRTiresUsed                     | How many right rear tires used so far                           |                      | int          |
| RRbrakeLinePress                | RR brake line pressure                                          | bar                  | float        |
| RRcoldPressure                  | RR tire cold pressure as set in the garage                      | kPa                  | float        |
| RRshockDefl                     | RR shock deflection                                             | m                    | float        |
| RRshockDefl_ST                  | RR shock deflection at 360 Hz                                   | m                    | float[6]     |
| RRshockVel                      | RR shock velocity                                               | m/s                  | float        |
| RRshockVel_ST                   | RR shock velocity at 360 Hz                                     | m/s                  | float[6]     |
| RRtempCL                        | RR tire left carcass temperature                                | C                    | float        |
| RRtempCM                        | RR tire middle carcass temperature                              | C                    | float        |
| RRtempCR                        | RR tire right carcass temperature                               | C                    | float        |
| RRwearL                         | RR tire left percent tread remaining                            | %                    | float        |
| RRwearM                         | RR tire middle percent tread remaining                          | %                    | float        |
| RRwearR                         | RR tire right percent tread remaining                           | %                    | float        |
| RaceLaps                        | Laps completed in race                                          |                      | int          |
| RadioTransmitCarIdx             | The car index of the current person speaking on the radio       |                      | int          |
| RadioTransmitFrequencyIdx       | The frequency index of the current person speaking on the radio |                      | int          |
| RadioTransmitRadioIdx           | The radio index of the current person speaking on the radio     |                      | int          |
| RearTireSetsAvailable           | How many rear tire sets are remaining 255 is unlimited          |                      | int          |
| RearTireSetsUsed                | How many rear tire sets used so far                             |                      | int          |
| RelativeHumidity                | Relative Humidity at start/finish line                          | %                    | float        |
| ReplayFrameNum                  | Integer replay frame number (60 per second)                     |                      | int          |
| ReplayFrameNumEnd               | Integer replay frame number from end of tape                    |                      | int          |
| ReplayPlaySlowMotion            | 0=not slow motion 1=replay is in slow motion                    |                      | bool         |
| ReplayPlaySpeed                 | Replay playback speed                                           |                      | int          |
| ReplaySessionNum                | Replay session number                                           |                      | int          |
| ReplaySessionTime               | Seconds since replay session start                              | s                    | double       |
| RightTireSetsAvailable          | How many right tire sets are remaining 255 is unlimited         |                      | int          |
| RightTireSetsUsed               | How many right tire sets used so far                            |                      | int          |
| Roll                            | Roll orientation                                                | rad                  | float        |
| RollRate                        | Roll rate                                                       | rad/s                | float        |
| RollRate_ST                     | Roll rate at 360 Hz                                             | rad/s                | float[6]     |
| SessionFlags                    | Session flags                                                   | irsdk_Flags          | bitfield     |
| SessionJokerLapsRemain          | Joker laps remaining to be taken                                |                      | int          |
| SessionLapsRemain               | Old laps left till session ends use SessionLapsRemainEx         |                      | int          |
| SessionLapsRemainEx             | New improved laps left till session ends                        |                      | int          |
| SessionLapsTotal                | Total number of laps in session                                 |                      | int          |
| SessionNum                      | Session number                                                  |                      | int          |
| SessionOnJokerLap               | Player is currently completing a joker lap                      |                      | bool         |
| SessionState                    | Session state                                                   | irsdk_SessionState   | int          |
| SessionTick                     | Current update number                                           |                      | int          |
| SessionTime                     | Seconds since session start                                     | s                    | double       |
| SessionTimeOfDay                | Time of day in seconds                                          | s                    | float        |
| SessionTimeRemain               | Seconds left till session ends                                  | s                    | double       |
| SessionTimeTotal                | Total number of seconds in session                              | s                    | double       |
| SessionUniqueID                 | Session ID                                                      |                      | int          |
| ShiftGrindRPM                   | RPM of shifter grinding noise                                   | RPM                  | float        |
| ShiftIndicatorPct               | DEPRECATED use DriverCarSLBlinkRPM instead                      | %                    | float        |
| ShiftPowerPct                   | Friction torque applied to gears when shifting or grinding      | %                    | float        |
| Skies                           | Skies (0=clear/1=p cloudy/2=m cloudy/3=overcast)                |                      | int          |
| SolarAltitude                   | Sun angle above horizon in radians                              | rad                  | float        |
| SolarAzimuth                    | Sun angle clockwise from north in radians                       | rad                  | float        |
| Speed                           | GPS vehicle speed                                               | m/s                  | float        |
| SteeringWheelAngle              | Steering wheel angle                                            | rad                  | float        |
| SteeringWheelAngleMax           | Steering wheel max angle                                        | rad                  | float        |
| SteeringWheelLimiter            | Force feedback limiter strength limits impacts and oscillation  | %                    | float        |
| SteeringWheelMaxForceNm         | Value of strength or max force slider in Nm for FFB             | N\*m                 | float        |
| SteeringWheelPctDamper          | Force feedback % max damping                                    | %                    | float        |
| SteeringWheelPctIntensity       | Force feedback % max intensity                                  | %                    | float        |
| SteeringWheelPctSmoothing       | Force feedback % max smoothing                                  | %                    | float        |
| SteeringWheelPctTorque          | Force feedback % max torque on steering shaft unsigned          | %                    | float        |
| SteeringWheelPctTorqueSign      | Force feedback % max torque on steering shaft signed            | %                    | float        |
| SteeringWheelPctTorqueSignStops | Force feedback % max torque on steering shaft signed stops      | %                    | float        |
| SteeringWheelPeakForceNm        | Peak torque mapping to direct input units for FFB               | N\*m                 | float        |
| SteeringWheelTorque             | Output torque on steering shaft                                 | N\*m                 | float        |
| SteeringWheelTorque_ST          | Output torque on steering shaft at 360 Hz                       | N\*m                 | float[6]     |
| SteeringWheelUseLinear          | True if steering wheel force is using linear mode               |                      | bool         |
| Throttle                        | 0=off throttle to 1=full throttle                               | %                    | float        |
| ThrottleRaw                     | Raw throttle input 0=off throttle to 1=full throttle            | %                    | float        |
| TireLF_RumblePitch              | Players LF Tire Sound rumblestrip pitch                         | Hz                   | float        |
| TireLR_RumblePitch              | Players LR Tire Sound rumblestrip pitch                         | Hz                   | float        |
| TireRF_RumblePitch              | Players RF Tire Sound rumblestrip pitch                         | Hz                   | float        |
| TireRR_RumblePitch              | Players RR Tire Sound rumblestrip pitch                         | Hz                   | float        |
| TireSetsAvailable               | How many tire sets are remaining 255 is unlimited               |                      | int          |
| TireSetsUsed                    | How many tire sets used so far                                  |                      | int          |
| TrackTemp                       | Deprecated set to TrackTempCrew                                 | C                    | float        |
| TrackTempCrew                   | Temperature of track measured by crew around track              | C                    | float        |
| TrackWetness                    | How wet is the average track surface                            | irsdk_TrackWetness   | int          |
| VelocityX                       | X velocity                                                      | m/s                  | float        |
| VelocityX_ST                    | X velocity                                                      | m/s at 360 Hz        | float[6]     |
| VelocityY                       | Y velocity                                                      | m/s                  | float        |
| VelocityY_ST                    | Y velocity                                                      | m/s at 360 Hz        | float[6]     |
| VelocityZ                       | Z velocity                                                      | m/s                  | float        |
| VelocityZ_ST                    | Z velocity                                                      | m/s at 360 Hz        | float[6]     |
| VertAccel                       | Vertical acceleration (including gravity)                       | m/s^2                | float        |
| VertAccel_ST                    | Vertical acceleration (including gravity) at 360 Hz             | m/s^2                | float[6]     |
| VidCapActive                    | True if video currently being captured                          |                      | bool         |
| VidCapEnabled                   | True if video capture system is enabled                         |                      | bool         |
| Voltage                         | Engine voltage                                                  | V                    | float        |
| WaterLevel                      | Engine coolant level                                            | l                    | float        |
| WaterTemp                       | Engine coolant temp                                             | C                    | float        |
| WeatherDeclaredWet              | The steward says rain tires can be used                         |                      | bool         |
| WindDir                         | Wind direction at start/finish line                             | rad                  | float        |
| WindVel                         | Wind velocity at start/finish line                              | m/s                  | float        |
| Yaw                             | Yaw orientation                                                 | rad                  | float        |
| YawNorth                        | Yaw orientation relative to north                               | rad                  | float        |
| YawRate                         | Yaw rate                                                        | rad/s                | float        |
| YawRate_ST                      | Yaw rate at 360 Hz                                              | rad/s                | float[6]     |
| dcAntiRollFront                 | In car front anti roll bar adjustment                           |                      | float        |
| dcAntiRollRear                  | In car rear anti roll bar adjustment                            |                      | float        |
| dcBrakeBias                     | In car brake bias adjustment                                    |                      | float        |
| dcDashPage                      | In car dash display page adjustment                             |                      | float        |
| dcFuelMixture                   | In car fuel mixture adjustment                                  |                      | float        |
| dcPitSpeedLimiterToggle         | In car traction control active                                  |                      | bool         |
| dcPushToPass                    | In car trigger push to pass                                     |                      | bool         |
| dcWeightJackerRight             | In car right wedge/weight jacker adjustment                     |                      | float        |
| dpFastRepair                    | Pitstop fast repair set                                         |                      | float        |
| dpFuelAddKg                     | Pitstop fuel add amount                                         | kg                   | float        |
| dpFuelAutoFillActive            | Pitstop auto fill fuel next stop flag                           |                      | float        |
| dpFuelAutoFillEnabled           | Pitstop auto fill fuel system enabled                           |                      | float        |
| dpFuelFill                      | Pitstop fuel fill flag                                          |                      | float        |
| dpLFTireColdPress               | Pitstop lf tire cold pressure adjustment                        | Pa                   | float        |
| dpLRTireColdPress               | Pitstop lr tire cold pressure adjustment                        | Pa                   | float        |
| dpRFTireColdPress               | Pitstop rf cold tire pressure adjustment                        | Pa                   | float        |
| dpRRTireColdPress               | Pitstop rr cold tire pressure adjustment                        | Pa                   | float        |
| dpTireChange                    | Pitstop all tire change request                                 |                      | float        |
| dpWindshieldTearoff             | Pitstop windshield tearoff                                      |                      | float        |
| dpWingFront                     | Pitstop front wing adjustment                                   |                      | float        |
| dpWingRear                      | Pitstop rear wing adjustment                                    |                      | float        |

