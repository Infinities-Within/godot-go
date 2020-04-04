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
#cgo CFLAGS: -I../godot_headers
#include "gdnative.gen.h"
// #include <godot_headers/gdnative/basis.h>
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

// NewEmptyBasis will return a pointer to an empty
// initialized Basis. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyBasis() Pointer {
	var obj C.godot_basis
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromBasis will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromBasis(obj Basis) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewBasisFromPointer will return a Basis from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewBasisFromPointer(ptr Pointer) Basis {

	return Basis{base: (*C.godot_basis)(ptr.getBase())}
}

type Basis struct {
	base *C.godot_basis
}

func (gdt Basis) getBase() *C.godot_basis {
	return gdt.base
}

// NewBasisWithRows godot_basis_new_with_rows [[godot_basis * r_dest] [const godot_vector3 * p_x_axis] [const godot_vector3 * p_y_axis] [const godot_vector3 * p_z_axis]] void
func NewBasisWithRows(xAxis Vector3, yAxis Vector3, zAxis Vector3) Basis {
	var dest C.godot_basis
	arg1 := xAxis.getBase()
	arg2 := yAxis.getBase()
	arg3 := zAxis.getBase()
	C.go_godot_basis_new_with_rows(GDNative.api, &dest, arg1, arg2, arg3)
	return Basis{base: &dest}
}

// NewBasisWithAxisAndAngle godot_basis_new_with_axis_and_angle [[godot_basis * r_dest] [const godot_vector3 * p_axis] [const godot_real p_phi]] void
func NewBasisWithAxisAndAngle(axis Vector3, phi Real) Basis {
	var dest C.godot_basis
	arg1 := axis.getBase()
	arg2 := phi.getBase()
	C.go_godot_basis_new_with_axis_and_angle(GDNative.api, &dest, arg1, arg2)
	return Basis{base: &dest}
}

// NewBasisWithEuler godot_basis_new_with_euler [[godot_basis * r_dest] [const godot_vector3 * p_euler]] void
func NewBasisWithEuler(euler Vector3) Basis {
	var dest C.godot_basis
	arg1 := euler.getBase()
	C.go_godot_basis_new_with_euler(GDNative.api, &dest, arg1)
	return Basis{base: &dest}
}

// AsString godot_basis_as_string [[const godot_basis * p_self]] godot_string
func (gdt *Basis) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_as_string(GDNative.api, arg0)

	utfStr := C.go_godot_string_utf8(GDNative.api, &ret)
	char := C.go_godot_char_string_get_data(GDNative.api, &utfStr)
	goStr := C.GoString(char)
	C.go_godot_char_string_destroy(GDNative.api, &utfStr)

	return String(goStr)

}

// Inverse godot_basis_inverse [[const godot_basis * p_self]] godot_basis
func (gdt *Basis) Inverse() Basis {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_inverse(GDNative.api, arg0)

	return Basis{base: &ret}

}

// Transposed godot_basis_transposed [[const godot_basis * p_self]] godot_basis
func (gdt *Basis) Transposed() Basis {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_transposed(GDNative.api, arg0)

	return Basis{base: &ret}

}

// Orthonormalized godot_basis_orthonormalized [[const godot_basis * p_self]] godot_basis
func (gdt *Basis) Orthonormalized() Basis {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_orthonormalized(GDNative.api, arg0)

	return Basis{base: &ret}

}

// Determinant godot_basis_determinant [[const godot_basis * p_self]] godot_real
func (gdt *Basis) Determinant() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_determinant(GDNative.api, arg0)

	return Real(ret)
}

// Rotated godot_basis_rotated [[const godot_basis * p_self] [const godot_vector3 * p_axis] [const godot_real p_phi]] godot_basis
func (gdt *Basis) Rotated(axis Vector3, phi Real) Basis {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()
	arg2 := phi.getBase()

	ret := C.go_godot_basis_rotated(GDNative.api, arg0, arg1, arg2)

	return Basis{base: &ret}

}

// Scaled godot_basis_scaled [[const godot_basis * p_self] [const godot_vector3 * p_scale]] godot_basis
func (gdt *Basis) Scaled(scale Vector3) Basis {
	arg0 := gdt.getBase()
	arg1 := scale.getBase()

	ret := C.go_godot_basis_scaled(GDNative.api, arg0, arg1)

	return Basis{base: &ret}

}

