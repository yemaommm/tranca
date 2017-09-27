package other

import (
	"container/list"
	// "fmt"
)

type MyList struct {
	*list.List
}

func NewList() *MyList {
	return &MyList{list.New()}
}

func (l *MyList) Find(v interface{}) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return e
		}
	}
	return nil
}

func (l *MyList) FindLast(v interface{}) *list.Element {
	for e := l.Back(); e != nil; e = e.Prev() {
		if e.Value == v {
			return e
		}
	}
	return nil
}
