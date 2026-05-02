package garmin

import (
	"encoding/json"
	"time"
)

// GarminWorkout is the full workout definition (header + segments).
type GarminWorkout struct {
	WorkoutID                 int64                  `json:"workoutId"`
	OwnerID                   int64                  `json:"ownerId"`
	WorkoutName               string                 `json:"workoutName"`
	Description               string                 `json:"description"`
	UpdateDate                LocalDateTime          `json:"updateDate"`
	CreatedDate               LocalDateTime          `json:"createdDate"`
	SportType                 GarminSportType        `json:"sportType"`
	TrainingPlanID            json.RawMessage        `json:"trainingPlanId"`
	Author                    GarminAuthor           `json:"author"`
	EstimatedDurationInSecs   int64                  `json:"estimatedDurationInSecs"`
	EstimatedDistanceInMeters float64                `json:"estimatedDistanceInMeters"`
	EstimateType              json.RawMessage        `json:"estimateType"`
	PoolLength                float64                `json:"poolLength"`
	PoolLengthUnit            GarminUnit             `json:"poolLengthUnit"`
	WorkoutProvider           string                 `json:"workoutProvider"`
	WorkoutSourceID           string                 `json:"workoutSourceId"`
	Consumer                  string                 `json:"consumer"`
	AtpPlanID                 *int64                 `json:"atpPlanId"`
	WorkoutNameI18NKey        string                 `json:"workoutNameI18nKey"`
	DescriptionI18NKey        string                 `json:"descriptionI18nKey"`
	EstimatedDistanceUnit     GarminUnit             `json:"estimatedDistanceUnit"`
	Shared                    bool                   `json:"shared"`
	Estimated                 bool                   `json:"estimated"`
	AvgTrainingSpeed          float64                `json:"avgTrainingSpeed"`
	SubSportType              json.RawMessage        `json:"subSportType"`
	SharedWithUsers           json.RawMessage        `json:"sharedWithUsers"`
	WorkoutSegments           []GarminWorkoutSegment `json:"workoutSegments"`
	UploadTimestamp           json.RawMessage        `json:"uploadTimestamp"`
}

// GarminAuthor is reused by GarminWorkout and the workout-update payload.
type GarminAuthor struct {
	UserProfilePk        *int64 `json:"userProfilePk"`
	DisplayName          string `json:"displayName"`
	FullName             string `json:"fullName"`
	ProfileImgNameLarge  string `json:"profileImgNameLarge"`
	ProfileImgNameMedium string `json:"profileImgNameMedium"`
	ProfileImgNameSmall  string `json:"profileImgNameSmall"`
	UserPro              bool   `json:"userPro"`
	VivokidUser          bool   `json:"vivokidUser"`
}

// GarminUnit is a measurement-unit descriptor reused across workout fields.
type GarminUnit struct {
	UnitID  *int64   `json:"unitId"`
	UnitKey string   `json:"unitKey"`
	Factor  *float64 `json:"factor"`
}

// GarminSportType describes a workout's sport.
type GarminSportType struct {
	SportTypeID  int64  `json:"sportTypeId"`
	SportTypeKey string `json:"sportTypeKey"`
	DisplayOrder int64  `json:"displayOrder"`
}

// GarminWorkoutSegment is one segment within a workout.
type GarminWorkoutSegment struct {
	SegmentOrder int64               `json:"segmentOrder"`
	SportType    GarminSportType     `json:"sportType"`
	WorkoutSteps []GarminWorkoutStep `json:"workoutSteps"`
}

