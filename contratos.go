package pug

import "syscall/js"

type TComponent interface {
	//Id() string
	//Render() *TComponent
	//Parent() *TComponent
	//Pointer() *TComponent
	Dom() js.Value
}
