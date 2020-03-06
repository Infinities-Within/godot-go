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

//func NewAnimationNodeTimeSeekFromPointer(ptr gdnative.Pointer) AnimationNodeTimeSeek {
func newAnimationNodeTimeSeekFromPointer(ptr gdnative.Pointer) AnimationNodeTimeSeek {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeTimeSeek{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationNodeTimeSeek struct {
	AnimationNode
	owner gdnative.Object
}

func (o *AnimationNodeTimeSeek) BaseClass() string {
	return "AnimationNodeTimeSeek"
}

// AnimationNodeTimeSeekImplementer is an interface that implements the methods
// of the AnimationNodeTimeSeek class.
type AnimationNodeTimeSeekImplementer interface {
	AnimationNodeImplementer
}