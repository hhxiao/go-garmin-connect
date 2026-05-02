package garmin

import (
	"encoding/json"
)

// GarminDevice is a registered Garmin device. The struct is wide because
// Garmin reports per-device capabilities as a flat boolean grid; many fields
// will be false for older watches.
type GarminDevice struct {
	AppSupport                          bool            `json:"appSupport"`
	ApplicationKey                      string          `json:"applicationKey"`
	DeviceTypePk                        int64           `json:"deviceTypePk"`
	BestInClassVideoLink                string          `json:"bestInClassVideoLink"`
	BluetoothClassicDevice              bool            `json:"bluetoothClassicDevice"`
	BluetoothLowEnergyDevice            bool            `json:"bluetoothLowEnergyDevice"`
	DeviceCategories                    []string        `json:"deviceCategories"`
	DeviceEmbedVideoLink                string          `json:"deviceEmbedVideoLink"`
	DeviceSettingsFile                  string          `json:"deviceSettingsFile"`
	GcmSettingsFile                     string          `json:"gcmSettingsFile"`
	DeviceVideoPageLink                 string          `json:"deviceVideoPageLink"`
	DisplayOrder                        int64           `json:"displayOrder"`
	GolfDisplayOrder                    int64           `json:"golfDisplayOrder"`
	HasOpticalHeartRate                 bool            `json:"hasOpticalHeartRate"`
	Highlighted                         bool            `json:"highlighted"`
	Hybrid                              bool            `json:"hybrid"`
	ImageUrl                            string          `json:"imageUrl"`
	MinGcmAndroidVersion                int64           `json:"minGCMAndroidVersion"`
	MinGcmWindowsVersion                int64           `json:"minGCMWindowsVersion"`
	MinGcMiOsVersion                    int64           `json:"minGCMiOSVersion"`
	MinGolfAppiOsVersion                int64           `json:"minGolfAppiOSVersion"`
	MinGolfAppAndroidVersion            int64           `json:"minGolfAppAndroidVersion"`
	PartNumber                          string          `json:"partNumber"`
	Primary                             bool            `json:"primary"`
	ProductDisplayName                  string          `json:"productDisplayName"`
	DeviceTags                          json.RawMessage `json:"deviceTags"`
	ProductSku                          string          `json:"productSku"`
	Wasp                                bool            `json:"wasp"`
	WeightScale                         bool            `json:"weightScale"`
	Wellness                            bool            `json:"wellness"`
	Wifi                                bool            `json:"wifi"`
	HasPowerButton                      bool            `json:"hasPowerButton"`
	SupportsSecondaryUsers              bool            `json:"supportsSecondaryUsers"`
	AbnormalHeartRateAlertCapable       bool            `json:"abnormalHeartRateAlertCapable"`
	ActivitySummFitFileCapable          bool            `json:"activitySummFitFileCapable"`
	AerobicTrainingEffectCapable        bool            `json:"aerobicTrainingEffectCapable"`
	AlarmDaysCapable                    bool            `json:"alarmDaysCapable"`
	AllDayStressCapable                 bool            `json:"allDayStressCapable"`
	AnaerobicTrainingEffectCapable      bool            `json:"anaerobicTrainingEffectCapable"`
	AtpWorkoutCapable                   bool            `json:"atpWorkoutCapable"`
	BodyBatteryCapable                  bool            `json:"bodyBatteryCapable"`
	BrickWorkoutCapable                 bool            `json:"brickWorkoutCapable"`
	CardioCapable                       bool            `json:"cardioCapable"`
	CardioOptionCapable                 bool            `json:"cardioOptionCapable"`
	CardioSportsCapable                 bool            `json:"cardioSportsCapable"`
	CardioWorkoutCapable                bool            `json:"cardioWorkoutCapable"`
	CellularCapable                     bool            `json:"cellularCapable"`
	ChangeLogCapable                    bool            `json:"changeLogCapable"`
	ContactManagementCapable            bool            `json:"contactManagementCapable"`
	CourseCapable                       bool            `json:"courseCapable"`
	CourseFileType                      string          `json:"courseFileType"`
	CoursePromptCapable                 bool            `json:"coursePromptCapable"`
	CustomIntensityMinutesCapable       bool            `json:"customIntensityMinutesCapable"`
	CustomWorkoutCapable                bool            `json:"customWorkoutCapable"`
	CyclingSegmentCapable               bool            `json:"cyclingSegmentCapable"`
	CyclingSportsCapable                bool            `json:"cyclingSportsCapable"`
	CyclingWorkoutCapable               bool            `json:"cyclingWorkoutCapable"`
	DefaultSettingCapable               bool            `json:"defaultSettingCapable"`
	DeviceSettingCapable                bool            `json:"deviceSettingCapable"`
	DisplayFieldsExtCapable             bool            `json:"displayFieldsExtCapable"`
	DivingCapable                       bool            `json:"divingCapable"`
	EllipticalOptionCapable             bool            `json:"ellipticalOptionCapable"`
	FloorsClimbedGoalCapable            bool            `json:"floorsClimbedGoalCapable"`
	FtpCapable                          bool            `json:"ftpCapable"`
	Gcj02CourseCapable                  bool            `json:"gcj02CourseCapable"`
	GlonassCapable                      bool            `json:"glonassCapable"`
	GoalCapable                         bool            `json:"goalCapable"`
	GoalFileType                        string          `json:"goalFileType"`
	GolfAppSyncCapable                  bool            `json:"golfAppSyncCapable"`
	GpsRouteCapable                     bool            `json:"gpsRouteCapable"`
	HandednessCapable                   bool            `json:"handednessCapable"`
	HrZoneCapable                       bool            `json:"hrZoneCapable"`
	HrvStressCapable                    bool            `json:"hrvStressCapable"`
	IntensityMinutesGoalCapable         bool            `json:"intensityMinutesGoalCapable"`
	LactateThresholdCapable             bool            `json:"lactateThresholdCapable"`
	LanguageSettingCapable              bool            `json:"languageSettingCapable"`
	LowHrAlertCapable                   bool            `json:"lowHrAlertCapable"`
	MaxHrCapable                        bool            `json:"maxHRCapable"`
	MaxWorkoutCount                     int64           `json:"maxWorkoutCount"`
	MetricsFitFileReceiveCapable        bool            `json:"metricsFitFileReceiveCapable"`
	MetricsUploadCapable                bool            `json:"metricsUploadCapable"`
	MilitaryTimeCapable                 bool            `json:"militaryTimeCapable"`
	ModerateIntensityMinutesGoalCapable bool            `json:"moderateIntensityMinutesGoalCapable"`
	NfcCapable                          bool            `json:"nfcCapable"`
	OtherOptionCapable                  bool            `json:"otherOptionCapable"`
	OtherSportsCapable                  bool            `json:"otherSportsCapable"`
	PersonalRecordCapable               bool            `json:"personalRecordCapable"`
	PersonalRecordFileType              string          `json:"personalRecordFileType"`
	PoolSwimOptionCapable               bool            `json:"poolSwimOptionCapable"`
	PowerCurveCapable                   bool            `json:"powerCurveCapable"`
	PowerZonesCapable                   bool            `json:"powerZonesCapable"`
	PulseOxAllDayCapable                bool            `json:"pulseOxAllDayCapable"`
	PulseOxOnDemandCapable              bool            `json:"pulseOxOnDemandCapable"`
	PulseOxSleepCapable                 bool            `json:"pulseOxSleepCapable"`
	RemCapable                          bool            `json:"remCapable"`
	ReminderAlarmCapable                bool            `json:"reminderAlarmCapable"`
	ReorderablePagesCapable             bool            `json:"reorderablePagesCapable"`
	RestingHrCapable                    bool            `json:"restingHRCapable"`
	RideOptionsCapable                  bool            `json:"rideOptionsCapable"`
	RunOptionIndoorCapable              bool            `json:"runOptionIndoorCapable"`
	RunOptionsCapable                   bool            `json:"runOptionsCapable"`
	RunningSegmentCapable               bool            `json:"runningSegmentCapable"`
	RunningSportsCapable                bool            `json:"runningSportsCapable"`
	RunningWorkoutCapable               bool            `json:"runningWorkoutCapable"`
	ScheduleCapable                     bool            `json:"scheduleCapable"`
	ScheduleFileType                    string          `json:"scheduleFileType"`
	SegmentCapable                      bool            `json:"segmentCapable"`
	SegmentPointCapable                 bool            `json:"segmentPointCapable"`
	SettingCapable                      bool            `json:"settingCapable"`
	SettingFileType                     string          `json:"settingFileType"`
	SleepTimeCapable                    bool            `json:"sleepTimeCapable"`
	SmallFitFileOnlyCapable             bool            `json:"smallFitFileOnlyCapable"`
	SportCapable                        bool            `json:"sportCapable"`
	SportFileType                       string          `json:"sportFileType"`
	StairStepperOptionCapable           bool            `json:"stairStepperOptionCapable"`
	StrengthOptionsCapable              bool            `json:"strengthOptionsCapable"`
	StrengthWorkoutCapable              bool            `json:"strengthWorkoutCapable"`
	SupportedHrZones                    []string        `json:"supportedHrZones"`
	SwimWorkoutCapable                  bool            `json:"swimWorkoutCapable"`
	TrainingPlanCapable                 bool            `json:"trainingPlanCapable"`
	TrainingStatusCapable               bool            `json:"trainingStatusCapable"`
	TrainingStatusPauseCapable          bool            `json:"trainingStatusPauseCapable"`
	UserProfileCapable                  bool            `json:"userProfileCapable"`
	UserTcxExportCapable                bool            `json:"userTcxExportCapable"`
	Vo2MaxBikeCapable                   bool            `json:"vo2MaxBikeCapable"`
	Vo2MaxRunCapable                    bool            `json:"vo2MaxRunCapable"`
	WalkOptionCapable                   bool            `json:"walkOptionCapable"`
	WalkingSportsCapable                bool            `json:"walkingSportsCapable"`
	WeatherAlertsCapable                bool            `json:"weatherAlertsCapable"`
	WeatherSettingsCapable              bool            `json:"weatherSettingsCapable"`
	WorkoutCapable                      bool            `json:"workoutCapable"`
	WorkoutFileType                     string          `json:"workoutFileType"`
	YogaCapable                         bool            `json:"yogaCapable"`
	YogaOptionCapable                   bool            `json:"yogaOptionCapable"`
	HeatAndAltitudeAcclimationCapable   bool            `json:"heatAndAltitudeAcclimationCapable"`
	TrainingLoadBalanceCapable          bool            `json:"trainingLoadBalanceCapable"`
	IndoorTrackOptionsCapable           bool            `json:"indoorTrackOptionsCapable"`
	IndoorBikeOptionsCapable            bool            `json:"indoorBikeOptionsCapable"`
	IndoorWalkOptionsCapable            bool            `json:"indoorWalkOptionsCapable"`
	TrainingEffectLabelCapable          bool            `json:"trainingEffectLabelCapable"`
	PacebandCapable                     bool            `json:"pacebandCapable"`
	RespirationCapable                  bool            `json:"respirationCapable"`
	OpenWaterSwimOptionCapable          bool            `json:"openWaterSwimOptionCapable"`
	PhoneVerificationCheckRequired      bool            `json:"phoneVerificationCheckRequired"`
	WeightGoalCapable                   bool            `json:"weightGoalCapable"`
	YogaWorkoutCapable                  bool            `json:"yogaWorkoutCapable"`
	PilatesWorkoutCapable               bool            `json:"pilatesWorkoutCapable"`
	ConnectedGpsCapable                 bool            `json:"connectedGPSCapable"`
	DiveAppSyncCapable                  bool            `json:"diveAppSyncCapable"`
	GolfLiveScoringCapable              bool            `json:"golfLiveScoringCapable"`
	BloodEfficiencySleepCapable         bool            `json:"bloodEfficiencySleepCapable"`
	BloodEfficiencyAllDayCapable        bool            `json:"bloodEfficiencyAllDayCapable"`
	BloodEfficiencyOnDemandCapable      bool            `json:"bloodEfficiencyOnDemandCapable"`
	SolarPanelUtilizationCapable        bool            `json:"solarPanelUtilizationCapable"`
	SweatLossCapable                    bool            `json:"sweatLossCapable"`
	DiveAlertCapable                    bool            `json:"diveAlertCapable"`
	RequiresInitialDeviceNickname       bool            `json:"requiresInitialDeviceNickname"`
	DefaultSettingsHbaseMigrated        bool            `json:"defaultSettingsHbaseMigrated"`
	SleepScoreCapable                   bool            `json:"sleepScoreCapable"`
	FitnessAgeV2Capable                 bool            `json:"fitnessAgeV2Capable"`
	IntensityMinutesV2Capable           bool            `json:"intensityMinutesV2Capable"`
	CollapsibleControlMenuCapable       bool            `json:"collapsibleControlMenuCapable"`
	MeasurementUnitSettingCapable       bool            `json:"measurementUnitSettingCapable"`
	OnDeviceSleepCalculationCapable     bool            `json:"onDeviceSleepCalculationCapable"`
	HiitWorkoutCapable                  bool            `json:"hiitWorkoutCapable"`
	RunningHeartRateZoneCapable         bool            `json:"runningHeartRateZoneCapable"`
	CyclingHeartRateZoneCapable         bool            `json:"cyclingHeartRateZoneCapable"`
	SwimmingHeartRateZoneCapable        bool            `json:"swimmingHeartRateZoneCapable"`
	DefaultHeartRateZoneCapable         bool            `json:"defaultHeartRateZoneCapable"`
	CyclingPowerZonesCapable            bool            `json:"cyclingPowerZonesCapable"`
	XcSkiPowerZonesCapable              bool            `json:"xcSkiPowerZonesCapable"`
	SwimAlgorithmCapable                bool            `json:"swimAlgorithmCapable"`
	BenchmarkExerciseCapable            bool            `json:"benchmarkExerciseCapable"`
	SpectatorMessagingCapable           bool            `json:"spectatorMessagingCapable"`
	EcgCapable                          bool            `json:"ecgCapable"`
	LteLiveEventSharingCapable          bool            `json:"lteLiveEventSharingCapable"`
	SleepFitFileReceiveCapable          bool            `json:"sleepFitFileReceiveCapable"`
	SecondaryWorkoutStepTargetCapable   bool            `json:"secondaryWorkoutStepTargetCapable"`
	AssistancePlusCapable               bool            `json:"assistancePlusCapable"`
	PowerGuidanceCapable                bool            `json:"powerGuidanceCapable"`
	AirIntegrationCapable               bool            `json:"airIntegrationCapable"`
	HealthSnapshotCapable               bool            `json:"healthSnapshotCapable"`
	RacePredictionsRunCapable           bool            `json:"racePredictionsRunCapable"`
	VivohubCompatible                   bool            `json:"vivohubCompatible"`
	StepsTrueUpChartCapable             bool            `json:"stepsTrueUpChartCapable"`
	SportingEventCapable                bool            `json:"sportingEventCapable"`
	SolarChargeCapable                  bool            `json:"solarChargeCapable"`
	RealTimeSettingsCapable             bool            `json:"realTimeSettingsCapable"`
	EmergencyCallingCapable             bool            `json:"emergencyCallingCapable"`
	PersonalRepRecordCapable            bool            `json:"personalRepRecordCapable"`
	HrvStatusCapable                    bool            `json:"hrvStatusCapable"`
	TrainingReadinessCapable            bool            `json:"trainingReadinessCapable"`
	PublicBetaSoftwareCapable           bool            `json:"publicBetaSoftwareCapable"`
	WorkoutAudioPromptsCapable          bool            `json:"workoutAudioPromptsCapable"`
	ActualStepRecordingCapable          bool            `json:"actualStepRecordingCapable"`
	GroupTrack2Capable                  bool            `json:"groupTrack2Capable"`
	GolfAppPairingCapable               bool            `json:"golfAppPairingCapable"`
	LocalWindConditionsCapable          bool            `json:"localWindConditionsCapable"`
	MultipleGolfCourseCapable           bool            `json:"multipleGolfCourseCapable"`
	BeaconTrackingCapable               bool            `json:"beaconTrackingCapable"`
	BatteryStatusCapable                bool            `json:"batteryStatusCapable"`
	Datasource                          string          `json:"datasource"`
	DeviceStatus                        string          `json:"deviceStatus"`
	RegisteredDate                      int64           `json:"registeredDate"`
	ActualProductSku                    string          `json:"actualProductSku"`
	VivohubConfigurable                 json.RawMessage `json:"vivohubConfigurable"`
	SerialNumber                        string          `json:"serialNumber"`
	ShortName                           json.RawMessage `json:"shortName"`
	DisplayName                         string          `json:"displayName"`
	UnitID                              int64           `json:"unitId"`
	DeviceID                            int64           `json:"deviceId"`
	WifiSetup                           bool            `json:"wifiSetup"`
	CurrentFirmwareVersionMajor         int64           `json:"currentFirmwareVersionMajor"`
	CurrentFirmwareVersionMinor         int64           `json:"currentFirmwareVersionMinor"`
	ActiveInd                           int64           `json:"activeInd"`
	PrimaryActivityTrackerIndicator     bool            `json:"primaryActivityTrackerIndicator"`
	CurrentFirmwareVersion              string          `json:"currentFirmwareVersion"`
	IsPrimaryUser                       bool            `json:"isPrimaryUser"`
	CorporateDevice                     bool            `json:"corporateDevice"`
	PrePairedWithHrm                    bool            `json:"prePairedWithHRM"`
	UnRetirable                         bool            `json:"unRetirable"`
	OtherAssociation                    bool            `json:"otherAssociation"`
	MinCustomIntensityMinutesVersion    string          `json:"minCustomIntensityMinutesVersion"`
}

