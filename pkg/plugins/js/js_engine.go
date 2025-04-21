package js

import (
	"github.com/robertkrimen/otto"
)

// JSEngine 表示 JavaScript 引擎结构体
type JSEngine struct {
	vm *otto.Otto
}

var jsEngine *JSEngine

func GetInstance() *JSEngine {
	if jsEngine == nil {
		jsEngine = &JSEngine{
			otto.New(),
		}
	}
	return jsEngine
}

// NewJSEngine 创建一个新的 JavaScript 引擎实例
func NewJSEngine() *JSEngine {
	return &JSEngine{
		vm: otto.New(),
	}
}

// ExecuteScript 执行 JavaScript 脚本
func (e *JSEngine) ExecuteScript(script string) (any, error) {
	// 假设这里调用 generate_data 函数，添加 ctx 参数
	val, err := e.vm.Run(script)
	if err != nil {
		return nil, err
	}
	return jsValueToGoValue(val), nil
}

func jsValueToGoValue(val otto.Value) interface{} {
	switch val.Class() {
	case "Boolean":
		ret, _ := val.ToBoolean()
		return ret
	case "String":
		ret, _ := val.ToString()
		return ret
	case "Array", "Number", "Object":
		obj, _ := val.Export()
		return obj
	default:
		return nil
	}
}
