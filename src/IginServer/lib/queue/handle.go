package queue

import (
	// "fmt"
	"reflect"
	"server/lib/other"
	"server/lib/redis"
)

var Work = &Lwork{map[string]func(string){}, map[string]int64{}}

type Lwork struct {
	mywork map[string]func(string)
	mynum  map[string]int64
}

func (w *Lwork) CreateWork(key string, work func(string)) {
	w.mynum[key] = 0
	w.mywork[key] = work
}

func Add(name, data string) {
	defer other.Thow()
	r := redis.Get()
	defer r.Close()
	r.RPUSH(name, data)

	if nil == Work.mywork[name] {
		// fmt.Println("1")
		return
	} else if 0 == Work.mynum[name] {
		// fmt.Println("2")
		Work.mynum[name] = Work.mynum[name] + 1
		go run(name, Work.mywork[name])
	} else {
		// fmt.Println("3")
		Work.mynum[name] = Work.mynum[name] + 1
	}

}

func Get(name string) (string, error) {
	r := redis.Get()
	defer r.Close()
	return r.LPOP(name)
}

func run(name string, f func(string)) {
	defer other.Thow()
	for {
		data, err := Get(name)
		if err != nil || data == "" {
			Work.mynum[name] = 0
			return
		}
		f(data)
	}
}

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}
