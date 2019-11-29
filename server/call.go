// MIT License

// Copyright (c) 2019 gonethopper

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// * @Author: ankye
// * @Date: 2019-06-12 15:53:22
// * @Last Modified by:   ankye
// * @Last Modified time: 2019-06-12 15:53:22

package server

import (
	"fmt"
	"reflect"
)

// CallUserFunc simply to dynamically call a function or a method on an object
// Calls the callback given by the first parameter and passes the remaining parameters as arguments.
// Zero or more parameters to be passed to the callback.
// Returns the return value of the callback.
func CallUserFunc(f interface{}, v ...interface{}) []reflect.Value {
	valueFunc := reflect.ValueOf(f)
	paramsList := []reflect.Value{}
	if len(v) > 0 {
		for i := 0; i < len(v); i++ {
			paramsList = append(paramsList, reflect.ValueOf(v[i]))
		}
	}
	return valueFunc.Call(paramsList)

}

// CallUserMethod simply to dynamically call a method on an object
// Calls the instance given by the first parameter and method name as the second parameter
// and passes the remaining parameters as arguments.
// Zero or more parameters to be passed to the method.
// Returns the return value of the method.
func CallUserMethod(instance interface{}, method string, v ...interface{}) []reflect.Value {
	valueS := reflect.ValueOf(instance)
	m := valueS.MethodByName(method)
	paramsList := []reflect.Value{}
	if len(v) > 0 {
		for i := 0; i < len(v); i++ {
			paramsList = append(paramsList, reflect.ValueOf(v[i]))
		}
	}
	return m.Call(paramsList)
}

// GO wapper exec goruntine and stat count
func GO(v ...interface{}) {
	f := v[0]
	WG.Add(1)
	App.ModifyGoCount(1)
	go func() {
		CallUserFunc(f, v[1:]...)
		App.ModifyGoCount(-1)
		WG.Done()
	}()
}

// GOWithContext wapper exec goruntine and use context to manager goruntine
func GOWithContext(v ...interface{}) {
	f := v[0]
	App.ModifyGoCount(1)
	go func() {
		CallUserFunc(f, v[1:]...)
		App.ModifyGoCount(-1)
	}()
}

func Future(f func() (interface{}, error)) func() (interface{}, error) {
	var result interface{}
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		result, err = f()
	}()

	return func() (interface{}, error) {
		<-c
		return result, err
	}
}

// CallObject call struct
type CallObject struct {
	Cmd     string
	Args    []interface{}
	ChanRet chan RetObject
}

// RetObject call return object
type RetObject struct {
	Ret interface{}
	Err error
}

// NewCallObject create call object
func NewCallObject(cmd string, args ...interface{}) *CallObject {
	return &CallObject{
		Cmd:     cmd,
		Args:    args,
		ChanRet: make(chan RetObject, 1),
	}
}

// Processor goruntine process pre call
func Processor(s Service, obj *CallObject) {

	var err error

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	var ret = RetObject{
		Ret: nil,
		Err: nil,
	}
	f := s.GetHandler(obj.Cmd)
	if f == nil {
		err = fmt.Errorf("function id %v: function not registered", obj.Cmd)
	} else {
		args := []interface{}{s, obj}
		args = append(args, obj.Args...)
		values := CallUserFunc(f, args...)
		ret.Ret = values[0].Interface()
		if values[1].Interface() != nil {
			err = values[1].Interface().(error)
			ret.Err = err
		}
	}

	obj.ChanRet <- ret

}
