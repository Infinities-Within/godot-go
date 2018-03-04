package class

import (
	"log"

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

//func NewGrooveJoint2DFromPointer(ptr gdnative.Pointer) GrooveJoint2D {
func NewGrooveJoint2DFromPointer(ptr gdnative.Pointer) GrooveJoint2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GrooveJoint2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Groove constraint for 2D physics. This is useful for making a body "slide" through a segment placed in another.
*/
type GrooveJoint2D struct {
	Joint2D
	owner gdnative.Object
}

func (o *GrooveJoint2D) BaseClass() string {
	return "GrooveJoint2D"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *GrooveJoint2D) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *GrooveJoint2D) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *GrooveJoint2D) GetInitialOffset() gdnative.Float {
	log.Println("Calling GrooveJoint2D.GetInitialOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GrooveJoint2D", "get_initial_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *GrooveJoint2D) GetLength() gdnative.Float {
	log.Println("Calling GrooveJoint2D.GetLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GrooveJoint2D", "get_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false offset float}], Returns: void
*/
func (o *GrooveJoint2D) SetInitialOffset(offset gdnative.Float) {
	log.Println("Calling GrooveJoint2D.SetInitialOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GrooveJoint2D", "set_initial_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *GrooveJoint2D) SetLength(length gdnative.Float) {
	log.Println("Calling GrooveJoint2D.SetLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GrooveJoint2D", "set_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}