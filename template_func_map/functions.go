package template_func_map

import (
	"fmt"
	"time"
)

func DateFormat(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%2d/%2d", year, month, day)
}
