package gstsdp

// #include "gst.go.h"
import "C"

type Connection struct {
	ptr *C.GstSDPConnection
}

func (c *Connection) Nettype() string {
	return C.GoString(c.ptr.nettype)
}

func (c *Connection) Addrtype() string {
	return C.GoString(c.ptr.addrtype)
}

func (c *Connection) Address() string {
	return C.GoString(c.ptr.address)
}

func (c *Connection) TTL() uint {
	return uint(c.ptr.ttl)
}

func (c *Connection) AddrNumber() uint {
	return uint(c.ptr.addr_number)
}
