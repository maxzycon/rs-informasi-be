package timeutil

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func FromString(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func FromStringDateOnly(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func ToString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func ToStringReverse(date time.Time) string {
	return date.Format("02-01-2006 15:04:05")
}

func ToStringDateOnlyWithoutSlash(date time.Time) string {
	return date.Format("20060102")
}

func ToTimestamp(date time.Time) int64 {
	return date.Unix()
}

func ToStringDateOnly(date time.Time) string {
	return date.Format("2006-01-02")
}

func ToStringDateOnlyReverse(date time.Time) string {
	return date.Format("02-01-2006")
}

func DateTimeFromString(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", str)
}

func ParseDate(t string) (local time.Time) {
	local, _ = time.Parse("2006-01-02", t)
	return
}

func StringUnixToTime(str string) (time.Time, error) {
	end, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		log.Errorf("error parse int timeUtil %v", err)
	}

	return time.Unix(end, 0), nil
}

func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func GenerateScheduledDates(startDate time.Time, endDate time.Time, interval int) []time.Time {
	var scheduledDates []time.Time

	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		scheduledDates = append(scheduledDates, currentDate)
		currentDate = currentDate.AddDate(0, 0, interval)
	}

	return scheduledDates
}

func GenerateScheduledDatesTax(startDate time.Time, endDate time.Time, interval int) []time.Time {
	var scheduledDates []time.Time

	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		scheduledDates = append(scheduledDates, currentDate)
		currentDate = currentDate.AddDate(0, interval, -14) // ---- minus 7 day
	}

	return scheduledDates
}
