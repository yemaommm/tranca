package py

import (
	"fmt"
	"github.com/qiniu/py"
	"strings"
	"sync"
	// "github.com/qiniu/py/pyutil"
)

var mutex sync.Mutex
var code *py.Code
var mod *py.Module

func Do(fun string, arg ...string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	rem := make([]*py.Base, 0)
	defer remove(rem)

	n := strings.Split(fun, ".")

	w, err := mod.GetAttrString(n[0])
	if err != nil {
		return "", err
	}
	rem = append(rem, w)
	for _, i := range n[1:] {
		w, err = w.GetAttrString(i)
		if err != nil {
			return "", err
		}
		rem = append(rem, w)
	}
	ret, err := w.CallObject(nil)
	if err != nil {
		return "", err
	}
	defer ret.Decref()
	return ret.String(), nil
}

func remove(rem []*py.Base) {
	for _, i := range rem {
		i.Decref()
	}
}

func init() {
	py.AddToPath("./py/")
	var err error
	code, err = py.CompileFile("py/init.py", py.FileInput)
	if err != nil {
		fmt.Println("py:", err)
		return
	}
	// defer code.Decref()
	mod, err = py.ExecCodeModule("go", code.Obj())
	if err != nil {
		fmt.Println("py:", err)
		return
	}
	// defer mod.Decref()
	fmt.Println(mod.Dir())
}
