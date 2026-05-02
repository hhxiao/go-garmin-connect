package garmin

import (
	"encoding/json"
)

// GarminStats is the daily activity/wellness summary returned by
// `GetUserSummary`.
type GarminStats struct {
	UserProfileID                        int64           `json:"userProfileId"`
	TotalKilocalories                    float64         `json:"totalKilocalories"`
	ActiveKilocalories                   float64         `json:"activeKilocalories"`
	BmrKilocalories                      float64         `json:"bmrKilocalories"`
	WellnessKilocalories                 float64         `json:"wellnessKilocalories"`
	BurnedKilocalories                   json.RawMessage `json:"burnedKilocalories"`
	ConsumedKilocalories                 float64         `json:"consumedKilocalories"`
	RemainingKilocalories                float64         `json:"remainingKilocalories"`
	TotalSteps                           int64           `json:"totalSteps"`
	NetCalorieGoal                       int64           `json:"netCalorieGoal"`
	TotalDistanceMeters                  int64           `json:"totalDistanceMeters"`
	WellnessDistanceMeters               int64           `json:"wellnessDistanceMeters"`
	WellnessActiveKilocalories           float64         `json:"wellnessActiveKilocalories"`
	NetRemainingKilocalories             float64         `json:"netRemainingKilocalories"`
	UserDailySummaryID                   int64           `json:"userDailySummaryId"`
	CalendarDate                         Date            `json:"calendarDate"`
	Rule                                 Rule            `json:"rule"`
	Uuid                                 string          `json:"uuid"`
	DailyStepGoal                        int64           `json:"dailyStepGoal"`
	WellnessStartTimeGmt                 LocalDateTime   `json:"wellnessStartTimeGmt"`
	WellnessStartTimeLocal               LocalDateTime   `json:"wellnessStartTimeLocal"`
	WellnessEndTimeGmt                   LocalDateTime   `json:"wellnessEndTimeGmt"`
	WellnessEndTimeLocal                 LocalDateTime   `json:"wellnessEndTimeLocal"`
	DurationInMilliseconds               int64           `json:"durationInMilliseconds"`
	WellnessDescription                  json.RawMessage `json:"wellnessDescription"`
	HighlyActiveSeconds                  int64           `json:"highlyActiveSeconds"`
	ActiveSeconds                        int64           `json:"activeSeconds"`
	SedentarySeconds                     int64           `json:"sedentarySeconds"`
	SleepingSeconds                      int64           `json:"sleepingSeconds"`
	IncludesWellnessData                 bool            `json:"includesWellnessData"`
	IncludesActivityData                 bool            `json:"includesActivityData"`
	IncludesCalorieConsumedData          bool            `json:"includesCalorieConsumedData"`
	PrivacyProtected                     bool            `json:"privacyProtected"`
	ModerateIntensityMinutes             int64           `json:"moderateIntensityMinutes"`
	VigorousIntensityMinutes             int64           `json:"vigorousIntensityMinutes"`
	FloorsAscendedInMeters               float64         `json:"floorsAscendedInMeters"`
	FloorsDescendedInMeters              float64         `json:"floorsDescendedInMeters"`
	FloorsAscended                       float64         `json:"floorsAscended"`
	FloorsDescended                      float64         `json:"floorsDescended"`
	IntensityMinutesGoal                 int64           `json:"intensityMinutesGoal"`
	UserFloorsAscendedGoal               int64           `json:"userFloorsAscendedGoal"`
	MinHeartRate                         int64           `json:"minHeartRate"`
	MaxHeartRate                         int64           `json:"maxHeartRate"`
	RestingHeartRate                     int64           `json:"restingHeartRate"`
	LastSevenDaysAvgRestingHeartRate     int64           `json:"lastSevenDaysAvgRestingHeartRate"`
	Source                               string          `json:"source"`
	AverageStressLevel                   int64           `json:"averageStressLevel"`
	MaxStressLevel                       int64           `json:"maxStressLevel"`
	StressDuration                       int64           `json:"stressDuration"`
	RestStressDuration                   int64           `json:"restStressDuration"`
	ActivityStressDuration               int64           `json:"activityStressDuration"`
	UncategorizedStressDuration          int64           `json:"uncategorizedStressDuration"`
	TotalStressDuration                  int64           `json:"totalStressDuration"`
	LowStressDuration                    int64           `json:"lowStressDuration"`
	MediumStressDuration                 int64           `json:"mediumStressDuration"`
	HighStressDuration                   int64           `json:"highStressDuration"`
	StressPercentage                     float64         `json:"stressPercentage"`
	RestStressPercentage                 float64         `json:"restStressPercentage"`
	ActivityStressPercentage             float64         `json:"activityStressPercentage"`
	UncategorizedStressPercentage        float64         `json:"uncategorizedStressPercentage"`
	LowStressPercentage                  float64         `json:"lowStressPercentage"`
	MediumStressPercentage               float64         `json:"mediumStressPercentage"`
	HighStressPercentage                 float64         `json:"highStressPercentage"`
	StressQualifier                      string          `json:"stressQualifier"`
	MeasurableAwakeDuration              int64           `json:"measurableAwakeDuration"`
	MeasurableAsleepDuration             int64           `json:"measurableAsleepDuration"`
	LastSyncTimestampGmt                 json.RawMessage `json:"lastSyncTimestampGMT"`
	MinAvgHeartRate                      int64           `json:"minAvgHeartRate"`
	MaxAvgHeartRate                      int64           `json:"maxAvgHeartRate"`
	BodyBatteryChargedValue              int64           `json:"bodyBatteryChargedValue"`
	BodyBatteryDrainedValue              int64           `json:"bodyBatteryDrainedValue"`
	BodyBatteryHighestValue              int64           `json:"bodyBatteryHighestValue"`
	BodyBatteryLowestValue               int64           `json:"bodyBatteryLowestValue"`
	BodyBatteryMostRecentValue           int64           `json:"bodyBatteryMostRecentValue"`
	BodyBatteryVersion                   float64         `json:"bodyBatteryVersion"`
	AbnormalHeartRateAlertsCount         json.RawMessage `json:"abnormalHeartRateAlertsCount"`
	AverageSpo2                          json.RawMessage `json:"averageSpo2"`
	LowestSpo2                           json.RawMessage `json:"lowestSpo2"`
	LatestSpo2                           json.RawMessage `json:"latestSpo2"`
	LatestSpo2ReadingTimeGmt             json.RawMessage `json:"latestSpo2ReadingTimeGmt"`
	LatestSpo2ReadingTimeLocal           json.RawMessage `json:"latestSpo2ReadingTimeLocal"`
	AverageMonitoringEnvironmentAltitude float64         `json:"averageMonitoringEnvironmentAltitude"`
	AvgWakingRespirationValue            float64         `json:"avgWakingRespirationValue"`
	HighestRespirationValue              float64         `json:"highestRespirationValue"`
	LowestRespirationValue               float64         `json:"lowestRespirationValue"`
	LatestRespirationValue               float64         `json:"latestRespirationValue"`
	LatestRespirationTimeGmt             LocalDateTime   `json:"latestRespirationTimeGMT"`
}

