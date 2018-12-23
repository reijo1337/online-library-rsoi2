package sheduler

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

func CreateTask(function Function, params ...Param) (*Task, error) {
	funcMeta, err := createFuncMeta(function)
	if err != nil {
		return nil, err
	}

	return &Task{
		Func:   funcMeta,
		Params: params,
	}, nil
}

func createFuncMeta(function Function) (FunctionMeta, error) {
	funcValue := reflect.ValueOf(function)
	if funcValue.Kind() != reflect.Func {
		return FunctionMeta{}, fmt.Errorf("Provided function value is not an actual function")
	}

	name := runtime.FuncForPC(funcValue.Pointer()).Name()

	ret := FunctionMeta{
		Name:     name,
		function: function,
		params:   resolveParamTypes(function),
	}
	return ret, nil
}

func resolveParamTypes(function Function) map[string]reflect.Type {
	paramTypes := make(map[string]reflect.Type)
	funcType := reflect.TypeOf(function)
	for idx := 0; idx < funcType.NumIn(); idx++ {
		in := funcType.In(idx)
		paramTypes[in.Name()] = in
	}
	return paramTypes
}

func ExecuteTask(ch chan *Task) {
	time.Sleep(5 * time.Second)
	log.Println("Sheduler: waiting task")
	task := <-ch
	go ExecuteTask(ch)
	log.Println("Sheduler: execute task", task.Func.Name)
	if !task.Run() {
		log.Println("Sheduler: task not finushed, sending back to queue")
		ch <- task
	}
}

func CreateSheduler() chan *Task {
	ch := make(chan *Task)

	go ExecuteTask(ch)

	return ch
}
