package constructor

import "syscall/js"

func NewObject() *js.Value {

	obj := js.Global().Call("eval", "new Object()")

	return &obj

}

func NewDate() *js.Value {

	obj := js.Global().Call("eval", "new Date()")

	return &obj

}
