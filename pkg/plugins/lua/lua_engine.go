package lua

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

// LuaEngine 表示 Lua 引擎结构体
type LuaEngine struct {
	vm *lua.LState
}

var luaEngine *LuaEngine

func GetInstance() *LuaEngine {
	if luaEngine == nil {
		luaEngine = &LuaEngine{
			lua.NewState(),
		}
	}
	return luaEngine
}

// ExecuteScript 执行 Lua 脚本
func (e *LuaEngine) ExecuteScript(script string) (any, error) {
	err := e.vm.DoString(script)
	if err != nil {
		return nil, err
	}
	// 获取 Lua 函数
	generateDataFunc := e.vm.GetGlobal("generate_data")
	if generateDataFunc.Type() != lua.LTFunction {
		log.Fatalf("failed to get generate_data function")
	}

	// 调用 Lua 函数
	if err := e.vm.CallByParam(lua.P{
		Fn:      generateDataFunc,
		NRet:    1,
		Protect: true,
	}); err != nil {
		return nil, err
	}
	// 获取函数返回值
	result := e.vm.Get(-1)
	e.vm.Pop(1)
	goResult := LValueToGoValue(result)
	return goResult, nil
}

// Helper functions to convert Go values to Lua values and vice-versa
func GoValueToLValue(L *lua.LState, val interface{}) lua.LValue {
	switch v := val.(type) {
	case bool:
		return lua.LBool(v)
	case float64:
		return lua.LNumber(v)
	case float32:
		return lua.LNumber(float64(v))
	case int:
		return lua.LNumber(float64(v))
	case int8:
		return lua.LNumber(float64(v))
	case int16:
		return lua.LNumber(float64(v))
	case int32:
		return lua.LNumber(float64(v))
	case int64:
		return lua.LNumber(float64(v))
	case uint:
		return lua.LNumber(float64(v))
	case uint8:
		return lua.LNumber(float64(v))
	case uint16:
		return lua.LNumber(float64(v))
	case uint32:
		return lua.LNumber(float64(v))
	case uint64:
		return lua.LNumber(float64(v))
	case string:
		return lua.LString(v)
	case map[string]interface{}:
		table := L.NewTable()
		for key, value := range v {
			table.RawSetString(key, GoValueToLValue(L, value))
		}
		return table
	case []interface{}:
		table := L.NewTable()
		for i, value := range v {
			table.RawSetInt(i+1, GoValueToLValue(L, value))
		}
		return table
	default:
		return lua.LNil
	}
}

func LValueToGoValue(lv lua.LValue) interface{} {
	switch v := lv.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case *lua.LTable:
		// Simple conversion for demonstration, handle nested tables as needed
		m := make(map[string]interface{})
		v.ForEach(func(key lua.LValue, value lua.LValue) {
			if strKey, ok := key.(lua.LString); ok {
				m[string(strKey)] = LValueToGoValue(value)
			}
		})
		return m
	default:
		return nil
	}
}
