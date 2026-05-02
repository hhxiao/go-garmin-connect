package garmin

import (
	"encoding/json"
	"time"
)

// GarminSocialProfile is the public-facing user profile. Multiple wellness
// endpoints key off DisplayName, so the Client lazy-fetches and caches it on
// first use.
type GarminSocialProfile struct {
	ID                            int64             `json:"id"`
	ProfileID                     int64             `json:"profileId"`
	GarminGUID                    string            `json:"garminGUID"`
	DisplayName                   string            `json:"displayName"`
	FullName                      string            `json:"fullName"`
	UserName                      string            `json:"userName"`
	ProfileImageUrlLarge          string            `json:"profileImageUrlLarge"`
	ProfileImageUrlMedium         string            `json:"profileImageUrlMedium"`
	ProfileImageUrlSmall          string            `json:"profileImageUrlSmall"`
	Location                      string            `json:"location"`
	FacebookUrl                   string            `json:"facebookUrl"`
	TwitterUrl                    string            `json:"twitterUrl"`
	PersonalWebsite               string            `json:"personalWebsite"`
	Motivation                    *int64            `json:"motivation"`
	Bio                           json.RawMessage   `json:"bio"`
	PrimaryActivity               string            `json:"primaryActivity"`
	FavoriteActivityTypes         []string          `json:"favoriteActivityTypes"`
	RunningTrainingSpeed          float64           `json:"runningTrainingSpeed"`
	CyclingTrainingSpeed          float64           `json:"cyclingTrainingSpeed"`
	FavoriteCyclingActivityTypes  []json.RawMessage `json:"favoriteCyclingActivityTypes"`
	CyclingClassification         json.RawMessage   `json:"cyclingClassification"`
	CyclingMaxAvgPower            float64           `json:"cyclingMaxAvgPower"`
	SwimmingTrainingSpeed         float64           `json:"swimmingTrainingSpeed"`
	ProfileVisibility             string            `json:"profileVisibility"`
	ActivityStartVisibility       string            `json:"activityStartVisibility"`
	ActivityMapVisibility         string            `json:"activityMapVisibility"`
	CourseVisibility              string            `json:"courseVisibility"`
	ActivityHeartRateVisibility   string            `json:"activityHeartRateVisibility"`
	ActivityPowerVisibility       string            `json:"activityPowerVisibility"`
	BadgeVisibility               string            `json:"badgeVisibility"`
	ShowAge                       bool              `json:"showAge"`
	ShowWeight                    bool              `json:"showWeight"`
	ShowHeight                    bool              `json:"showHeight"`
	ShowWeightClass               bool              `json:"showWeightClass"`
	ShowAgeRange                  bool              `json:"showAgeRange"`
	ShowGender                    bool              `json:"showGender"`
	ShowActivityClass             bool              `json:"showActivityClass"`
	ShowVo2Max                    bool              `json:"showVO2Max"`
	ShowPersonalRecords           bool              `json:"showPersonalRecords"`
	ShowLast12Months              bool              `json:"showLast12Months"`
	ShowLifetimeTotals            bool              `json:"showLifetimeTotals"`
	ShowUpcomingEvents            bool              `json:"showUpcomingEvents"`
	ShowRecentFavorites           bool              `json:"showRecentFavorites"`
	ShowRecentDevice              bool              `json:"showRecentDevice"`
	ShowRecentGear                bool              `json:"showRecentGear"`
	ShowBadges                    bool              `json:"showBadges"`
	OtherActivity                 string            `json:"otherActivity"`
	OtherPrimaryActivity          json.RawMessage   `json:"otherPrimaryActivity"`
	OtherMotivation               json.RawMessage   `json:"otherMotivation"`
	UserRoles                     []string          `json:"userRoles"`
	NameApproved                  bool              `json:"nameApproved"`
	UserProfileFullName           string            `json:"userProfileFullName"`
	MakeGolfScorecardsPrivate     bool              `json:"makeGolfScorecardsPrivate"`
	AllowGolfLiveScoring          bool              `json:"allowGolfLiveScoring"`
	AllowGolfScoringByConnections bool              `json:"allowGolfScoringByConnections"`
	UserLevel                     int64             `json:"userLevel"`
	UserPoint                     int64             `json:"userPoint"`
	LevelUpdateDate               time.Time         `json:"levelUpdateDate"`
	LevelIsViewed                 bool              `json:"levelIsViewed"`
	LevelPointThreshold           int64             `json:"levelPointThreshold"`
	UserPointOffset               int64             `json:"userPointOffset"`
	UserPro                       bool              `json:"userPro"`
}

// GarminUserSettings is the device/user preferences blob.
type GarminUserSettings struct {
	ID          int64           `json:"id"`
	UserData    UserData        `json:"userData"`
	UserSleep   UserSleep       `json:"userSleep"`
	ConnectDate json.RawMessage `json:"connectDate"`
	SourceType  json.RawMessage `json:"sourceType"`
}

