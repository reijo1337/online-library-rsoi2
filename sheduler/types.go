package sheduler

import "reflect"

type Function interface{}

type Param interface{}

type FunctionMeta struct {
	Name     string
	function Function
	params   map[string]reflect.Type
}
