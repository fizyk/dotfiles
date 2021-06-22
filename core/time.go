package core

import (
	"fmt"
	"time"
)

func MeasureTime(start time.Time, banner string) {
	end := time.Now()
	fmt.Printf("%s took %.2f seconds\n", banner, end.Sub(start).Seconds())
}
