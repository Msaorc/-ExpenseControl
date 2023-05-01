package date

import "time"

var dateLayout = "2006-01-02"

func ConvertDateToString(date time.Time) string {
	return date.Format(dateLayout)
}

func ConvertDateToTime(date string) time.Time {
	convertedDate, _ := time.Parse(dateLayout, date)
	return convertedDate
}

func Validate(date string) error {
	_, err := time.Parse(dateLayout, date)
	return err
}
