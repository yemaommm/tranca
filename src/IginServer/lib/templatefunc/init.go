package templatefunc

import (
	"fmt"
	// "time"
	// "html/template"
	"IginServer/conf"
	"IginServer/lib/other"
	"encoding/json"
	"github.com/flosch/pongo2"
	"strconv"
)

// var Temfunc = []template.FuncMap{{}}
var Temfunc = pongo2.Context{}

// func unhtml(x string) interface{} { return template.HTML(x) }

func add(num ...interface{}) interface{} {
	var sum float64
	sum = 0
	for _, i := range num {
		s, _ := strconv.ParseFloat(fmt.Sprintf("%v", i), 64)
		sum += s
	}
	return sum
}
func sub(isum interface{}, num ...interface{}) interface{} {
	var sum float64
	sum, _ = strconv.ParseFloat(fmt.Sprintf("%v", isum), 64)
	for _, i := range num {
		s, _ := strconv.ParseFloat(fmt.Sprintf("%v", i), 64)
		sum -= s
	}
	return sum
}
func mult(num ...interface{}) interface{} {
	var sum float64
	sum = 1
	for _, i := range num {
		s, _ := strconv.ParseFloat(fmt.Sprintf("%v", i), 64)
		sum *= s
	}
	return sum
}
func div(isum interface{}, num ...interface{}) interface{} {
	var sum float64
	sum, _ = strconv.ParseFloat(fmt.Sprintf("%v", isum), 64)
	for _, i := range num {
		s, _ := strconv.ParseFloat(fmt.Sprintf("%v", i), 64)
		sum /= s
	}
	return sum
}

func init() {
	Temfunc["int"] = other.ToInt
	Temfunc["string"] = other.ToString
	Temfunc["conf"] = conf.GET
	Temfunc["add"] = add
	Temfunc["sub"] = sub
	Temfunc["mult"] = mult
	Temfunc["div"] = div
	Temfunc["log"] = ilog
	// Temfunc["unhtml"] = unhtml
	Temfunc["todate"] = todate
	Temfunc["conf"] = config
	Temfunc["mstr"] = mstr
	Temfunc["include"] = include
	Temfunc["jsondecode"] = func(sf string) map[string]interface{} {
		var info map[string]interface{}
		json.Unmarshal([]byte(sf), &info)
		// fmt.Println(info)
		return info
	}
	// Temfunc["gotime"] = time
}
