package log

import (
	"context"
	"time"
)

//Handles ...
type Handles struct {
	handles []Handle
}

func newHandles(hans ...Handle) Handle {
	th := Handles{
		handles: hans,
	}
	return th
}

//Log ...
func (hs Handles) Log(ctx context.Context, l Level, d ...D) {
	//assample the D slice
	ds := make([]D, 0, 5)
	if _accessLevel == l {
		ds = append(ds, d...)
	} else {
		ds = append(ds, KVString(_level, l.String()), KV(_time, time.Now()))
		ds = append(ds, d...)
		if logConfig.Source {
			fn := funcName(3)
			ds = append(ds, KVString(_source, fn))
		}
	}

	for _, f := range hs.handles {
		f.Log(ctx, l, ds...)
	}
}

//Close ...
func (hs Handles) Close() error {
	for _, f := range hs.handles {
		if err := f.Close; nil != err {

		}
	}
	return nil
}
