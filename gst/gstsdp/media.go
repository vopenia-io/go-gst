package gstsdp

// #include "gst.go.h"
import "C"
import (
	"errors"
	"iter"
	"runtime"
	"unsafe"

	"github.com/go-gst/go-gst/gst"
)

type Media struct {
	ptr *C.GstSDPMedia
}

// NewMedia allocates a new empty GstSDPMedia.
func NewMedia() (*Media, error) {
	var media *C.GstSDPMedia
	res := SDPResult(C.gst_sdp_media_new(&media))
	if res != SDPResultOk || media == nil {
		return nil, errors.New("failed to create new SDP media")
	}
	m := &Media{ptr: media}
	runtime.SetFinalizer(m, func(m *Media) {
		m.Free()
	})
	return m, nil
}

// FromCaps sets the media information from caps.
func MediaSetFromCaps(caps *gst.Caps, media *Media) SDPResult {
	return SDPResult(C.gst_sdp_media_set_media_from_caps((*C.GstCaps)(caps.Unsafe()), media.ptr))
}

// FromCaps sets the media information from caps.
func MediaAddMediaFromStructure(structure *gst.Structure, media *Media) SDPResult {
	return SDPResult(C.gst_sdp_media_add_media_from_structure((*C.GstStructure)(structure.Unsafe()), media.ptr))
}

// Free frees the SDP media. This is called automatically when the object is garbage collected
// for media created with NewMedia or Copy.
func (m *Media) Free() {
	if m == nil || m.ptr == nil {
		return
	}
	C.gst_sdp_media_free(m.ptr)
	m.ptr = nil
}

// Copy creates a new copy of the media.
func (m *Media) Copy() (*Media, error) {
	var cp *C.GstSDPMedia
	res := SDPResult(C.gst_sdp_media_copy(m.ptr, &cp))
	if res != SDPResultOk || cp == nil {
		return nil, errors.New("failed to copy SDP media")
	}
	mc := &Media{ptr: cp}
	runtime.SetFinalizer(mc, func(mc *Media) {
		mc.Free()
	})
	return mc, nil
}

