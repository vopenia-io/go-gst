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

type SDPResult C.GstSDPResult

const (
	SDPResultOk SDPResult = C.GST_SDP_OK
	SDPEinval   SDPResult = C.GST_SDP_EINVAL
)

type Message struct {
	ptr *C.GstSDPMessage
}

func wrapSDPMessageAndFinalize(sdp *C.GstSDPMessage) *Message {
	msg := &Message{
		ptr: sdp,
	}

	// this requires that we copy the SDP message before passing it to any transfer-ownership function
	runtime.SetFinalizer(msg, func(msg *Message) {
		msg.Free()
	})

	return msg
}

// NewMessage creates a new empty SDP message
func NewMessage() (*Message, error) {
	var msg *C.GstSDPMessage

	res := SDPResult(C.gst_sdp_message_new(&msg))

	if res != SDPResultOk || msg == nil {
		return nil, ErrSDPInvalid
	}

	return wrapSDPMessageAndFinalize(msg), nil
}

// NewMessageFromUnsafe creates a new SDP message from a pointer and does not finalize it
func NewMessageFromUnsafe(ptr unsafe.Pointer) *Message {
	return &Message{
		ptr: (*C.GstSDPMessage)(ptr),
	}
}

var ErrSDPInvalid = errors.New("invalid SDP")

func ParseSDPMessage(sdp string) (*Message, error) {
	cstr := C.CString(sdp)
	defer C.free(unsafe.Pointer(cstr))

	var msg *C.GstSDPMessage

	res := SDPResult(C.gst_sdp_message_new_from_text(cstr, &msg))

	if res != SDPResultOk || msg == nil {
		return nil, ErrSDPInvalid
	}

	return wrapSDPMessageAndFinalize(msg), nil
}

// ParseSDPBuffer parses the data and stores the result in a new message.
func ParseSDPBuffer(data []byte) (*Message, error) {
	msg, err := NewMessage()
	if err != nil {
		return nil, err
	}

	res := SDPResult(C.gst_sdp_message_parse_buffer((*C.guint8)(unsafe.Pointer(&data[0])), C.guint(len(data)), msg.ptr))
	if res != SDPResultOk {
		msg.Free()
		return nil, ErrSDPInvalid
	}

	return msg, nil
}

// ParseSDPURI parses the given URI and stores the result in a new message.
func ParseSDPURI(uri string) (*Message, error) {
	msg, err := NewMessage()
	if err != nil {
		return nil, err
	}

	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))

	res := SDPResult(C.gst_sdp_message_parse_uri(curi, msg.ptr))
	if res != SDPResultOk {
		msg.Free()
		return nil, ErrSDPInvalid
	}

	return msg, nil
}

// AddressIsMulticast checks if the given address is a multicast address.
func AddressIsMulticast(nettype, addrtype, addr string) bool {
	cnettype := C.CString(nettype)
	defer C.free(unsafe.Pointer(cnettype))
	caddrtype := C.CString(addrtype)
	defer C.free(unsafe.Pointer(caddrtype))
	caddr := C.CString(addr)
	defer C.free(unsafe.Pointer(caddr))
	return C.gst_sdp_address_is_multicast(cnettype, caddrtype, caddr) != 0
}

func (msg *Message) String() string {
	return msg.AsText()
}

func (msg *Message) AsText() string {
	if msg == nil {
		return ""
	}
	cstr := C.gst_sdp_message_as_text(msg.native())
	defer C.free(unsafe.Pointer(cstr))

	return C.GoString(cstr)
}

// AsURI creates a URI from msg with the given scheme.
func (msg *Message) AsURI(scheme string) string {
	cscheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(cscheme))
	cstr := C.gst_sdp_message_as_uri(cscheme, msg.native())
	if cstr == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// UnownedCopy creates a new copy of the SDP message that will not be finalized
//
// this is needed to pass the message back to C where C takes ownership of the message
//
// the returned SDP message will leak memory if not freed manually
func (msg *Message) UnownedCopy() *Message {
	if msg == nil {
		return nil
	}

	var newMsg *C.GstSDPMessage
	res := C.gst_sdp_message_copy(msg.native(), &newMsg)

	if res != C.GST_SDP_OK || newMsg == nil {
		return nil
	}

	return &Message{
		ptr: newMsg,
	}
}

