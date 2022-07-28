package expression

import "fmt"

func (c *CronExpression) toStringMinuteExpression() string {
	values, err := getEvaluatedExpresssionValues(c.MinuteExpression, 0, 59, "minute")
	if err != nil {
		fmt.Println("error: minute values out of bounds (0 - 59) ", err)
	}
	return values
}
