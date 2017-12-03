package ecaas

import "errors"

var (
	//ErrorInvalidDateFormat is a an error that should be used to handle invalid date formats
	ErrorInvalidDateFormat = errors.New("The provided date format is not valid, expecting a unix date")
)

//MoveType describes the kind of move being conducted, and it's multiplier on the total job cost
type MoveType interface {
	GetName() string
	GetMultiplier() string
	GetTaxRate() string
}

//The JobDetails struct is an entity the contains critical values for a move, hourly rate and a preferred move
type JobDetails struct {
	Hours             float64
	HourlyRate        string
	PreferredMoveDate string
}

//NewJobDetails is a constructor for a JobDetails struct
func NewJobDetails(hours float64, hourlyRate string, preferredMoveDate string) *JobDetails {
	return &JobDetails{
		Hours:             hours,
		HourlyRate:        hourlyRate,
		PreferredMoveDate: preferredMoveDate,
	}
}

//EstimateRange models a low to high range of possible cost for an estimate
type EstimateRange struct {
	Low  string
	High string
}
