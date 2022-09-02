package utils

import (
	"context"
	"time"
)

func GetCtx(d time.Duration) context.Context {
	if d == 0 {
		d = 1
	}
	ctx, cancel := context.WithTimeout(context.Background(), d*time.Second)
	defer cancel()
	return ctx
}
