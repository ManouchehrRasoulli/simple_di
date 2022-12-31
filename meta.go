package simpledi

import "reflect"

type input struct {
	name  string
	typ   reflect.Type
	value reflect.Value
}

type output struct {
	name  string
	typ   reflect.Type
	value reflect.Value
}
