package gst

// #include "gst.go.h"
import "C"
import "unsafe"

// StaticPadTemplate is a go representation of a GstStaticPadTemplate.
type StaticPadTemplate struct {
	ptr *C.GstStaticPadTemplate
}

// FromGstStaticPadTemplateUnsafe wraps the given GstStaticPadTemplate pointer.
func FromGstStaticPadTemplateUnsafe(tmpl unsafe.Pointer) *StaticPadTemplate {
	if tmpl == nil {
		return nil
	}
	return &StaticPadTemplate{ptr: (*C.GstStaticPadTemplate)(tmpl)}
}

// Instance returns the underlying C GstStaticPadTemplate.
func (s *StaticPadTemplate) Instance() *C.GstStaticPadTemplate {
	if s == nil {
		return nil
	}
	return s.ptr
}

// Name returns the name template of the pad template.
func (s *StaticPadTemplate) Name() string {
	return C.GoString(s.Instance().name_template)
}

// Direction returns the direction of the pad template.
func (s *StaticPadTemplate) Direction() PadDirection {
	return PadDirection(s.Instance().direction)
}

// Presence returns the presence of the pad template.
func (s *StaticPadTemplate) Presence() PadPresence {
	return PadPresence(s.Instance().presence)
}

// Caps converts the static caps of the pad template to a GstCaps.
func (s *StaticPadTemplate) Caps() *Caps {
	return FromGstCapsUnsafeFull(unsafe.Pointer(C.gst_static_pad_template_get_caps(s.Instance())))
}
