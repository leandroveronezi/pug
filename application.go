package pug

import (
	"github.com/leandroveronezi/pug/console"
	"syscall/js"
)

var MainApplication TApplication

func init() {

	js.Global().Set("Pug", js.Global().Get("Object").New())

	MainApplication = *TApplication{}.New()
}

type TApplication struct {
	el js.Value
}

func (_this TApplication) Dom() js.Value {
	return _this.el
}

func (_this TApplication) Append(Elements ...TComponent) {

	for _, element := range Elements {
		_this.el.Call("appendChild", element.Dom())
	}

}

func (_this TApplication) New() *TApplication {

	_this.el = js.Global().Get("document").Call("getElementById", _this.Id())
	return &_this

}

func (_this TApplication) Id() string {
	return "application_main"
}

func (_this TApplication) Start() {

	js.Global().Set("Binding", js.FuncOf(Binding))

	_this.el = js.Global().Get("document").Call("getElementById", _this.Id())

	if _this.el.Type() == js.TypeNull || _this.el.Type() == js.TypeUndefined {

		var div js.Value = js.Global().Get("document").Call("createElement", "DIV")
		_this.el = div
		_this.el.Set("innerHTML", ``)
		_this.el.Call("setAttribute", "id", "application_main")

		js.Global().Get("document").Get("body").Call("appendChild", _this.el)

	} else {
		_this.el.Set("innerHTML", "")
	}

	console.Log("Pug started!")

}

func (_this TApplication) Run() {

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