// GetScale godot_basis_get_scale [[const godot_basis * p_self]] godot_vector3
func (gdt *Basis) GetScale() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_get_scale(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// GetEuler godot_basis_get_euler [[const godot_basis * p_self]] godot_vector3
func (gdt *Basis) GetEuler() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_get_euler(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// Tdotx godot_basis_tdotx [[const godot_basis * p_self] [const godot_vector3 * p_with]] godot_real
func (gdt *Basis) Tdotx(with Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := with.getBase()

	ret := C.go_godot_basis_tdotx(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Tdoty godot_basis_tdoty [[const godot_basis * p_self] [const godot_vector3 * p_with]] godot_real
func (gdt *Basis) Tdoty(with Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := with.getBase()

	ret := C.go_godot_basis_tdoty(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Tdotz godot_basis_tdotz [[const godot_basis * p_self] [const godot_vector3 * p_with]] godot_real
func (gdt *Basis) Tdotz(with Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := with.getBase()

	ret := C.go_godot_basis_tdotz(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Xform godot_basis_xform [[const godot_basis * p_self] [const godot_vector3 * p_v]] godot_vector3
func (gdt *Basis) Xform(v Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_basis_xform(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// XformInv godot_basis_xform_inv [[const godot_basis * p_self] [const godot_vector3 * p_v]] godot_vector3
func (gdt *Basis) XformInv(v Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_basis_xform_inv(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// GetOrthogonalIndex godot_basis_get_orthogonal_index [[const godot_basis * p_self]] godot_int
func (gdt *Basis) GetOrthogonalIndex() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_basis_get_orthogonal_index(GDNative.api, arg0)

	return Int(ret)
}

// NewBasis godot_basis_new [[godot_basis * r_dest]] void
func NewBasis() Basis {
	var dest C.godot_basis
	C.go_godot_basis_new(GDNative.api, &dest)
	return Basis{base: &dest}
}

// NewBasisWithEulerQuat godot_basis_new_with_euler_quat [[godot_basis * r_dest] [const godot_quat * p_euler]] void
func NewBasisWithEulerQuat(euler Quat) Basis {
	var dest C.godot_basis
	arg1 := euler.getBase()
	C.go_godot_basis_new_with_euler_quat(GDNative.api, &dest, arg1)
	return Basis{base: &dest}
}

// GetElements godot_basis_get_elements [[const godot_basis * p_self] [godot_vector3 * p_elements]] void
func (gdt *Basis) GetElements(elements Vector3) {
	arg0 := gdt.getBase()
	arg1 := elements.getBase()

	C.go_godot_basis_get_elements(GDNative.api, arg0, arg1)
}

// GetAxis godot_basis_get_axis [[const godot_basis * p_self] [const godot_int p_axis]] godot_vector3
func (gdt *Basis) GetAxis(axis Int) Vector3 {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()

	ret := C.go_godot_basis_get_axis(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// SetAxis godot_basis_set_axis [[godot_basis * p_self] [const godot_int p_axis] [const godot_vector3 * p_value]] void
func (gdt *Basis) SetAxis(axis Int, value Vector3) {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()
	arg2 := value.getBase()

	C.go_godot_basis_set_axis(GDNative.api, arg0, arg1, arg2)
}

// GetRow godot_basis_get_row [[const godot_basis * p_self] [const godot_int p_row]] godot_vector3
func (gdt *Basis) GetRow(row Int) Vector3 {
	arg0 := gdt.getBase()
	arg1 := row.getBase()

	ret := C.go_godot_basis_get_row(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// SetRow godot_basis_set_row [[godot_basis * p_self] [const godot_int p_row] [const godot_vector3 * p_value]] void
func (gdt *Basis) SetRow(row Int, value Vector3) {
	arg0 := gdt.getBase()
	arg1 := row.getBase()
	arg2 := value.getBase()

	C.go_godot_basis_set_row(GDNative.api, arg0, arg1, arg2)
}

// OperatorEqual godot_basis_operator_equal [[const godot_basis * p_self] [const godot_basis * p_b]] godot_bool
func (gdt *Basis) OperatorEqual(b Basis) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_basis_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorAdd godot_basis_operator_add [[const godot_basis * p_self] [const godot_basis * p_b]] godot_basis
func (gdt *Basis) OperatorAdd(b Basis) Basis {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_basis_operator_add(GDNative.api, arg0, arg1)

	return Basis{base: &ret}

}

// OperatorSubtract godot_basis_operator_subtract [[const godot_basis * p_self] [const godot_basis * p_b]] godot_basis
func (gdt *Basis) OperatorSubtract(b Basis) Basis {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_basis_operator_subtract(GDNative.api, arg0, arg1)

	return Basis{base: &ret}

}

// OperatorMultiplyVector godot_basis_operator_multiply_vector [[const godot_basis * p_self] [const godot_basis * p_b]] godot_basis
func (gdt *Basis) OperatorMultiplyVector(b Basis) Basis {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_basis_operator_multiply_vector(GDNative.api, arg0, arg1)

	return Basis{base: &ret}

}

// OperatorMultiplyScalar godot_basis_operator_multiply_scalar [[const godot_basis * p_self] [const godot_real p_b]] godot_basis
func (gdt *Basis) OperatorMultiplyScalar(b Real) Basis {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_basis_operator_multiply_scalar(GDNative.api, arg0, arg1)

	return Basis{base: &ret}

}
