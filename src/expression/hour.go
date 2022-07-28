package expression

import "fmt"

func (c *CronExpression) toStringHourExpression() string {
	values, err := getEvaluatedExpresssionValues(c.HourExpression, 0, 23, "hour")
	if err != nil {
		fmt.Println("error:hours values out of bounds (1 - 23) ", err)
	}

	return values
}
