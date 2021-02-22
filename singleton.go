package pug

import (
	"sync"
	"syscall/js"

	"github.com/leandroveronezi/pug/console"
)

type application struct {
	TagId   string
	TagMain js.Value
}

var instance *application
var onceApplicationSingleton sync.Once

func init() {

	if js.Global().Get("Pug").IsUndefined() {
		js.Global().Set("Pug", js.Global().Call("eval", "new Object()"))
	}

}

func GetApplication() *application {

	onceApplicationSingleton.Do(func() {

		instance = &application{}
		instance.TagId = "application_main"
		instance.TagMain = js.Global().Get("document").Call("getElementById", instance.TagId)

	})

	return instance
}

func (_this *application) Start() {

	_this.TagMain = js.Global().Get("document").Call("getElementById", instance.TagId)

	/*
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

		}
	*/

	console.Log("%cPug started!", "color: blue; font-size: 20px")

}

func (_this *application) Run() {

	defer func() {
		if err := recover(); err != nil {
			console.Group("recover")
			console.Warn("recovered...")
			console.Error(err)
			console.GroupEnd()
		}
	}()

	c := make(chan bool)
	<-c

}
