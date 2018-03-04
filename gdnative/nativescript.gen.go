package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <nativescript/godot_nativescript.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// MethodRpcMode is a Go wrapper for the C.godot_method_rpc_mode enum type.
type MethodRpcMode int

func (e MethodRpcMode) getBase() C.godot_method_rpc_mode {
	return C.godot_method_rpc_mode(e)
}

const (
	MethodRpcModeDisabled MethodRpcMode = iota
	MethodRpcModeRemote
	MethodRpcModeSync
	MethodRpcModeMaster
	MethodRpcModeSlave
)

// PropertyHint is a Go wrapper for the C.godot_property_hint enum type.
type PropertyHint int

func (e PropertyHint) getBase() C.godot_property_hint {
	return C.godot_property_hint(e)
}

const (
	PropertyHintNone PropertyHint = iota
	PropertyHintRange
	PropertyHintExpRange
	PropertyHintEnum
	PropertyHintSpriteFrame
	PropertyHintLayers2DRender
	PropertyHintLayers2DPhysics
	PropertyHintLayers3DRender
	PropertyHintLayers3DPhysics
	PropertyHintDir
	PropertyHintGlobalDir
	PropertyHintResourceType
	PropertyHintMultilineText
	PropertyHintColorNoAlpha
	PropertyHintImageCompressLossy
	PropertyHintImageCompressLossless
	PropertyHintObjectId
	PropertyHintTypeString
	PropertyHintMethodOfVariantType
	PropertyHintMethodOfBaseType
	PropertyHintMethodOfInstance
	PropertyHintMethodOfScript
	PropertyHintPropertyOfVariantType
	PropertyHintPropertyOfBaseType
	PropertyHintPropertyOfInstance
	PropertyHintPropertyOfScript
	PropertyHintMax
)

// NewEmptyPropertyAttributes will return a pointer to an empty
// initialized PropertyAttributes. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyPropertyAttributes() Pointer {
	var obj C.godot_property_attributes
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromPropertyAttributes will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromPropertyAttributes(obj PropertyAttributes) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewPropertyAttributesFromPointer will return a PropertyAttributes from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewPropertyAttributesFromPointer(ptr Pointer) PropertyAttributes {
	return PropertyAttributes{base: (*C.godot_property_attributes)(ptr.getBase())}
}

type PropertyAttributes struct {
	base *C.godot_property_attributes

	RsetType     MethodRpcMode
	Type         Int
	Hint         PropertyHint
	HintString   String
	Usage        PropertyUsageFlags
	DefaultValue Variant
}

func (gdt PropertyAttributes) getBase() *C.godot_property_attributes {
	return gdt.base
}

// NewEmptySignalArgument will return a pointer to an empty
// initialized SignalArgument. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptySignalArgument() Pointer {
	var obj C.godot_signal_argument
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromSignalArgument will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromSignalArgument(obj SignalArgument) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewSignalArgumentFromPointer will return a SignalArgument from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewSignalArgumentFromPointer(ptr Pointer) SignalArgument {
	return SignalArgument{base: (*C.godot_signal_argument)(ptr.getBase())}
}

type SignalArgument struct {
	base *C.godot_signal_argument

	Name         String
	Type         Int
	Hint         PropertyHint
	HintString   String
	Usage        PropertyUsageFlags
	DefaultValue Variant
}

func (gdt SignalArgument) getBase() *C.godot_signal_argument {
	return gdt.base
}

// NewEmptySignal will return a pointer to an empty
// initialized Signal. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptySignal() Pointer {
	var obj C.godot_signal
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromSignal will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromSignal(obj Signal) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewSignalFromPointer will return a Signal from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewSignalFromPointer(ptr Pointer) Signal {
	return Signal{base: (*C.godot_signal)(ptr.getBase())}
}

type Signal struct {
	base *C.godot_signal

	Name           String
	NumArgs        Int
	Args           SignalArgument
	NumDefaultArgs Int
	DefaultArgs    Variant
}

func (gdt Signal) getBase() *C.godot_signal {
	return gdt.base
}

// NewEmptyMethodArg will return a pointer to an empty
// initialized MethodArg. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyMethodArg() Pointer {
	var obj C.godot_method_arg
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromMethodArg will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromMethodArg(obj MethodArg) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewMethodArgFromPointer will return a MethodArg from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewMethodArgFromPointer(ptr Pointer) MethodArg {
	return MethodArg{base: (*C.godot_method_arg)(ptr.getBase())}
}

type MethodArg struct {
	base *C.godot_method_arg

	Name       String
	Type       VariantType
	Hint       PropertyHint
	HintString String
}

func (gdt MethodArg) getBase() *C.godot_method_arg {
	return gdt.base
}