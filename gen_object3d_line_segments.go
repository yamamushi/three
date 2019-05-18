package three
// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName LineSegments -typeSlug line_segments

import "github.com/gopherjs/gopherjs/js"

func (obj *LineSegments) ApplyMatrix(matrix *Matrix4) {
	obj.Call("applyMatrix", matrix)
}

func (obj *LineSegments) Add(m Object3D) {
	obj.Object.Call("add", m)
}

func (obj *LineSegments) Remove(m *js.Object) {
	obj.Object.Call("remove", m)
}

func (obj *LineSegments) GetObjectById(id int) *js.Object {
	return obj.Call("getObjectById", id)
}

// func (obj *LineSegments) Copy() *LineSegments {
// 	return &LineSegments{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *LineSegments) ToJSON() interface{} {
	return obj.Object.Call("toJSON").Interface()
}

func (obj *LineSegments) getInternalObject() *js.Object {
	return obj.Object
}

func (obj *LineSegments) UpdateMatrix() {
	obj.Call("updateMatrix")
}
