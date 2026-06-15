package logic


import (
	"time"
	"fmt"
)


func StartTimer() {
	time.Since(time.Now()).Truncate(time.Second)

}

// func Chrono(start time.Time) string {
// 	elapsed := time.Since(start)

// 	hours := int(elapsed.Hours())
// 	minutes := int(elapsed.Minutes()) % 60
// 	seconds := int(elapsed.Seconds()) % 60

// 	if hours > 0 {
// 		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
// 	}

// 	return fmt.Sprintf("%02d:%02d", minutes, seconds)
// }

func Chrono(d time.Duration) string {

	hours 	:= int(d.Hours())
    minutes := int(d.Minutes())
    seconds := int(d.Seconds()) % 60

    if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
