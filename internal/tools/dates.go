package tools

import "time"

func ParseFromToDates(from *string, to *string) (*time.Time, *time.Time) {
	layout := "2006-01-02"

	var fromTime *time.Time
	var toTime *time.Time

	if from != nil {
		if parsedTime, err := time.Parse(layout, *from); err == nil {
			fromTimeUTC := parsedTime.UTC()
			fromTime = &fromTimeUTC
		}
	}

	if to != nil {
		if parsedTime, err := time.Parse(layout, *to); err == nil {
			toTimeUTC := parsedTime.UTC()
			toTime = &toTimeUTC
		}
	}

	return fromTime, toTime
}
