package router

import (
	"reflect"
	"strings"
	"syscall/js"

	"github.com/leandroveronezi/pug/console"
)

func elBind(el js.Value, ptr reflect.Value) {

	kids := el.Get("childNodes")

	if kids.Get("length").Int() > 0 {

		for i := 0; i < kids.Get("length").Int(); i++ {
			elBind(kids.Index(i), ptr)
		}

	}

	for attr, attrvalue := range attributes(el) {

		if attr[:1] == "$" {

			attrvalue = attrvalue + "," + attr[1:]

			attr = "@bind"

		}

		switch attr {

		case "@click":

			a := strings.Index(attrvalue, "(")
			b := strings.Index(attrvalue, ")")

			function := attrvalue
			prt := "[]"

			if a >= 0 {
				prt = "[" + attrvalue[a+1:b] + "]"
				function = attrvalue[:a]
			}

			if !hasMethod(ptr, function) {
				continue
			}

			arr := js.Global().Get("JSON").Call("parse", prt)

			args := make([]js.Value, 0)

			for i := 0; i < arr.Length(); i++ {
				args = append(args, js.ValueOf(arr.Index(i)))
			}

			console.Dir(arr)

			ev := getMethodByName(ptr, function, args)

			el.Call("addEventListener", "click", js.FuncOf(ev))

		case "@bind":

			a := strings.Split(attrvalue, ",")

			var VarName = ""
			var Property = "value"
			var Event = "change"

			if len(a) < 1 {
				continue
			}

			if len(a) >= 1 {
				VarName = a[0]
			}

			Property = tagGetDefaultProperty(el)

			if len(a) >= 2 {
				Property = a[1]
			}

			if len(a) >= 3 {
				Event = a[2]
			}

			f := ptr.Elem().FieldByName(VarName).Addr()

			if f.IsValid() {

				arr := make([]reflect.Value, 0)

				arr = append(arr, reflect.ValueOf(el))

				arr = append(arr, reflect.ValueOf(Property))

				if Event != "" {
					arr = append(arr, reflect.ValueOf(Event))
				}

				f.MethodByName("Bind").Call(arr)

			}

			//Bind(el,Property,Event) {

		default:
			continue
		}

	}

}

func tagGetDefaultProperty(el js.Value) string {

	tag := el.Get("tagName").String()

	switch tag {

	case "LABEL":
		return "innerHTML"

	case "INPUT":
		return "value"

	default:
		return "value"
	}

}
