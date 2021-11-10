# Cron

Package cron provides functionality for generate signals at intervals timed in January 1, 1970 UTC. According to most cases, cron can work with [gracefulexit](https://github.com/godump/gracefulexit).

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
