package expression

import "fmt"

func (c *CronExpression) toStringMonthExpression() string {
	values, err := getEvaluatedExpresssionValues(c.MonthExpression, 1, 12, "month")
	if err != nil {
		fmt.Println("error: month values out of bounds (1- 12) ", err)
	}
	return values
}
