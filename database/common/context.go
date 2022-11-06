package common

import (
	"context"
	"time"
)

func MakeContext() (context.Context, func()) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