// GarminWorkoutStep is one step within a segment.
type GarminWorkoutStep struct {
	Type                      string                     `json:"type"`
	StepID                    int64                      `json:"stepId"`
	StepOrder                 int64                      `json:"stepOrder"`
	StepType                  GarminWorkoutStepType      `json:"stepType"`
	ChildStepID               json.RawMessage            `json:"childStepId"`
	Description               json.RawMessage            `json:"description"`
	EndCondition              GarminWorkoutEndCondition  `json:"endCondition"`
	EndConditionValue         *float64                   `json:"endConditionValue"`
	PreferredEndConditionUnit json.RawMessage            `json:"preferredEndConditionUnit"`
	EndConditionCompare       json.RawMessage            `json:"endConditionCompare"`
	TargetType                GarminWorkoutTargetType    `json:"targetType"`
	TargetValueOne            *float64                   `json:"targetValueOne"`
	TargetValueTwo            *float64                   `json:"targetValueTwo"`
	TargetValueUnit           json.RawMessage            `json:"targetValueUnit"`
	ZoneNumber                *int64                     `json:"zoneNumber"`
	SecondaryTargetType       json.RawMessage            `json:"secondaryTargetType"`
	SecondaryTargetValueOne   json.RawMessage            `json:"secondaryTargetValueOne"`
	SecondaryTargetValueTwo   json.RawMessage            `json:"secondaryTargetValueTwo"`
	SecondaryTargetValueUnit  json.RawMessage            `json:"secondaryTargetValueUnit"`
	SecondaryZoneNumber       json.RawMessage            `json:"secondaryZoneNumber"`
	EndConditionZone          json.RawMessage            `json:"endConditionZone"`
	StrokeType                GarminWorkoutStrokeType    `json:"strokeType"`
	EquipmentType             GarminWorkoutEquipmentType `json:"equipmentType"`
	Category                  json.RawMessage            `json:"category"`
	ExerciseName              json.RawMessage            `json:"exerciseName"`
	WorkoutProvider           json.RawMessage            `json:"workoutProvider"`
	ProviderExerciseSourceID  json.RawMessage            `json:"providerExerciseSourceId"`
	WeightValue               json.RawMessage            `json:"weightValue"`
	WeightUnit                json.RawMessage            `json:"weightUnit"`
}

// GarminWorkoutEndCondition is the end-condition descriptor.
type GarminWorkoutEndCondition struct {
	ConditionTypeID  int64  `json:"conditionTypeId"`
	ConditionTypeKey string `json:"conditionTypeKey"`
	DisplayOrder     int64  `json:"displayOrder"`
	Displayable      bool   `json:"displayable"`
}

// GarminWorkoutEquipmentType is one of the workout equipment options.
type GarminWorkoutEquipmentType struct {
	EquipmentTypeID  int64           `json:"equipmentTypeId"`
	EquipmentTypeKey json.RawMessage `json:"equipmentTypeKey"`
	DisplayOrder     int64           `json:"displayOrder"`
}

// GarminWorkoutStepType identifies an interval step's role (work, rest, etc.).
type GarminWorkoutStepType struct {
	StepTypeID   int64  `json:"stepTypeId"`
	StepTypeKey  string `json:"stepTypeKey"`
	DisplayOrder int64  `json:"displayOrder"`
}

// GarminWorkoutStrokeType is a swim-stroke selector.
type GarminWorkoutStrokeType struct {
	StrokeTypeID  int64           `json:"strokeTypeId"`
	StrokeTypeKey json.RawMessage `json:"strokeTypeKey"`
	DisplayOrder  int64           `json:"displayOrder"`
}

// GarminWorkoutTargetType is a target-type selector (heart rate, pace, etc.).
type GarminWorkoutTargetType struct {
	WorkoutTargetTypeID  int64  `json:"workoutTargetTypeId"`
	WorkoutTargetTypeKey string `json:"workoutTargetTypeKey"`
	DisplayOrder         int64  `json:"displayOrder"`
}

