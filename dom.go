package pug

import (
	"errors"
	"github.com/leandroveronezi/casandra/sui"
	"syscall/js"
)

type function func(this js.Value, args []js.Value) interface{}

type TDomElementType string

const (
	TDomElemenInput  TDomElementType = "input"
	TDomElemenScript TDomElementType = "script"
)

func (_this TDomElementType) String() string {
	return string(_this)
}

type TDomElement struct {
	elementType TDomElementType
	el          js.Value
	id          string

	classBind js.Value
	classes   map[string]sui.TSemanticClass
}

func (_this *TDomElement) prepare() {

	_this.classes = make(map[string]sui.TSemanticClass, 0)
	_this.classBind = js.Global().Get("Object").New()
	_this.classBind.Set("class", "")

}

func (_this *TDomElement) ElementType(val TDomElementType) {
	_this.elementType = val
}

func (_this TDomElement) Dom() js.Value {
	return _this.el
}

func (_this *TDomElement) Create() {

	_this.prepare()

	var el js.Value = js.Global().Get("document").Call("createElement", _this.elementType.String())
	_this.el = el
	_this.el.Call("setAttribute", "id", _this.id)

	bind := TDataBind{}.New(&_this.classBind, "class")
	bind.Add(_this.el, "className", nil)

}

func (_this *TDomElement) InnerHTML(Val string) {
	_this.el.Set("innerHTML", Val)
}

func (_this *TDomElement) AddEventListener(Event string, Fn function) {
	_this.el.Call("addEventListener", Event, js.FuncOf(Fn))
}

func (_this *TDomElement) OnClick(Fn function) {
	_this.el.Set("onclick", js.FuncOf(Fn))
}

func (_this *TDomElement) Append(Elements ...TComponent) *TDomElement {

	for _, element := range Elements {
		_this.el.Call("appendChild", element.Dom())
	}

	return _this

}

func (_this *TDomElement) loadClass() {

	auxClass := ""

	for _, class := range _this.classes {
		auxClass += " " + class.String()
	}

	_this.classBind.Set("class", auxClass)
}

func (_this *TDomElement) AddClass(Classes ...sui.TSemanticClass) *TDomElement {

	for _, class := range Classes {
		_this.classes[class.String()] = class
	}

	_this.loadClass()

	return _this

}

func (_this *TDomElement) RemoveClass(Classes ...sui.TSemanticClass) *TDomElement {

	for _, class := range Classes {

		_, ok := _this.classes[class.String()]
		if ok {
			delete(_this.classes, class.String())
		}

	}

	_this.loadClass()

	return _this

}

func (_this *TDomElement) SetAttribute(atr string, val interface{}) *TDomElement {

	_this.el.Call("setAttribute", atr, val)

	return _this

}

func (_this *TDomElement) RemoveAttribute(atr string) *TDomElement {

	_this.el.Call("removeAttribute", atr)

	return _this

}

func (_this *TDomElement) HasAttribute(atr string) (bool, error) {

	has := _this.el.Call("hasAttribute", atr)

	if has.Type().String() != "boolean" {
		return false, errors.New("type invalid")
	}

	return has.Bool(), nil

}
