package pug

import (
	"github.com/leandroveronezi/casandra/sui"
	"strings"
	"syscall/js"
)

func GetDomElement(Id string) *TDomElement {

	var el js.Value = js.Global().Get("document").Call("getElementById", Id)

	if el.Type() == js.TypeNull || el.Type() == js.TypeUndefined {
		return nil
	}

	_this := TDomElement{}
	_this.id = Id

	_this.prepare()

	_this.el = el
	classes := el.Get("className").String()

	for _, class := range strings.Split(classes, " ") {

		str := sui.TSemanticClass(class)

		_this.AddClass(sui.TSemanticClass(str))
	}

	bind := TDataBind{}.New(&_this.classBind, "class")
	bind.Add(_this.el, "className", nil)

	return &_this
}
