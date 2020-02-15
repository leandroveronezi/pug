package pug

import (
	"reflect"
	"syscall/js"
)

type TDataBind struct {
	bindObj *js.Value
	bind    *js.Value
}

func (_this TDataBind) New(Obj *js.Value, Property string) *TDataBind {

	jsObj := js.Global().Get("Object").New()
	_this.bindObj = &jsObj

	Set(_this.bindObj, "object", Obj)
	Set(_this.bindObj, "property", Property)

	//console.Log(_this.bindObj)

	aux := js.Global().Get("Binding").New(_this.bindObj)
	_this.bind = &aux

	return &_this
}

func (_this *TDataBind) Add(El js.Value, Property string, Event *string) *TDataBind {

	if Event != nil {
		_this.bind.Call("addBinding", El, Property, *Event)
	} else {
		_this.bind.Call("addBinding", El, Property)
	}

	return _this

}

func Binding(this js.Value, args []js.Value) interface{} {

	b := &args[0]

	_this := &this

	existe := b.Get("object").Call("hasOwnProperty", "___"+b.Get("property").String()+"__elementBindings")

	if !existe.Bool() {
		b.Get("object").Set("___"+b.Get("property").String()+"__elementBindings", js.Global().Get("Array").New())
	}

	this.Set("value", b.Get("object").Get(b.Get("property").String()))

	this.Set("valueBind", js.FuncOf(func(t js.Value, args []js.Value) interface{} {
		return _this.Get("value")
	}))

	this.Set("valueGetter", js.FuncOf(func(t js.Value, args []js.Value) interface{} {
		return _this.Get("value")
	}))

	this.Set("valueSetter", js.FuncOf(func(t js.Value, args []js.Value) interface{} {

		Set(_this, "value", &args[0])

		for i := 0; i < b.Get("object").Get("___"+b.Get("property").String()+"__elementBindings").Get("length").Int(); i++ {

			binding := b.Get("object").Get("___" + b.Get("property").String() + "__elementBindings").Index(i)

			el := binding.Get("element")
			Set(&el, binding.Get("attribute").String(), &args[0])

		}

		return nil

	}))

	this.Set("addBinding", js.FuncOf(func(t js.Value, args []js.Value) interface{} {

		binding := js.Global().Get("Object").New()

		Set(&binding, "element", &args[0])
		Set(&binding, "attribute", &args[1])

		if len(args) >= 3 {

			if args[2].Type() != js.TypeUndefined && args[2].Type() != js.TypeNull {

				args[0].Call("addEventListener", args[2], js.FuncOf(func(t js.Value, args2 []js.Value) interface{} {

					_this.Call("valueSetter", args[0].Get(args[1].String()))
					return nil

				}))

				Set(&binding, "event", &args[2])

			}

		}

		b.Get("object").Get("___"+b.Get("property").String()+"__elementBindings").Call("push", binding)

		Set(&args[0], args[1].String(), _this.Get("value"))

		return _this

	}))

	existe = js.Global().Get("Object").Call("getOwnPropertyDescriptor", b.Get("object"), b.Get("property")).Call("hasOwnProperty", "get")

	if !existe.Bool() {

		aux := js.Global().Get("Object").New()

		Set(&aux, "get", this.Get("valueGetter"))
		Set(&aux, "set", this.Get("valueSetter"))

		js.Global().Get("Object").Call("defineProperty", b.Get("object"), b.Get("property"), aux)

		aa := b.Get("object")
		Set(&aa, b.Get("property").String(), this.Get("value"))

	}

	return nil

}

func Set(obj *js.Value, prop string, value interface{}) {

	s := reflect.ValueOf(value)

	if !s.IsValid() {
		return
	}

	switch s.Kind() {

	case reflect.Ptr:

		if s.IsNil() {
			return
		}

		Set(obj, prop, s.Elem().Interface())

	case reflect.Struct:
		obj.Set(prop, s.Interface())

	case reflect.Bool:
		obj.Set(prop, s.Bool())
	case reflect.Int:
		obj.Set(prop, s.Int())
	case reflect.Int8:
		obj.Set(prop, s.Int())
	case reflect.Int16:
		obj.Set(prop, s.Int())
	case reflect.Int32:
		obj.Set(prop, s.Int())
	case reflect.Int64:
		obj.Set(prop, s.Int())
	case reflect.Uint:
		obj.Set(prop, s.Uint())
	case reflect.Uint8:
		obj.Set(prop, s.Uint())
	case reflect.Uint16:
		obj.Set(prop, s.Uint())
	case reflect.Uint32:
		obj.Set(prop, s.Uint())
	case reflect.Uint64:
		obj.Set(prop, s.Uint())
	case reflect.Uintptr:
		obj.Set(prop, s.Uint())
	case reflect.Float32:
		obj.Set(prop, s.Float())
	case reflect.Float64:
		obj.Set(prop, s.Float())
	case reflect.String:
		obj.Set(prop, s.String())
	}

}
