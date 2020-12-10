package log

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

type pattern struct {
	bufPool sync.Pool
	funcMap map[string]func(D) string
}

// newPatternRender new pattern render
func newPatternRender() Render {
	pa := &pattern{
		bufPool: sync.Pool{New: func() interface{} { return &strings.Builder{} }},
	}
	pa.funcMap = map[string]func(D) string{
		"default": renderDefault,
		_level:    renderLevel,
		_time:     renderLongTime,
	}
	return pa
}

// Render implemet Formater
func (p *pattern) Render(w io.Writer, d ...D) error {
	builder := p.bufPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		p.bufPool.Put(builder)
	}()

	for _, t := range d {
		if nil != p.funcMap[t.Key] {
			builder.WriteString(p.funcMap[t.Key](t))
		} else {
			builder.WriteString(p.funcMap["default"](t))
		}
		builder.WriteString(" ")
	}
	builder.WriteString("\n")
	_, err := w.Write([]byte(builder.String()))
	return err
}

func renderLevel(d D) string {
	return d.Value.(string)
}

func renderLongTime(d D) string {
	return time.Now().Format("15:04:05.000")
}

func renderDefault(d D) string {
	v, ok := d.Value.(string)
	if ok {
		return v
	}
	return fmt.Sprintf("%v", d.Value)
}
