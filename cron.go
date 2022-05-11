// Package cron provides functionality for generate signals at intervals timed in January 1, 1970 UTC.
package cron

import (
	"time"
)

// Loop generate a signal every d duration timed in January 1, 1970 UTC.
func CronWait(d time.Duration, s time.Duration) <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now().UnixNano()
			m := s + time.Duration(int64(d)-n%int64(d))*time.Nanosecond
			time.Sleep(m)
			r <- struct{}{}
		}
	}()
	return r
}

// Loop generate a signal every d duration timed in January 1, 1970 UTC.
func Cron(d time.Duration) <-chan struct{} {
	return CronWait(d, 0)
}