// GarminDeviceMessages is the device message inbox.
type GarminDeviceMessages struct {
	ServiceHost   string                `json:"serviceHost"`
	NumOfMessages int64                 `json:"numOfMessages"`
	Messages      []GarminDeviceMessage `json:"messages"`
}

// GarminDeviceMessage is one entry in the device message inbox.
type GarminDeviceMessage struct {
	MessageID         int64                       `json:"messageId"`
	MessageType       string                      `json:"messageType"`
	MessageStatus     string                      `json:"messageStatus"`
	DeviceID          int64                       `json:"deviceId"`
	DeviceName        string                      `json:"deviceName"`
	ApplicationKey    string                      `json:"applicationKey"`
	FirmwareVersion   string                      `json:"firmwareVersion"`
	WifiSetup         bool                        `json:"wifiSetup"`
	DeviceXmlDataType string                      `json:"deviceXmlDataType"`
	Hidden            bool                        `json:"hidden"`
	CreatedTimeStamp  int64                       `json:"createdTimeStamp"`
	UpdatedTimeStamp  json.RawMessage             `json:"updatedTimeStamp"`
	MetaData          GarminDeviceMessageMetaData `json:"metaData"`
}

// GarminDeviceMessageMetaData is the metadata payload of a device message.
type GarminDeviceMessageMetaData struct {
	Type                 string            `json:"type"`
	FileType             string            `json:"fileType"`
	MessageUrl           string            `json:"messageUrl"`
	Absolute             *bool             `json:"absolute"`
	MessageName          string            `json:"messageName"`
	GroupName            string            `json:"groupName"`
	Priority             *int64            `json:"priority"`
	MetaDataID           *int64            `json:"metaDataId"`
	AppDetails           json.RawMessage   `json:"appDetails"`
	PartNumber           string            `json:"partNumber"`
	Version              string            `json:"version"`
	Path                 string            `json:"path"`
	FileSize             *int64            `json:"fileSize"`
	ServerPath           string            `json:"serverPath"`
	SecureServerPath     string            `json:"secureServerPath"`
	FileName             string            `json:"fileName"`
	FilenameOnDevice     string            `json:"filenameOnDevice"`
	ProductName          string            `json:"productName"`
	DataType             string            `json:"dataType"`
	Notes                json.RawMessage   `json:"notes"`
	Instructions         json.RawMessage   `json:"instructions"`
	InstallationOrder    *int64            `json:"installationOrder"`
	DeliveryRestrictions []string          `json:"deliveryRestrictions"`
	ChangeLog            []json.RawMessage `json:"changeLog"`
}

