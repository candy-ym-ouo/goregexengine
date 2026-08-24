package server

import "context"

func startWorkers(ctx context.Context, count int) {
	for i := 0; i < count; i++ {
		go func() {
			defer func() { _ = recover() }()
			<-ctx.Done()
		}()
	}
}
