
package main
import (
	"time"
	"context"
)
func sleep(ctx context.Context, duration time.Duration) {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return
	case <-timer.C:
		return
	}
}

func doJob(ctx context.Context) {
	

	for {
		// ... do something
		select {
		case <-ctx.Done():
			return
		//...
		default:
			sleep(ctx, time.Minute) // wait for a minute
		}
	}
}