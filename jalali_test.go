/**
 * Author: Amin Zamani
 * File: jalali_test.go
 */

package gojalali

import (
	"testing"
	"time"
)

func TestJalali(t *testing.T) {
	jalali := Now()
	t.Log(jalali)
	t.Log(jalali.Year())
	t.Log(jalali.Month())
	t.Log(jalali.Day())
	t.Log(jalali.Hour())
	t.Log(jalali.Minute())
	t.Log(jalali.Second())
	t.Log(jalali.Weekday())
	t.Log(jalali.YearDay())
	t.Log(jalali.Location())
	t.Log(jalali.Zone())
	t.Log(jalali.Unix())
	t.Log(jalali.UnixNano())
	t.Log(jalali.Date())
	t.Log(jalali.Clock())
	t.Log(jalali.ISOWeek())
	t.Log(jalali.In(time.UTC))
	t.Log(jalali.AddDate(1, 2, 3))
	t.Log(jalali.Add(time.Hour * 24))
	t.Log(jalali.Sub(time.Now()))
	t.Log(jalali.String())
	t.Log(jalali.Georgian())
}

func TestJalaliFrom(t *testing.T) {
	jalali := From(time.Now())
	t.Log(jalali)
	t.Log(jalali.Year())
	t.Log(jalali.Month())
	t.Log(jalali.Day())
	t.Log(jalali.Hour())
	t.Log(jalali.Minute())
	t.Log(jalali.Second())
	t.Log(jalali.Weekday())
	t.Log(jalali.YearDay())
	t.Log(jalali.Location())
	t.Log(jalali.Zone())
	t.Log(jalali.Unix())
	t.Log(jalali.UnixNano())
	t.Log(jalali.Date())
	t.Log(jalali.Clock())
	t.Log(jalali.ISOWeek())
	t.Log(jalali.In(time.UTC))
	t.Log(jalali.AddDate(1, 2, 3))
	t.Log(jalali.Add(time.Hour * 24))
	t.Log(jalali.Sub(time.Now()))
	t.Log(jalali.String())
	t.Log(jalali.Georgian())
}
