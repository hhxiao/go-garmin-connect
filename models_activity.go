package garmin

import (
	"encoding/json"
	"time"
)

// ActivityDownloadFormat selects the file format DownloadActivity returns.
type ActivityDownloadFormat int

const (
	// FormatOriginal returns the FIT zip bundle.
	FormatOriginal ActivityDownloadFormat = iota
	// FormatTCX returns a Training Center XML file.
	FormatTCX
	// FormatGPX returns a GPS Exchange XML file.
	FormatGPX
	// FormatKML returns a Keyhole Markup Language file (Google Earth).
	FormatKML
	// FormatCSV returns a CSV of splits.
	FormatCSV
)

// GarminActivity is the wide flat record returned by the activity-list and
// activity-search endpoints. Many fields are zero for activity types that
// don't produce them (e.g. swim metrics on a run).
type GarminActivity struct {
	ActivityID                            int64              `json:"activityId"`
	ActivityName                          string             `json:"activityName"`
	Description                           string             `json:"description"`
	StartTimeLocal                        LocalDateTime      `json:"startTimeLocal"`
	StartTimeGmt                          LocalDateTime      `json:"startTimeGMT"`
	ActivityType                          ActivityType       `json:"activityType"`
	EventType                             EventType          `json:"eventType"`
	Comments                              string             `json:"comments"`
	ParentID                              string             `json:"parentId"`
	Distance                              float64            `json:"distance"`
	Duration                              float64            `json:"duration"`
	ElapsedDuration                       float64            `json:"elapsedDuration"`
	MovingDuration                        float64            `json:"movingDuration"`
	ElevationGain                         float64            `json:"elevationGain"`
	ElevationLoss                         float64            `json:"elevationLoss"`
	AverageSpeed                          float64            `json:"averageSpeed"`
	MaxSpeed                              float64            `json:"maxSpeed"`
	StartLatitude                         float64            `json:"startLatitude"`
	StartLongitude                        float64            `json:"startLongitude"`
	HasPolyline                           bool               `json:"hasPolyline"`
	OwnerID                               int64              `json:"ownerId"`
	OwnerDisplayName                      string             `json:"ownerDisplayName"`
	OwnerFullName                         string             `json:"ownerFullName"`
	OwnerProfileImageUrlSmall             string             `json:"ownerProfileImageUrlSmall"`
	OwnerProfileImageUrlMedium            string             `json:"ownerProfileImageUrlMedium"`
	OwnerProfileImageUrlLarge             string             `json:"ownerProfileImageUrlLarge"`
	Calories                              float64            `json:"calories"`
	AverageHr                             float64            `json:"averageHR"`
	MaxHr                                 float64            `json:"maxHR"`
	AverageRunningCadenceInStepsPerMinute float64            `json:"averageRunningCadenceInStepsPerMinute"`
	MaxRunningCadenceInStepsPerMinute     float64            `json:"maxRunningCadenceInStepsPerMinute"`
	AverageBikingCadenceInRevPerMinute    float64            `json:"averageBikingCadenceInRevPerMinute"`
	MaxBikingCadenceInRevPerMinute        float64            `json:"maxBikingCadenceInRevPerMinute"`
	AverageSwimCadenceInStrokesPerMinute  float64            `json:"averageSwimCadenceInStrokesPerMinute"`
	MaxSwimCadenceInStrokesPerMinute      float64            `json:"maxSwimCadenceInStrokesPerMinute"`
	AverageSwolf                          float64            `json:"averageSwolf"`
	ActiveLengths                         float64            `json:"activeLengths"`
	Steps                                 int                `json:"steps"`
	NumberOfActivityLikes                 int                `json:"numberOfActivityLikes"`
	NumberOfActivityComments              int                `json:"numberOfActivityComments"`
	LikedByUser                           int                `json:"likedByUser"`
	UserRoles                             []string           `json:"userRoles"`
	Privacy                               Privacy            `json:"privacy"`
	UserPro                               bool               `json:"userPro"`
	HasVideo                              bool               `json:"hasVideo"`
	VideoUrl                              string             `json:"videoUrl"`
	TimeZoneID                            int64              `json:"timeZoneId"`
	BeginTimestamp                        int64              `json:"beginTimestamp"`
	SportTypeID                           int64              `json:"sportTypeId"`
	AvgPower                              float64            `json:"avgPower"`
	MaxPower                              float64            `json:"maxPower"`
	AerobicTrainingEffect                 float64            `json:"aerobicTrainingEffect"`
	AnaerobicTrainingEffect               float64            `json:"anaerobicTrainingEffect"`
	Strokes                               float64            `json:"strokes"`
	NormPower                             float64            `json:"normPower"`
	LeftBalance                           float64            `json:"leftBalance"`
	RightBalance                          float64            `json:"rightBalance"`
	AvgLeftBalance                        float64            `json:"avgLeftBalance"`
	Max20MinPower                         float64            `json:"max20MinPower"`
	AvgVerticalOscillation                float64            `json:"avgVerticalOscillation"`
	AvgGroundContactTime                  float64            `json:"avgGroundContactTime"`
	AvgStrideLength                       float64            `json:"avgStrideLength"`
	AvgFractionalCadence                  float64            `json:"avgFractionalCadence"`
	MaxFractionalCadence                  float64            `json:"maxFractionalCadence"`
	TrainingStressScore                   float64            `json:"trainingStressScore"`
	IntensityFactor                       float64            `json:"intensityFactor"`
	VO2MaxValue                           float64            `json:"vO2MaxValue"`
	AvgVerticalRatio                      float64            `json:"avgVerticalRatio"`
	AvgGroundContactBalance               float64            `json:"avgGroundContactBalance"`
	LactateThresholdBpm                   float64            `json:"lactateThresholdBpm"`
	LactateThresholdSpeed                 float64            `json:"lactateThresholdSpeed"`
	MaxFtp                                float64            `json:"maxFtp"`
	AvgStrokeDistance                     float64            `json:"avgStrokeDistance"`
	AvgStrokeCadence                      float64            `json:"avgStrokeCadence"`
	MaxStrokeCadence                      float64            `json:"maxStrokeCadence"`
	WorkoutID                             float64            `json:"workoutId"`
	AvgStrokes                            float64            `json:"avgStrokes"`
	MinStrokes                            float64            `json:"minStrokes"`
	DeviceID                              int64              `json:"deviceId"`
	MinTemperature                        float64            `json:"minTemperature"`
	MaxTemperature                        float64            `json:"maxTemperature"`
	MinElevation                          float64            `json:"minElevation"`
	MaxElevation                          float64            `json:"maxElevation"`
	AvgDoubleCadence                      float64            `json:"avgDoubleCadence"`
	MaxDoubleCadence                      float64            `json:"maxDoubleCadence"`
	MaxDepth                              float64            `json:"maxDepth"`
	AvgDepth                              float64            `json:"avgDepth"`
	SummarizedDiveInfo                    SummarizedDiveInfo `json:"summarizedDiveInfo"`
	AvgVerticalSpeed                      float64            `json:"avgVerticalSpeed"`
	MaxVerticalSpeed                      float64            `json:"maxVerticalSpeed"`
	Manufacturer                          string             `json:"manufacturer"`
	LocationName                          string             `json:"locationName"`
	LapCount                              int64              `json:"lapCount"`
	EndLatitude                           float64            `json:"endLatitude"`
	EndLongitude                          float64            `json:"endLongitude"`
	MinAirSpeed                           float64            `json:"minAirSpeed"`
	MaxAirSpeed                           float64            `json:"maxAirSpeed"`
	AvgAirSpeed                           float64            `json:"avgAirSpeed"`
	AvgWindYawAngle                       float64            `json:"avgWindYawAngle"`
	MinCda                                float64            `json:"minCda"`
	MaxCda                                float64            `json:"maxCda"`
	AvgCda                                float64            `json:"avgCda"`
	AvgWattsPerCda                        float64            `json:"avgWattsPerCda"`
	MaxAvgPower1                          float64            `json:"maxAvgPower_1"`
	MaxAvgPower2                          float64            `json:"maxAvgPower_2"`
	MaxAvgPower5                          float64            `json:"maxAvgPower_5"`
	MaxAvgPower10                         float64            `json:"maxAvgPower_10"`
	MaxAvgPower20                         float64            `json:"maxAvgPower_20"`
	MaxAvgPower30                         float64            `json:"maxAvgPower_30"`
	MaxAvgPower60                         float64            `json:"maxAvgPower_60"`
	MaxAvgPower120                        float64            `json:"maxAvgPower_120"`
	MaxAvgPower300                        float64            `json:"maxAvgPower_300"`
	MaxAvgPower600                        float64            `json:"maxAvgPower_600"`
	MaxAvgPower1200                       float64            `json:"maxAvgPower_1200"`
	MaxAvgPower1800                       float64            `json:"maxAvgPower_1800"`
	MaxAvgPower3600                       float64            `json:"maxAvgPower_3600"`
	MaxAvgPower7200                       float64            `json:"maxAvgPower_7200"`
	MaxAvgPower18000                      float64            `json:"maxAvgPower_18000"`
	ExcludeFromPowerCurveReports          bool               `json:"excludeFromPowerCurveReports"`
	TotalSets                             float64            `json:"totalSets"`
	ActiveSets                            float64            `json:"activeSets"`
	TotalReps                             float64            `json:"totalReps"`
	MinRespirationRate                    float64            `json:"minRespirationRate"`
	MaxRespirationRate                    float64            `json:"maxRespirationRate"`
	AvgRespirationRate                    float64            `json:"avgRespirationRate"`
	ActivityTrainingLoad                  float64            `json:"activityTrainingLoad"`
	AvgFlow                               float64            `json:"avgFlow"`
	AvgGrit                               float64            `json:"avgGrit"`
	MinActivityLapDuration                float64            `json:"minActivityLapDuration"`
	AvgStress                             float64            `json:"avgStress"`
	StartStress                           float64            `json:"startStress"`
	EndStress                             float64            `json:"endStress"`
	MaxStress                             float64            `json:"maxStress"`
	SplitSummaries                        []SplitSummary     `json:"splitSummaries"`
	HasSplits                             bool               `json:"hasSplits"`
	ModerateIntensityMinutes              int64              `json:"moderateIntensityMinutes"`
	VigorousIntensityMinutes              int64              `json:"vigorousIntensityMinutes"`
	Purposeful                            bool               `json:"purposeful"`
	Favorite                              bool               `json:"favorite"`
	Pr                                    bool               `json:"pr"`
	AutoCalcCalories                      bool               `json:"autoCalcCalories"`
	AtpActivity                           bool               `json:"atpActivity"`
	ManualActivity                        bool               `json:"manualActivity"`
	ElevationCorrected                    bool               `json:"elevationCorrected"`
	DecoDive                              bool               `json:"decoDive"`
	Parent                                bool               `json:"parent"`
	// Untyped passthrough fields kept as raw JSON so callers can decode them
	// case-by-case if a future Garmin schema change starts populating them.
	ConversationUuid               json.RawMessage `json:"conversationUuid"`
	ConversationPk                 json.RawMessage `json:"conversationPk"`
	CommentedByUser                json.RawMessage `json:"commentedByUser"`
	ActivityLikeDisplayNames       json.RawMessage `json:"activityLikeDisplayNames"`
	ActivityLikeFullNames          json.RawMessage `json:"activityLikeFullNames"`
	ActivityLikeProfileImageUrls   json.RawMessage `json:"activityLikeProfileImageUrls"`
	RequestorRelationship          json.RawMessage `json:"requestorRelationship"`
	CourseID                       json.RawMessage `json:"courseId"`
	PoolLength                     json.RawMessage `json:"poolLength"`
	UnitOfPoolLength               json.RawMessage `json:"unitOfPoolLength"`
	SummarizedExerciseSets         json.RawMessage `json:"summarizedExerciseSets"`
	SurfaceInterval                json.RawMessage `json:"surfaceInterval"`
	StartN2                        json.RawMessage `json:"startN2"`
	EndN2                          json.RawMessage `json:"endN2"`
	StartCns                       json.RawMessage `json:"startCns"`
	EndCns                         json.RawMessage `json:"endCns"`
	ActivityLikeAuthors            json.RawMessage `json:"activityLikeAuthors"`
	FloorsClimbed                  json.RawMessage `json:"floorsClimbed"`
	FloorsDescended                json.RawMessage `json:"floorsDescended"`
	DiveNumber                     json.RawMessage `json:"diveNumber"`
	BottomTime                     json.RawMessage `json:"bottomTime"`
	Flow                           json.RawMessage `json:"flow"`
	Grit                           json.RawMessage `json:"grit"`
	JumpCount                      json.RawMessage `json:"jumpCount"`
	CaloriesEstimated              json.RawMessage `json:"caloriesEstimated"`
	CaloriesConsumed               json.RawMessage `json:"caloriesConsumed"`
	WaterEstimated                 json.RawMessage `json:"waterEstimated"`
	WaterConsumed                  json.RawMessage `json:"waterConsumed"`
	TrainingEffectLabel            json.RawMessage `json:"trainingEffectLabel"`
	DifferenceStress               json.RawMessage `json:"differenceStress"`
	AerobicTrainingEffectMessage   json.RawMessage `json:"aerobicTrainingEffectMessage"`
	AnaerobicTrainingEffectMessage json.RawMessage `json:"anaerobicTrainingEffectMessage"`
	MaxBottomTime                  json.RawMessage `json:"maxBottomTime"`
	HasSeedFirstbeatProfile        json.RawMessage `json:"hasSeedFirstbeatProfile"`
}

