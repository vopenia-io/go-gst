package gstsdp

// #include "gst.go.h"
import "C"

type Origin struct {
	ptr *C.GstSDPOrigin
}

func (o *Origin) Username() string {
	return C.GoString(o.ptr.username)
}

func (o *Origin) SessID() string {
	return C.GoString(o.ptr.sess_id)
}

func (o *Origin) SessVersion() string {
	return C.GoString(o.ptr.sess_version)
}

func (o *Origin) Nettype() string {
	return C.GoString(o.ptr.nettype)
}

func (o *Origin) Addrtype() string {
	return C.GoString(o.ptr.addrtype)
}

func (o *Origin) Addr() string {
	return C.GoString(o.ptr.addr)
}
