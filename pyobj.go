package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

type PO *python.PyObject

type Py struct {
	obj *python.PyObject
}

type Fn struct {
	Py
}

type Class struct {
	Py
}

type Object struct {
	Py
}

type Module struct {
	Py
}

func (m *Module) Attr(attr string) *python.PyObject {
	return m.Py.obj.GetAttrString(attr)
}

func (m *Module) Fn(attr string) *Fn {
	return &Fn{Py{m.Attr(attr)}}
}

func (p *Py) makeArgs(args []*python.PyObject, kwargs map[string]*python.PyObject) *python.PyObject {
	/* py_all := python.PyTuple_New(2)
	py_args := python.PyTuple_New(len(args))
	for i, a := range args {
	python.PyTuple_SET_ITEM(py_args, i , a)
	}
	python.PyTuple_SET_ITEM(py_all, 0, py_args)
	python.PyTuple_SET_ITEM(py_all, 1, python.PyDict_New())
	return py_args
	*/
	return (&List{topy(args)}).ToTuple()
}

func (p *Py) Call(args ...*python.PyObject) *python.PyObject {
	py_args := p.makeArgs(args, nil)
	result := p.obj.CallObject(py_args)
	fmt.Println("Call Res:", togo(result))
	return result
}