// ActivityType is a sub-record describing the activity's primary type.
type ActivityType struct {
	TypeID       int64           `json:"typeId"`
	TypeKey      string          `json:"typeKey"`
	ParentTypeID int64           `json:"parentTypeId"`
	IsHidden     bool            `json:"isHidden"`
	SortOrder    json.RawMessage `json:"sortOrder"`
	Trimmable    bool            `json:"trimmable"`
	Restricted   bool            `json:"restricted"`
}

// EventType describes the activity's recorded event (e.g. race, training).
type EventType struct {
	TypeID    int64  `json:"typeId"`
	TypeKey   string `json:"typeKey"`
	SortOrder int64  `json:"sortOrder"`
}

// Privacy is a sub-record describing visibility flags.
type Privacy struct {
	TypeID  int64  `json:"typeId"`
	TypeKey string `json:"typeKey"`
}

// SplitSummary is a sub-record (also reused inside GarminActivity).
type SplitSummary struct {
	NoOfSplits           json.RawMessage `json:"noOfSplits"`
	MaxGradeValue        json.RawMessage `json:"maxGradeValue"`
	TotalAscent          float64         `json:"totalAscent"`
	Duration             float64         `json:"duration"`
	SplitType            string          `json:"splitType"`
	NumClimbSends        float64         `json:"numClimbSends"`
	MaxElevationGain     float64         `json:"maxElevationGain"`
	AverageElevationGain float64         `json:"averageElevationGain"`
	MaxDistance          float64         `json:"maxDistance"`
	Distance             float64         `json:"distance"`
	AverageSpeed         float64         `json:"averageSpeed"`
	MaxSpeed             float64         `json:"maxSpeed"`
	Mode                 json.RawMessage `json:"mode"`
	NumFalls             int64           `json:"numFalls"`
}

