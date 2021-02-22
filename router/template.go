package router

import (
	"bytes"
	"reflect"
	"syscall/js"
	"text/template"
)

func compile(component reflect.Value) (string, error) {

	var finalMethod reflect.Value

	method := component.MethodByName("Template")
	if method.IsValid() {
		finalMethod = method
	}

	str := finalMethod.Call([]reflect.Value{})[0].String()

	name := reflect.TypeOf(component).Name()
	t := template.New(name)

	t, err := t.Parse(str)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, component)
	if err != nil {
		return "", err
	}

	return buf.String(), nil

}

func stringToHtml(str string) js.Value {

	return js.Global().Get("DOMParser").
		New().
		Call("parseFromString", str, "text/html").
		Get("body").
		Get("childNodes").
		Index(0)

}

func attributes(el js.Value) map[string]string {

	if el.Get("hasAttributes").IsUndefined() {
		return nil
	}

	if !el.Call("hasAttributes").Bool() {
		//return nil
	}

	m := make(map[string]string)

	attrs := el.Get("attributes")
	for i := 0; i < attrs.Get("length").Int(); i++ {
		attr := attrs.Index(i)
		m[attr.Get("name").String()] = attr.Get("value").String()
	}

	return m
}

func uuidv4() string {

	str := js.Global().Call("eval", `
		(() => {
			return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
				var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
				return v.toString(16);
			});
		})()
	`)

	return str.String()

}

func listMethod(v interface{}) {

	y := reflect.TypeOf(v)

	for i := 0; i < y.NumMethod(); i++ {

		//fmt.Println(y.Method(i).Name)

	}

	for i := 0; i < y.NumField(); i++ {

		//fmt.Println(y.Field(i).Name)

	}

}

func hasMethod(component reflect.Value, name string) bool {

	method := component.MethodByName(name)
	if method.IsValid() {
		return true
	}

	return false

}

func callMethod(component reflect.Value, name string, val []reflect.Value) {

	var finalMethod reflect.Value

	method := component.MethodByName(name)
	if method.IsValid() {
		finalMethod = method
	}

	finalMethod.Call(val)

}

func getMethodByName(component reflect.Value, name string, args2 []js.Value) func(this js.Value, args []js.Value) interface{} {

	return func(this js.Value, args []js.Value) interface{} {

		var colection []js.Value

		if len(args) > 0 {
			colection = append(colection, args...)
		}

		if len(args2) > 0 {
			colection = append(colection, args2...)
		}

		inputs := make([]reflect.Value, 2)

		inputs[0] = reflect.ValueOf(this)

		inputs[1] = reflect.ValueOf(colection)

		var result []reflect.Value

		result = component.MethodByName(name).Call(inputs)

		if len(result) == 0 {
			return nil
		}

		return js.ValueOf(result[0].Interface())

	}

}
