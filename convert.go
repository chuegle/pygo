package main

import (
	"github.com/sbinet/go-python"
	"reflect"
	"log"
	"fmt"
)

var G_int func(*python.PyObject) int = python.PyInt_AsLong
var G_float func(*python.PyObject) float32 = python.PyFloat_AsDouble
var G_str func(*python.PyObject) string = python.PyString_AsString

func GI_int(p *python.PyObject) interface{} {
	return G_int(p)
}
func GI_str(p *python.PyObject) interface{} {
	return G_str(p)
}

var P_int func(int) *python.PyObject = python.PyLong_FromLong
var P_float func(float32) *python.PyObject = python.PyFloat_FromDouble
var P_str func(string) *python.PyObject = python.PyString_FromString

func PyType(p *python.PyObject) string {
	return G_str(p.Type().GetAttrString("__name__"))
}

func togo(p *python.PyObject) interface{} {
	object_type := PyType(p)
	switch object_type {
	case "str", "unicode":
		return G_str(p)
	case "int":
		return G_int(p)
	case "list":
		return &List{p}
	case "tuple":
		return (&List{}).FromTuple(p)
	case "float":
		return G_float(p)
	case "dict":
		return Dict{p}
	case "NoneType":
		return nil
	default:
		log.Fatal("Unsupported python type", object_type)
	}
	fmt.Println(object_type)
	return object_type
}

func topy(g interface{}) *python.PyObject {
	switch g := g.(type) {
	case string:
		return P_str(g)
	case int, int8, int16, int32, int64:
		return P_int(g.(int))
	case []interface{}:
		return (&List{}).FromGo(g).list
		/*case []*python.PyObject:
		for i :=
		return tester(g...).list */
	case map[interface{}]interface{}:
		return (&Dict{}).FromGo(g).dict
	case float32, float64:
		return P_float(g.(float32))
	case Dict:
		return g.dict
	case List:
		return g.list
	case *python.PyObject:
		return g
		// bool
		// Uint
		//
	default:
		v := reflect.ValueOf(g)
		kind := v.Kind().String()
		if kind == "slice" || kind == "array" {
			x := make([]interface{}, v.Len())
			for i := 0; i < v.Len(); i++ {
				x[i] = v.Index(i).Interface()
			}
			return (&List{}).FromGo(x).list
		}
		fmt.Println("G", g)
		fmt.Println("TOG", reflect.TypeOf(g))
		fmt.Println("V", v)
		fmt.Println("VK", v.Kind())
		log.Fatal("Unsupported go type", g, reflect.TypeOf(g), v.Kind().String())
	}
	return nil
}

