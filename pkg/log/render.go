package log

import (
	"io"
)

//Render render the string
type Render interface {
	Render(io.Writer, ...D) error
}
