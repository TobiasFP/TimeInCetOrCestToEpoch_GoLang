package CetCest

import (
	"fmt"
	"strconv"
	"time"
)

func GetCETOrCEST(timeAsString string) string {
	layout := "2006-01-02T15:04:00"
	ti, err := time.Parse(layout, timeAsString)
	if err != nil {
		fmt.Println("Didnt Expect to receive:" + err.Error() + " Expected: a time'")
	}

	if IsCET(ti) {
		return timeAsString + " CET"
	}
	return timeAsString + " CEST"

}

func IsCET(timeToConvert time.Time) bool {
	switch month := timeToConvert.Month(); {
	// If month is November to February, then time is CET
	case month < 3 || month > 10:
		return true
	case month == 3:
		//If month is march and day is later than last sunday,
		// We are no longer in CET
		if isLaterThanLastSundayOfMonth(timeToConvert) {
			return false
		}
		return true
	case month == 10:
		//If month is october and day is later than last sunday,
		// We are back in CET
		if isLaterThanLastSundayOfMonth(timeToConvert) {
			return true
		}
		return false
	default:
		return false
	}
}

func isLaterThanLastSundayOfMonth(timeToConvert time.Time) bool {
	return timeToConvert.Day() > getLastSundayOfMonth(timeToConvert)
}

func getLastSundayOfMonth(timeToConvert time.Time) int {
	month := timeToConvert.Month()
	lastDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	lastDay := lastDays[month-1]
	lastSundayOfMonth := 1

	layout := "2006-January-02"

	for i := 1; i <= lastDay; i++ {
		var day string
		if i < 10 {
			day = "0" + strconv.Itoa(i)
		} else {
			day = strconv.Itoa(i)
		}
		dateAsString := strconv.Itoa(timeToConvert.Year()) + "-" + month.String() + "-" + day
		ti, _ := time.Parse(layout, dateAsString)
		if ti.Weekday() == time.Sunday {
			lastSundayOfMonth = i
		}
	}
	return lastSundayOfMonth
}