// SummarizedDiveInfo is a sub-record of GarminActivity.
type SummarizedDiveInfo struct {
	Weight              float64           `json:"weight"`
	WeightUnit          json.RawMessage   `json:"weightUnit"`
	Visibility          json.RawMessage   `json:"visibility"`
	VisibilityUnit      json.RawMessage   `json:"visibilityUnit"`
	SurfaceCondition    json.RawMessage   `json:"surfaceCondition"`
	Current             json.RawMessage   `json:"current"`
	WaterType           json.RawMessage   `json:"waterType"`
	WaterDensity        json.RawMessage   `json:"waterDensity"`
	SummarizedDiveGases []json.RawMessage `json:"summarizedDiveGases"`
	TotalSurfaceTime    json.RawMessage   `json:"totalSurfaceTime"`
}

// GarminActivityDetails is the per-activity metric/polyline payload.
type GarminActivityDetails struct {
	ActivityID            int64                  `json:"activityId"`
	MeasurementCount      int64                  `json:"measurementCount"`
	MetricsCount          int64                  `json:"metricsCount"`
	MetricDescriptors     []MetricDescriptor     `json:"metricDescriptors"`
	ActivityDetailMetrics []ActivityDetailMetric `json:"activityDetailMetrics"`
	GeoPolylineDto        GeoPolylineDto         `json:"geoPolylineDTO"`
	HeartRateDtos         json.RawMessage        `json:"heartRateDTOs"`
	DetailsAvailable      bool                   `json:"detailsAvailable"`
}

