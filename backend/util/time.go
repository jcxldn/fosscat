package util

import (
	"time"
)

func GetTimeOrZero(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	} else {
		return *t
	}
}
