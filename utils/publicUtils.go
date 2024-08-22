package utils

import (
	"reflect"
	"runtime"
)

func GetFunctionName(f interface{}) string {
	// 通过反射获取函数指针（以确保传入参数的类型
	funcValue := reflect.ValueOf(f)

	if funcValue.Kind() != reflect.Func {
		return "not a function"
	}

	// 获取程序计数器
	pc := funcValue.Pointer()

	// 使用 runtime.FuncForPC 获取函数信息
	return runtime.FuncForPC(pc).Name()
}
