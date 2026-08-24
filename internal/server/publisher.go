package server

import "context"

func publishResult(ctx context.Context, value string) {
	queue := make(chan string)
	go func() {
		select {
		case queue <- value:
		case <-ctx.Done():
		}
	}()
}
