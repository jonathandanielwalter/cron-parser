package expression

import (
	"fmt"
	"strconv"
	"strings"
)

type CronExpression struct {
	MinuteExpression     string
	HourExpression       string
	DayOfMonthExpression string
	MonthExpression      string
	DayOfWeekExpression  string
	Command              string
}

func (c *CronExpression) CreateCronOutput() string {
	var sb strings.Builder

	sb.WriteString(c.toStringMinuteExpression())
	sb.WriteString("\n")
	sb.WriteString(c.toStringHourExpression())
	sb.WriteString("\n")
	sb.WriteString(c.toStringDayOfMonthExpression())
	sb.WriteString("\n")
	sb.WriteString(c.toStringMonthExpression())
	sb.WriteString("\n")
	sb.WriteString(c.toStringDayOfWeekExpression())
	sb.WriteString("\n")
	sb.WriteString(c.toStringCommandExpression())

	return sb.String()
}

func isWildcard(expression string) bool {
	return expression == "*"
}

func isSeperator(expression string) bool {
	return strings.Contains(expression, ",")
}

func isStepValue(expression string) bool {
	return strings.Contains(expression, "/")
}

func isRange(expression string) bool {
	return strings.Contains(expression, "-")
}

func getValuesFromSeperator(expression string, lowValue int, highValue int) (string, error) {
	speratedValues := strings.Split(expression, ",")
	var sb strings.Builder
	for _, value := range speratedValues {
		intValue, _ := strconv.Atoi(value)
		if outOfBounds(intValue, lowValue, highValue) {
			return "", fmt.Errorf("value %v is out of bounds", intValue)
		}

		sb.WriteString(" " + value)
	}
	return sb.String(), nil
}

func getNumbersInRange(expression string, highestValue int, lowestValue int) (string, error) {
	var sb strings.Builder
	ranges := strings.Split(expression, "-")

	leftValue, err := strconv.Atoi(ranges[0])
	if err != nil {
		fmt.Println("Error parsing range left value")
	}

	rightValue, err := strconv.Atoi(ranges[1])
	if err != nil {
		fmt.Println("Error parsing range right value")
	}

	if outOfBounds(leftValue, lowestValue, highestValue) || outOfBounds(rightValue, lowestValue, highestValue) {
		return "", fmt.Errorf("values for range %v-%v are out of bounds", leftValue, rightValue)
	}

	for leftValue <= rightValue {
		sb.WriteString(" " + strconv.Itoa(leftValue))
		leftValue++
	}

	return sb.String(), nil
}

func getEvaluatedExpresssionValues(expression string, lowestValue int, hightestValue int, expressionPrefix string) (string, error) {
	var sb strings.Builder

	sb.WriteString(expressionPrefix)

	//if its just a wildcard return all possible values
	if isWildcard(expression) {
		for i := lowestValue; i <= hightestValue; i++ {
			sb.WriteString(" " + strconv.Itoa(i))
		}
		return sb.String(), nil
	}

	if isSeperator(expression) {
		values, err := getValuesFromSeperator(expression, lowestValue, hightestValue)
		if err != nil {
			return "", err
		}
		sb.WriteString(values)
		return sb.String(), nil
	}

	if isStepValue(expression) {
		arguments := strings.Split(expression, "/")

		//check both sides of expression for being OOB
		for _, arugment := range arguments {
			intArgument, _ := strconv.Atoi(arugment)
			if intArgument > hightestValue || intArgument < lowestValue {
				return "", fmt.Errorf("value %v is out of bounds", intArgument)
			}
		}

		increment, err := strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println("Error parsing step value for minute")
		}

		start := lowestValue
		for start <= hightestValue {
			sb.WriteString(" " + strconv.Itoa(start))
			start = start + increment
		}
		return sb.String(), nil
	}

	if isRange(expression) {
		rangeValues, err := getNumbersInRange(expression, hightestValue, lowestValue)
		if err != nil {
			return "", err
		}
		sb.WriteString(rangeValues)
		return sb.String(), nil
	}

	sb.WriteString(" " + expression)
	return sb.String(), nil
}

func outOfBounds(checkNumber int, lowValue int, highValue int) bool {
	if checkNumber > highValue {
		return true
	}

	if checkNumber < lowValue {
		return true
	}

	return false
}
