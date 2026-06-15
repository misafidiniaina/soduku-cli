package logic


import (
	"time"
	"fmt"
)


func StartTimer() {
	time.Since(time.Now()).Truncate(time.Second)

}

func Chrono(start time.Time) string {
	elapsed := time.Since(start)

	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}