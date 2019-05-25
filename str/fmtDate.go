package str

import (
	"strings"
	"time"
)

//FmtDate formats date with basic java pattern instead golang pattern
func FmtDate(format string, t time.Time) string {
	//2006-01-02T15:04:05Z07:00
	newFormat := strings.Replace(format, "dd", "02", -1)
	newFormat = strings.Replace(newFormat, "MMM", "Jan", -1)
	newFormat = strings.Replace(newFormat, "MM", "01", -1)
	newFormat = strings.Replace(newFormat, "yyyy", "2006", -1)
	newFormat = strings.Replace(newFormat, "yy", "06", -1)
	newFormat = strings.Replace(newFormat, "HH", "15", -1)
	newFormat = strings.Replace(newFormat, "mm", "04", -1)
	newFormat = strings.Replace(newFormat, "ss", "05", -1)

	return t.Format(newFormat)
}
