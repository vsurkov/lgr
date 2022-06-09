package lgr

import (
	"fmt"
	"time"
)


func Logger(text string) {
	currentTime := time.Now()
	fmt.Printf("INFO: %d-%d-%d %d:%d:%d - %s\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second(),
		text)
}