// AsText converts the contents of media to a text string.
func (m *Media) AsText() string {
	cstr := C.gst_sdp_media_as_text(m.ptr)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetMedia returns the media description (e.g. "audio", "video").
func (m *Media) GetMedia() string {
	cstr := C.gst_sdp_media_get_media(m.ptr)
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetPort returns the port number for media.
func (m *Media) GetPort() uint {
	return uint(C.gst_sdp_media_get_port(m.ptr))
}

// GetNumPorts returns the number of ports for media.
func (m *Media) GetNumPorts() uint {
	return uint(C.gst_sdp_media_get_num_ports(m.ptr))
}

// GetProto returns the transport protocol of media.
func (m *Media) GetProto() string {
	cstr := C.gst_sdp_media_get_proto(m.ptr)
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetInformation returns the information of media.
func (m *Media) GetInformation() string {
	cstr := C.gst_sdp_media_get_information(m.ptr)
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetKey returns the encryption information from media.
func (m *Media) GetKey() *Key {
	k := C.gst_sdp_media_get_key(m.ptr)
	if k == nil {
		return nil
	}
	key := &Key{ptr: k}
	runtime.SetFinalizer(key, func(_ *Key) {
		runtime.KeepAlive(m)
	})
	return key
}

// FormatsLen returns the number of formats in media.
func (m *Media) FormatsLen() int {
	return int(C.gst_sdp_media_formats_len(m.ptr))
}

// Format returns the format at index idx.
func (m *Media) Format(idx int) string {
	cstr := C.gst_sdp_media_get_format(m.ptr, C.guint(idx))
	return C.GoString(cstr)
}

// Formats returns an iterator over all formats in the media.
func (m *Media) Formats() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i := 0; i < m.FormatsLen(); i++ {
			if !yield(i, m.Format(i)) {
				return
			}
		}
	}
}

// AttributesLen returns the number of attributes in media.
func (m *Media) AttributesLen() int {
	return int(C.gst_sdp_media_attributes_len(m.ptr))
}

func (m *Media) HasAttribute(key string) bool {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	return C.gst_sdp_media_get_attribute_val(m.ptr, ckey) != nil
}

// GetAttribute returns the attribute at index idx.
func (m *Media) GetAttribute(idx int) *Attribute {
	attr := C.gst_sdp_media_get_attribute(m.ptr, C.guint(idx))
	if attr == nil {
		return nil
	}
	a := &Attribute{ptr: attr}
	runtime.SetFinalizer(a, func(_ *Attribute) {
		runtime.KeepAlive(m)
	})
	return a
}

// GetAttributeVal returns the first attribute value for key.
func (m *Media) GetAttributeVal(key string) string {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cstr := C.gst_sdp_media_get_attribute_val(m.ptr, ckey)
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetAttributeValN returns the nth attribute value for key.
func (m *Media) GetAttributeValN(key string, nth int) string {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cstr := C.gst_sdp_media_get_attribute_val_n(m.ptr, ckey, C.guint(nth))
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// Attributes returns an iterator over all attributes in the media.
func (m *Media) Attributes() iter.Seq2[int, *Attribute] {
	return func(yield func(int, *Attribute) bool) {
		for i := 0; i < m.AttributesLen(); i++ {
			if !yield(i, m.GetAttribute(i)) {
				return
			}
		}
	}
}

// BandwidthsLen returns the number of bandwidth fields in media.
func (m *Media) BandwidthsLen() int {
	return int(C.gst_sdp_media_bandwidths_len(m.ptr))
}

// GetBandwidth returns the bandwidth at index idx.
func (m *Media) GetBandwidth(idx int) *Bandwidth {
	bw := C.gst_sdp_media_get_bandwidth(m.ptr, C.guint(idx))
	if bw == nil {
		return nil
	}
	b := &Bandwidth{ptr: bw}
	runtime.SetFinalizer(b, func(_ *Bandwidth) {
		runtime.KeepAlive(m)
	})
	return b
}

// Bandwidths returns an iterator over all bandwidths in the media.
func (m *Media) Bandwidths() iter.Seq2[int, *Bandwidth] {
	return func(yield func(int, *Bandwidth) bool) {
		for i := 0; i < m.BandwidthsLen(); i++ {
			if !yield(i, m.GetBandwidth(i)) {
				return
			}
		}
	}
}

// ConnectionsLen returns the number of connection fields in media.
func (m *Media) ConnectionsLen() int {
	return int(C.gst_sdp_media_connections_len(m.ptr))
}

// GetConnection returns the connection at index idx.
func (m *Media) GetConnection(idx int) *Connection {
	conn := C.gst_sdp_media_get_connection(m.ptr, C.guint(idx))
	if conn == nil {
		return nil
	}
	c := &Connection{ptr: conn}
	runtime.SetFinalizer(c, func(_ *Connection) {
		runtime.KeepAlive(m)
	})
	return c
}

// Connections returns an iterator over all connections in the media.
func (m *Media) Connections() iter.Seq2[int, *Connection] {
	return func(yield func(int, *Connection) bool) {
		for i := 0; i < m.ConnectionsLen(); i++ {
			if !yield(i, m.GetConnection(i)) {
				return
			}
		}
	}
}

var ErrCouldNotGetCaps = errors.New("could not get caps")

// GetCaps returns the GstCaps for the given payload type.
func (m *Media) GetCaps(pt int) (*gst.Caps, error) {
	ccaps := C.gst_sdp_media_get_caps_from_media(m.ptr, C.gint(pt))
	if ccaps == nil {
		return nil, ErrCouldNotGetCaps
	}
	return gst.FromGstCapsUnsafeFull(unsafe.Pointer(ccaps)), nil
}

// AttributesToCaps maps attributes of media to caps.
func (m *Media) AttributesToCaps(caps *gst.Caps) SDPResult {
	return SDPResult(C.gst_sdp_media_attributes_to_caps(m.ptr, (*C.GstCaps)(caps.Unsafe())))
}

// SetMedia sets the media description of media.
func (m *Media) SetMedia(med string) SDPResult {
	cmed := C.CString(med)
	defer C.free(unsafe.Pointer(cmed))
	return SDPResult(C.gst_sdp_media_set_media(m.ptr, cmed))
}

// SetPortInfo sets the port information of media.
func (m *Media) SetPortInfo(port, numPorts uint) SDPResult {
	return SDPResult(C.gst_sdp_media_set_port_info(m.ptr, C.guint(port), C.guint(numPorts)))
}

// SetProto sets the media transport protocol.
func (m *Media) SetProto(proto string) SDPResult {
	cproto := C.CString(proto)
	defer C.free(unsafe.Pointer(cproto))
	return SDPResult(C.gst_sdp_media_set_proto(m.ptr, cproto))
}

// SetInformation sets the media information.
func (m *Media) SetInformation(information string) SDPResult {
	cinfo := C.CString(information)
	defer C.free(unsafe.Pointer(cinfo))
	return SDPResult(C.gst_sdp_media_set_information(m.ptr, cinfo))
}

// SetKey sets the encryption information.
func (m *Media) SetKey(typ, data string) SDPResult {
	ctyp := C.CString(typ)
	defer C.free(unsafe.Pointer(ctyp))
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	return SDPResult(C.gst_sdp_media_set_key(m.ptr, ctyp, cdata))
}

// AddAttribute adds an attribute with key and value to media.
func (m *Media) AddAttribute(key, value string) SDPResult {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	return SDPResult(C.gst_sdp_media_add_attribute(m.ptr, ckey, cval))
}

// AddBandwidth adds bandwidth information to media.
func (m *Media) AddBandwidth(bwtype string, bandwidth uint) SDPResult {
	cbwtype := C.CString(bwtype)
	defer C.free(unsafe.Pointer(cbwtype))
	return SDPResult(C.gst_sdp_media_add_bandwidth(m.ptr, cbwtype, C.guint(bandwidth)))
}

// AddConnection adds connection parameters to media.
func (m *Media) AddConnection(nettype, addrtype, address string, ttl, addrNumber uint) SDPResult {
	cnettype := C.CString(nettype)
	defer C.free(unsafe.Pointer(cnettype))
	caddrtype := C.CString(addrtype)
	defer C.free(unsafe.Pointer(caddrtype))
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return SDPResult(C.gst_sdp_media_add_connection(m.ptr, cnettype, caddrtype, caddress, C.guint(ttl), C.guint(addrNumber)))
}

// AddFormat adds format information to media.
func (m *Media) AddFormat(format string) SDPResult {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))
	return SDPResult(C.gst_sdp_media_add_format(m.ptr, cfmt))
}

// RemoveAttribute removes the attribute at index idx.
func (m *Media) RemoveAttribute(idx int) SDPResult {
	return SDPResult(C.gst_sdp_media_remove_attribute(m.ptr, C.guint(idx)))
}

// RemoveBandwidth removes the bandwidth at index idx.
func (m *Media) RemoveBandwidth(idx int) SDPResult {
	return SDPResult(C.gst_sdp_media_remove_bandwidth(m.ptr, C.guint(idx)))
}

// RemoveConnection removes the connection at index idx.
func (m *Media) RemoveConnection(idx int) SDPResult {
	return SDPResult(C.gst_sdp_media_remove_connection(m.ptr, C.guint(idx)))
}

// RemoveFormat removes the format at index idx.
func (m *Media) RemoveFormat(idx int) SDPResult {
	return SDPResult(C.gst_sdp_media_remove_format(m.ptr, C.guint(idx)))
}

// InsertAttribute inserts an attribute at index idx. When idx is -1, the attribute is appended.
func (m *Media) InsertAttribute(idx int, key, value string) SDPResult {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	var attr C.GstSDPAttribute
	C.gst_sdp_attribute_set(&attr, ckey, cval)
	return SDPResult(C.gst_sdp_media_insert_attribute(m.ptr, C.gint(idx), &attr))
}

// InsertBandwidth inserts bandwidth at index idx. When idx is -1, the bandwidth is appended.
func (m *Media) InsertBandwidth(idx int, bwtype string, bandwidth uint) SDPResult {
	cbwtype := C.CString(bwtype)
	defer C.free(unsafe.Pointer(cbwtype))
	var bw C.GstSDPBandwidth
	C.gst_sdp_bandwidth_set(&bw, cbwtype, C.guint(bandwidth))
	return SDPResult(C.gst_sdp_media_insert_bandwidth(m.ptr, C.gint(idx), &bw))
}

// InsertConnection inserts connection at index idx. When idx is -1, the connection is appended.
func (m *Media) InsertConnection(idx int, nettype, addrtype, address string, ttl, addrNumber uint) SDPResult {
	cnettype := C.CString(nettype)
	defer C.free(unsafe.Pointer(cnettype))
	caddrtype := C.CString(addrtype)
	defer C.free(unsafe.Pointer(caddrtype))
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	var conn C.GstSDPConnection
	C.gst_sdp_connection_set(&conn, cnettype, caddrtype, caddress, C.guint(ttl), C.guint(addrNumber))
	return SDPResult(C.gst_sdp_media_insert_connection(m.ptr, C.gint(idx), &conn))
}

// InsertFormat inserts format at index idx. When idx is -1, the format is appended.
func (m *Media) InsertFormat(idx int, format string) SDPResult {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))
	return SDPResult(C.gst_sdp_media_insert_format(m.ptr, C.gint(idx), cfmt))
}

