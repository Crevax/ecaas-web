package ecaas

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

//CalculateTotalCost is a function used to total and provide an estimate range
//for a move based upon the following business rules:
//	1. Multiply total estimated hours by the provided hourly rate for an initial subtotal
//  2. Apply the cost multiplier - effectively a service fee - to the subtotal
//  3. Apply tax to the subtotal plus service fee to generate the low end of the etimate
//  4. To estimate the high end of the estimate, if the preferred move date is on Friday or Saturday
//		- add a complexity factor of 30% of the low estimate before tax - otherwise add a standard 15% weekday
//		complexity factor to get a estimate high end subtotal
//  5. Apply tax to the high end of the estimate, end return the range formatted in USD
func CalculateTotalCost(details *JobDetails, moveType MoveType) *EstimateRange {

	totalHours := decimal.NewFromFloat(details.Hours)
	hourlyRate, _ := decimal.NewFromString(details.HourlyRate)

	taxRate, _ := decimal.NewFromString(moveType.GetTaxRate())
	jobMultiplier, _ := decimal.NewFromString(moveType.GetMultiplier())

	complexity, _ := jobComplexityForDate(details.PreferredMoveDate)
	complexityFactor := decimal.NewFromFloat(complexity)

	lowTotal := addPercentageToTotal(totalHours.Mul(hourlyRate), jobMultiplier)
	lowTotalWithTax := addTax(lowTotal, taxRate)
	highTotal := addPercentageToTotal(lowTotal, complexityFactor)
	highTotalWithTax := addTax(highTotal, taxRate)

	return &EstimateRange{
		Low:  fmt.Sprintf("$%s", lowTotalWithTax.StringFixedBank(2)),
		High: fmt.Sprintf("$%s", highTotalWithTax.StringFixedBank(2)),
	}
}

func addPercentageToTotal(total decimal.Decimal, percentage decimal.Decimal) decimal.Decimal {
	return total.Mul(percentage.Add(decimal.NewFromFloat(1)))
}

func addTax(total decimal.Decimal, taxRate decimal.Decimal) decimal.Decimal {
	return addPercentageToTotal(total, taxRate)
}

func jobComplexityForDate(dateString string) (float64, error) {
	date, err := time.Parse(time.UnixDate, dateString)
	if err != nil {
		return 0, ErrorInvalidDateFormat
	}
	complexity := 0.15
	if date.Weekday() == time.Friday ||
		date.Weekday() == time.Saturday {
		complexity = 0.3
	}
	return complexity, nil
}
