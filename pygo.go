package main

import (
	"fmt"
	"github.com/sbinet/go-python"
	"log"
)

func tester(args ...interface{}) *List {
	fmt.Println("TESTER", args)
	return &List{}
}

type Python struct {
	_path      *List
	_path_hash int64
	path       []string
	modules    map[string]*Module
}

func (py *Python) syncPath() {
	if py._path == nil {
		sys := py.Import("sys")
		py._path = &List{sys.Attr("path")}
	}
	if py._path.Hash() != py._path_hash {
		py.path = py._path.ToGoStr()
		/*
		   i := 0
		   for po := range py._path.Iterate() {
		   path[i] = G_str(po)
		   i++
		   }
		   py.path = path
		*/
		py._path_hash = py._path.Hash()
	}
}

func (py *Python) Import(module string) *Module {
	var m *Module
	var ok bool
	if m, ok = py.modules[module]; !ok {
		mod := python.PyImport_ImportModuleNoBlock(module)
		if mod == nil {
			log.Fatal("Could not import python module", m)
		}
		if py.modules == nil {
			py.modules = make(map[string]*Module)
		}
		m = &Module{Py{mod}}
		py.modules[module] = m
	}
	return m
}

func (py *Python) pathIndex(path string) (int, bool) {
	py.syncPath()
	for i, p := range py.path {
		if path == p {
			return i, true
		}
	}
	return -1, false
}

func (py *Python) prependPath(path string) {
	if _, ok := py.pathIndex(path); !ok {
		py.path = append([]string{path}, py.path...)
		py._path.Set(0, P_str(path))
	}
	fmt.Println(py.Import("sys"))
}

func main() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
	defer python.Finalize()

	py := Python{}
	py.prependPath("/home/test")

	// fmt.Println(py.path)
	// fmt.Println(G_str(python.PyList_GetItem(py.modules["sys"].GetAttrString("path"), 0)))

	// fmt.Println(python.PyRun_SimpleFile("/home/test/mytest.py"))
	// python.PyRun_SimpleFile("mytest.py")
	// fmt.Println(python.PyRun_SimpleString("print sys.path; B = 3"))

	fmt.Println(python.PyRun_SimpleString("import sys; print sys.path; B = 3"))

	po := py.Import("mytest")
	po.Fn("Test").Call()
	python.PyErr_Print()
	po.Fn("Foo").Call(P_str("X"), P_int(3), P_float(6))
	fmt.Println("OBJECT", python.PyMarshal_ReadObjectFromString("__name__"))

	py.syncPath()

	// fmt.Println(py.path)
	//fmt.Println(python.PyImport_GetModuleDict())
	//fmt.Println(po)
	//fmt.Println(po.GetAttrString("__file__"))

	gostr := "foo"
	pystr := P_str(gostr)
	str := G_str(pystr)
	fmt.Println("hello [", str, "]")
}