// Rule is a sub-record of GarminStats describing the wellness rule type.
type Rule struct {
	TypeID  int64  `json:"typeId"`
	TypeKey string `json:"typeKey"`
}

// GarminStepsData is one entry in the steps timeseries.
type GarminStepsData struct {
	StartGmt              LocalDateTime `json:"startGMT"`
	EndGmt                LocalDateTime `json:"endGMT"`
	Steps                 int64         `json:"steps"`
	PrimaryActivityLevel  string        `json:"primaryActivityLevel"`
	ActivityLevelConstant bool          `json:"activityLevelConstant"`
}

// GarminHr is the daily heart-rate timeseries.
type GarminHr struct {
	UserProfilePk                    int64                      `json:"userProfilePK"`
	CalendarDate                     Date                       `json:"calendarDate"`
	StartTimestampGmt                LocalDateTime              `json:"startTimestampGMT"`
	EndTimestampGmt                  LocalDateTime              `json:"endTimestampGMT"`
	StartTimestampLocal              LocalDateTime              `json:"startTimestampLocal"`
	EndTimestampLocal                LocalDateTime              `json:"endTimestampLocal"`
	MaxHeartRate                     int64                      `json:"maxHeartRate"`
	MinHeartRate                     int64                      `json:"minHeartRate"`
	RestingHeartRate                 int64                      `json:"restingHeartRate"`
	LastSevenDaysAvgRestingHeartRate int64                      `json:"lastSevenDaysAvgRestingHeartRate"`
	HeartRateValueDescriptors        []HeartRateValueDescriptor `json:"heartRateValueDescriptors"`
	HeartRateValues                  [][]int64                  `json:"heartRateValues"`
}