// GarminWorkoutTypes is the dictionary of workout-type lookup tables.
type GarminWorkoutTypes struct {
	WorkoutStepTypes            []WorkoutStepType            `json:"workoutStepTypes"`
	WorkoutSportTypes           []WorkoutSportType           `json:"workoutSportTypes"`
	WorkoutConditionTypes       []WorkoutConditionType       `json:"workoutConditionTypes"`
	WorkoutIntensityTypes       []WorkoutIntensityType       `json:"workoutIntensityTypes"`
	WorkoutTargetTypes          []WorkoutTargetType          `json:"workoutTargetTypes"`
	WorkoutEquipmentTypes       []WorkoutEquipmentType       `json:"workoutEquipmentTypes"`
	WorkoutStrokeTypes          []WorkoutStrokeType          `json:"workoutStrokeTypes"`
	WorkoutSwimInstructionTypes []WorkoutSwimInstructionType `json:"workoutSwimInstructionTypes"`
	WorkoutDrillTypes           []WorkoutDrillType           `json:"workoutDrillTypes"`
}

// WorkoutConditionType is one entry in GarminWorkoutTypes.
type WorkoutConditionType struct {
	ConditionTypeID  int64  `json:"conditionTypeId"`
	ConditionTypeKey string `json:"conditionTypeKey"`
	DisplayOrder     int64  `json:"displayOrder"`
	Displayable      bool   `json:"displayable"`
}

// WorkoutDrillType is one entry in GarminWorkoutTypes.
type WorkoutDrillType struct {
	DrillTypeID  int64  `json:"drillTypeId"`
	DrillTypeKey string `json:"drillTypeKey"`
	DisplayOrder int64  `json:"displayOrder"`
}

// WorkoutEquipmentType is one entry in GarminWorkoutTypes.
type WorkoutEquipmentType struct {
	EquipmentTypeID  int64  `json:"equipmentTypeId"`
	EquipmentTypeKey string `json:"equipmentTypeKey"`
	DisplayOrder     int64  `json:"displayOrder"`
}

// WorkoutIntensityType is one entry in GarminWorkoutTypes.
type WorkoutIntensityType struct {
	IntensityTypeID  int64  `json:"intensityTypeId"`
	IntensityTypeKey string `json:"intensityTypeKey"`
	DisplayOrder     int64  `json:"displayOrder"`
}

// WorkoutSportType is one entry in GarminWorkoutTypes.
type WorkoutSportType struct {
	SportTypeID  int64  `json:"sportTypeId"`
	SportTypeKey string `json:"sportTypeKey"`
	DisplayOrder int64  `json:"displayOrder"`
}

// WorkoutStepType is one entry in GarminWorkoutTypes.
type WorkoutStepType struct {
	StepTypeID   int64  `json:"stepTypeId"`
	StepTypeKey  string `json:"stepTypeKey"`
	DisplayOrder int64  `json:"displayOrder"`
}

// WorkoutStrokeType is one entry in GarminWorkoutTypes.
type WorkoutStrokeType struct {
	StrokeTypeID  int64  `json:"strokeTypeId"`
	StrokeTypeKey string `json:"strokeTypeKey"`
	DisplayOrder  int64  `json:"displayOrder"`
}

// WorkoutSwimInstructionType is one entry in GarminWorkoutTypes.
type WorkoutSwimInstructionType struct {
	SwimInstructionTypeID  int64  `json:"swimInstructionTypeId"`
	SwimInstructionTypeKey string `json:"swimInstructionTypeKey"`
	DisplayOrder           int64  `json:"displayOrder"`
}

// WorkoutTargetType is one entry in GarminWorkoutTypes.
type WorkoutTargetType struct {
	WorkoutTargetTypeID  int64  `json:"workoutTargetTypeId"`
	WorkoutTargetTypeKey string `json:"workoutTargetTypeKey"`
	DisplayOrder         int64  `json:"displayOrder"`
}

