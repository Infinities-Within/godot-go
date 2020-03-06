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

//func NewAnimationNodeBlendTreeFromPointer(ptr gdnative.Pointer) AnimationNodeBlendTree {
func newAnimationNodeBlendTreeFromPointer(ptr gdnative.Pointer) AnimationNodeBlendTree {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeBlendTree{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationNodeBlendTree struct {
	AnimationRootNode
	owner gdnative.Object
}

func (o *AnimationNodeBlendTree) BaseClass() string {
	return "AnimationNodeBlendTree"
}

/*
        Undocumented
	Args: [{ false node String}], Returns: void
*/
func (o *AnimationNodeBlendTree) X_NodeChanged(node gdnative.String) {
	//log.Println("Calling AnimationNodeBlendTree.X_NodeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(node)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "_node_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationNodeBlendTree) X_TreeChanged() {
	//log.Println("Calling AnimationNodeBlendTree.X_TreeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "_tree_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false node AnimationNode} {(0, 0) true position Vector2}], Returns: void
*/
func (o *AnimationNodeBlendTree) AddNode(name gdnative.String, node AnimationNodeImplementer, position gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendTree.AddNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "add_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false input_node String} { false input_index int} { false output_node String}], Returns: void
*/
func (o *AnimationNodeBlendTree) ConnectNode(inputNode gdnative.String, inputIndex gdnative.Int, outputNode gdnative.String) {
	//log.Println("Calling AnimationNodeBlendTree.ConnectNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(inputNode)
	ptrArguments[1] = gdnative.NewPointerFromInt(inputIndex)
	ptrArguments[2] = gdnative.NewPointerFromString(outputNode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "connect_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false input_node String} { false input_index int}], Returns: void
*/
func (o *AnimationNodeBlendTree) DisconnectNode(inputNode gdnative.String, inputIndex gdnative.Int) {
	//log.Println("Calling AnimationNodeBlendTree.DisconnectNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(inputNode)
	ptrArguments[1] = gdnative.NewPointerFromInt(inputIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "disconnect_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *AnimationNodeBlendTree) GetGraphOffset() gdnative.Vector2 {
	//log.Println("Calling AnimationNodeBlendTree.GetGraphOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "get_graph_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: AnimationNode
*/
func (o *AnimationNodeBlendTree) GetNode(name gdnative.String) AnimationNodeImplementer {
	//log.Println("Calling AnimationNodeBlendTree.GetNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "get_node")

	// Call the parent method.
	// AnimationNode
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAnimationNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AnimationNodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AnimationNode" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AnimationNodeImplementer)
	}

	return &ret
}

/*

	Args: [{ false name String}], Returns: Vector2
*/
func (o *AnimationNodeBlendTree) GetNodePosition(name gdnative.String) gdnative.Vector2 {
	//log.Println("Calling AnimationNodeBlendTree.GetNodePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "get_node_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: bool
*/
func (o *AnimationNodeBlendTree) HasNode(name gdnative.String) gdnative.Bool {
	//log.Println("Calling AnimationNodeBlendTree.HasNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "has_node")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNodeBlendTree) RemoveNode(name gdnative.String) {
	//log.Println("Calling AnimationNodeBlendTree.RemoveNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "remove_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false new_name String}], Returns: void
*/
func (o *AnimationNodeBlendTree) RenameNode(name gdnative.String, newName gdnative.String) {
	//log.Println("Calling AnimationNodeBlendTree.RenameNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(newName)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "rename_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *AnimationNodeBlendTree) SetGraphOffset(offset gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendTree.SetGraphOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "set_graph_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false position Vector2}], Returns: void
*/
func (o *AnimationNodeBlendTree) SetNodePosition(name gdnative.String, position gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendTree.SetNodePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendTree", "set_node_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeBlendTreeImplementer is an interface that implements the methods
// of the AnimationNodeBlendTree class.
type AnimationNodeBlendTreeImplementer interface {
	AnimationRootNodeImplementer
	X_NodeChanged(node gdnative.String)
	X_TreeChanged()
	AddNode(name gdnative.String, node AnimationNodeImplementer, position gdnative.Vector2)
	ConnectNode(inputNode gdnative.String, inputIndex gdnative.Int, outputNode gdnative.String)
	DisconnectNode(inputNode gdnative.String, inputIndex gdnative.Int)
	GetGraphOffset() gdnative.Vector2
	GetNode(name gdnative.String) AnimationNodeImplementer
	GetNodePosition(name gdnative.String) gdnative.Vector2
	HasNode(name gdnative.String) gdnative.Bool
	RemoveNode(name gdnative.String)
	RenameNode(name gdnative.String, newName gdnative.String)
	SetGraphOffset(offset gdnative.Vector2)
	SetNodePosition(name gdnative.String, position gdnative.Vector2)
}