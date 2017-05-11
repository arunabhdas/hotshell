package interpreter

import "strconv"

const DESC_KEY = "desc"
const CMD_KEY = "cmd"
const KEY_KEY = "key"
const ITEMS_KEY = "items"

type Ast struct {
	Key   string
	Desc  string
	Cmd   string
	Items []Ast
}

func NewAst(value interface{}) []Ast {

	val, validType := value.([]map[string]interface{})
	if !validType {
		return []Ast{}
	}

	list := make([]Ast, len(val))
	for ix, item := range val {
		list[ix] = (&astBuilder{item}).build()
	}

	return list
}

type astBuilder struct {
	value map[string]interface{}
}

func (a *astBuilder) build() Ast {
	ast := Ast{}
	ast.Key = a.getKey()
	ast.Desc = a.getDesc()
	ast.Cmd = a.getCmd()
	ast.Items = NewAst(a.value[ITEMS_KEY])
	return ast
}

func (a *astBuilder) getDesc() string {
	return a.getScalar(DESC_KEY)
}

func (a *astBuilder) getKey() string {
	return a.getScalar(KEY_KEY)
}

func (a *astBuilder) getCmd() string {
	return a.getScalar(CMD_KEY)
}

func (a *astBuilder) getScalar(key string) string {
	value := a.value[key]
	if value == nil {
		return ""
	}
	if intValue, isInt := value.(int); isInt {
		return strconv.Itoa(intValue)
	}
	if intValue, isInt := value.(int64); isInt {
		return strconv.FormatInt(intValue, 10)
	}
	if floatValue, isFloat := value.(float64); isFloat {
		return strconv.FormatFloat(floatValue, 'f', 0, 64)
	}
	return a.value[key].(string)
}
