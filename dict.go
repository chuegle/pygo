package main

import (
	"github.com/sbinet/go-python"
	"log"
)

type Dict struct {
	dict PO
}

func (d *Dict) Len() int {
	return python.PyDict_Size(d.dict)

}

func (d *Dict) Values() *List {
	return &List{python.PyDict_Values(d.dict)}
}

func (d *Dict) Keys() *List {
	return &List{python.PyDict_Keys(d.dict)}
}

func (d *Dict) Get(key string) *python.PyObject {
	return python.PyDict_GetItemString(d.dict, key)
}

func (d *Dict) Copy() *Dict {
	return &Dict{python.PyDict_Copy(d.dict)}
}

func (d *Dict) HasPy(key *python.PyObject) bool {
	v, err := python.PyDict_Contains(d.dict, key)
	if err != nil {
		log.Fatal("Python dict error", err)
	}
	return v
}

func (d *Dict) Has(key string) bool {
	return d.HasPy(P_str(key))
}

func (d *Dict) RemovePy(key *python.PyObject) {
	err := python.PyDict_DelItem(d.dict, key)
	if err != nil {
		log.Fatal("Python delete error", err)
	}
}

func (d *Dict) Remove(key string) {
	d.RemovePy(P_str(key))
}

func (d *Dict) FromGo(mp map[interface{}]interface{}) *Dict {
	return d
}
