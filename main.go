package ifc

import (
	"fmt"
	"time"
)

// check if year is leap
func ifLeapYear(year int) bool {
	leapYear := year%4 == 0 && !(year%100 == 0) || year%400 == 0
	return leapYear
}

// calculate year day according to Gregorian calendar
func calcYearDay(month int, monthDay int, leapYear bool) int {
	yearDay := 0
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leapYear {
		monthDays[1] = 29
	}
	for _, monthDays := range monthDays[:month-1] {
		yearDay += monthDays
	}
	yearDay += monthDay
	return yearDay
}

// calculate date according to IFC
func calcDateIFC(day int, leapYear bool) (int, int) {
	const leapDay = 169
	yearDay := 365
	// calculate date
	month := (day / 28) + 1 // 1 2 ... 13
	monthDay := day % 28    // 1 2 ... 0 (<- 28 [will be reassigned])
	// recalculate date according to the leap day
	if leapYear {
		yearDay = 366
		if day == leapDay {
			// 29.06 (169 day)
			// leap day is irregular
			month = 6
			monthDay = 29
		} else if day > leapDay {
			// 01.07+ (170+ day)
			// all days after the leap day are regular
			afterdays := day - leapDay
			month = 7 + (afterdays / 28)
			monthDay = afterdays % 28
		}
	}
	// recalculate date according to the year day
	if day == yearDay {
		month = 13
		monthDay = 29
	}
	// reassign the start and the end of the cycle
	if monthDay == 0 {
		month--
		monthDay = 28
	}
	return month, monthDay
}

func ToStringDateIFC(date [3]int) string {
	monthNamesIFC_EO := [13]string{"januaro", "februaro", "marto",
		"aprilo", "majo", "junio", "sunio", "julio", "aŭgusto",
		"septembro", "oktobro", "novembro", "decembro"}
	year, monthIFC, monthDayIFC := date[0], date[1], date[2]
	description := fmt.Sprintf("Jaro %d: la %s de %s", year, monthDayIFC, monthNamesIFC_EO[monthIFC-1])
	return description
}

func GetStringDateIFC(timezoneShiftMinutes) string {
	const minutesPerHour = 60
	monthNamesIFC_EO := [13]string{"januaro", "februaro", "marto",
		"aprilo", "majo", "junio", "sunio", "julio", "aŭgusto",
		"septembro", "oktobro", "novembro", "decembro"}
	place := time.FixedZone("UTC", timezoneShiftMinutes*minutesPerHour)
	timestamp := time.Now().In(place)
	// get Gregorian date
	year, month, monthDay := timestamp.Year(), int(timestamp.Month()), timestamp.Day()
	leapYear := ifLeapYear(year)
	// calculate IFC date
	yearDay := calcYearDay(month, monthDay, leapYear)
	monthIFC, monthDayIFC := calcDateIFC(yearDay, leapYear)
	// IFC date string
	stringDateIFC := fmt.Sprintf("Jaro %d: la %s de %s", year, monthDayIFC, monthNamesIFC_EO[monthIFC-1])
	return stringDateIFC
}

func GetNumericDateIFC(timezoneShiftMinutes int) [3]int {
	const minutesPerHour = 60
	place := time.FixedZone("UTC", timezoneShiftMinutes*minutesPerHour)
	timestamp := time.Now().In(place)
	// get Gregorian date
	year, month, monthDay := timestamp.Year(), int(timestamp.Month()), timestamp.Day()
	leapYear := ifLeapYear(year)
	// calculate IFC date
	yearDay := calcYearDay(month, monthDay, leapYear)
	monthIFC, monthDayIFC := calcDateIFC(yearDay, leapYear)
	// IFC date array
	numericDateIFC = [3]int{year, monthIFC, monthDayIFC}
	return numericDateIFC
}
