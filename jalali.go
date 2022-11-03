/**
 * Author: Amin Zamani
 * File: jalali.go
 */

package gojalali

import "time"

type Jalali struct {
	time.Time
}

func (j Jalali) Year() int {
	return j.Time.Year()
}

func (j Jalali) Month() time.Month {
	return j.Time.Month()
}

func (j Jalali) Day() int {
	return j.Time.Day()
}

func (j Jalali) Hour() int {
	return j.Time.Hour()
}

func (j Jalali) Minute() int {
	return j.Time.Minute()
}

func (j Jalali) Second() int {
	return j.Time.Second()
}

func (j Jalali) Nanosecond() int {
	return j.Time.Nanosecond()
}

func (j Jalali) Location() *time.Location {
	return j.Time.Location()
}

func (j Jalali) Weekday() time.Weekday {
	return j.Time.Weekday()
}

func (j Jalali) YearDay() int {
	return j.Time.YearDay()
}

func (j Jalali) Unix() int64 {
	return j.Time.Unix()
}

func (j Jalali) UnixNano() int64 {
	return j.Time.UnixNano()
}

func (j Jalali) String() string {
	return j.Time.String()
}

func (j Jalali) Format(layout string) string {
	return j.Time.Format(layout)
}

func (j Jalali) Add(d time.Duration) time.Time {
	return j.Time.Add(d)
}

func (j Jalali) AddDate(years int, months int, days int) time.Time {
	return j.Time.AddDate(years, months, days)
}

func (j Jalali) Sub(u time.Time) time.Duration {
	return j.Time.Sub(u)
}

func (j Jalali) In(loc *time.Location) time.Time {
	return j.Time.In(loc)
}

func (j Jalali) Equal(u time.Time) bool {
	return j.Time.Equal(u)
}

func (j Jalali) Before(u time.Time) bool {
	return j.Time.Before(u)
}

func (j Jalali) After(u time.Time) bool {
	return j.Time.After(u)
}

func (j Jalali) IsZero() bool {
	return j.Time.IsZero()
}

func (j Jalali) UTC() time.Time {
	return j.Time.UTC()
}

func (j Jalali) Local() time.Time {
	return j.Time.Local()
}