// MetricDescriptor describes a single column of ActivityDetailMetric.
type MetricDescriptor struct {
	MetricsIndex         int64  `json:"metricsIndex"`
	Key                  string `json:"key"`
	Unit                 Unit   `json:"unit"`
	AppID                string `json:"appID"`
	DeveloperFieldNumber int64  `json:"developerFieldNumber"`
}

// ActivityDetailMetric is one row of metric values, indexed by the parallel
// MetricDescriptor list.
type ActivityDetailMetric struct {
	Metrics []float64 `json:"metrics"`
}

// Unit is a measurement unit descriptor.
type Unit struct {
	ID     int64   `json:"id"`
	Key    string  `json:"key"`
	Factor float64 `json:"factor"`
}

// GeoPolylineDto is the route polyline for an activity.
type GeoPolylineDto struct {
	StartPoint EndPoint   `json:"startPoint"`
	EndPoint   EndPoint   `json:"endPoint"`
	MinLat     float64    `json:"minLat"`
	MaxLat     float64    `json:"maxLat"`
	MinLon     float64    `json:"minLon"`
	MaxLon     float64    `json:"maxLon"`
	Polyline   []EndPoint `json:"polyline"`
}

// EndPoint is one polyline vertex.
type EndPoint struct {
	Lat                       float64         `json:"lat"`
	Lon                       float64         `json:"lon"`
	Altitude                  json.RawMessage `json:"altitude"`
	Time                      int64           `json:"time"`
	TimerStart                bool            `json:"timerStart"`
	TimerStop                 bool            `json:"timerStop"`
	DistanceFromPreviousPoint json.RawMessage `json:"distanceFromPreviousPoint"`
	DistanceInMeters          json.RawMessage `json:"distanceInMeters"`
	Speed                     float64         `json:"speed"`
	CumulativeAscent          json.RawMessage `json:"cumulativeAscent"`
	CumulativeDescent         json.RawMessage `json:"cumulativeDescent"`
	ExtendedCoordinate        bool            `json:"extendedCoordinate"`
	Valid                     bool            `json:"valid"`
}

// GarminActivityWeather is the weather snapshot for an activity.
type GarminActivityWeather struct {
	IssueDate                 LocalDateTime     `json:"issueDate"`
	Temp                      int64             `json:"temp"`
	ApparentTemp              int64             `json:"apparentTemp"`
	DewPoint                  int64             `json:"dewPoint"`
	RelativeHumidity          int64             `json:"relativeHumidity"`
	WindDirection             int64             `json:"windDirection"`
	WindDirectionCompassPoint string            `json:"windDirectionCompassPoint"`
	WindSpeed                 int64             `json:"windSpeed"`
	WindGust                  string            `json:"windGust"`
	Latitude                  float64           `json:"latitude"`
	Longitude                 float64           `json:"longitude"`
	WeatherStationDto         WeatherStationDto `json:"weatherStationDTO"`
	WeatherTypeDto            WeatherTypeDto    `json:"weatherTypeDTO"`
}

