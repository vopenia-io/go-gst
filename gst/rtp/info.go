package rtp

import "unsafe"

// #include <gst/rtp/gstrtppayloads.h>
import "C"

type PayloadInfo struct {
	ptr *C.GstRTPPayloadInfo
}

func (p *PayloadInfo) PayloadType() uint8 {
	if p == nil || p.ptr == nil {
		return 0
	}
	return uint8(p.ptr.payload_type)
}

func (p *PayloadInfo) Media() string {
	if p == nil || p.ptr == nil {
		return ""
	}
	return C.GoString((*C.char)(unsafe.Pointer(p.ptr.media)))
}

func (p *PayloadInfo) EncodingName() string {
	if p == nil || p.ptr == nil {
		return ""
	}
	return C.GoString((*C.char)(unsafe.Pointer(p.ptr.encoding_name)))
}

func (p *PayloadInfo) ClockRate() uint {
	if p == nil || p.ptr == nil {
		return 0
	}
	return uint(p.ptr.clock_rate)
}

func (p *PayloadInfo) EncodingParameters() string {
	if p == nil || p.ptr == nil {
		return ""
	}
	return C.GoString((*C.char)(unsafe.Pointer(p.ptr.encoding_parameters)))
}

func (p *PayloadInfo) BitRate() uint {
	if p == nil || p.ptr == nil {
		return 0
	}
	return uint(p.ptr.bitrate)
}

func PayloadInfoForName(media string, encodingName string) *PayloadInfo {
	cMedia := C.CString(media)
	defer C.free(unsafe.Pointer(cMedia))

	cEncodingName := C.CString(encodingName)
	defer C.free(unsafe.Pointer(cEncodingName))

	ptr := C.gst_rtp_payload_info_for_name(cMedia, cEncodingName)
	if ptr == nil {
		return nil
	}
	return &PayloadInfo{ptr: ptr}
}

func PayloadInfoForPt(payloadType uint8) *PayloadInfo {
	ptr := C.gst_rtp_payload_info_for_pt(C.guint8(payloadType))
	if ptr == nil {
		return nil
	}
	return &PayloadInfo{ptr: ptr}
}
