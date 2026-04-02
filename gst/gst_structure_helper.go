package gst

/*
#include "gst.go.h"
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/go-gst/go-glib/glib"
)

// SerializeFlags flags for gst_structure_serialize_full.
type SerializeFlags int

const (
	// SerializeFlagNone indicates no special flags.
	SerializeFlagNone SerializeFlags = C.GST_SERIALIZE_FLAG_NONE
	// SerializeFlagBackwardCompat serializes using the old format for nested structures.
	SerializeFlagBackwardCompat SerializeFlags = C.GST_SERIALIZE_FLAG_BACKWARD_COMPAT
	// SerializeFlagStrict causes serialization to fail if a value cannot be serialized.
	SerializeFlagStrict SerializeFlags = C.GST_SERIALIZE_FLAG_STRICT
)

// ---------------------------------------------------------------------------
// Typed getters
// ---------------------------------------------------------------------------

// GetString retrieves the string value at key. Returns an error if the field
// does not exist or is not a string.
func (s *Structure) GetString(key string) (string, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cStr := C.gst_structure_get_string(s.Instance(), cKey)
	if cStr == nil {
		return "", fmt.Errorf("no string value at %s", key)
	}
	return C.GoString(cStr), nil
}

// GetBool retrieves the boolean value at key. Returns an error if the field
// does not exist or is not a boolean.
func (s *Structure) GetBool(key string) (bool, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.gboolean
	if C.gst_structure_get_boolean(s.Instance(), cKey, &cVal) == 0 {
		return false, fmt.Errorf("no boolean value at %s", key)
	}
	return gobool(cVal), nil
}

// GetInt retrieves the int value at key. Returns an error if the field
// does not exist or is not an int.
func (s *Structure) GetInt(key string) (int, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.gint
	if C.gst_structure_get_int(s.Instance(), cKey, &cVal) == 0 {
		return 0, fmt.Errorf("no int value at %s", key)
	}
	return int(cVal), nil
}

// GetInt64 retrieves the int64 value at key. Returns an error if the field
// does not exist or is not an int64.
func (s *Structure) GetInt64(key string) (int64, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.gint64
	if C.gst_structure_get_int64(s.Instance(), cKey, &cVal) == 0 {
		return 0, fmt.Errorf("no int64 value at %s", key)
	}
	return int64(cVal), nil
}

// GetUint retrieves the uint value at key. Returns an error if the field
// does not exist or is not a uint.
func (s *Structure) GetUint(key string) (uint, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.guint
	if C.gst_structure_get_uint(s.Instance(), cKey, &cVal) == 0 {
		return 0, fmt.Errorf("no uint value at %s", key)
	}
	return uint(cVal), nil
}

// GetUint64 retrieves the uint64 value at key. Returns an error if the field
// does not exist or is not a uint64.
func (s *Structure) GetUint64(key string) (uint64, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.guint64
	if C.gst_structure_get_uint64(s.Instance(), cKey, &cVal) == 0 {
		return 0, fmt.Errorf("no uint64 value at %s", key)
	}
	return uint64(cVal), nil
}

// GetDouble retrieves the float64 value at key. Returns an error if the field
// does not exist or is not a double.
func (s *Structure) GetDouble(key string) (float64, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.gdouble
	if C.gst_structure_get_double(s.Instance(), cKey, &cVal) == 0 {
		return 0, fmt.Errorf("no double value at %s", key)
	}
	return float64(cVal), nil
}

// GetFraction retrieves the fraction value at key as numerator and denominator.
// Returns an error if the field does not exist or is not a fraction.
func (s *Structure) GetFraction(key string) (numerator, denominator int, err error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cNum, cDen C.gint
	if C.gst_structure_get_fraction(s.Instance(), cKey, &cNum, &cDen) == 0 {
		return 0, 0, fmt.Errorf("no fraction value at %s", key)
	}
	return int(cNum), int(cDen), nil
}

// GetClockTime retrieves the ClockTime value at key. Returns an error if the
// field does not exist or is not a ClockTime.
func (s *Structure) GetClockTime(key string) (ClockTime, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.GstClockTime
	if C.gst_structure_get_clock_time(s.Instance(), cKey, &cVal) == 0 {
		return 0, fmt.Errorf("no clock time value at %s", key)
	}
	return ClockTime(cVal), nil
}

// GetFieldType returns the GType of the field with the given name, or
// glib.TypeInvalid if the field is not found.
func (s *Structure) GetFieldType(key string) glib.Type {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return glib.Type(C.gst_structure_get_field_type(s.Instance(), cKey))
}

// ---------------------------------------------------------------------------
// Typed setters
// ---------------------------------------------------------------------------

// SetString sets a string value at key.
func (s *Structure) SetString(key, value string) error {
	return s.SetValue(key, value)
}

// SetBool sets a boolean value at key.
func (s *Structure) SetBool(key string, value bool) error {
	return s.SetValue(key, value)
}

// SetInt sets an int value at key.
func (s *Structure) SetInt(key string, value int) error {
	return s.SetValue(key, value)
}

// SetInt64 sets an int64 value at key.
func (s *Structure) SetInt64(key string, value int64) error {
	return s.SetValue(key, value)
}

// SetUint sets a uint value at key.
func (s *Structure) SetUint(key string, value uint) error {
	return s.SetValue(key, value)
}

// SetUint64 sets a uint64 value at key.
func (s *Structure) SetUint64(key string, value uint64) error {
	return s.SetValue(key, value)
}

// SetDouble sets a float64 value at key.
func (s *Structure) SetDouble(key string, value float64) error {
	return s.SetValue(key, value)
}

// SetFraction sets a fraction value at key.
func (s *Structure) SetFraction(key string, numerator, denominator int) error {
	return s.SetValue(key, &FractionValue{num: numerator, denom: denominator})
}

// ---------------------------------------------------------------------------
// Has / Is helpers
// ---------------------------------------------------------------------------

// HasField checks if the structure contains a field named key.
func (s *Structure) HasField(key string) bool {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return gobool(C.gst_structure_has_field(s.Instance(), cKey))
}

// HasFieldTyped checks if the structure contains a field named key with the given GType.
func (s *Structure) HasFieldTyped(key string, typ glib.Type) bool {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return gobool(C.gst_structure_has_field_typed(s.Instance(), cKey, C.GType(typ)))
}

// HasName checks if the structure has the given name.
func (s *Structure) HasName(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return gobool(C.gst_structure_has_name(s.Instance(), cName))
}

// IsEqual tests if the two structures are equal.
func (s *Structure) IsEqual(other *Structure) bool {
	return gobool(C.gst_structure_is_equal(s.Instance(), other.Instance()))
}

// IsSubset checks if this structure is a subset of superset, i.e. has the same
// name and for all fields that are existing in superset, this structure has a
// value that is a subset of the value in superset.
func (s *Structure) IsSubset(superset *Structure) bool {
	return gobool(C.gst_structure_is_subset(s.Instance(), superset.Instance()))
}

// ---------------------------------------------------------------------------
// Fixate helpers
// ---------------------------------------------------------------------------

// Fixate fixates all values in the structure in-place using gst_value_fixate.
func (s *Structure) Fixate() {
	C.gst_structure_fixate(s.Instance())
}

// FixateField fixates the given field by changing it to its fixated value.
func (s *Structure) FixateField(fieldName string) bool {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return gobool(C.gst_structure_fixate_field(s.Instance(), cField))
}

// FixateFieldBool fixates the given field to the given target boolean if not fixed yet.
func (s *Structure) FixateFieldBool(fieldName string, target bool) bool {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return gobool(C.gst_structure_fixate_field_boolean(s.Instance(), cField, gboolean(target)))
}

// FixateFieldNearestDouble fixates the given field to the nearest double to target.
func (s *Structure) FixateFieldNearestDouble(fieldName string, target float64) bool {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return gobool(C.gst_structure_fixate_field_nearest_double(s.Instance(), cField, C.gdouble(target)))
}

// FixateFieldNearestFraction fixates the given field to the nearest fraction to
// targetNum/targetDenom.
func (s *Structure) FixateFieldNearestFraction(fieldName string, targetNum, targetDenom int) bool {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return gobool(C.gst_structure_fixate_field_nearest_fraction(
		s.Instance(), cField, C.gint(targetNum), C.gint(targetDenom),
	))
}

// FixateFieldNearestInt fixates the given field to the nearest int to target.
func (s *Structure) FixateFieldNearestInt(fieldName string, target int) bool {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	return gobool(C.gst_structure_fixate_field_nearest_int(s.Instance(), cField, C.gint(target)))
}

// FixateFieldString fixates the given field to the given target string if not fixed yet.
func (s *Structure) FixateFieldString(fieldName string, target string) bool {
	cField := C.CString(fieldName)
	defer C.free(unsafe.Pointer(cField))
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	return gobool(C.gst_structure_fixate_field_string(s.Instance(), cField, cTarget))
}

// ---------------------------------------------------------------------------
// Nth helpers
// ---------------------------------------------------------------------------

// NthFieldName returns the name of the field at the given index (0-based).
func (s *Structure) NthFieldName(index uint) string {
	return C.GoString(C.gst_structure_nth_field_name(s.Instance(), C.guint(index)))
}

// ---------------------------------------------------------------------------
// Remove helpers
// ---------------------------------------------------------------------------

// RemoveAllFields removes all fields from the structure.
func (s *Structure) RemoveAllFields() {
	C.gst_structure_remove_all_fields(s.Instance())
}

// RemoveFields removes the fields with the given names. If a field does not
// exist, the argument is ignored.
func (s *Structure) RemoveFields(fields ...string) {
	for _, field := range fields {
		cField := C.CString(field)
		C.gst_structure_remove_field(s.Instance(), cField)
		C.free(unsafe.Pointer(cField))
	}
}

// ---------------------------------------------------------------------------
// Intersection
// ---------------------------------------------------------------------------

// Intersect intersects this structure with other and returns the intersection.
// Returns nil if the intersection is empty.
func (s *Structure) Intersect(other *Structure) *Structure {
	result := C.gst_structure_intersect(s.Instance(), other.Instance())
	if result == nil {
		return nil
	}
	return structureFromGlibFull(result)
}

// CanIntersect tries intersecting this structure with other and reports
// whether the result would not be empty.
func (s *Structure) CanIntersect(other *Structure) bool {
	return gobool(C.gst_structure_can_intersect(s.Instance(), other.Instance()))
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

// SerializeFull serializes the structure to a string using the given flags.
func (s *Structure) SerializeFull(flags SerializeFlags) string {
	cStr := C.gst_structure_serialize_full(s.Instance(), C.GstSerializeFlags(flags))
	defer C.g_free((C.gpointer)(cStr))
	return C.GoString(cStr)
}
