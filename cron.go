package cron

import (
	"time"
)

// Loop generate a signal at all time.
func Loop() <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			r <- struct{}{}
		}
	}()
	return r
}

// Second generate a signal every second.
func Second() <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for range time.NewTicker(time.Second).C {
			r <- struct{}{}
		}
	}()
	return r
}

// Minute generate a signal every minute.
func Minute() <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now()
			s := time.Duration(60-n.Second()) * time.Second
			time.Sleep(s)
			r <- struct{}{}
		}
	}()
	return r
}

// Hour generate a signal every hour.
func Hour() <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now()
			s := time.Duration(60-n.Second()) * time.Second
			m := time.Duration(59-n.Minute()) * time.Minute
			time.Sleep(s + m)
			r <- struct{}{}
		}
	}()
	return r
}

// Dayz generate a signal every dayz.
func Dayz() <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now()
			s := time.Duration(60-n.Second()) * time.Second
			m := time.Duration(59-n.Minute()) * time.Minute
			h := time.Duration(23-n.Hour()) * time.Hour
			time.Sleep(s + m + h)
			r <- struct{}{}
		}
	}()
	return r
}