// WeatherStationDto names the source station.
type WeatherStationDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
}

// WeatherTypeDto describes the weather phenomenon.
type WeatherTypeDto struct {
	WeatherTypePk string `json:"weatherTypePk"`
	Desc          string `json:"desc"`
	Image         string `json:"image"`
}

// GarminActivitySplits is the per-lap breakdown of an activity.
type GarminActivitySplits struct {
	ActivityID int64      `json:"activityId"`
	LapDtos    []LapDto   `json:"lapDTOs"`
	EventDtos  []EventDto `json:"eventDTOs"`
}

// EventDto is an in-activity event marker.
type EventDto struct {
	StartTimeGmt            time.Time      `json:"startTimeGMT"`
	StartTimeGmtDoubleValue float64        `json:"startTimeGMTDoubleValue"`
	SectionTypeDto          SectionTypeDto `json:"sectionTypeDTO"`
}

// SectionTypeDto describes the type of EventDto.
type SectionTypeDto struct {
	ID             int64  `json:"id"`
	Key            string `json:"key"`
	SectionTypeKey string `json:"sectionTypeKey"`
}

// LapDto is a single lap inside GarminActivitySplits.
type LapDto struct {
	StartTimeGmt             time.Time                    `json:"startTimeGMT"`
	StartLatitude            float64                      `json:"startLatitude"`
	StartLongitude           float64                      `json:"startLongitude"`
	Distance                 float64                      `json:"distance"`
	Duration                 float64                      `json:"duration"`
	MovingDuration           float64                      `json:"movingDuration"`
	ElapsedDuration          float64                      `json:"elapsedDuration"`
	ElevationGain            float64                      `json:"elevationGain"`
	ElevationLoss            float64                      `json:"elevationLoss"`
	MaxElevation             float64                      `json:"maxElevation"`
	MinElevation             float64                      `json:"minElevation"`
	AverageSpeed             float64                      `json:"averageSpeed"`
	AverageMovingSpeed       float64                      `json:"averageMovingSpeed"`
	MaxSpeed                 float64                      `json:"maxSpeed"`
	Calories                 float64                      `json:"calories"`
	BmrCalories              float64                      `json:"bmrCalories"`
	AverageHr                float64                      `json:"averageHR"`
	MaxHr                    float64                      `json:"maxHR"`
	AverageRunCadence        float64                      `json:"averageRunCadence"`
	MaxRunCadence            float64                      `json:"maxRunCadence"`
	AverageTemperature       float64                      `json:"averageTemperature"`
	MaxTemperature           float64                      `json:"maxTemperature"`
	MinTemperature           float64                      `json:"minTemperature"`
	GroundContactTime        float64                      `json:"groundContactTime"`
	GroundContactBalanceLeft float64                      `json:"groundContactBalanceLeft"`
	StrideLength             float64                      `json:"strideLength"`
	VerticalOscillation      float64                      `json:"verticalOscillation"`
	VerticalRatio            float64                      `json:"verticalRatio"`
	EndLatitude              float64                      `json:"endLatitude"`
	EndLongitude             float64                      `json:"endLongitude"`
	MaxVerticalSpeed         float64                      `json:"maxVerticalSpeed"`
	MaxRespirationRate       float64                      `json:"maxRespirationRate"`
	AvgRespirationRate       float64                      `json:"avgRespirationRate"`
	LapIndex                 int64                        `json:"lapIndex"`
	LengthDtos               []LengthDto                  `json:"lengthDTOs"`
	ConnectIqMeasurement     []GarminConnectIqMeasurement `json:"connectIQMeasurement"`
	IntensityType            string                       `json:"intensityType"`
	MessageIndex             int64                        `json:"messageIndex"`
}

// LengthDto is a swim length inside a LapDto.
type LengthDto struct {
	StartTimeGmt         time.Time `json:"startTimeGMT"`
	Distance             float64   `json:"distance"`
	Duration             float64   `json:"duration"`
	AverageSpeed         *float64  `json:"averageSpeed"`
	MaxSpeed             float64   `json:"maxSpeed"`
	Calories             float64   `json:"calories"`
	AverageHr            float64   `json:"averageHR"`
	MaxHr                float64   `json:"maxHR"`
	TotalNumberOfStrokes *int64    `json:"totalNumberOfStrokes"`
	AverageSwolf         float64   `json:"averageSWOLF"`
	LengthIndex          int64     `json:"lengthIndex"`
	SwimStroke           string    `json:"swimStroke"`
}

