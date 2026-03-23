package base

/*
#include "gst.go.h"
*/
import "C"

import (
	"github.com/go-gst/go-glib/glib"
	"github.com/go-gst/go-gst/gst"
)

// GstPushSrc represents a GstBaseSrc.
type GstPushSrc struct{ *GstBaseSrc }

// ToGstPushSrc returns a GstPushSrc object for the given object.
func ToGstPushSrc(obj interface{}) *GstPushSrc {
	switch obj := obj.(type) {
	case *GstPushSrc:
		return obj
	case *GstBaseSrc:
		return &GstPushSrc{obj}
	case *gst.Element:
		return &GstPushSrc{&GstBaseSrc{obj}}
	case *gst.Object:
		return &GstPushSrc{&GstBaseSrc{&gst.Element{Object: obj}}}
	case *glib.Object:
		return &GstPushSrc{&GstBaseSrc{gst.ToElement(obj)}}
	}
	return nil
}

// Instance returns the underlying C GstBaseSrc instance
func (g *GstPushSrc) Instance() *C.GstPushSrc {
	if g == nil {
		return nil
	}
	return C.toGstPushSrc(g.Unsafe())
}
