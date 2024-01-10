package ifc

import (
	"strconv"
	"time"
)

// checks if year is leap
func ifLeapYear(year int) bool {
	leapYear := year%4 == 0 && !(year%100 == 0) || year%400 == 0
	return leapYear
}

// calculates year day according to Gregorian calendar
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

// calculates date according to IFC
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

// returns IFC date as an array of ints
func GetNumsDateIFC(timezoneShiftMinutes int) [3]int {
	const minutesPerHour = 60
	place := time.FixedZone("UTC", timezoneShiftMinutes*minutesPerHour)
	timestamp := time.Now().In(place)
	// get Gregorian date
	year, month, monthDay := timestamp.Year(), int(timestamp.Month()), timestamp.Day()
	// calculate numerical IFC date
	leapYear := ifLeapYear(year)
	yearDay := calcYearDay(month, monthDay, leapYear)
	monthNumIFC, monthDayNumIFC := calcDateIFC(yearDay, leapYear)
	//
	numsDateIFC := [3]int{year, monthNumIFC, monthDayNumIFC}
	return numsDateIFC
}

// returns IFC date as an array of strings
func GetStrsDateIFC(timezoneShiftMinutes int) [3]string {
	monthNamesIFC_EO := [13]string{"januaro", "februaro", "marto",
		"aprilo", "majo", "junio", "sunio", "julio", "aŭgusto",
		"septembro", "oktobro", "novembro", "decembro"}
	// calculate numerical IFC date
	numsDateIFC := GetNumsDateIFC(timezoneShiftMinutes)
	year, monthIFC, monthDayIFC := numsDateIFC[0], numsDateIFC[1], numsDateIFC[2]
	// convert ints to strings
	yearStr := strconv.Itoa(year)
	monthStrIFC := monthNamesIFC_EO[monthIFC-1]
	monthDayStrIFC := strconv.Itoa(monthDayIFC)
	strsDateIFC := [3]string{yearStr, monthStrIFC, monthDayStrIFC}
	return strsDateIFC
}
