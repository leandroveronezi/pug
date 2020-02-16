package router

import (
	"github.com/gowebapi/webapi/dom"
)

type THTMLComponent interface {
	Render(el *dom.Element, Parameters ...interface{})
}
