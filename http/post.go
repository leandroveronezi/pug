package http

import (
	"encoding/json"
	"reflect"
	"syscall/js"
)

func Post(url string, value interface{}, fn func(this js.Value, args []js.Value) interface{}) error {

	xhr := js.Global().Call("eval", "new XMLHttpRequest()")

	xhr.Set("onload", js.FuncOf(fn))

	xhr.Call("open", "POST", url, true)
	xhr.Call("setRequestHeader", "Content-Type", "application/json")

	xhr.Set("timeout", 4000)

	/*
		xhr.ontimeout = function (e) {
			// Timeout na chamada XMLHttpRequest. Ação de timeout aqui.
		};
	*/

	vl := value

	if reflect.TypeOf(vl).Kind() == reflect.Struct {
		var err error
		bt, err := json.Marshal(vl)

		vl = string(bt)

		if err != nil {
			return err
		}
	}

	xhr.Call("send", vl)

	return nil

}
