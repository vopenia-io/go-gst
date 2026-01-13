package gst

/*
#include "gst.go.h"
#include <glib-object.h>
#include <stdio.h>

// Wrapper to call g_signal_newv with C types
guint
go_gst_signal_newv(GType type, const gchar *name, GSignalFlags flags, GType return_type, guint n_params, GType *param_types)
{
	return g_signal_newv(
		name,
		type,
		flags,
		NULL, // class_closure: usually NULL for dynamic signals
		NULL, // accumulator: NULL for standard signals
		NULL, // accu_data
		NULL, // c_marshaller: NULL requests g_cclosure_marshal_generic
		return_type,
		n_params,
		param_types
	);
}
*/
import "C"
import (
	"unsafe"

	"github.com/go-gst/go-glib/glib"
)

// SignalFlags corresponds to GSignalFlags
type SignalFlags int

const (
	SignalRunFirst   SignalFlags = C.G_SIGNAL_RUN_FIRST
	SignalRunLast    SignalFlags = C.G_SIGNAL_RUN_LAST
	SignalRunCleanup SignalFlags = C.G_SIGNAL_RUN_CLEANUP
	SignalNoRecurse  SignalFlags = C.G_SIGNAL_NO_RECURSE
	SignalDetailed   SignalFlags = C.G_SIGNAL_DETAILED
	SignalAction     SignalFlags = C.G_SIGNAL_ACTION
	SignalNoHooks    SignalFlags = C.G_SIGNAL_NO_HOOKS
)

// SignalNew registers a new signal on the given class type.
//
// type:       The GType of the class (use klass.Type())
// name:       The name of the signal (e.g. "my-event")
// flags:      Execution flags (usually SignalRunLast)
// returnType: The return type of the signal handler (usually glib.TypeNone)
// paramTypes: The types of the parameters passed to the signal handler
func SignalNew(gtype glib.Type, name string, flags SignalFlags, returnType glib.Type, paramTypes ...glib.Type) uint {
	// Convert Go slice of glib.Type to C array of GType
	var cParamTypes *C.GType
	if len(paramTypes) > 0 {
		cParams := make([]C.GType, len(paramTypes))
		for i, t := range paramTypes {
			cParams[i] = C.GType(t)
		}
		cParamTypes = (*C.GType)(unsafe.Pointer(&cParams[0]))
	}

	signalID := C.go_gst_signal_newv(
		C.GType(gtype),
		C.CString(name),
		C.GSignalFlags(flags),
		C.GType(returnType),
		C.guint(uint(len(paramTypes))),
		cParamTypes,
	)

	return uint(signalID)
}