// Copy creates a new copy of the SDP message with a finalizer.
func (msg *Message) Copy() (*Message, error) {
	if msg == nil {
		return nil, ErrSDPInvalid
	}

	var newMsg *C.GstSDPMessage
	res := SDPResult(C.gst_sdp_message_copy(msg.native(), &newMsg))

	if res != SDPResultOk || newMsg == nil {
		return nil, ErrSDPInvalid
	}

	return wrapSDPMessageAndFinalize(newMsg), nil
}

// Free frees the SDP message.
//
// This is called automatically when the object is garbage collected.
func (msg *Message) Free() {
	if msg == nil || msg.ptr == nil {
		return
	}
	C.gst_sdp_message_free(msg.ptr)
	msg.ptr = nil
}

func (msg *Message) native() *C.GstSDPMessage {
	if msg == nil {
		return nil
	}

	return msg.ptr
}

func (msg *Message) Instance() unsafe.Pointer {
	if msg == nil {
		return nil
	}
	return unsafe.Pointer(msg.native())
}

// GetVersion returns the version in msg.
func (msg *Message) GetVersion() string {
	cstr := C.gst_sdp_message_get_version(msg.native())
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetOrigin returns the origin of msg.
func (msg *Message) GetOrigin() *Origin {
	o := C.gst_sdp_message_get_origin(msg.native())
	if o == nil {
		return nil
	}
	origin := &Origin{ptr: o}
	runtime.SetFinalizer(origin, func(_ *Origin) {
		runtime.KeepAlive(msg)
	})
	return origin
}

// GetSessionName returns the session name in msg.
func (msg *Message) GetSessionName() string {
	cstr := C.gst_sdp_message_get_session_name(msg.native())
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetInformation returns the information in msg.
func (msg *Message) GetInformation() string {
	cstr := C.gst_sdp_message_get_information(msg.native())
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetURI returns the URI in msg.
func (msg *Message) GetURI() string {
	cstr := C.gst_sdp_message_get_uri(msg.native())
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetConnection returns the connection of msg.
func (msg *Message) GetConnection() *Connection {
	c := C.gst_sdp_message_get_connection(msg.native())
	if c == nil {
		return nil
	}
	conn := &Connection{ptr: c}
	runtime.SetFinalizer(conn, func(_ *Connection) {
		runtime.KeepAlive(msg)
	})
	return conn
}

// GetKey returns the encryption information from msg.
func (msg *Message) GetKey() *Key {
	k := C.gst_sdp_message_get_key(msg.native())
	if k == nil {
		return nil
	}
	key := &Key{ptr: k}
	runtime.SetFinalizer(key, func(_ *Key) {
		runtime.KeepAlive(msg)
	})
	return key
}

// MediasLen returns the number of media descriptions in msg.
func (msg *Message) MediasLen() int {
	return int(C.gst_sdp_message_medias_len(msg.native()))
}

// Media returns the media description at index i.
func (msg *Message) Media(i int) *Media {
	cmedia := C.gst_sdp_message_get_media(msg.native(), C.uint(i))

	if cmedia == nil {
		return nil
	}

	media := &Media{
		ptr: cmedia,
	}

	// keep the Message alive while we are handling the media
	runtime.SetFinalizer(media, func(_ *Media) {
		runtime.KeepAlive(msg)
	})

	return media
}

// Medias returns an iterator over all media descriptions in msg.
func (msg *Message) Medias() iter.Seq2[int, *Media] {
	return func(yield func(int, *Media) bool) {
		for i := 0; i < msg.MediasLen(); i++ {
			if !yield(i, msg.Media(i)) {
				return
			}
		}
	}
}

// AttributesLen returns the number of attributes in msg.
func (msg *Message) AttributesLen() int {
	return int(C.gst_sdp_message_attributes_len(msg.native()))
}

// GetAttribute returns the attribute at index idx.
func (msg *Message) GetAttribute(idx int) *Attribute {
	attr := C.gst_sdp_message_get_attribute(msg.native(), C.guint(idx))
	if attr == nil {
		return nil
	}
	a := &Attribute{ptr: attr}
	runtime.SetFinalizer(a, func(_ *Attribute) {
		runtime.KeepAlive(msg)
	})
	return a
}

// GetAttributeVal returns the first attribute value for key.
func (msg *Message) GetAttributeVal(key string) string {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cstr := C.gst_sdp_message_get_attribute_val(msg.native(), ckey)
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// GetAttributeValN returns the nth attribute value for key.
func (msg *Message) GetAttributeValN(key string, nth int) string {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cstr := C.gst_sdp_message_get_attribute_val_n(msg.native(), ckey, C.guint(nth))
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// Attributes returns an iterator over all attributes in msg.
func (msg *Message) Attributes() iter.Seq2[int, *Attribute] {
	return func(yield func(int, *Attribute) bool) {
		for i := 0; i < msg.AttributesLen(); i++ {
			if !yield(i, msg.GetAttribute(i)) {
				return
			}
		}
	}
}

// BandwidthsLen returns the number of bandwidth information in msg.
func (msg *Message) BandwidthsLen() int {
	return int(C.gst_sdp_message_bandwidths_len(msg.native()))
}

// GetBandwidth returns the bandwidth at index idx.
func (msg *Message) GetBandwidth(idx int) *Bandwidth {
	bw := C.gst_sdp_message_get_bandwidth(msg.native(), C.guint(idx))
	if bw == nil {
		return nil
	}
	b := &Bandwidth{ptr: bw}
	runtime.SetFinalizer(b, func(_ *Bandwidth) {
		runtime.KeepAlive(msg)
	})
	return b
}

// Bandwidths returns an iterator over all bandwidths in msg.
func (msg *Message) Bandwidths() iter.Seq2[int, *Bandwidth] {
	return func(yield func(int, *Bandwidth) bool) {
		for i := 0; i < msg.BandwidthsLen(); i++ {
			if !yield(i, msg.GetBandwidth(i)) {
				return
			}
		}
	}
}

// EmailsLen returns the number of emails in msg.
func (msg *Message) EmailsLen() int {
	return int(C.gst_sdp_message_emails_len(msg.native()))
}

// GetEmail returns the email at index idx.
func (msg *Message) GetEmail(idx int) string {
	cstr := C.gst_sdp_message_get_email(msg.native(), C.guint(idx))
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// Emails returns an iterator over all emails in msg.
func (msg *Message) Emails() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i := 0; i < msg.EmailsLen(); i++ {
			if !yield(i, msg.GetEmail(i)) {
				return
			}
		}
	}
}

// PhonesLen returns the number of phones in msg.
func (msg *Message) PhonesLen() int {
	return int(C.gst_sdp_message_phones_len(msg.native()))
}

// GetPhone returns the phone at index idx.
func (msg *Message) GetPhone(idx int) string {
	cstr := C.gst_sdp_message_get_phone(msg.native(), C.guint(idx))
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

// Phones returns an iterator over all phones in msg.
func (msg *Message) Phones() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i := 0; i < msg.PhonesLen(); i++ {
			if !yield(i, msg.GetPhone(i)) {
				return
			}
		}
	}
}

// TimesLen returns the number of time information entries in msg.
func (msg *Message) TimesLen() int {
	return int(C.gst_sdp_message_times_len(msg.native()))
}

// GetTime returns the time information at index idx.
func (msg *Message) GetTime(idx int) *Time {
	t := C.gst_sdp_message_get_time(msg.native(), C.guint(idx))
	if t == nil {
		return nil
	}
	tm := &Time{ptr: t}
	runtime.SetFinalizer(tm, func(_ *Time) {
		runtime.KeepAlive(msg)
	})
	return tm
}

// Times returns an iterator over all time entries in msg.
func (msg *Message) Times() iter.Seq2[int, *Time] {
	return func(yield func(int, *Time) bool) {
		for i := 0; i < msg.TimesLen(); i++ {
			if !yield(i, msg.GetTime(i)) {
				return
			}
		}
	}
}

// ZonesLen returns the number of time zone entries in msg.
func (msg *Message) ZonesLen() int {
	return int(C.gst_sdp_message_zones_len(msg.native()))
}

// GetZone returns the time zone information at index idx.
func (msg *Message) GetZone(idx int) *Zone {
	z := C.gst_sdp_message_get_zone(msg.native(), C.guint(idx))
	if z == nil {
		return nil
	}
	zone := &Zone{ptr: z}
	runtime.SetFinalizer(zone, func(_ *Zone) {
		runtime.KeepAlive(msg)
	})
	return zone
}

// Zones returns an iterator over all time zone entries in msg.
func (msg *Message) Zones() iter.Seq2[int, *Zone] {
	return func(yield func(int, *Zone) bool) {
		for i := 0; i < msg.ZonesLen(); i++ {
			if !yield(i, msg.GetZone(i)) {
				return
			}
		}
	}
}

// AttributesToCaps maps attributes of msg to caps.
func (msg *Message) AttributesToCaps(caps *gst.Caps) SDPResult {
	return SDPResult(C.gst_sdp_message_attributes_to_caps(msg.native(), (*C.GstCaps)(caps.Unsafe())))
}

// SetVersion sets the version in msg.
func (msg *Message) SetVersion(version string) SDPResult {
	cver := C.CString(version)
	defer C.free(unsafe.Pointer(cver))
	return SDPResult(C.gst_sdp_message_set_version(msg.native(), cver))
}

// SetOrigin configures the SDP origin in msg.
func (msg *Message) SetOrigin(username, sessID, sessVersion, nettype, addrtype, addr string) SDPResult {
	cusername := C.CString(username)
	defer C.free(unsafe.Pointer(cusername))
	csessID := C.CString(sessID)
	defer C.free(unsafe.Pointer(csessID))
	csessVersion := C.CString(sessVersion)
	defer C.free(unsafe.Pointer(csessVersion))
	cnettype := C.CString(nettype)
	defer C.free(unsafe.Pointer(cnettype))
	caddrtype := C.CString(addrtype)
	defer C.free(unsafe.Pointer(caddrtype))
	caddr := C.CString(addr)
	defer C.free(unsafe.Pointer(caddr))
	return SDPResult(C.gst_sdp_message_set_origin(msg.native(), cusername, csessID, csessVersion, cnettype, caddrtype, caddr))
}

// SetSessionName sets the session name in msg.
func (msg *Message) SetSessionName(name string) SDPResult {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return SDPResult(C.gst_sdp_message_set_session_name(msg.native(), cname))
}

// SetInformation sets the information in msg.
func (msg *Message) SetInformation(information string) SDPResult {
	cinfo := C.CString(information)
	defer C.free(unsafe.Pointer(cinfo))
	return SDPResult(C.gst_sdp_message_set_information(msg.native(), cinfo))
}

// SetURI sets the URI in msg.
func (msg *Message) SetURI(uri string) SDPResult {
	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))
	return SDPResult(C.gst_sdp_message_set_uri(msg.native(), curi))
}

// SetConnection configures the SDP connection in msg.
func (msg *Message) SetConnection(nettype, addrtype, address string, ttl, addrNumber uint) SDPResult {
	cnettype := C.CString(nettype)
	defer C.free(unsafe.Pointer(cnettype))
	caddrtype := C.CString(addrtype)
	defer C.free(unsafe.Pointer(caddrtype))
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	return SDPResult(C.gst_sdp_message_set_connection(msg.native(), cnettype, caddrtype, caddress, C.guint(ttl), C.guint(addrNumber)))
}

// SetKey sets the encryption information in msg.
func (msg *Message) SetKey(typ, data string) SDPResult {
	ctyp := C.CString(typ)
	defer C.free(unsafe.Pointer(ctyp))
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	return SDPResult(C.gst_sdp_message_set_key(msg.native(), ctyp, cdata))
}

// AddAttribute adds an attribute with key and value to msg.
func (msg *Message) AddAttribute(key, value string) SDPResult {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	return SDPResult(C.gst_sdp_message_add_attribute(msg.native(), ckey, cval))
}

// AddBandwidth adds bandwidth information to msg.
func (msg *Message) AddBandwidth(bwtype string, bandwidth uint) SDPResult {
	cbwtype := C.CString(bwtype)
	defer C.free(unsafe.Pointer(cbwtype))
	return SDPResult(C.gst_sdp_message_add_bandwidth(msg.native(), cbwtype, C.guint(bandwidth)))
}

// AddEmail adds email to the list of emails in msg.
func (msg *Message) AddEmail(email string) SDPResult {
	cemail := C.CString(email)
	defer C.free(unsafe.Pointer(cemail))
	return SDPResult(C.gst_sdp_message_add_email(msg.native(), cemail))
}

// AddPhone adds phone to the list of phones in msg.
func (msg *Message) AddPhone(phone string) SDPResult {
	cphone := C.CString(phone)
	defer C.free(unsafe.Pointer(cphone))
	return SDPResult(C.gst_sdp_message_add_phone(msg.native(), cphone))
}

// AddTime adds time information start and stop to msg.
func (msg *Message) AddTime(start, stop string, repeat []string) SDPResult {
	cstart := C.CString(start)
	defer C.free(unsafe.Pointer(cstart))
	cstop := C.CString(stop)
	defer C.free(unsafe.Pointer(cstop))

	var crepeat **C.gchar
	if len(repeat) > 0 {
		arr := make([]*C.gchar, len(repeat)+1)
		for i, r := range repeat {
			arr[i] = C.CString(r)
			defer C.free(unsafe.Pointer(arr[i]))
		}
		arr[len(repeat)] = nil
		crepeat = &arr[0]
	}

	return SDPResult(C.gst_sdp_message_add_time(msg.native(), cstart, cstop, crepeat))
}

// AddZone adds time zone information to msg.
func (msg *Message) AddZone(adjTime, typedTime string) SDPResult {
	cadjTime := C.CString(adjTime)
	defer C.free(unsafe.Pointer(cadjTime))
	ctypedTime := C.CString(typedTime)
	defer C.free(unsafe.Pointer(ctypedTime))
	return SDPResult(C.gst_sdp_message_add_zone(msg.native(), cadjTime, ctypedTime))
}

// AddMedia adds media to the array of medias in msg.
// This function takes ownership of the contents of media.
func (msg *Message) AddMedia(media *Media) SDPResult {
	return SDPResult(C.gst_sdp_message_add_media(msg.native(), media.ptr))
}

// AddMediaCopy adds a copy of media to the array of medias in msg.
func (msg *Message) AddMediaCopy(media *Media) SDPResult {
	mediaCopy, err := media.Copy()
	if err != nil {
		return SDPResult(C.GST_SDP_EINVAL)
	}
	return msg.AddMedia(mediaCopy)
}

// RemoveAttribute removes the attribute at index idx.
func (msg *Message) RemoveAttribute(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_attribute(msg.native(), C.guint(idx)))
}

// RemoveBandwidth removes the bandwidth at index idx.
func (msg *Message) RemoveBandwidth(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_bandwidth(msg.native(), C.guint(idx)))
}

