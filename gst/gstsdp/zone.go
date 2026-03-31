package gstsdp

// #include "gst.go.h"
import "C"

type Zone struct {
	ptr *C.GstSDPZone
}

func (z *Zone) Time() string {
	return C.GoString(z.ptr.time)
}

func (z *Zone) TypedTime() string {
	return C.GoString(z.ptr.typed_time)
}
