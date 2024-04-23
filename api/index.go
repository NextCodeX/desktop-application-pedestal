package api

import (
	. "github.com/changlie/go-common/a"
	"reflect"
)

type ModuleInfo = map[string]*FunctionExecutor
type ModMap = map[string]ModuleInfo

var allModInfo ModMap = make(map[string]ModuleInfo)

type Req struct {
	Val string `json:"val"`
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func success1(data any) *Resp {
	return &Resp{Code: 0, Msg: "success", Data: data}
}
func success() *Resp {
	return &Resp{Code: 0, Msg: "success", Data: nil}
}
func fail1(msg string) *Resp {
	return &Resp{Code: 500, Msg: msg, Data: nil}
}
func fail2(code int, msg string) *Resp {
	return &Resp{Code: code, Msg: msg, Data: nil}
}

func Entry(c Ctx) {
	mod := c.PathVar("mod")
	action := c.PathVar("action")

	if modInfo, ok := allModInfo[mod]; ok {
		if executor, ok := modInfo[action]; ok {
			res := executor.Run([]reflect.Value{
				reflect.ValueOf(c),
			})
			c.JsonBody(res)
			return
		}
	}
	c.JsonBody(fail1("api not found"))
}