// RemoveEmail removes the email at index idx.
func (msg *Message) RemoveEmail(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_email(msg.native(), C.guint(idx)))
}

// RemovePhone removes the phone at index idx.
func (msg *Message) RemovePhone(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_phone(msg.native(), C.guint(idx)))
}

// RemoveTime removes the time information at index idx.
func (msg *Message) RemoveTime(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_time(msg.native(), C.guint(idx)))
}

// RemoveZone removes the zone information at index idx.
func (msg *Message) RemoveZone(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_zone(msg.native(), C.guint(idx)))
}

// RemoveMedia removes the media at index idx.
func (msg *Message) RemoveMedia(idx int) SDPResult {
	return SDPResult(C.gst_sdp_message_remove_media(msg.native(), C.guint(idx)))
}

// InsertAttribute inserts an attribute at index idx. When idx is -1, it is appended.
func (msg *Message) InsertAttribute(idx int, key, value string) SDPResult {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	var attr C.GstSDPAttribute
	C.gst_sdp_attribute_set(&attr, ckey, cval)
	return SDPResult(C.gst_sdp_message_insert_attribute(msg.native(), C.gint(idx), &attr))
}

// InsertBandwidth inserts bandwidth at index idx. When idx is -1, it is appended.
func (msg *Message) InsertBandwidth(idx int, bwtype string, bandwidth uint) SDPResult {
	cbwtype := C.CString(bwtype)
	defer C.free(unsafe.Pointer(cbwtype))
	var bw C.GstSDPBandwidth
	C.gst_sdp_bandwidth_set(&bw, cbwtype, C.guint(bandwidth))
	return SDPResult(C.gst_sdp_message_insert_bandwidth(msg.native(), C.gint(idx), &bw))
}

