package gstsdp

// #include "gst.go.h"
import "C"

type Attribute struct {
	ptr *C.GstSDPAttribute
}

func (a *Attribute) Key() string {
	return C.GoString(a.ptr.key)
}

func (a *Attribute) Value() string {
	return C.GoString(a.ptr.value)
}
