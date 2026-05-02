package garmin

import (
	"errors"
	"time"
)

// WeightUnit selects which unit AddWeight submits.
type WeightUnit int

const (
	// WeightUnitKg measures in kilograms.
	WeightUnitKg WeightUnit = iota
	// WeightUnitLbs measures in pounds.
	WeightUnitLbs
)

// String returns the lowercase form Garmin expects ("kg" / "lbs").
func (w WeightUnit) String() string {
	switch w {
	case WeightUnitKg:
		return "kg"
	case WeightUnitLbs:
		return "lbs"
	}
	return ""
}

// GarminWeight is the input record for AddWeight.
type GarminWeight struct {
	MeasurementDateTime time.Time
	UnitKey             WeightUnit
	Value               float64
}

// IsValid reports whether the weight passes Garmin's plausibility checks
// (1–453 kg covers human + most pet scales).
func (w GarminWeight) IsValid() bool { return w.Value >= 1.0 && w.Value <= 453.0 }

// GarminWeightRange is the multi-day weight summary returned by GetWeightRange.
type GarminWeightRange struct {
	DailyWeightSummaries []GarminWeightDailyWeightSummary `json:"dailyWeightSummaries"`
	TotalAverage         GarminWeightTotalAverage         `json:"totalAverage"`
}

// GarminWeightDailyWeightSummary is the per-day rollup inside GarminWeightRange.
type GarminWeightDailyWeightSummary struct {
	SummaryDate        Date                      `json:"summaryDate"`
	NumOfWeightEntries int64                     `json:"numOfWeightEntries"`
	MinWeight          float64                   `json:"minWeight"`
	MaxWeight          float64                   `json:"maxWeight"`
	LatestWeight       GarminWeightMeasurement   `json:"latestWeight"`
	AllWeightMetrics   []GarminWeightMeasurement `json:"allWeightMetrics"`
}

// GarminWeightIdentifier is the (date, samplePk) pair that uniquely names a
// stored weight measurement.
type GarminWeightIdentifier struct {
	SamplePk     int64 `json:"samplePk"`
	CalendarDate Date  `json:"calendarDate"`
}

// GarminWeightMeasurement extends GarminWeightIdentifier with the actual
// metrics. Embedding mirrors the C# inheritance.
type GarminWeightMeasurement struct {
	GarminWeightIdentifier
	Date         int64    `json:"date"`
	Weight       float64  `json:"weight"`
	Bmi          *float64 `json:"bmi"`
	BodyFat      *float64 `json:"bodyFat"`
	BodyWater    *float64 `json:"bodyWater"`
	BoneMass     *int64   `json:"boneMass"`
	MuscleMass   *int64   `json:"muscleMass"`
	SourceType   string   `json:"sourceType"`
	TimestampGmt int64    `json:"timestampGMT"`
	WeightDelta  *float64 `json:"weightDelta"`
}

// GarminWeightTotalAverage is the multi-day average inside GarminWeightRange.
type GarminWeightTotalAverage struct {
	From       int64   `json:"from"`
	Until      int64   `json:"until"`
	Weight     float64 `json:"weight"`
	Bmi        float64 `json:"bmi"`
	BodyFat    float64 `json:"bodyFat"`
	BodyWater  float64 `json:"bodyWater"`
	BoneMass   int64   `json:"boneMass"`
	MuscleMass int64   `json:"muscleMass"`
}

// GarminBloodPressure is the input record for AddBloodPressure.
type GarminBloodPressure struct {
	Systolic            int64
	Diastolic           int64
	Pulse               int64
	Notes               string
	MeasurementDateTime time.Time
}

// IsValid mirrors the original C# bounds: diastolic 30–200, systolic 40–300,
// pulse 1–300. Returns nil when the record passes all checks.
func (b GarminBloodPressure) Validate() error {
	if b.Diastolic < 30 || b.Diastolic > 200 {
		return errors.New("diastolic must be between 30 and 200")
	}
	if b.Systolic < 40 || b.Systolic > 300 {
		return errors.New("systolic must be between 40 and 300")
	}
	if b.Pulse < 1 || b.Pulse > 300 {
		return errors.New("pulse must be between 1 and 300")
	}
	return nil
}

// IsValid is the boolean shortcut for Validate(); kept for parity with the C#
// API.
func (b GarminBloodPressure) IsValid() bool { return b.Validate() == nil }