// InsertEmail inserts email at index idx. When idx is -1, it is appended.
func (msg *Message) InsertEmail(idx int, email string) SDPResult {
	cemail := C.CString(email)
	defer C.free(unsafe.Pointer(cemail))
	return SDPResult(C.gst_sdp_message_insert_email(msg.native(), C.gint(idx), cemail))
}

// InsertPhone inserts phone at index idx. When idx is -1, it is appended.
func (msg *Message) InsertPhone(idx int, phone string) SDPResult {
	cphone := C.CString(phone)
	defer C.free(unsafe.Pointer(cphone))
	return SDPResult(C.gst_sdp_message_insert_phone(msg.native(), C.gint(idx), cphone))
}

// InsertTime inserts time at index idx. When idx is -1, it is appended.
func (msg *Message) InsertTime(idx int, start, stop string) SDPResult {
	cstart := C.CString(start)
	defer C.free(unsafe.Pointer(cstart))
	cstop := C.CString(stop)
	defer C.free(unsafe.Pointer(cstop))
	var t C.GstSDPTime
	C.gst_sdp_time_set(&t, cstart, cstop, nil)
	return SDPResult(C.gst_sdp_message_insert_time(msg.native(), C.gint(idx), &t))
}

// InsertZone inserts zone at index idx. When idx is -1, it is appended.
func (msg *Message) InsertZone(idx int, adjTime, typedTime string) SDPResult {
	cadjTime := C.CString(adjTime)
	defer C.free(unsafe.Pointer(cadjTime))
	ctypedTime := C.CString(typedTime)
	defer C.free(unsafe.Pointer(ctypedTime))
	var zone C.GstSDPZone
	C.gst_sdp_zone_set(&zone, cadjTime, ctypedTime)
	return SDPResult(C.gst_sdp_message_insert_zone(msg.native(), C.gint(idx), &zone))
}

