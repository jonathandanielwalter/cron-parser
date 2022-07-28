package expression

import "fmt"

func (c *CronExpression) toStringDayOfWeekExpression() string {
	values, err := getEvaluatedExpresssionValues(c.DayOfWeekExpression, 1, 7, "day of week")
	if err != nil {
		fmt.Println("error: day of week values out of bounds (1 - 7) ", err)
	}
	return values
}
