package log

import (
	"context"
	"os"
)

type stdoutHandle struct {
	render Render
}

func newStdoutHandle() stdoutHandle {
	return stdoutHandle{
		render: newPatternRender(),
	}
}

//Log ...
func (st stdoutHandle) Log(ctx context.Context, l Level, d ...D) {
	st.render.Render(os.Stderr, d...)
}

//Close ...
func (st stdoutHandle) Close() error {
	return nil
}