// ReplaceAttribute replaces the attribute at index idx.
func (msg *Message) ReplaceAttribute(idx int, key, value string) SDPResult {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))
	var attr C.GstSDPAttribute
	C.gst_sdp_attribute_set(&attr, ckey, cval)
	return SDPResult(C.gst_sdp_message_replace_attribute(msg.native(), C.guint(idx), &attr))
}

// ReplaceBandwidth replaces the bandwidth at index idx.
func (msg *Message) ReplaceBandwidth(idx int, bwtype string, bandwidth uint) SDPResult {
	cbwtype := C.CString(bwtype)
	defer C.free(unsafe.Pointer(cbwtype))
	var bw C.GstSDPBandwidth
	C.gst_sdp_bandwidth_set(&bw, cbwtype, C.guint(bandwidth))
	return SDPResult(C.gst_sdp_message_replace_bandwidth(msg.native(), C.guint(idx), &bw))
}

// ReplaceEmail replaces the email at index idx.
func (msg *Message) ReplaceEmail(idx int, email string) SDPResult {
	cemail := C.CString(email)
	defer C.free(unsafe.Pointer(cemail))
	return SDPResult(C.gst_sdp_message_replace_email(msg.native(), C.guint(idx), cemail))
}

// ReplacePhone replaces the phone at index idx.
func (msg *Message) ReplacePhone(idx int, phone string) SDPResult {
	cphone := C.CString(phone)
	defer C.free(unsafe.Pointer(cphone))
	return SDPResult(C.gst_sdp_message_replace_phone(msg.native(), C.guint(idx), cphone))
}

// ReplaceTime replaces the time at index idx.
func (msg *Message) ReplaceTime(idx int, start, stop string) SDPResult {
	cstart := C.CString(start)
	defer C.free(unsafe.Pointer(cstart))
	cstop := C.CString(stop)
	defer C.free(unsafe.Pointer(cstop))
	var t C.GstSDPTime
	C.gst_sdp_time_set(&t, cstart, cstop, nil)
	return SDPResult(C.gst_sdp_message_replace_time(msg.native(), C.guint(idx), &t))
}

// ReplaceZone replaces the zone at index idx.
func (msg *Message) ReplaceZone(idx int, adjTime, typedTime string) SDPResult {
	cadjTime := C.CString(adjTime)
	defer C.free(unsafe.Pointer(cadjTime))
	ctypedTime := C.CString(typedTime)
	defer C.free(unsafe.Pointer(ctypedTime))
	var zone C.GstSDPZone
	C.gst_sdp_zone_set(&zone, cadjTime, ctypedTime)
	return SDPResult(C.gst_sdp_message_replace_zone(msg.native(), C.guint(idx), &zone))
}