// UserData is a sub-record of GarminUserSettings.
type UserData struct {
	Gender                         string               `json:"gender"`
	Weight                         float64              `json:"weight"`
	Height                         float64              `json:"height"`
	TimeFormat                     string               `json:"timeFormat"`
	BirthDate                      LocalDateTime        `json:"birthDate"`
	MeasurementSystem              string               `json:"measurementSystem"`
	ActivityLevel                  int64                `json:"activityLevel"`
	Handedness                     string               `json:"handedness"`
	PowerFormat                    GarminFormat         `json:"powerFormat"`
	HeartRateFormat                GarminFormat         `json:"heartRateFormat"`
	FirstDayOfWeek                 GarminFirstDayOfWeek `json:"firstDayOfWeek"`
	Vo2MaxRunning                  float64              `json:"vo2MaxRunning"`
	Vo2MaxCycling                  float64              `json:"vo2MaxCycling"`
	LactateThresholdSpeed          float64              `json:"lactateThresholdSpeed"`
	LactateThresholdHeartRate      json.RawMessage      `json:"lactateThresholdHeartRate"`
	DiveNumber                     json.RawMessage      `json:"diveNumber"`
	IntensityMinutesCalcMethod     string               `json:"intensityMinutesCalcMethod"`
	ModerateIntensityMinutesHrZone int64                `json:"moderateIntensityMinutesHrZone"`
	VigorousIntensityMinutesHrZone int64                `json:"vigorousIntensityMinutesHrZone"`
	HydrationMeasurementUnit       string               `json:"hydrationMeasurementUnit"`
	HydrationContainers            []json.RawMessage    `json:"hydrationContainers"`
	HydrationAutoGoalEnabled       bool                 `json:"hydrationAutoGoalEnabled"`
	FirstbeatMaxStressScore        int64                `json:"firstbeatMaxStressScore"`
	FirstbeatCyclingLtTimestamp    int64                `json:"firstbeatCyclingLtTimestamp"`
	FirstbeatRunningLtTimestamp    int64                `json:"firstbeatRunningLtTimestamp"`
	ThresholdHeartRateAutoDetected bool                 `json:"thresholdHeartRateAutoDetected"`
	FtpAutoDetected                bool                 `json:"ftpAutoDetected"`
	TrainingStatusPausedDate       json.RawMessage      `json:"trainingStatusPausedDate"`
	WeatherLocation                WeatherLocation      `json:"weatherLocation"`
	GolfDistanceUnit               string               `json:"golfDistanceUnit"`
	GolfElevationUnit              json.RawMessage      `json:"golfElevationUnit"`
	GolfSpeedUnit                  json.RawMessage      `json:"golfSpeedUnit"`
	ExternalBottomTime             json.RawMessage      `json:"externalBottomTime"`
}

// WeatherLocation is a sub-record of UserData.
type WeatherLocation struct {
	UseFixedLocation json.RawMessage `json:"useFixedLocation"`
	Latitude         json.RawMessage `json:"latitude"`
	Longitude        json.RawMessage `json:"longitude"`
	LocationName     json.RawMessage `json:"locationName"`
	IsoCountryCode   json.RawMessage `json:"isoCountryCode"`
	PostalCode       json.RawMessage `json:"postalCode"`
}

// UserSleep is a sub-record of GarminUserSettings shared with the
// SetUserSleepTimes write request.
type UserSleep struct {
	SleepTime        int64 `json:"sleepTime,omitempty"`
	DefaultSleepTime bool  `json:"defaultSleepTime"`
	WakeTime         int64 `json:"wakeTime,omitempty"`
	DefaultWakeTime  bool  `json:"defaultWakeTime"`
}

// GarminFormat is the time/date/power format descriptor reused across
// settings and preferences.
type GarminFormat struct {
	FormatID      int64  `json:"formatId"`
	FormatKey     string `json:"formatKey"`
	MinFraction   int64  `json:"minFraction"`
	MaxFraction   int64  `json:"maxFraction"`
	GroupingUsed  bool   `json:"groupingUsed"`
	DisplayFormat string `json:"displayFormat"`
}

// GarminFirstDayOfWeek is shared by user settings and preferences.
type GarminFirstDayOfWeek struct {
	DayID              int64  `json:"dayId"`
	DayName            string `json:"dayName"`
	SortOrder          int64  `json:"sortOrder"`
	IsPossibleFirstDay bool   `json:"isPossibleFirstDay"`
}

// GarminUserPreferences is a thin wrapper around the user settings flag set.
type GarminUserPreferences struct {
	DisplayName              string               `json:"displayName"`
	PreferredLocale          string               `json:"preferredLocale"`
	MeasurementSystem        string               `json:"measurementSystem"`
	FirstDayOfWeek           GarminFirstDayOfWeek `json:"firstDayOfWeek"`
	NumberFormat             string               `json:"numberFormat"`
	TimeFormat               GarminFormat         `json:"timeFormat"`
	DateFormat               GarminFormat         `json:"dateFormat"`
	PowerFormat              GarminFormat         `json:"powerFormat"`
	HeartRateFormat          GarminFormat         `json:"heartRateFormat"`
	TimeZone                 string               `json:"timeZone"`
	HydrationMeasurementUnit string               `json:"hydrationMeasurementUnit"`
	HydrationContainers      []json.RawMessage    `json:"hydrationContainers"`
	GolfDistanceUnit         string               `json:"golfDistanceUnit"`
	GolfElevationUnit        json.RawMessage      `json:"golfElevationUnit"`
	GolfSpeedUnit            json.RawMessage      `json:"golfSpeedUnit"`
}
