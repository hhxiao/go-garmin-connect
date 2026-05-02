package garmin

import (
	"encoding/json"
)

// GarminMonth is a 0-based month index used by GetCalendarByMonth (Garmin's
// API treats January as 0). Use NewGarminMonth(time.Month) to convert from
// the standard library.
type GarminMonth int

const (
	MonthJanuary GarminMonth = iota
	MonthFebruary
	MonthMarch
	MonthApril
	MonthMay
	MonthJune
	MonthJuly
	MonthAugust
	MonthSeptember
	MonthOctober
	MonthNovember
	MonthDecember
)

// GarminCalendarYear is the year-overview blob.
type GarminCalendarYear struct {
	StartDayOfJanuary int                 `json:"startDayOfJanuary"`
	LeapYear          bool                `json:"leapYear"`
	YearItems         []GarminYearItem    `json:"yearItems"`
	YearSummaries     []GarminYearSummary `json:"yearSummaries"`
}

// GarminYearItem is one day inside the year-overview heatmap.
type GarminYearItem struct {
	Date    Date  `json:"date"`
	Display int64 `json:"display"`
}

// GarminYearSummary is per-activity-type aggregate data.
type GarminYearSummary struct {
	ActivityTypeID     int64   `json:"activityTypeId"`
	NumberOfActivities int     `json:"numberOfActivities"`
	TotalDistance      float64 `json:"totalDistance"`
	TotalDuration      float64 `json:"totalDuration"`
	TotalCalories      float64 `json:"totalCalories"`
}

// GarminCalendarMonth is the month-overview blob.
type GarminCalendarMonth struct {
	StartDayOfMonth      int                  `json:"startDayOfMonth"`
	NumOfDaysInMonth     int                  `json:"numOfDaysInMonth"`
	NumOfDaysInPrevMonth int                  `json:"numOfDaysInPrevMonth"`
	Month                GarminMonth          `json:"month"`
	Year                 int                  `json:"year"`
	CalendarItems        []GarminCalendarItem `json:"calendarItems"`
}

// GarminCalendarItem is one entry on the calendar (activity, badge, etc.).
type GarminCalendarItem struct {
	ID                         int64                  `json:"id"`
	GroupID                    *int64                 `json:"groupId"`
	TrainingPlanID             *int64                 `json:"trainingPlanId"`
	ItemType                   string                 `json:"itemType"`
	ActivityTypeID             *int64                 `json:"activityTypeId"`
	WellnessActivityUuid       string                 `json:"wellnessActivityUuid"`
	Title                      string                 `json:"title"`
	Date                       Date                   `json:"date"`
	Duration                   *float64               `json:"duration"`
	Distance                   *float64               `json:"distance"`
	Calories                   *float64               `json:"calories"`
	FloorsClimbed              json.RawMessage        `json:"floorsClimbed"`
	AvgRespirationRate         *float64               `json:"avgRespirationRate"`
	UnitOfPoolLength           GarminUnitOfPoolLength `json:"unitOfPoolLength"`
	Weight                     *float64               `json:"weight"`
	Difference                 *float64               `json:"difference"`
	CourseID                   json.RawMessage        `json:"courseId"`
	CourseName                 json.RawMessage        `json:"courseName"`
	SportTypeKey               string                 `json:"sportTypeKey"`
	Url                        string                 `json:"url"`
	IsStart                    bool                   `json:"isStart"`
	IsRace                     bool                   `json:"isRace"`
	RecurrenceID               json.RawMessage        `json:"recurrenceId"`
	IsParent                   *bool                  `json:"isParent"`
	ParentID                   *int64                 `json:"parentId"`
	UserBadgeID                *int64                 `json:"userBadgeId"`
	BadgeCategoryTypeID        *int64                 `json:"badgeCategoryTypeId"`
	BadgeCategoryTypeDesc      string                 `json:"badgeCategoryTypeDesc"`
	BadgeAwardedDate           json.RawMessage        `json:"badgeAwardedDate"`
	BadgeViewed                json.RawMessage        `json:"badgeViewed"`
	HideBadge                  json.RawMessage        `json:"hideBadge"`
	StartTimestampLocal        *LocalDateTime         `json:"startTimestampLocal"`
	EventTimeLocal             json.RawMessage        `json:"eventTimeLocal"`
	DiveNumber                 json.RawMessage        `json:"diveNumber"`
	MaxDepth                   json.RawMessage        `json:"maxDepth"`
	AvgDepth                   json.RawMessage        `json:"avgDepth"`
	SurfaceInterval            json.RawMessage        `json:"surfaceInterval"`
	ElapsedDuration            *float64               `json:"elapsedDuration"`
	LapCount                   *int64                 `json:"lapCount"`
	BottomTime                 json.RawMessage        `json:"bottomTime"`
	AtpPlanID                  json.RawMessage        `json:"atpPlanId"`
	WorkoutID                  *int64                 `json:"workoutId"`
	ProtectedWorkoutSchedule   bool                   `json:"protectedWorkoutSchedule"`
	ActiveSets                 json.RawMessage        `json:"activeSets"`
	Strokes                    *float64               `json:"strokes"`
	NoOfSplits                 json.RawMessage        `json:"noOfSplits"`
	MaxGradeValue              json.RawMessage        `json:"maxGradeValue"`
	TotalAscent                *float64               `json:"totalAscent"`
	DifferenceStress           json.RawMessage        `json:"differenceStress"`
	ClimbDuration              *float64               `json:"climbDuration"`
	MaxSpeed                   *float64               `json:"maxSpeed"`
	AverageHr                  *float64               `json:"averageHR"`
	ActiveSplitSummaryDuration *float64               `json:"activeSplitSummaryDuration"`
	MaxSplitDistance           *int64                 `json:"maxSplitDistance"`
	MaxSplitSpeed              *float64               `json:"maxSplitSpeed"`
	Location                   json.RawMessage        `json:"location"`
	ShareableEventUuid         json.RawMessage        `json:"shareableEventUuid"`
	SplitSummaryMode           json.RawMessage        `json:"splitSummaryMode"`
	CompletionTarget           json.RawMessage        `json:"completionTarget"`
	DecoDive                   *bool                  `json:"decoDive"`
	AutoCalcCalories           *bool                  `json:"autoCalcCalories"`
	ShareableEvent             bool                   `json:"shareableEvent"`
	PhasedTrainingPlan         *bool                  `json:"phasedTrainingPlan"`
	PrimaryEvent               json.RawMessage        `json:"primaryEvent"`
	Subscribed                 json.RawMessage        `json:"subscribed"`
}

// GarminUnitOfPoolLength describes a swim-pool's length unit.
type GarminUnitOfPoolLength struct {
	UnitID  int64   `json:"unitId"`
	UnitKey string  `json:"unitKey"`
	Factor  float64 `json:"factor"`
}

// GarminCalendarWeek is the week-overview blob.
type GarminCalendarWeek struct {
	StartDate        Date                 `json:"startDate"`
	EndDate          Date                 `json:"endDate"`
	NumOfDaysInMonth int                  `json:"numOfDaysInMonth"`
	CalendarItems    []GarminCalendarItem `json:"calendarItems"`
}
