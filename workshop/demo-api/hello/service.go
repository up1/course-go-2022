package hello

import (
	"context"
	"time"
)

func todoNext(ctx context.Context) {
	_, span := tracer.Start(ctx, "todoNext")
	defer span.End()

	time.Sleep(20 * time.Nanosecond)

	CallExternalApi(ctx)
}
