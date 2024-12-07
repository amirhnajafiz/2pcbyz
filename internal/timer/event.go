package timer

import "time"

type event struct {
	expiresAt time.Time
	label     int
}