// garminUpdateWorkoutPayload is the body used by UpdateWorkout. It diverges
// from GarminWorkout: `updatedDate` (not `updateDate`), an explicit
// consumerName, and a few other bookkeeping fields.
type garminUpdateWorkoutPayload struct {
	WorkoutID                 int64                  `json:"workoutId"`
	OwnerID                   int64                  `json:"ownerId"`
	WorkoutName               string                 `json:"workoutName"`
	Description               string                 `json:"description"`
	UpdatedDate               LocalDateTime          `json:"updatedDate"`
	CreatedDate               LocalDateTime          `json:"createdDate"`
	SportType                 GarminSportType        `json:"sportType"`
	SubSportType              json.RawMessage        `json:"subSportType"`
	TrainingPlanID            json.RawMessage        `json:"trainingPlanId"`
	Author                    GarminAuthor           `json:"author"`
	SharedWithUsers           json.RawMessage        `json:"sharedWithUsers"`
	WorkoutSegments           []GarminWorkoutSegment `json:"workoutSegments"`
	PoolLength                *float64               `json:"poolLength"`
	PoolLengthUnit            GarminUnit             `json:"poolLengthUnit"`
	Locale                    json.RawMessage        `json:"locale"`
	WorkoutProvider           string                 `json:"workoutProvider"`
	WorkoutSourceID           string                 `json:"workoutSourceId"`
	UploadTimestamp           json.RawMessage        `json:"uploadTimestamp"`
	AtpPlanID                 *int64                 `json:"atpPlanId"`
	Consumer                  string                 `json:"consumer"`
	ConsumerName              string                 `json:"consumerName"`
	ConsumerImageUrl          string                 `json:"consumerImageURL"`
	ConsumerWebsiteUrl        string                 `json:"consumerWebsiteURL"`
	WorkoutNameI18NKey        string                 `json:"workoutNameI18nKey"`
	DescriptionI18NKey        string                 `json:"descriptionI18nKey"`
	AvgTrainingSpeed          float64                `json:"avgTrainingSpeed"`
	EstimateType              json.RawMessage        `json:"estimateType"`
	EstimatedDistanceUnit     GarminUnit             `json:"estimatedDistanceUnit"`
	Shared                    bool                   `json:"shared"`
	EstimatedDurationInSecs   int64                  `json:"estimatedDurationInSecs"`
	EstimatedDistanceInMeters float64                `json:"estimatedDistanceInMeters"`
}

func newUpdateWorkoutPayload(w GarminWorkout) garminUpdateWorkoutPayload {
	var poolLength *float64
	if w.PoolLength != 0 {
		v := w.PoolLength
		poolLength = &v
	}
	return garminUpdateWorkoutPayload{
		WorkoutID:                 w.WorkoutID,
		OwnerID:                   w.OwnerID,
		WorkoutName:               w.WorkoutName,
		Description:               w.Description,
		UpdatedDate:               LocalDateTime{Time: time.Now()},
		CreatedDate:               w.CreatedDate,
		SportType:                 w.SportType,
		SubSportType:              w.SubSportType,
		TrainingPlanID:            w.TrainingPlanID,
		Author:                    w.Author,
		SharedWithUsers:           w.SharedWithUsers,
		WorkoutSegments:           w.WorkoutSegments,
		PoolLength:                poolLength,
		PoolLengthUnit:            w.PoolLengthUnit,
		WorkoutProvider:           w.WorkoutProvider,
		WorkoutSourceID:           w.WorkoutSourceID,
		UploadTimestamp:           w.UploadTimestamp,
		AtpPlanID:                 w.AtpPlanID,
		Consumer:                  w.Consumer,
		ConsumerName:              w.Consumer,
		WorkoutNameI18NKey:        w.WorkoutNameI18NKey,
		DescriptionI18NKey:        w.DescriptionI18NKey,
		AvgTrainingSpeed:          w.AvgTrainingSpeed,
		EstimateType:              w.EstimateType,
		EstimatedDistanceUnit:     w.EstimatedDistanceUnit,
		Shared:                    w.Shared,
		EstimatedDurationInSecs:   w.EstimatedDurationInSecs,
		EstimatedDistanceInMeters: w.EstimatedDistanceInMeters,
	}
}