// GarminConnectIqMeasurement is reused across activity records.
type GarminConnectIqMeasurement struct {
	AppID                string `json:"appID"`
	DeveloperFieldNumber int64  `json:"developerFieldNumber"`
	Value                string `json:"value"`
}

// GarminActivityUuid wraps the activity UUID.
type GarminActivityUuid struct {
	Uuid string `json:"uuid"`
}

// GarminHrTimeInZones is one zone in the time-in-zones breakdown.
type GarminHrTimeInZones struct {
	ZoneNumber      int64   `json:"zoneNumber"`
	SecsInZone      float64 `json:"secsInZone"`
	ZoneLowBoundary int64   `json:"zoneLowBoundary"`
}

// GarminSplitSummary is a richer split-summary record returned by the
// dedicated split-summaries endpoint.
type GarminSplitSummary struct {
	ActivityID                            int64              `json:"activityId"`
	ActivityName                          string             `json:"activityName"`
	StartTimeLocal                        LocalDateTime      `json:"startTimeLocal"`
	StartTimeGmt                          LocalDateTime      `json:"startTimeGMT"`
	ActivityType                          ActivityType       `json:"activityType"`
	EventType                             EventType          `json:"eventType"`
	Distance                              float64            `json:"distance"`
	Duration                              float64            `json:"duration"`
	ElapsedDuration                       float64            `json:"elapsedDuration"`
	MovingDuration                        float64            `json:"movingDuration"`
	ElevationGain                         float64            `json:"elevationGain"`
	ElevationLoss                         float64            `json:"elevationLoss"`
	AverageSpeed                          float64            `json:"averageSpeed"`
	MaxSpeed                              float64            `json:"maxSpeed"`
	StartLatitude                         float64            `json:"startLatitude"`
	StartLongitude                        float64            `json:"startLongitude"`
	HasPolyline                           bool               `json:"hasPolyline"`
	HasImages                             bool               `json:"hasImages"`
	OwnerID                               int64              `json:"ownerId"`
	OwnerDisplayName                      string             `json:"ownerDisplayName"`
	OwnerFullName                         string             `json:"ownerFullName"`
	OwnerProfileImageUrlSmall             string             `json:"ownerProfileImageUrlSmall"`
	OwnerProfileImageUrlMedium            string             `json:"ownerProfileImageUrlMedium"`
	OwnerProfileImageUrlLarge             string             `json:"ownerProfileImageUrlLarge"`
	Calories                              int64              `json:"calories"`
	BmrCalories                           int64              `json:"bmrCalories"`
	AverageHr                             int64              `json:"averageHR"`
	MaxHr                                 int64              `json:"maxHR"`
	AverageRunningCadenceInStepsPerMinute float64            `json:"averageRunningCadenceInStepsPerMinute"`
	MaxRunningCadenceInStepsPerMinute     int64              `json:"maxRunningCadenceInStepsPerMinute"`
	Steps                                 int64              `json:"steps"`
	UserRoles                             []string           `json:"userRoles"`
	Privacy                               Privacy            `json:"privacy"`
	UserPro                               bool               `json:"userPro"`
	HasVideo                              bool               `json:"hasVideo"`
	TimeZoneID                            int64              `json:"timeZoneId"`
	BeginTimestamp                        int64              `json:"beginTimestamp"`
	SportTypeID                           int64              `json:"sportTypeId"`
	AerobicTrainingEffect                 float64            `json:"aerobicTrainingEffect"`
	AnaerobicTrainingEffect               int64              `json:"anaerobicTrainingEffect"`
	AvgStrideLength                       float64            `json:"avgStrideLength"`
	VO2MaxValue                           int64              `json:"vO2MaxValue"`
	DeviceID                              int64              `json:"deviceId"`
	MinTemperature                        int64              `json:"minTemperature"`
	MaxTemperature                        int64              `json:"maxTemperature"`
	MinElevation                          float64            `json:"minElevation"`
	MaxElevation                          float64            `json:"maxElevation"`
	MaxDoubleCadence                      int64              `json:"maxDoubleCadence"`
	SummarizedDiveInfo                    SummarizedDiveInfo `json:"summarizedDiveInfo"`
	MaxVerticalSpeed                      float64            `json:"maxVerticalSpeed"`
	Manufacturer                          string             `json:"manufacturer"`
	LocationName                          string             `json:"locationName"`
	LapCount                              int64              `json:"lapCount"`
	EndLatitude                           float64            `json:"endLatitude"`
	EndLongitude                          float64            `json:"endLongitude"`
	WaterEstimated                        int64              `json:"waterEstimated"`
	TrainingEffectLabel                   string             `json:"trainingEffectLabel"`
	ActivityTrainingLoad                  float64            `json:"activityTrainingLoad"`
	MinActivityLapDuration                float64            `json:"minActivityLapDuration"`
	AerobicTrainingEffectMessage          string             `json:"aerobicTrainingEffectMessage"`
	AnaerobicTrainingEffectMessage        string             `json:"anaerobicTrainingEffectMessage"`
	SplitSummaries                        []json.RawMessage  `json:"splitSummaries"`
	HasSplits                             bool               `json:"hasSplits"`
	ModerateIntensityMinutes              int64              `json:"moderateIntensityMinutes"`
	VigorousIntensityMinutes              int64              `json:"vigorousIntensityMinutes"`
	ElevationCorrected                    bool               `json:"elevationCorrected"`
	AtpActivity                           bool               `json:"atpActivity"`
	Parent                                bool               `json:"parent"`
	Favorite                              bool               `json:"favorite"`
	DecoDive                              bool               `json:"decoDive"`
	Pr                                    bool               `json:"pr"`
	Purposeful                            bool               `json:"purposeful"`
	ManualActivity                        bool               `json:"manualActivity"`
	AutoCalcCalories                      bool               `json:"autoCalcCalories"`
}

