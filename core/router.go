package core

import (
	"reflect"
	"strings"
)

var (
	routerTable  = make(map[string]interface{}, 0)
)

func Router(ifc interface{}) {
	t := reflect.TypeOf(ifc)
	ts := t.String()

	module := ts[strings.LastIndex(ts, ".")+1:]
	module = strings.Replace(module, "Controller", "", 1)
	module = strings.ToLower(module)

	routerTable[module] = ifc
}