// HeartRateValueDescriptor maps an index into HeartRateValues to a key.
type HeartRateValueDescriptor struct {
	Key   string `json:"key"`
	Index int64  `json:"index"`
}

// GarminHydrationData is the daily hydration summary.
type GarminHydrationData struct {
	UserID                  int64          `json:"userId"`
	CalendarDate            Date           `json:"calendarDate"`
	ValueInMl               float64        `json:"valueInML"`
	GoalInMl                float64        `json:"goalInML"`
	DailyAverageInMl        float64        `json:"dailyAverageinML"`
	LastEntryTimestampLocal *LocalDateTime `json:"lastEntryTimestampLocal"`
	SweatLossInMl           float64        `json:"sweatLossInML"`
	ActivityIntakeInMl      float64        `json:"activityIntakeInML"`
}

// GarminSleepData is the daily sleep summary.
type GarminSleepData struct {
	DailySleepDto                       DailySleepDto                     `json:"dailySleepDTO"`
	SleepMovement                       []Sleep                           `json:"sleepMovement"`
	RemSleepData                        bool                              `json:"remSleepData"`
	SleepLevels                         []Sleep                           `json:"sleepLevels"`
	WellnessEpochRespirationDataDtoList []WellnessEpochRespirationDtoList `json:"wellnessEpochRespirationDataDTOList"`
	SleepStress                         []SleepStress                     `json:"sleepStress"`
}

// DailySleepDto is the headline daily sleep summary.
type DailySleepDto struct {
	ID                          int64           `json:"id"`
	UserProfilePk               int64           `json:"userProfilePK"`
	CalendarDate                Date            `json:"calendarDate"`
	SleepTimeSeconds            int64           `json:"sleepTimeSeconds"`
	NapTimeSeconds              int64           `json:"napTimeSeconds"`
	SleepWindowConfirmed        bool            `json:"sleepWindowConfirmed"`
	SleepWindowConfirmationType string          `json:"sleepWindowConfirmationType"`
	SleepStartTimestampGmt      int64           `json:"sleepStartTimestampGMT"`
	SleepEndTimestampGmt        int64           `json:"sleepEndTimestampGMT"`
	SleepStartTimestampLocal    int64           `json:"sleepStartTimestampLocal"`
	SleepEndTimestampLocal      int64           `json:"sleepEndTimestampLocal"`
	AutoSleepStartTimestampGmt  json.RawMessage `json:"autoSleepStartTimestampGMT"`
	AutoSleepEndTimestampGmt    json.RawMessage `json:"autoSleepEndTimestampGMT"`
	SleepQualityTypePk          json.RawMessage `json:"sleepQualityTypePK"`
	SleepResultTypePk           json.RawMessage `json:"sleepResultTypePK"`
	UnmeasurableSleepSeconds    int64           `json:"unmeasurableSleepSeconds"`
	DeepSleepSeconds            int64           `json:"deepSleepSeconds"`
	LightSleepSeconds           int64           `json:"lightSleepSeconds"`
	RemSleepSeconds             int64           `json:"remSleepSeconds"`
	AwakeSleepSeconds           int64           `json:"awakeSleepSeconds"`
	DeviceRemCapable            bool            `json:"deviceRemCapable"`
	Retro                       bool            `json:"retro"`
	SleepFromDevice             bool            `json:"sleepFromDevice"`
	AverageRespirationValue     float64         `json:"averageRespirationValue"`
	LowestRespirationValue      float64         `json:"lowestRespirationValue"`
	HighestRespirationValue     float64         `json:"highestRespirationValue"`
	AwakeCount                  int64           `json:"awakeCount"`
	AvgSleepStress              float64         `json:"avgSleepStress"`
	AgeGroup                    string          `json:"ageGroup"`
	SleepScoreFeedback          string          `json:"sleepScoreFeedback"`
	SleepScoreInsight           string          `json:"sleepScoreInsight"`
	SleepScores                 SleepScores     `json:"sleepScores"`
}

