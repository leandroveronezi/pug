package router

import (
	"github.com/leandroveronezi/pug"
	"reflect"
	"strings"
	"syscall/js"
)

var Router = TRouter{}.New()

func init() {

	js.Global().Get("Casandra").Set("To", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		Router.To(args[0].String())
		return nil
	}))

	js.Global().Get("window").Set("onhashchange", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		Router.CheckRoute()

		return nil
	}))

}

type TRoute struct {
	Path      string
	Component reflect.Type
	Meta      interface{}
}

type TRouter struct {
	RouterView *casandra.TDomElement
	Routes     map[string]TRoute
}

func (_this TRouter) CheckRoute() {

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		hash := js.Global().Get("window").Get("location").Get("hash").String()
		hash = strings.TrimSpace(hash)

		if len(hash) == 0 {
			hash = "/"
		}

		if hash[0:1] == "#" {
			hash = hash[1:]
		}

		for idx, val := range Router.Routes {

			if val.Path == hash {
				_this.To(idx)
				return nil
			}

		}

		_this.To("404")

		return nil
	})

	js.Global().Call("$", js.Global().Get("document")).Call("ready", fn)

}

func (_this TRouter) New() TRouter {
	_this.Routes = make(map[string]TRoute, 0)
	return _this
}

func (_this *TRouter) Add(Routes map[string]TRoute) {

	for idx, route := range Routes {
		_this.Routes[idx] = route
	}

}

func (_this *TRouter) To(Name string, Parameters ...interface{}) {

	if _this.RouterView == nil {
		return
	}

	_, ok := _this.Routes[Name]

	var componentType reflect.Type

	routeName := Name

	if !ok {
		routeName = "404"
	} else {

	}

	componentType = _this.Routes[routeName].Component

	ms := reflect.New(componentType)

	_this.RouterView.Dom().Set("innerHTML", "")

	Method := ms.MethodByName("Render")

	in := make([]reflect.Value, len(Parameters)+1)
	in[0] = reflect.ValueOf(_this.RouterView)

	for k, param := range Parameters {
		in[k+1] = reflect.ValueOf(param)
	}

	if Method.IsValid() {
		Method.Call(in)
		history := js.Global().Get("Object").New()
		js.Global().Get("history").Call("pushState", history, js.Null(), "#"+_this.Routes[routeName].Path)
	}

}

/*
const link = $(`a[href$='${window.location.pathname}']`);
link.addClass('active');






$('a').on('click', (event) => {
  // Block browser page load
  event.preventDefault();

  // Highlight Active Menu on Click
  const target = $(event.target);
  $('.item').removeClass('active');
  target.addClass('active');

  // Navigate to clicked url
  const href = target.attr('href');
  const path = href.substr(href.lastIndexOf('/'));
  router.navigateTo(path);
});




*/
