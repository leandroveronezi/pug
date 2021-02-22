package router

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
	"syscall/js"

	"github.com/gowebapi/webapi"
	"github.com/leandroveronezi/pug/console"
)

type TRoute struct {
	Path      string
	Component reflect.Type
	Meta      interface{}
}

type TRouter struct {
	Routes       map[string]TRoute
	TagMain      *js.Value
	CurrentRoute struct {
		Name       string
		Parameters map[string]js.Value
	}
	NotFoundRouteName string
}

var instance *TRouter
var onceRouterSingleton sync.Once

func init() {

	if js.Global().Get("Pug").IsUndefined() {
		js.Global().Set("Pug", js.Global().Call("eval", "new Object()"))
	}

	if js.Global().Get("Pug").Get("Router").IsUndefined() {
		js.Global().Get("Pug").Set("Router", js.Global().Call("eval", "new Object()"))
	}

	js.Global().Get("Pug").Get("Router").Set("Push", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if instance == nil {
			return nil
		}

		Parameters := make(map[string]js.Value)

		if len(args) == 2 {

			if args[1].Type() == js.TypeObject {

				keys := js.Global().Get("Object").Call("keys", args[1])

				for i := 0; i < keys.Get("length").Int(); i++ {

					Parameters[keys.Index(i).String()] = args[1].Get(keys.Index(i).String())

				}

			}

		}

		instance.Push(args[0].String(), Parameters)

		return nil
	}))

	js.Global().Get("window").Set("onhashchange", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if instance == nil {
			return nil
		}

		instance.CheckRoute()

		return nil
	}))

	script := webapi.GetDocument().CreateElement("script", nil)
	script.SetInnerHTML(`

function getRouteParams(url,routes) {

	var argsVal,
		argsNames,
		params = {};

	for(var x = 0; x < routes.length; x++) {

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

			params.url = currRoute;
	
		  	return {
				params : params
		  	};
	
		}

	}
	
	return null;

}

`)

	js.Global().Get("document").Get("body").Call("appendChild", script.JSValue())

}

func GetRouter() *TRouter {

	onceRouterSingleton.Do(func() {

		instance = &TRouter{}
		instance.Routes = make(map[string]TRoute, 0)
		//instance.TagMain = TagMain

	})

	return instance
}

func (_this *TRouter) CheckRoute() {

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
			_this.Push(idx, nil)
			return
		}

	}

	params := _this.getRouteParams(hash)

	Parameters := make(map[string]js.Value)

	if params.Type() == js.TypeObject {

		existe := params.Call("hasOwnProperty", "params")

		if existe.Bool() {
			params = params.Get("params")

			keys := js.Global().Get("Object").Call("keys", params)

			for i := 0; i < keys.Get("length").Int(); i++ {

				if keys.Index(i).String() == "url" {
					hash = params.Get("url").String()
					continue
				}

				Parameters[keys.Index(i).String()] = params.Get(keys.Index(i).String())

			}

		}

	}

	for idx, val := range _this.Routes {

		if val.Path == hash {
			_this.Push(idx, Parameters)
			return
		}

	}
	//{params:{uid:"1024", avatar: "ok", url}
	//url = strings.ReplaceAll(url, ":"+idx, jsToString(val))

}

func (_this *TRouter) Add(Routes map[string]TRoute) {

	for idx, route := range Routes {
		_this.Routes[idx] = route
	}

}

func (_this *TRouter) Push(Name string, Parameters map[string]js.Value) {

	if _this.TagMain == nil {
		console.Error("TagMain is not defined")
		return
	}

	_, ok := _this.Routes[Name]

	routeName := Name

	if !ok {
		console.Error("route not found")

		if _this.NotFoundRouteName == "" {
			return
		}

		routeName = _this.NotFoundRouteName
	}

	url := _this.Routes[routeName].Path

	for idx, val := range Parameters {
		url = strings.ReplaceAll(url, ":"+idx, jsToString(val))
	}

	_this.CurrentRoute.Name = routeName
	_this.CurrentRoute.Parameters = Parameters

	componentType := _this.Routes[routeName].Component

	_this.Render(componentType)

	history := js.Global().Get("Object").New()
	js.Global().Get("history").Call("pushState", history, js.Null(), "#"+url)

}

func (_this *TRouter) Render(Component reflect.Type) {

	ptr := reflect.New(Component)

	if hasMethod(ptr, "Created") {
		callMethod(ptr, "Created", []reflect.Value{})
	}

	h, err := compile(ptr)

	if err != nil {
		console.Error(err.Error())
		return
	}

	el := stringToHtml(h)

	elBind(el, ptr)

	_this.TagMain.Set("innerHTML", "")
	_this.TagMain.Call("appendChild", el.JSValue())

	if hasMethod(ptr, "Mounted") {
		callMethod(ptr, "Mounted", []reflect.Value{})
	}

}

func (_this *TRouter) getRouteParams(url string) js.Value {

	routes := js.Global().Get("Array").New()

	for _, reg := range _this.Routes {

		obj := js.Global().Get("Object").New()
		obj.Set("url", reg.Path)

		routes.Call("push", obj)

	}

	return js.Global().Call("getRouteParams", url, routes)

}

func jsToString(val js.Value) string {

	switch val.Type() {
	case js.TypeString:
		return val.String()
	case js.TypeUndefined:
		return "<undefined>"
	case js.TypeNull:
		return "<null>"
	case js.TypeBoolean:

		if val.Bool() {
			return "true"
		}
		return "false"

	case js.TypeNumber:
		return strconv.Itoa(val.Int())
	case js.TypeSymbol:
		return "<symbol>"
	case js.TypeObject:
		return "<object>"
	case js.TypeFunction:
		return "<function>"
	default:
		panic("bad type")
	}

}
