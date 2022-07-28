package expression

import "strings"

func (c *CronExpression) toStringCommandExpression() string {
	var sb strings.Builder
	sb.WriteString("command")
	sb.WriteString(" " + c.Command)
	return sb.String()
}
