package router

import "github.com/leandroveronezi/pug"

type TComponent interface {
	Render(el *casandra.TDomElement, Parameters ...interface{})
}
