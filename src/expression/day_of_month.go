package expression

import "fmt"

func (c *CronExpression) toStringDayOfMonthExpression() string {
	values, err := getEvaluatedExpresssionValues(c.DayOfMonthExpression, 1, 31, "day of month")
	if err != nil {
		fmt.Println("error: day of month values out of bounds (1 - 31) ", err)
	}
	return values
}
