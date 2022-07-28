package ingestor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngestCorrect(t *testing.T) {
	os.Args[1] = "*/15 0 1,15 * 1-5 /usr/bin/find" //example from document

	expression, _ := Ingest()

	assert.Equal(t, "*/15", expression.MinuteExpression)
	assert.Equal(t, "0", expression.HourExpression)
	assert.Equal(t, "1,15", expression.DayOfMonthExpression)
	assert.Equal(t, "*", expression.MonthExpression)
	assert.Equal(t, "1-5", expression.DayOfWeekExpression)
	assert.Equal(t, "/usr/bin/find", expression.Command)
}

func TestErrorOnBadInput(t *testing.T) {

	os.Args[1] = "*/15 0 1,15 * 1-5 /usr/bin/find thestraw"

	_, err := Ingest()

	assert.NotNil(t, err)
	assert.Equal(t, "error: expected 5 values in argument but got: 7", err.Error())
}
