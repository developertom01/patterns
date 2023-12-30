package patterns

import (
	"context"
	"time"
)

//Try to do somethings and if you cannot do it in time drop the ball

// Tries to read message into chan in a duration and if the duration is done, we  quite
// Implementation with context timeout
func ToChannelTimedContext(ctx context.Context, d time.Duration, message any, c chan<- any) bool {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	select {
	case c <- message:
		return true
	case <-ctx.Done():
		return false
	}
}

// Tries to read message into chan in a duration and if the duration is done, we  quite
// Implementation with timer
func ToChannelTimedTimer(d time.Duration, message any, c chan<- any) bool {
	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case c <- message:
		return true
	case <-t.C:
		return false
	}
}
