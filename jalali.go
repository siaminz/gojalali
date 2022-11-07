/**
 * Author: Amin Zamani
 * File: jalali.go
 */

package gojalali

import (
	"time"
)

func Now() Jalali {
	return Jalali{time.Now()}
}

func From(time time.Time) Jalali {
	return Jalali{time}
}
