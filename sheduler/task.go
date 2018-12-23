package sheduler

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Task struct {
	Func   FunctionMeta
	Params []Param
}

func (task *Task) Run() bool {
	function := reflect.ValueOf(task.Func.function)
	params := make([]reflect.Value, len(task.Params))
	for i, param := range task.Params {
		params[i] = reflect.ValueOf(param)
	}
	rets := function.Call(params)
	err := reflect.ValueOf(rets[len(rets)-1])

	errErr := err.Interface()

	log.Println("Task: err:", errErr)
	if errErr != nil {
		errStr := fmt.Sprint(errErr)
		normalErr := errors.New(errStr)
		log.Println("Task: normal err:", normalErr, errStr)
		if strings.Contains(errStr, "Unavailable") {
			return false
		}
	}

	return true
}
