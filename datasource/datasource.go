package datasource

import (
	"errors"
	"syscall/js"
)

type DataSource struct {
	datasource *js.Value
}

func (_this *DataSource) Field(Name string) *Field {

	if _this.datasource.Call("hasOwnProperty", Name).Bool() {
		_, field := _this.GetFieldByName(Name)
		return field
	}

	_this.datasource.Set(Name, js.Null())
	_this.datasource.Get("$bind").Set(Name, js.Null())

	_, field := _this.GetFieldByName(Name)
	return field

}

func (_this DataSource) GetFieldByName(Name string) (error, *Field) {

	if !_this.datasource.Call("hasOwnProperty", Name).Bool() {
		return errors.New("Field not found"), &Field{}
	}

	aux := _this.datasource.Get(Name)

	return nil, &Field{&aux}

}

func New() DataSource {

	ds := DataSource{}
	obj := js.Global().Get("Object").New()
	ds.datasource = &obj
	ds.datasource.Set("$bind", js.Global().Get("Object").New())

	return ds

}
