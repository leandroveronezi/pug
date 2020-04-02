package router

import (
	"github.com/gowebapi/webapi/dom"
	"syscall/js"
)

type THTMLComponent interface {
	Render(el *dom.Element, Parameters map[string]js.Value)
}
