package gojalali

import "time"

func toJalali(month time.Month, day int) (int, int) {
	if month == 1 && day <= 20 {
		return 10, day + 10
	} else if month == 1 && day <= 31 {
		return 11, day - 20
	} else if month == 2 && day <= 20 {
		return 11, day + 11
	} else if month == 2 && day <= 29 {
		return 12, day - 20
	} else if month == 3 && day <= 20 {
		return 12, day + 9
	} else if month == 3 && day <= 31 {
		return 1, day - 20
	} else if month == 4 && day <= 20 {
		return 1, day + 11
	} else if month == 4 && day <= 30 {
		return 2, day - 20
	} else if month == 5 && day <= 21 {
		return 2, day + 10
	} else if month == 5 && day <= 31 {
		return 3, day - 21
	} else if month == 6 && day <= 21 {
		return 3, day + 10
	} else if month == 6 && day <= 30 {
		return 4, day - 21
	} else if month == 7 && day <= 22 {
		return 4, day + 9
	} else if month == 7 && day <= 31 {
		return 5, day - 22
	} else if month == 8 && day <= 22 {
		return 5, day + 9
	} else if month == 8 && day <= 31 {
		return 6, day - 22
	} else if month == 9 && day <= 22 {
		return 6, day + 9
	} else if month == 9 && day <= 30 {
		return 7, day - 22
	} else if month == 10 && day <= 21 {
		return 7, day + 8
	} else if month == 10 && day <= 30 {
		return 8, day - 21
	} else if month == 11 && day <= 20 {
		return 8, day + 9
	} else if month == 11 && day <= 30 {
		return 9, day - 20
	} else if month == 12 && day <= 20 {
		return 9, day + 9
	}
	return 0, 0
}

func monthDays(month int, year int) int {
	if month <= 6 {
		return 31
	} else if month <= 11 {
		return 30
	} else if month == 12 {
		if isLeapYear(year) {
			return 30
		}
		return 29
	}
	return 0
}

func isLeapYear(year int) bool {
	return year%33 == 1 || year%33 == 5 || year%33 == 9 || year%33 == 13 || year%33 == 17 || year%33 == 22 || year%33 == 26 || year%33 == 30
}
