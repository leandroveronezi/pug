package router

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/html/htmlevent"
	"reflect"
	"strings"
	"sync"
	"syscall/js"
)

type TRoute struct {
	Path      string
	Component reflect.Type
	Meta      interface{}
}

type TRouter struct {
	RouterView *dom.Element
	Routes     map[string]TRoute
}

var instance *TRouter
var onceRouterSingleton sync.Once

func GetRouter() *TRouter {

	onceRouterSingleton.Do(func() {

		instance = &TRouter{}
		instance.Routes = make(map[string]TRoute, 0)

		js.Global().Get("Pug").Set("Router", js.Global().Call("eval", "new Object()"))

		js.Global().Get("Pug").Get("Router").Set("Push", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			//_args[0:_end]
			instance.Push(args[0].String())
			return nil

		}))

		webapi.GetWindow().AddEventHashChange(func(event *htmlevent.HashChangeEvent, currentTarget *webapi.Window) {
			instance.CheckRoute()
		})

	})
	return instance
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

		for idx, val := range _this.Routes {

			if val.Path == hash {
				_this.Push(idx)
				return nil
			}

		}

		_this.Push("404")

		return nil
	})

	js.Global().Get("document").Call("addEventListener", "DOMContentLoaded", fn, false)

}

func (_this *TRouter) Add(Routes map[string]TRoute) {

	for idx, route := range Routes {
		_this.Routes[idx] = route
	}

}

func (_this *TRouter) Push(Name string, Parameters ...interface{}) {

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

	_this.RouterView.SetInnerHTML("")

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

/*

var routes = [
  {url: "/users/:uid/pictures/:avatar"},
  {url: "/users/:uid"},
  {url: "/home"}
];


function getRouteParams(url) {

  var argsVal,
      argsNames,
      params = {};

  for(var x = 0; x < routes.length; x++){

    var currRoute = routes[x].url;
    var routeMatcher = new RegExp(currRoute.replace(/(:\w+)/g, '([\\w-]+)'));
    argsVal = url.match(routeMatcher);

    if(argsVal) {

      argsVal.shift();
      argsNames = currRoute.match(/(:\w+)/g);

      if(argsNames) {
        for(var y = 0; y < argsNames.length; y++){
          params[argsNames[y].slice(1)] = argsVal[y];
        }
      }

      return {
        params : params
      };

    }

  }
  return null;
}


console.log(getRouteParams("/users/1024/pictures/ok")); // {params:{uid:"1024", avatar: "ok"}
console.log(getRouteParams("/users/zezen")); // {params: {uid: "zezen"}}
console.log(getRouteParams("/home")); // {params: {}}
console.log(getRouteParams("/zezen")); // null

*/
