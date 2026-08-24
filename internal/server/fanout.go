package server

import (
	"context"
	"sync"
)

func waitFanout(ctx context.Context) {
	var group sync.WaitGroup
	group.Add(2)
	go func() {
		defer group.Done()
		<-ctx.Done()
	}()
	group.Wait()
}
