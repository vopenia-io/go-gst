package gstsdp

// #include "gst.go.h"
import "C"

type Bandwidth struct {
	ptr *C.GstSDPBandwidth
}

func (b *Bandwidth) BWType() string {
	return C.GoString(b.ptr.bwtype)
}

func (b *Bandwidth) Value() uint {
	return uint(b.ptr.bandwidth)
}
