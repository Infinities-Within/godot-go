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

//func NewAudioStreamGeneratorPlaybackFromPointer(ptr gdnative.Pointer) AudioStreamGeneratorPlayback {
func newAudioStreamGeneratorPlaybackFromPointer(ptr gdnative.Pointer) AudioStreamGeneratorPlayback {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamGeneratorPlayback{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AudioStreamGeneratorPlayback struct {
	AudioStreamPlaybackResampled
	owner gdnative.Object
}

func (o *AudioStreamGeneratorPlayback) BaseClass() string {
	return "AudioStreamGeneratorPlayback"
}

/*

	Args: [{ false amount int}], Returns: bool
*/
func (o *AudioStreamGeneratorPlayback) CanPushBuffer(amount gdnative.Int) gdnative.Bool {
	//log.Println("Calling AudioStreamGeneratorPlayback.CanPushBuffer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamGeneratorPlayback", "can_push_buffer")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: void
*/
func (o *AudioStreamGeneratorPlayback) ClearBuffer() {
	//log.Println("Calling AudioStreamGeneratorPlayback.ClearBuffer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamGeneratorPlayback", "clear_buffer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: int
*/
func (o *AudioStreamGeneratorPlayback) GetFramesAvailable() gdnative.Int {
	//log.Println("Calling AudioStreamGeneratorPlayback.GetFramesAvailable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamGeneratorPlayback", "get_frames_available")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: int
*/
func (o *AudioStreamGeneratorPlayback) GetSkips() gdnative.Int {
	//log.Println("Calling AudioStreamGeneratorPlayback.GetSkips()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamGeneratorPlayback", "get_skips")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false frames PoolVector2Array}], Returns: bool
*/
func (o *AudioStreamGeneratorPlayback) PushBuffer(frames gdnative.PoolVector2Array) gdnative.Bool {
	//log.Println("Calling AudioStreamGeneratorPlayback.PushBuffer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(frames)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamGeneratorPlayback", "push_buffer")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false frame Vector2}], Returns: bool
*/
func (o *AudioStreamGeneratorPlayback) PushFrame(frame gdnative.Vector2) gdnative.Bool {
	//log.Println("Calling AudioStreamGeneratorPlayback.PushFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamGeneratorPlayback", "push_frame")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// AudioStreamGeneratorPlaybackImplementer is an interface that implements the methods
// of the AudioStreamGeneratorPlayback class.
type AudioStreamGeneratorPlaybackImplementer interface {
	AudioStreamPlaybackResampledImplementer
	CanPushBuffer(amount gdnative.Int) gdnative.Bool
	ClearBuffer()
	GetFramesAvailable() gdnative.Int
	GetSkips() gdnative.Int
	PushBuffer(frames gdnative.PoolVector2Array) gdnative.Bool
	PushFrame(frame gdnative.Vector2) gdnative.Bool
}