// GarminBloodPressureDaily is one day's worth of blood-pressure entries.
type GarminBloodPressureDaily struct {
	StartDate                 Date                             `json:"startDate"`
	EndDate                   Date                             `json:"endDate"`
	BloodPressureMeasurements []GarminBloodPressureMeasurement `json:"bloodPressureMeasurements"`
	TotalMeasurementCount     *int64                           `json:"totalMeasurementCount"`
	ElevatedMeasurementCount  *int64                           `json:"elevatedMeasurementCount"`
}

// GarminBloodPressureIdentifier is the (timestamp, version) pair that
// uniquely names a stored blood-pressure measurement.
type GarminBloodPressureIdentifier struct {
	Version                 int64         `json:"version"`
	MeasurementTimestampGmt LocalDateTime `json:"measurementTimestampGMT"`
}

// GarminBloodPressureMeasurement extends the identifier with the actual
// readings.
type GarminBloodPressureMeasurement struct {
	GarminBloodPressureIdentifier
	Systolic                  int64         `json:"systolic"`
	Diastolic                 int64         `json:"diastolic"`
	Pulse                     int64         `json:"pulse"`
	MultiMeasurement          bool          `json:"multiMeasurement"`
	Notes                     string        `json:"notes"`
	SourceType                string        `json:"sourceType"`
	MeasurementTimestampLocal LocalDateTime `json:"measurementTimestampLocal"`
	Category                  string        `json:"category"`
	CategoryName              string        `json:"categoryName"`
}

// GarminGear describes one piece of gear the user has registered.
type GarminGear struct {
	Uuid            string         `json:"uuid"`
	GearPk          int            `json:"gearPk"`
	UserProfileID   int64          `json:"userProfilePk"`
	GearMakeName    string         `json:"gearMakeName"`
	GearModelName   string         `json:"gearModelName"`
	GearTypeName    string         `json:"gearTypeName"`
	DisplayName     string         `json:"displayName"`
	CustomMakeModel string         `json:"customMakeModel"`
	ImageNameLarge  string         `json:"imageNameLarge"`
	ImageNameMedium string         `json:"imageNameMedium"`
	ImageNameSmall  string         `json:"imageNameSmall"`
	DateBegin       LocalDateTime  `json:"dateBegin"`
	DateEnd         *LocalDateTime `json:"dateEnd"`
	MaximumMeters   float64        `json:"maximumMeters"`
	Notified        bool           `json:"notified"`
	CreateDate      LocalDateTime  `json:"createDate"`
	UpdateDate      LocalDateTime  `json:"updateDate"`
}

// GarminGearType is one available gear-type lookup entry.
type GarminGearType struct {
	GearTypePk   int            `json:"gearTypePk"`
	GearTypeName string         `json:"gearTypeName"`
	CreateDate   LocalDateTime  `json:"createDate"`
	UpdateData   *LocalDateTime `json:"updateData"`
}

// GarminPersonalRecord is one entry in the personal-records list.
type GarminPersonalRecord struct {
	ID                                  int64          `json:"id"`
	TypeID                              int64          `json:"typeId"`
	ActivityID                          int64          `json:"activityId"`
	ActivityName                        string         `json:"activityName"`
	ActivityStartDateTimeInGmt          int64          `json:"activityStartDateTimeInGMT"`
	ActStartDateTimeInGmtFormatted      *LocalDateTime `json:"actStartDateTimeInGMTFormatted"`
	ActivityStartDateTimeLocal          *int64         `json:"activityStartDateTimeLocal"`
	ActivityStartDateTimeLocalFormatted *LocalDateTime `json:"activityStartDateTimeLocalFormatted"`
	Value                               float64        `json:"value"`
	PrStartTimeGmt                      int64          `json:"prStartTimeGmt"`
	PrStartTimeGmtFormatted             LocalDateTime  `json:"prStartTimeGmtFormatted"`
	PrStartTimeLocal                    int64          `json:"prStartTimeLocal"`
	PrStartTimeLocalFormatted           *LocalDateTime `json:"prStartTimeLocalFormatted"`
	PrTypeLabelKey                      string         `json:"prTypeLabelKey"`
	PoolLengthUnit                      PoolLengthUnit `json:"poolLengthUnit"`
}

// PoolLengthUnit is the swim-pool unit descriptor used by personal records.
type PoolLengthUnit struct {
	ID     int64   `json:"id"`
	Key    string  `json:"key"`
	Factor float64 `json:"factor"`
}
