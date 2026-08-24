package server

import "context"

func awaitStream(ctx context.Context) {
	ready := make(chan struct{})
	go func() {
		<-ctx.Done()
	}()
	<-ready
}
