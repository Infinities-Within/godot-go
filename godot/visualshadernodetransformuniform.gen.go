package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualShaderNodeTransformUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeTransformUniform {
func newVisualShaderNodeTransformUniformFromPointer(ptr gdnative.Pointer) VisualShaderNodeTransformUniform {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeTransformUniform{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeTransformUniform struct {
	VisualShaderNodeUniform
	owner gdnative.Object
}

func (o *VisualShaderNodeTransformUniform) BaseClass() string {
	return "VisualShaderNodeTransformUniform"
}

// VisualShaderNodeTransformUniformImplementer is an interface that implements the methods
// of the VisualShaderNodeTransformUniform class.
type VisualShaderNodeTransformUniformImplementer interface {
	VisualShaderNodeUniformImplementer
}