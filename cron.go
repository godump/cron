// Package cron provides functionality for generate signals at intervals timed in January 1, 1970 UTC.
package cron

import (
	"time"
)

// Loop generate a signal every d duration timed in January 1, 1970 UTC.
func Loop(d time.Duration) <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now().UnixNano()
			s := time.Duration(int64(d)-n%int64(d)) * time.Nanosecond
			time.Sleep(s)
			r <- struct{}{}
		}
	}()
	return r
}

// Second generate a signal every second timed in January 1, 1970 UTC.
func Second() <-chan struct{} {
	return Loop(time.Second)
}

// Minute generate a signal every minute timed in January 1, 1970 UTC.
func Minute() <-chan struct{} {
	return Loop(time.Minute)
}

// Hour generate a signal every hour timed in January 1, 1970 UTC.
func Hour() <-chan struct{} {
	return Loop(time.Hour)
}

// Dayz generate a signal every dayz timed in January 1, 1970 UTC.
func Dayz() <-chan struct{} {
	return Loop(time.Hour * 24)
}
