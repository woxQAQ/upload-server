package plugins

import (
	"github.com/woxQAQ/upload-server/pkg/plugins/js"
	"github.com/woxQAQ/upload-server/pkg/plugins/lua"
)

// ScriptEngine 定义执行脚本的接口
type ScriptEngine interface {
	ExecuteScript(script string) (any, error)
}

// GetEngine 根据类型获取对应的脚本引擎
func GetEngine(engineType string) ScriptEngine {
	switch engineType {
	case "js":
		return js.GetInstance()
	case "lua":
		return lua.GetInstance()
	default:
		return nil
	}
}
