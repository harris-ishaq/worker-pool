package helpers

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func GetLocalDate(tz string) time.Time {
	var location, err = time.LoadLocation(tz)
	if err != nil {
		return time.Now()
	}

	return time.Now().In(location)
}

// TimeHostNow gets time now in jakarta GMT+7
func TimeHostNow() time.Time {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Print(err)
		panic(err)
	}
	now := time.Now()
	timeInLoc := now.In(location)
	return timeInLoc
}

func CvtJulianDay(day, month, year int) string {
	monthLengths := [11]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

	// Increase february by 1 if it's a leap year
	monthLengths[1] = monthLengths[1] + checkIsLeapYear(year)

	var totalDays = day

	for iMonth := 0; iMonth < month-1; iMonth++ {
		totalDays += monthLengths[iMonth]
	}

	stringYear := strconv.Itoa(year)

	if totalDays < 100 {
		return fmt.Sprintf("%s0%v", stringYear[2:], totalDays)
	}

	return fmt.Sprintf("%s%v", stringYear[2:], totalDays)
}

func checkIsLeapYear(year int) int {
	var leap = 0

	if year%4 == 0 {
		leap = 1
		if year%100 == 0 {
			leap = 0
			if year%400 == 0 {
				leap = 1
			}
		}
	}

	return leap
}