// SleepScores groups the various sub-scores Garmin reports.
type SleepScores struct {
	TotalDuration   AwakeCount   `json:"totalDuration"`
	Stress          AwakeCount   `json:"stress"`
	AwakeCount      AwakeCount   `json:"awakeCount"`
	Overall         Overall      `json:"overall"`
	RemPercentage   Percentage   `json:"remPercentage"`
	Restlessness    Restlessness `json:"restlessness"`
	LightPercentage Percentage   `json:"lightPercentage"`
	DeepPercentage  Percentage   `json:"deepPercentage"`
}

// AwakeCount is the qualifier+optimal-range tuple Garmin uses for several
// sleep sub-scores.
type AwakeCount struct {
	QualifierKey string  `json:"qualifierKey"`
	OptimalStart float64 `json:"optimalStart"`
	OptimalEnd   float64 `json:"optimalEnd"`
}

// Percentage is similar to AwakeCount but adds an ideal-window range.
type Percentage struct {
	QualifierKey        string  `json:"qualifierKey"`
	OptimalStart        float64 `json:"optimalStart"`
	OptimalEnd          float64 `json:"optimalEnd"`
	IdealStartInSeconds float64 `json:"idealStartInSeconds"`
	IdealEndInSeconds   float64 `json:"idealEndInSeconds"`
}

// Overall is the headline sleep score.
type Overall struct {
	Value        int64  `json:"value"`
	QualifierKey string `json:"qualifierKey"`
}

// Restlessness is the qualifier-only restlessness score.
type Restlessness struct {
	QualifierKey string `json:"qualifierKey"`
}

// Sleep is one bin in the sleep movement/stage timeseries.
type Sleep struct {
	StartGmt      LocalDateTime `json:"startGMT"`
	EndGmt        LocalDateTime `json:"endGMT"`
	ActivityLevel float64       `json:"activityLevel"`
}

// SleepStress is one bin in the sleep stress timeseries.
type SleepStress struct {
	Value    int64 `json:"value"`
	StartGmt int64 `json:"startGMT"`
}

// WellnessEpochRespirationDtoList is one bin in the respiration timeseries.
type WellnessEpochRespirationDtoList struct {
	StartTimeGmt     int64   `json:"startTimeGMT"`
	RespirationValue float64 `json:"respirationValue"`
}

// GarminBodyBatteryData is the daily body-battery timeseries + events.
type GarminBodyBatteryData struct {
	Date                                    Date                                `json:"date"`
	Charged                                 int64                               `json:"charged"`
	Drained                                 int64                               `json:"drained"`
	StartTimestampGmt                       LocalDateTime                       `json:"startTimestampGMT"`
	EndTimestampGmt                         LocalDateTime                       `json:"endTimestampGMT"`
	StartTimestampLocal                     LocalDateTime                       `json:"startTimestampLocal"`
	EndTimestampLocal                       LocalDateTime                       `json:"endTimestampLocal"`
	BodyBatteryValuesArray                  [][]int64                           `json:"bodyBatteryValuesArray"`
	BodyBatteryValueDescriptorDtoList       []BodyBatteryValueDescriptorDtoList `json:"bodyBatteryValueDescriptorDTOList"`
	BodyBatteryDynamicFeedbackEvent         BodyBatteryDynamicFeedbackEvent     `json:"bodyBatteryDynamicFeedbackEvent"`
	BodyBatteryActivityEvent                []BodyBatteryActivityEvent          `json:"bodyBatteryActivityEvent"`
	EndOfDayBodyBatteryDynamicFeedbackEvent BodyBatteryDynamicFeedbackEvent     `json:"endOfDayBodyBatteryDynamicFeedbackEvent"`
}

// BodyBatteryActivityEvent is one event in the body-battery activity log.
type BodyBatteryActivityEvent struct {
	EventType              string        `json:"eventType"`
	EventStartTimeGmt      LocalDateTime `json:"eventStartTimeGmt"`
	TimezoneOffset         int64         `json:"timezoneOffset"`
	DurationInMilliseconds int64         `json:"durationInMilliseconds"`
	BodyBatteryImpact      int64         `json:"bodyBatteryImpact"`
	FeedbackType           string        `json:"feedbackType"`
	ShortFeedback          string        `json:"shortFeedback"`
}

