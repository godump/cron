package main

import (
	"log"
	"time"

	"github.com/godump/cron"
)

func main() {
	log.Println("main: loop")
	for range cron.CronWait(time.Minute, time.Second*10) {
		log.Println("main:", time.Now())
	}
}
