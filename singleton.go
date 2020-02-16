package pug

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/leandroveronezi/pug/console"
	"github.com/leandroveronezi/pug/router"
	"sync"
	"syscall/js"
)

type application struct {
	TagId   string
	TagMain *dom.Element
	Router  *router.TRouter
}

var instance *application
var onceApplicationSingleton sync.Once

func GetApplication() *application {

	onceApplicationSingleton.Do(func() {

		js.Global().Set("Pug", js.Global().Call("eval", "new Object()"))

		instance = &application{}
		instance.TagId = "application_main"
		instance.TagMain = webapi.GetWindow().Document().GetElementById(instance.TagId)

	})

	if instance.Router == nil {
		instance.Router = router.GetRouter()
	}

	return instance
}

func (_this *application) Start() {

	_this.TagMain = webapi.GetWindow().Document().GetElementById(_this.TagId)

	if _this.TagMain == nil {

		div := webapi.GetDocument().CreateElement("DIV", nil)
		div.SetAttribute("id", _this.TagId)
		div.SetInnerHTML("")
		_this.TagMain = div

		js.Global().Get("document").Get("body").Call("appendChild", div.JSValue())

	} else if _this.TagMain.JSValue().Type() == js.TypeNull || _this.TagMain.JSValue().Type() == js.TypeUndefined {

		div := webapi.GetDocument().CreateElement("DIV", nil)
		div.SetAttribute("id", _this.TagId)
		div.SetInnerHTML("")
		_this.TagMain = div

		js.Global().Get("document").Get("body").Call("appendChild", div.JSValue())

	} else {
		_this.TagMain.SetInnerHTML("")
	}

	console.Log("%cPug started!", "color: blue; font-size: 20px")

}

func (_this *application) Run() {

	defer func() {
		if err := recover(); err != nil {
			console.Group("recover")
			console.Warn("Recuperado...")
			console.Error(err)
			console.GroupEnd()
		}
	}()

	c := make(chan bool)
	<-c

}