// GarminExerciseSets is the strength-training exercise breakdown.
type GarminExerciseSets struct {
	ActivityID            int64                        `json:"activityId"`
	ActivityUuid          GarminActivityUuid           `json:"activityUUID"`
	ActivityName          string                       `json:"activityName"`
	UserProfileID         int64                        `json:"userProfileId"`
	IsMultiSportParent    bool                         `json:"isMultiSportParent"`
	ActivityTypeDto       ActivityTypeDto              `json:"activityTypeDTO"`
	EventTypeDto          EventTypeDto                 `json:"eventTypeDTO"`
	AccessControlRuleDto  AccessControlRuleDto         `json:"accessControlRuleDTO"`
	TimeZoneUnitDto       TimeZoneUnitDto              `json:"timeZoneUnitDTO"`
	MetadataDto           MetadataDto                  `json:"metadataDTO"`
	SummaryDto            SummaryDto                   `json:"summaryDTO"`
	ConnectIqMeasurements []GarminConnectIqMeasurement `json:"connectIQMeasurements"`
	LocationName          string                       `json:"locationName"`
}

// AccessControlRuleDto is part of GarminExerciseSets.
type AccessControlRuleDto struct {
	TypeID  int64  `json:"typeId"`
	TypeKey string `json:"typeKey"`
}

// ActivityTypeDto is the verbose activity-type sub-record.
type ActivityTypeDto struct {
	TypeID       int64           `json:"typeId"`
	TypeKey      string          `json:"typeKey"`
	ParentTypeID int64           `json:"parentTypeId"`
	IsHidden     bool            `json:"isHidden"`
	SortOrder    json.RawMessage `json:"sortOrder"`
	Restricted   bool            `json:"restricted"`
	Trimmable    bool            `json:"trimmable"`
}

// EventTypeDto is the verbose event-type sub-record.
type EventTypeDto struct {
	TypeID    int64  `json:"typeId"`
	TypeKey   string `json:"typeKey"`
	SortOrder int64  `json:"sortOrder"`
}

// MetadataDto carries device + upload metadata.
type MetadataDto struct {
	IsOriginal                      bool              `json:"isOriginal"`
	DeviceApplicationInstallationID int64             `json:"deviceApplicationInstallationId"`
	AgentApplicationInstallationID  string            `json:"agentApplicationInstallationId"`
	AgentString                     string            `json:"agentString"`
	FileFormat                      FileFormat        `json:"fileFormat"`
	AssociatedCourseID              string            `json:"associatedCourseId"`
	LastUpdateDate                  time.Time         `json:"lastUpdateDate"`
	UploadedDate                    time.Time         `json:"uploadedDate"`
	VideoUrl                        string            `json:"videoUrl"`
	HasPolyline                     bool              `json:"hasPolyline"`
	HasChartData                    bool              `json:"hasChartData"`
	HasHrTimeInZones                bool              `json:"hasHrTimeInZones"`
	HasPowerTimeInZones             bool              `json:"hasPowerTimeInZones"`
	UserInfoDto                     UserInfoDto       `json:"userInfoDto"`
	ChildIDs                        []json.RawMessage `json:"childIds"`
	ChildActivityTypes              []json.RawMessage `json:"childActivityTypes"`
	Sensors                         []Sensor          `json:"sensors"`
	ActivityImages                  []string          `json:"activityImages"`
	Manufacturer                    string            `json:"manufacturer"`
	LapCount                        int64             `json:"lapCount"`
	DeviceMetaDataDto               DeviceMetaDataDto `json:"deviceMetaDataDTO"`
	HasIntensityIntervals           bool              `json:"hasIntensityIntervals"`
	HasSplits                       bool              `json:"hasSplits"`
	ManualActivity                  bool              `json:"manualActivity"`
	AutoCalcCalories                bool              `json:"autoCalcCalories"`
	PersonalRecord                  bool              `json:"personalRecord"`
	Gcj02                           bool              `json:"gcj02"`
	Favorite                        bool              `json:"favorite"`
	Trimmed                         bool              `json:"trimmed"`
	ElevationCorrected              bool              `json:"elevationCorrected"`
}

