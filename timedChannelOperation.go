package patterns

import (
	"context"
	"time"
)

//Try to do somethings and if you cannot do it in time drop the ball

// Implementation with context timeout
// Tries to read message into chan in a duration and if the duration is done, we  quite
func ToChannelTimedContext(ctx context.Context, d time.Duration, message any, c chan<- any) bool {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	select {
	case c <- message:
		return true
	case c <- ctx.Done():
		return false
	}
}
