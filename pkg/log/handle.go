package log

import (
	"context"
)

//Handle ...
type Handle interface {
	Log(context.Context, Level, ...D)
	Close() error
}