// DeviceMetaDataDto identifies the device an activity was recorded on.
type DeviceMetaDataDto struct {
	DeviceID        string `json:"deviceId"`
	DeviceTypePk    int64  `json:"deviceTypePk"`
	DeviceVersionPk int64  `json:"deviceVersionPk"`
}

// FileFormat is the original-file format of an upload.
type FileFormat struct {
	FormatID  int64  `json:"formatId"`
	FormatKey string `json:"formatKey"`
}

// Sensor lists hardware that contributed to the activity.
type Sensor struct {
	Manufacturer      string  `json:"manufacturer"`
	SerialNumber      int64   `json:"serialNumber"`
	Sku               string  `json:"sku"`
	SourceType        string  `json:"sourceType"`
	AntplusDeviceType string  `json:"antplusDeviceType"`
	SoftwareVersion   float64 `json:"softwareVersion"`
	BatteryStatus     string  `json:"batteryStatus"`
}

// UserInfoDto duplicates a tiny slice of the social profile inside metadata.
type UserInfoDto struct {
	UserProfilePk         int64  `json:"userProfilePk"`
	Displayname           string `json:"displayname"`
	Fullname              string `json:"fullname"`
	ProfileImageUrlLarge  string `json:"profileImageUrlLarge"`
	ProfileImageUrlMedium string `json:"profileImageUrlMedium"`
	ProfileImageUrlSmall  string `json:"profileImageUrlSmall"`
	UserPro               bool   `json:"userPro"`
}

// SummaryDto is the wide summary block inside GarminExerciseSets.
type SummaryDto struct {
	StartTimeLocal                 time.Time `json:"startTimeLocal"`
	StartTimeGmt                   time.Time `json:"startTimeGMT"`
	StartLatitude                  float64   `json:"startLatitude"`
	StartLongitude                 float64   `json:"startLongitude"`
	Distance                       float64   `json:"distance"`
	Duration                       float64   `json:"duration"`
	MovingDuration                 float64   `json:"movingDuration"`
	ElapsedDuration                float64   `json:"elapsedDuration"`
	ElevationGain                  float64   `json:"elevationGain"`
	ElevationLoss                  float64   `json:"elevationLoss"`
	MaxElevation                   float64   `json:"maxElevation"`
	MinElevation                   float64   `json:"minElevation"`
	AverageSpeed                   float64   `json:"averageSpeed"`
	AverageMovingSpeed             float64   `json:"averageMovingSpeed"`
	MaxSpeed                       float64   `json:"maxSpeed"`
	Calories                       float64   `json:"calories"`
	AverageHr                      float64   `json:"averageHR"`
	MaxHr                          float64   `json:"maxHR"`
	AverageRunCadence              float64   `json:"averageRunCadence"`
	MaxRunCadence                  float64   `json:"maxRunCadence"`
	AverageTemperature             float64   `json:"averageTemperature"`
	MaxTemperature                 float64   `json:"maxTemperature"`
	MinTemperature                 float64   `json:"minTemperature"`
	GroundContactTime              float64   `json:"groundContactTime"`
	GroundContactBalanceLeft       float64   `json:"groundContactBalanceLeft"`
	StrideLength                   float64   `json:"strideLength"`
	VerticalOscillation            float64   `json:"verticalOscillation"`
	TrainingEffect                 float64   `json:"trainingEffect"`
	AnaerobicTrainingEffect        float64   `json:"anaerobicTrainingEffect"`
	AerobicTrainingEffectMessage   string    `json:"aerobicTrainingEffectMessage"`
	AnaerobicTrainingEffectMessage string    `json:"anaerobicTrainingEffectMessage"`
	VerticalRatio                  float64   `json:"verticalRatio"`
	EndLatitude                    float64   `json:"endLatitude"`
	EndLongitude                   float64   `json:"endLongitude"`
	MaxVerticalSpeed               float64   `json:"maxVerticalSpeed"`
	WaterEstimated                 float64   `json:"waterEstimated"`
	MinRespirationRate             float64   `json:"minRespirationRate"`
	MaxRespirationRate             float64   `json:"maxRespirationRate"`
	AvgRespirationRate             float64   `json:"avgRespirationRate"`
	TrainingEffectLabel            string    `json:"trainingEffectLabel"`
	ActivityTrainingLoad           float64   `json:"activityTrainingLoad"`
	MinActivityLapDuration         float64   `json:"minActivityLapDuration"`
	ModerateIntensityMinutes       int64     `json:"moderateIntensityMinutes"`
	VigorousIntensityMinutes       int64     `json:"vigorousIntensityMinutes"`
}

// TimeZoneUnitDto is the timezone unit descriptor.
type TimeZoneUnitDto struct {
	UnitID   int64   `json:"unitId"`
	UnitKey  string  `json:"unitKey"`
	Factor   float64 `json:"factor"`
	TimeZone string  `json:"timeZone"`
}