// BodyBatteryDynamicFeedbackEvent is the start/end-of-day feedback record.
type BodyBatteryDynamicFeedbackEvent struct {
	EventTimestampGmt LocalDateTime `json:"eventTimestampGmt"`
	BodyBatteryLevel  string        `json:"bodyBatteryLevel"`
	FeedbackShortType string        `json:"feedbackShortType"`
	FeedbackLongType  string        `json:"feedbackLongType"`
}

// BodyBatteryValueDescriptorDtoList maps the inner array index to a key.
type BodyBatteryValueDescriptorDtoList struct {
	BodyBatteryValueDescriptorIndex int64  `json:"bodyBatteryValueDescriptorIndex"`
	BodyBatteryValueDescriptorKey   string `json:"bodyBatteryValueDescriptorKey"`
}

// GarminReportHrvStatus is the multi-day HRV summary.
type GarminReportHrvStatus struct {
	HrvSummaries  []GarminHrvSummary `json:"hrvSummaries"`
	UserProfilePk int64              `json:"userProfilePk"`
}

// GarminHrvSummary is one day of HRV data.
type GarminHrvSummary struct {
	CalendarDate      Date              `json:"calendarDate"`
	WeeklyAvg         int               `json:"weeklyAvg"`
	LastNightAvg      int               `json:"lastNightAvg"`
	LastNight5MinHigh int               `json:"lastNight5MinHigh"`
	Baseline          GarminHrvBaseline `json:"baseline"`
	Status            string            `json:"status"`
	FeedbackPhrase    string            `json:"feedbackPhrase"`
	CreateTimeStamp   LocalDateTime     `json:"createTimeStamp"`
}

// GarminHrvBaseline is the baseline range used to compute HRV status.
type GarminHrvBaseline struct {
	LowUpper      int     `json:"lowUpper"`
	BalancedLow   int     `json:"balancedLow"`
	BalancedUpper int     `json:"balancedUpper"`
	MarkerValue   float64 `json:"markerValue"`
}

// GarminBodyComposition is the multi-day weight / body composition snapshot.
type GarminBodyComposition struct {
	StartDate      Date               `json:"startDate"`
	EndDate        Date               `json:"endDate"`
	DateWeightList []GarminDateWeight `json:"dateWeightList"`
	TotalAverage   GarminTotalAverage `json:"totalAverage"`
}

// GarminTotalAverage is the multi-day average inside GarminBodyComposition.
type GarminTotalAverage struct {
	From           int64   `json:"from"`
	Until          int64   `json:"until"`
	Weight         float64 `json:"weight"`
	Bmi            float64 `json:"bmi"`
	BodyFat        float64 `json:"bodyFat"`
	BodyWater      float64 `json:"bodyWater"`
	BoneMass       float64 `json:"boneMass"`
	MuscleMass     float64 `json:"muscleMass"`
	PhysiqueRating float64 `json:"physiqueRating"`
	VisceralFat    float64 `json:"visceralFat"`
	MetabolicAge   float64 `json:"metabolicAge"`
}

// GarminDateWeight is one weight entry inside GarminBodyComposition.
type GarminDateWeight struct {
	SamplePk       int64   `json:"samplePk"`
	Date           int64   `json:"date"`
	CalendarDate   Date    `json:"calendarDate"`
	Weight         float64 `json:"weight"`
	Bmi            float64 `json:"bmi"`
	BodyFat        float64 `json:"bodyFat"`
	BodyWater      float64 `json:"bodyWater"`
	BoneMass       int64   `json:"boneMass"`
	MuscleMass     int64   `json:"muscleMass"`
	PhysiqueRating float64 `json:"physiqueRating"`
	VisceralFat    float64 `json:"visceralFat"`
	MetabolicAge   float64 `json:"metabolicAge"`
	SourceType     string  `json:"sourceType"`
	TimestampGmt   int64   `json:"timestampGMT"`
	WeightDelta    float64 `json:"weightDelta"`
}
