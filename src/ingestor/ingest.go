package ingestor

import (
	"cron-parser/expression"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Ingest reads the inputted argument and trys to parse it into a
// CronExpression. Errors if there are invalid numbers of arguments.
func Ingest() (expression.CronExpression, error) {

	if len(os.Args) <= 1 {
		return expression.CronExpression{}, errors.New("error: an expression must be supplied")
	}

	input := os.Args[1]

	fields := strings.Fields(input)

	if len(fields) != 6 {
		return expression.CronExpression{}, fmt.Errorf("error: expected 5 values in argument but got: %v", len(fields))
	}

	cron := expression.CronExpression{
		MinuteExpression:     fields[0],
		HourExpression:       fields[1],
		DayOfMonthExpression: fields[2],
		MonthExpression:      fields[3],
		DayOfWeekExpression:  fields[4],
		Command:              fields[5],
	}

	return cron, nil
}