// GarminDeviceLastUsed describes the most recently synced device.
type GarminDeviceLastUsed struct {
	UserDeviceID                 int64  `json:"userDeviceId"`
	UserProfileNumber            int64  `json:"userProfileNumber"`
	ApplicationNumber            int64  `json:"applicationNumber"`
	LastUsedDeviceApplicationKey string `json:"lastUsedDeviceApplicationKey"`
	LastUsedDeviceName           string `json:"lastUsedDeviceName"`
	LastUsedDeviceUploadTime     int64  `json:"lastUsedDeviceUploadTime"`
	ImageUrl                     string `json:"imageUrl"`
	Released                     bool   `json:"released"`
}

// GarminDeviceSettings is the per-device settings blob. Many fields are
// `json.RawMessage` because Garmin returns them with type that varies by
// device family.
type GarminDeviceSettings struct {
	DeviceID                   int64                      `json:"deviceId"`
	TimeFormat                 string                     `json:"timeFormat"`
	DateFormat                 string                     `json:"dateFormat"`
	MeasurementUnits           string                     `json:"measurementUnits"`
	AllUnits                   string                     `json:"allUnits"`
	VisibleScreens             json.RawMessage            `json:"visibleScreens"`
	EnabledScreens             json.RawMessage            `json:"enabledScreens"`
	ScreenLists                json.RawMessage            `json:"screenLists"`
	IsVivohubEnabled           json.RawMessage            `json:"isVivohubEnabled"`
	Alarms                     []json.RawMessage          `json:"alarms"`
	SupportedAlarmModes        json.RawMessage            `json:"supportedAlarmModes"`
	MultipleAlarmEnabled       bool                       `json:"multipleAlarmEnabled"`
	MaxAlarm                   json.RawMessage            `json:"maxAlarm"`
	ActivityTracking           json.RawMessage            `json:"activityTracking"`
	KeyTonesEnabled            json.RawMessage            `json:"keyTonesEnabled"`
	KeyVibrationEnabled        json.RawMessage            `json:"keyVibrationEnabled"`
	AlertTonesEnabled          json.RawMessage            `json:"alertTonesEnabled"`
	UserNoticeTonesEnabled     json.RawMessage            `json:"userNoticeTonesEnabled"`
	GlonassEnabled             json.RawMessage            `json:"glonassEnabled"`
	TurnPromptEnabled          json.RawMessage            `json:"turnPromptEnabled"`
	SegmentPromptEnabled       json.RawMessage            `json:"segmentPromptEnabled"`
	SupportedLanguages         []SupportedLanguage        `json:"supportedLanguages"`
	Language                   int64                      `json:"language"`
	StartOfWeek                string                     `json:"startOfWeek"`
	ConnectIq                  ConnectIq                  `json:"connectIQ"`
	AutoUploadEnabled          bool                       `json:"autoUploadEnabled"`
	MultipleSupportedWatchFace MultipleSupportedWatchFace `json:"multipleSupportedWatchFace"`
	MetricsFileTrueupEnabled   bool                       `json:"metricsFileTrueupEnabled"`
	// Many additional fields are returned by Garmin but typed as `object` in
	// the C# library. They are exposed here as RawMessage so callers can
	// project specific fields without forcing the library to track every one.
	Extra map[string]json.RawMessage `json:"-"`
}

// ConnectIq is a sub-record of GarminDeviceSettings.
type ConnectIq struct {
	AutoUpdate bool `json:"autoUpdate"`
}

// MultipleSupportedWatchFace is intentionally empty on most devices. Kept as
// a struct so future fields can be added without breaking consumers.
type MultipleSupportedWatchFace struct{}

// SupportedLanguage is one entry in the GarminDeviceSettings language list.
type SupportedLanguage struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
