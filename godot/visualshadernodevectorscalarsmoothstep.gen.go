package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualShaderNodeVectorScalarSmoothStepFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorScalarSmoothStep {
func newVisualShaderNodeVectorScalarSmoothStepFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorScalarSmoothStep {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorScalarSmoothStep{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeVectorScalarSmoothStep struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorScalarSmoothStep) BaseClass() string {
	return "VisualShaderNodeVectorScalarSmoothStep"
}

// VisualShaderNodeVectorScalarSmoothStepImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorScalarSmoothStep class.
type VisualShaderNodeVectorScalarSmoothStepImplementer interface {
	VisualShaderNodeImplementer
}
