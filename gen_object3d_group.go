package three
// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Group -typeSlug group

import "github.com/gopherjs/gopherjs/js"
	
func (obj *Group) ApplyMatrix(matrix *Matrix4) {
	obj.Call("applyMatrix", matrix)
}

func (obj *Group) Add(m Object3D) {
	obj.Object.Call("add", m.getInternalObject())
}

func (obj *Group) Remove(m *js.Object) {
	//obj.Object.Call("remove", m.getInternalObject())
	obj.Object.Call("remove", m)
}

// func (obj *Group) Copy() *Group {
// 	return &Group{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *Group) ToJSON() interface{} {
	return obj.Object.Call("toJSON").Interface()
}

func (obj *Group) getInternalObject() *js.Object {
	return obj.Object
}
