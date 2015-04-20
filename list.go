package main

import (
	"github.com/sbinet/go-python"
	"log"
)

type List struct {
	list PO
}

func NewList(size int) *List {
	return &List{python.PyList_New(size)}
}

func (l *List) FromTuple(p *python.PyObject) *List {
	size := python.PyTuple_Size(p)
	l.list = python.PyList_New(size)
	for i := 0; i < size; i++ {
		o := python.PyTuple_GetItem(p, i)
		python.PyList_SetItem(l.list, i, o)
	}
	return l
}

func (l *List) ToTuple() *python.PyObject {
	size := l.Len()
	tup := python.PyTuple_New(size)
	for i := 0; i < size; i++ {
		python.PyTuple_SetItem(tup, i, l.Get(i))
	}
	return tup
}

func (l *List) FromGo(arr []interface{}) *List {
	l.list = python.PyList_New(len(arr))
	for i, v := range arr {
		l.Set(i, topy(v))
	}
	return l
}

func (l *List) Len() int {
	return python.PyList_Size(l.list)
}

func (l *List) Insert(p *python.PyObject, i int) {
	if i < 0 {
		i = l.Len() + i
	}
	if err := python.PyList_Insert(l.list, i, p); err != nil {
		log.Fatal("List insert error", err)
	}
}

func (l *List) Append(p *python.PyObject) {
	python.PyList_Append(l.list, p)
}

func (l *List) Set(i int, p *python.PyObject) {
	python.PyList_SetItem(l.list, i, p)
}

func (l *List) Get(i int) *python.PyObject {
	return python.PyList_GetItem(l.list, i)
}

func (l *List) Hash() int64 {
	return int64(l.Len())
}

func (l *List) Iterate() <-chan *python.PyObject {
	c := make(chan *python.PyObject, 10)
	go func() {
		size := l.Len()
		for i := 0; i < size; i++ {
			c <- python.PyList_GetItem(l.list, i)
		}
		close(c)
	}()
	return c
}

func (l *List) ToGo() []interface{} {
	i := 0
	arr := make([]interface{}, l.Len())
	for po := range l.Iterate() {
		arr[i] = togo(po)
		i++
	}
	return arr
}

func (l *List) ToGoStr() []string {
	arr := make([]string, l.Len())
	for i, v := range l.ToGo() {
		arr[i] = v.(string)
	}
	return arr
}

func (l *List) ToSlice() []*python.PyObject {
	arr := make([]*python.PyObject, l.Len())
	i := 0
	for po := range l.Iterate() {
		arr[i] = po
		i++
	}
	return arr
}
