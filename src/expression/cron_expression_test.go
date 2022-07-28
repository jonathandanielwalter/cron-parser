package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	expected string
	input    string
}

func TestEntireCronOutput(t *testing.T) {
	expectedOutput := "minute 1 2 3\nhour 1 2 3\nday of month 1 6 11\nmonth 1\nday of week 1 2 3 4 5 6 7\ncommand a/command/here"

	cron := CronExpression{
		MinuteExpression:     "1,2,3",
		HourExpression:       "1,2,3",
		DayOfMonthExpression: "1,6,11",
		MonthExpression:      "1",
		DayOfWeekExpression:  "*",
		Command:              "a/command/here",
	}

	assert.Equal(t, expectedOutput, cron.CreateCronOutput())
}

func TestMinutesToString(t *testing.T) {
	tests := []TestCase{
		{
			expected: "minute 15",
			input:    "15",
		},
		{
			expected: "minute 0 15 30 45",
			input:    "*/15",
		},
		{
			expected: "minute 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59",
			input:    "*",
		},
		{
			expected: "minute 1 2 3",
			input:    "1,2,3",
		},
		{
			expected: "minute 5 6 7 8",
			input:    "5-8",
		},
	}

	for _, test := range tests {
		cron := CronExpression{
			MinuteExpression: test.input,
		}

		assert.Equal(t, test.expected, cron.toStringMinuteExpression())
	}
}

func TestHoursToString(t *testing.T) {
	tests := []TestCase{
		{
			expected: "hour 0 4 8 12 16 20",
			input:    "*/4",
		},
		{
			expected: "hour 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
			input:    "*",
		},
		{
			expected: "hour 1 2 3",
			input:    "1,2,3",
		},
		{
			expected: "hour 5 6 7 8",
			input:    "5-8",
		},
	}

	for _, test := range tests {
		cron := CronExpression{
			HourExpression: test.input,
		}

		assert.Equal(t, test.expected, cron.toStringHourExpression())
	}
}

func TestDayOfMonthToString(t *testing.T) {
	tests := []TestCase{
		{
			expected: "day of month 1 6 11 16 21 26 31",
			input:    "*/5",
		},
		{
			expected: "day of month 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
			input:    "*",
		},
		{
			expected: "day of month 1 2 3",
			input:    "1,2,3",
		},
		{
			expected: "day of month 5 6 7 8",
			input:    "5-8",
		},
	}

	for _, test := range tests {
		cron := CronExpression{
			DayOfMonthExpression: test.input,
		}

		assert.Equal(t, test.expected, cron.toStringDayOfMonthExpression())
	}
}

func TestMonthToString(t *testing.T) {
	tests := []TestCase{
		{
			expected: "month 1 4 7 10",
			input:    "*/3",
		},
		{
			expected: "month 1 2 3 4 5 6 7 8 9 10 11 12",
			input:    "*",
		},
		{
			expected: "month 1 2 3 4 5 6",
			input:    "1,2,3,4,5,6",
		},
		{
			expected: "month 9 10 11 12",
			input:    "9-12",
		},
	}

	for _, test := range tests {
		cron := CronExpression{
			MonthExpression: test.input,
		}

		assert.Equal(t, test.expected, cron.toStringMonthExpression())
	}
}

func TestDayOfWeekToString(t *testing.T) {
	tests := []TestCase{
		{
			expected: "day of week 1 4 7",
			input:    "*/3",
		},
		{
			expected: "day of week 1 2 3 4 5 6 7",
			input:    "*",
		},
		{
			expected: "day of week 1 2 3 4 5 6",
			input:    "1,2,3,4,5,6",
		},
		{
			expected: "day of week 5 6 7",
			input:    "5-7",
		},
	}

	for _, test := range tests {
		cron := CronExpression{
			DayOfWeekExpression: test.input,
		}

		assert.Equal(t, test.expected, cron.toStringDayOfWeekExpression())
	}
}

func TestCommandExpression(t *testing.T) {
	tests := []TestCase{
		{
			expected: "command here/is/a/command",
			input:    "here/is/a/command",
		},
	}

	for _, test := range tests {
		cron := CronExpression{
			Command: test.input,
		}

		assert.Equal(t, test.expected, cron.toStringCommandExpression())
	}
}

func TestOutOfBounds(t *testing.T) {
	cron := CronExpression{
		MinuteExpression:    "1-61",
		DayOfWeekExpression: "1,8",
		HourExpression:      "*/34",
	}

	_, minErr := getEvaluatedExpresssionValues(cron.MinuteExpression, 0, 59, "minute")
	assert.NotNil(t, minErr)

	_, dayErr := getEvaluatedExpresssionValues(cron.DayOfWeekExpression, 1, 7, "day of week")
	assert.NotNil(t, dayErr)

	_, hourErr := getEvaluatedExpresssionValues(cron.HourExpression, 1, 23, "hour")
	assert.NotNil(t, hourErr)

}
