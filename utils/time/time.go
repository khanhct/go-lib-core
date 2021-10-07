package time

import "time"


const (
	YYYYMMDD string 	= "2006-01-02"
)

func ConvertStrToTime(timeStr string, format string) (time.Time, error) {
	t, err := time.Parse(format, timeStr)
	return t, err
}