// ReplaceAttribute replaces the attribute at index idx.
func (m *Media) ReplaceAttribute(idx int, key, value string) SDPResult {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	var attr C.GstSDPAttribute
	C.gst_sdp_attribute_set(&attr, ckey, cval)
	return SDPResult(C.gst_sdp_media_replace_attribute(m.ptr, C.guint(idx), &attr))
}

// ReplaceBandwidth replaces the bandwidth at index idx.
func (m *Media) ReplaceBandwidth(idx int, bwtype string, bandwidth uint) SDPResult {
	cbwtype := C.CString(bwtype)
	defer C.free(unsafe.Pointer(cbwtype))
	var bw C.GstSDPBandwidth
	C.gst_sdp_bandwidth_set(&bw, cbwtype, C.guint(bandwidth))
	return SDPResult(C.gst_sdp_media_replace_bandwidth(m.ptr, C.guint(idx), &bw))
}

// ReplaceConnection replaces the connection at index idx.
func (m *Media) ReplaceConnection(idx int, nettype, addrtype, address string, ttl, addrNumber uint) SDPResult {
	cnettype := C.CString(nettype)
	defer C.free(unsafe.Pointer(cnettype))
	caddrtype := C.CString(addrtype)
	defer C.free(unsafe.Pointer(caddrtype))
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	var conn C.GstSDPConnection
	C.gst_sdp_connection_set(&conn, cnettype, caddrtype, caddress, C.guint(ttl), C.guint(addrNumber))
	return SDPResult(C.gst_sdp_media_replace_connection(m.ptr, C.guint(idx), &conn))
}

// ReplaceFormat replaces the format at index idx.
func (m *Media) ReplaceFormat(idx int, format string) SDPResult {
	cfmt := C.CString(format)
	defer C.free(unsafe.Pointer(cfmt))
	return SDPResult(C.gst_sdp_media_replace_format(m.ptr, C.guint(idx), cfmt))
}
