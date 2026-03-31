package gstsdp

// #include "gst.go.h"
import "C"

type Key struct {
	ptr *C.GstSDPKey
}

func (k *Key) Type() string {
	return C.GoString(k.ptr._type)
}

func (k *Key) Data() string {
	return C.GoString(k.ptr.data)
}
