package data

import (
	"github.com/gowebapi/webapi"
	"reflect"
	"syscall/js"
)

type FnChange *func(value *js.Value)

type DataField struct {
	JsValue         *js.Value
	Valid           bool // Valid is true if JsValue is not NULL
	someObject      *js.Value
	proxiedObject   *js.Value
	elementBindings *js.Value
	OnChange        FnChange
}

func (_this DataField) Value() *js.Value {

	if !_this.Valid {
		return nil
	}

	return _this.JsValue

}

func (_this *DataField) SetA(value interface{}) {

	obj := js.Global().Get("Object").New()
	obj.Set("value", value)

	aux := obj.Get("value")

	_this.Set(&aux)

}

func (_this *DataField) SetB(value *interface{}) {

	obj := js.Global().Get("Object").New()
	obj.Set("value", value)

	aux := obj.Get("value")

	_this.Set(&aux)

}

func (_this *DataField) Set(value *js.Value) {

	if _this.someObject == nil {
		obj := js.Global().Get("Object").New()
		_this.someObject = &obj
		_this.newProxy()
	}

	if value == nil {
		_this.JsValue, _this.Valid = nil, false

		set(_this.proxiedObject, "value", nil)
	}

	_this.Valid = true
	_this.JsValue = value

	set(_this.proxiedObject, "value", value)

}

func (_this *DataField) BindA(Id string, Property string, Event ...string) {

	el := webapi.GetWindow().Document().GetElementById(Id)

	if el == nil {
		return
	}

	_this.Bind(el.JSValue(), Property, Event...)

}

func (_this *DataField) Bind(El js.Value, Property string, Event ...string) {

	if _this.elementBindings == nil {

		arr := js.Global().Get("Array").New()
		_this.elementBindings = &arr

	}

	binding := js.Global().Get("Object").New()
	set(&binding, "element", El)
	set(&binding, "attribute", Property)

	_this.elementBindings.Call("push", binding)

	if len(Event) == 1 {

		ev := Event[0]

		El.Call("addEventListener", ev, js.FuncOf(func(t js.Value, args2 []js.Value) interface{} {

			_this.SetA(El.Get(Property))

			return nil

		}))

	}

}

func (_this *DataField) newProxy() {

	if _this.proxiedObject != nil {
		return
	}

	obj := js.Global().Get("Object").New()

	obj.Set("get", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		target := args[0]
		key := args[1]
		receiver := args[2]

		return js.Global().Get("Reflect").Call("get", target, key, receiver)

	}))

	obj.Set("set", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		target := args[0]
		key := args[1]
		val := args[2]
		receiver := args[3]

		for i := 0; i < _this.elementBindings.Get("length").Int(); i++ {

			binding := _this.elementBindings.Index(i)

			el := binding.Get("element")
			set(&el, binding.Get("attribute").String(), val)

		}

		if _this.OnChange != nil {
			auxFN := *_this.OnChange
			auxFN(&val)
		}

		return js.Global().Get("Reflect").Call("set", target, key, val, receiver)

	}))

	proxy := js.Global().Get("Proxy").New(_this.someObject, obj)

	_this.proxiedObject = &proxy

}

func set(obj *js.Value, prop string, value interface{}) {

	s := reflect.ValueOf(value)

	if !s.IsValid() {
		return
	}

	switch s.Kind() {

	case reflect.Ptr:

		if s.IsNil() {
			return
		}

		set(obj, prop, s.Elem().Interface())

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
