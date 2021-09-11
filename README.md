# Cron

Execute scheduled jobs.

Cron generates a signal every second, the 0th second of every minute, the 0th second of every hour, or the 0th second of every day. We can use it with the library [gracefulexit](https://github.com/godump/gracefulexit) to form a basic scheduler:

```go
package main

import (
	"log"
	"time"

	"github.com/godump/cron"
	"github.com/godump/gracefulexit"
)

func main() {
	chanSpin := cron.Minute()
	chanExit := gracefulexit.Chan()
	done := 0
	log.Println("main: loop")
	for {
		select {
		case <-chanSpin:
			log.Println(time.Now())
		case <-chanExit:
			done = 1
		}
		if done != 0 {
			break
		}
	}
	log.Println("main: done")
}
```

Doc: [https://godoc.org/github.com/godump/cron](https://godoc.org/github.com/godump/cron).
