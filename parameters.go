package garmin

import (
	"fmt"
	"strings"
)

// WorkoutsOrderBy is the sort key for GetWorkouts.
type WorkoutsOrderBy int

const (
	WorkoutsOrderByName WorkoutsOrderBy = iota
	WorkoutsOrderByUpdateDate
	WorkoutsOrderByCreatedDate
	WorkoutsOrderBySportPk
)

// String returns the constant Garmin's API expects in the orderBy query
// parameter.
func (o WorkoutsOrderBy) String() string {
	switch o {
	case WorkoutsOrderByName:
		return "WORKOUT_NAME"
	case WorkoutsOrderByUpdateDate:
		return "UPDATE_DATE"
	case WorkoutsOrderByCreatedDate:
		return "CREATED_DATE"
	case WorkoutsOrderBySportPk:
		return "WORKOUT_SPORT_PK"
	}
	return ""
}

// OrderSeq is the sort direction.
type OrderSeq int

const (
	OrderSeqAsc OrderSeq = iota
	OrderSeqDesc
)

// String returns "ASC" or "DESC".
func (o OrderSeq) String() string {
	if o == OrderSeqAsc {
		return "ASC"
	}
	return "DESC"
}

// WorkoutsParameters narrows GetWorkouts. Zero value defaults to the same
// values the C# library uses.
type WorkoutsParameters struct {
	Start              uint
	Limit              uint
	OrderBy            WorkoutsOrderBy
	OrderSeq           OrderSeq
	IncludeAtp         bool
	MyWorkoutsOnly     bool
	SharedWorkoutsOnly bool
}

// DefaultWorkoutsParameters returns the same defaults the C# library applies.
func DefaultWorkoutsParameters() WorkoutsParameters {
	return WorkoutsParameters{
		Start:              1,
		Limit:              50,
		OrderBy:            WorkoutsOrderByCreatedDate,
		OrderSeq:           OrderSeqDesc,
		IncludeAtp:         true,
		MyWorkoutsOnly:     true,
		SharedWorkoutsOnly: false,
	}
}

// queryString renders the parameters as a `key=value&…` string. Values are
// emitted in lowercase for booleans to match the casing produced by C#'s
// default `ToString` behaviour for those types.
func (p WorkoutsParameters) queryString() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "start=%d&limit=%d&orderBy=%s&orderSeq=%s&includeAtp=%s&myWorkoutsOnly=%s&sharedWorkoutsOnly=%s",
		p.Start, p.Limit, p.OrderBy, p.OrderSeq,
		boolStr(p.IncludeAtp), boolStr(p.MyWorkoutsOnly), boolStr(p.SharedWorkoutsOnly))
	return sb.String()
}

func boolStr(b bool) string {
	if b {
		return "True"
	}
	return "False"
}
