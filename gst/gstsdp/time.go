package gstsdp

// #include "gst.go.h"
import "C"

type Time struct {
	ptr *C.GstSDPTime
}

func (t *Time) Start() string {
	return C.GoString(t.ptr.start)
}

func (t *Time) Stop() string {
	return C.GoString(t.ptr.stop)
}
