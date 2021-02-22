package router

import (
	"syscall/js"
)

type TPugComponent struct {
	Router TRouter
	Route  struct {
		Parameters map[string]js.Value
	}
}
