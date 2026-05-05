package gst

/*
#include "gst.go.h"
*/
import "C"

import (
	"unsafe"

	"github.com/go-gst/go-glib/glib"
)

func init() { registerMarshalers() }

// Object wrappers

func safeWrap[From any, To any](from *From, wrapFunc func(*From) *To) *To {
	if from == nil {
		return nil
	}
	return wrapFunc(from)
}

func wrapAllocator(obj *glib.Object) *Allocator  { return safeWrap(obj, _wrapAllocator) }
func _wrapAllocator(obj *glib.Object) *Allocator { return &Allocator{wrapObject(obj)} }
func wrapBin(obj *glib.Object) *Bin              { return safeWrap(obj, _wrapBin) }
func _wrapBin(obj *glib.Object) *Bin             { return &Bin{wrapElement(obj)} }
func wrapBuffer(buf *C.GstBuffer) *Buffer        { return safeWrap(buf, _wrapBuffer) }
func _wrapBuffer(buf *C.GstBuffer) *Buffer       { return &Buffer{ptr: buf} }
func wrapBufferList(bufList *C.GstBufferList) *BufferList {
	if bufList == nil {
		return nil
	}
	return &BufferList{ptr: bufList}
}
func wrapBufferPool(obj *glib.Object) *BufferPool  { return safeWrap(obj, _wrapBufferPool) }
func _wrapBufferPool(obj *glib.Object) *BufferPool { return &BufferPool{wrapObject(obj)} }
func wrapBus(obj *glib.Object) *Bus                { return safeWrap(obj, _wrapBus) }
func _wrapBus(obj *glib.Object) *Bus               { return &Bus{Object: wrapObject(obj)} }
func wrapCaps(caps *C.GstCaps) *Caps               { return safeWrap(caps, _wrapCaps) }
func _wrapCaps(caps *C.GstCaps) *Caps              { return &Caps{native: caps} }
func wrapClock(obj *glib.Object) *Clock            { return safeWrap(obj, _wrapClock) }
func _wrapClock(obj *glib.Object) *Clock           { return &Clock{wrapObject(obj)} }
func wrapContext(ctx *C.GstContext) *Context {
	if ctx == nil {
		return nil
	}
	return &Context{ptr: ctx}
}
func wrapElement(obj *glib.Object) *Element              { return safeWrap(obj, _wrapElement) }
func _wrapElement(obj *glib.Object) *Element             { return &Element{wrapObject(obj)} }
func wrapEvent(ev *C.GstEvent) *Event                    { return safeWrap(ev, _wrapEvent) }
func _wrapEvent(ev *C.GstEvent) *Event                   { return &Event{ptr: ev} }
func wrapGhostPad(obj *glib.Object) *GhostPad            { return safeWrap(obj, _wrapGhostPad) }
func _wrapGhostPad(obj *glib.Object) *GhostPad           { return &GhostPad{wrapProxyPad(obj)} }
func wrapMapInfo(mapInfo *C.GstMapInfo) *MapInfo         { return safeWrap(mapInfo, _wrapMapInfo) }
func _wrapMapInfo(mapInfo *C.GstMapInfo) *MapInfo        { return &MapInfo{ptr: mapInfo} }
func wrapMemory(mem *C.GstMemory) *Memory                { return safeWrap(mem, _wrapMemory) }
func _wrapMemory(mem *C.GstMemory) *Memory               { return &Memory{ptr: mem} }
func wrapMessage(msg *C.GstMessage) *Message             { return safeWrap(msg, _wrapMessage) }
func _wrapMessage(msg *C.GstMessage) *Message            { return &Message{msg: msg} }
func wrapMeta(meta *C.GstMeta) *Meta                     { return safeWrap(meta, _wrapMeta) }
func _wrapMeta(meta *C.GstMeta) *Meta                    { return &Meta{ptr: meta} }
func wrapMetaInfo(info *C.GstMetaInfo) *MetaInfo         { return safeWrap(info, _wrapMetaInfo) }
func _wrapMetaInfo(info *C.GstMetaInfo) *MetaInfo        { return &MetaInfo{ptr: info} }
func wrapPad(obj *glib.Object) *Pad                      { return safeWrap(obj, _wrapPad) }
func _wrapPad(obj *glib.Object) *Pad                     { return &Pad{wrapObject(obj)} }
func wrapPadTemplate(obj *glib.Object) *PadTemplate      { return safeWrap(obj, _wrapPadTemplate) }
func _wrapPadTemplate(obj *glib.Object) *PadTemplate     { return &PadTemplate{wrapObject(obj)} }
func wrapPipeline(obj *glib.Object) *Pipeline            { return safeWrap(obj, _wrapPipeline) }
func _wrapPipeline(obj *glib.Object) *Pipeline           { return &Pipeline{Bin: wrapBin(obj)} }
func wrapPluginFeature(obj *glib.Object) *PluginFeature  { return safeWrap(obj, _wrapPluginFeature) }
func _wrapPluginFeature(obj *glib.Object) *PluginFeature { return &PluginFeature{wrapObject(obj)} }
func wrapPlugin(obj *glib.Object) *Plugin                { return safeWrap(obj, _wrapPlugin) }
func _wrapPlugin(obj *glib.Object) *Plugin               { return &Plugin{wrapObject(obj)} }
func wrapProxyPad(obj *glib.Object) *ProxyPad            { return safeWrap(obj, _wrapProxyPad) }
func _wrapProxyPad(obj *glib.Object) *ProxyPad           { return &ProxyPad{wrapPad(obj)} }
func wrapQuery(query *C.GstQuery) *Query                 { return safeWrap(query, _wrapQuery) }
func _wrapQuery(query *C.GstQuery) *Query                { return &Query{ptr: query} }
func wrapSample(sample *C.GstSample) *Sample {
	if sample == nil {
		return nil
	}
	return &Sample{sample: sample}
}
func wrapSegment(segment *C.GstSegment) *Segment  { return safeWrap(segment, _wrapSegment) }
func _wrapSegment(segment *C.GstSegment) *Segment { return &Segment{ptr: segment} }
func wrapStream(obj *glib.Object) *Stream         { return safeWrap(obj, _wrapStream) }
func _wrapStream(obj *glib.Object) *Stream        { return &Stream{wrapObject(obj)} }
func wrapTagList(tagList *C.GstTagList) *TagList  { return safeWrap(tagList, _wrapTagList) }
func _wrapTagList(tagList *C.GstTagList) *TagList { return &TagList{ptr: tagList} }
func wrapTOC(toc *C.GstToc) *TOC {
	if toc == nil {
		return nil
	}
	return &TOC{ptr: toc}
}
func wrapTOCEntry(toc *C.GstTocEntry) *TOCEntry {
	if toc == nil {
		return nil
	}
	return &TOCEntry{ptr: toc}
}
func wrapCapsFeatures(f *C.GstCapsFeatures) *CapsFeatures {
	if f == nil {
		return nil
	}
	return &CapsFeatures{native: f}
}
func wrapObject(obj *glib.Object) *Object { return safeWrap(obj, _wrapObject) }
func _wrapObject(obj *glib.Object) *Object {
	return &Object{InitiallyUnowned: &glib.InitiallyUnowned{Object: obj}}
}
func wrapElementFactory(obj *glib.Object) *ElementFactory { return safeWrap(obj, _wrapElementFactory) }
func _wrapElementFactory(obj *glib.Object) *ElementFactory {
	return &ElementFactory{wrapPluginFeature(obj)}
}
func wrapAllocationParams(p *C.GstAllocationParams) *AllocationParams {
	return safeWrap(p, _wrapAllocationParams)
}
func _wrapAllocationParams(p *C.GstAllocationParams) *AllocationParams {
	return &AllocationParams{ptr: p}
}

// ToObject wraps the given *glib.Object or *gst.Object without changing reference counts.
func ToObject(obj interface{}) *Object {
	switch obj := obj.(type) {
	case *Object:
		return obj
	case *glib.Object:
		return wrapObject(obj)
	}
	return nil
}

// Marshallers

func registerMarshalers() {
	tm := []glib.TypeMarshaler{
		{
			T: glib.Type(C.gst_buffering_mode_get_type()),
			F: marshalBufferingMode,
		},
		{
			T: glib.Type(C.gst_format_get_type()),
			F: marshalFormat,
		},
		{
			T: glib.Type(C.gst_message_type_get_type()),
			F: marshalMessageType,
		},
		{
			T: glib.Type(C.gst_pad_link_return_get_type()),
			F: marshalPadLinkReturn,
		},
		{
			T: glib.Type(C.gst_state_get_type()),
			F: marshalState,
		},
		{
			T: glib.Type(C.gst_seek_flags_get_type()),
			F: marshalSeekFlags,
		},
		{
			T: glib.Type(C.gst_seek_type_get_type()),
			F: marshalSeekType,
		},
		{
			T: glib.Type(C.gst_state_change_return_get_type()),
			F: marshalStateChangeReturn,
		},
		{
			T: glib.Type(C.gst_buffer_get_type()),
			F: marshalBuffer,
		},
		{
			T: glib.Type(C.gst_pipeline_get_type()),
			F: marshalPipeline,
		},
		{
			T: glib.Type(C.gst_bin_get_type()),
			F: marshalBin,
		},
		{
			T: glib.Type(C.gst_bus_get_type()),
			F: marshalBus,
		},
		{
			T: glib.Type(C.gst_element_get_type()),
			F: marshalElement,
		},
		{
			T: glib.Type(C.gst_element_factory_get_type()),
			F: marshalElementFactory,
		},
		{
			T: glib.Type(C.gst_proxy_pad_get_type()),
			F: marshalProxyPad,
		},
		{
			T: glib.Type(C.gst_ghost_pad_get_type()),
			F: marshalGhostPad,
		},
		{
			T: glib.Type(C.gst_object_get_type()),
			F: marshalObject,
		},
		{
			T: glib.Type(C.gst_pad_get_type()),
			F: marshalPad,
		},
		{
			T: glib.Type(C.gst_plugin_feature_get_type()),
			F: marshalPluginFeature,
		},
		{
			T: glib.Type(C.gst_allocation_params_get_type()),
			F: marshalAllocationParams,
		},
		{
			T: glib.Type(C.gst_memory_get_type()),
			F: marshalMemory,
		},
		{
			T: glib.Type(C.gst_buffer_list_get_type()),
			F: marshalBufferList,
		},
		{
			T: TypeCaps,
			F: marshalCaps,
		},
		{
			T: TypeCapsFeatures,
			F: marshalCapsFeatures,
		},
		{
			T: glib.Type(C.gst_context_get_type()),
			F: marshalContext,
		},
		{
			T: glib.Type(C.gst_toc_entry_get_type()),
			F: marshalTOCEntry,
		},
		{
			T: glib.Type(C.gst_toc_get_type()),
			F: marshalTOC,
		},
		{
			T: glib.Type(C.gst_tag_list_get_type()),
			F: marsalTagList,
		},
		{
			T: glib.Type(C.gst_event_get_type()),
			F: marshalEvent,
		},
		{
			T: glib.Type(C.gst_segment_get_type()),
			F: marshalSegment,
		},
		{
			T: glib.Type(C.gst_query_get_type()),
			F: marshalQuery,
		},
		{
			T: glib.Type(C.gst_message_get_type()),
			F: marshalMessage,
		},
		{
			T: TypeBitmask,
			F: marshalBitmask,
		},
		{
			T: TypeFraction,
			F: marshalFraction,
		},
		{
			T: TypeFractionRange,
			F: marshalFractionRange,
		},
		{
			T: TypeStructure,
			F: marshalStructure,
		},
		{
			T: TypeFloat64Range,
			F: marshalDoubleRange,
		},
		{
			T: TypeFlagset,
			F: marshalFlagset,
		},
		{
			T: TypeInt64Range,
			F: marshalInt64Range,
		},
		{
			T: TypeIntRange,
			F: marshalIntRange,
		},
		{
			T: TypeValueArray,
			F: marshalValueArray,
		},
		{
			T: TypeValueList,
			F: marshalValueList,
		},
		{
			T: TypePromise,
			F: marshalPromise,
		},
		{
			T: glib.Type(C.gst_sample_get_type()),
			F: marshalSample,
		},
	}

	glib.RegisterGValueMarshalers(tm)
}

func toGValue(p unsafe.Pointer) *C.GValue {
	return (*C.GValue)(p)
}

func marshalValueArray(p unsafe.Pointer) (interface{}, error) {
	val := toGValue(p)
	out := ValueArrayValue(*glib.ValueFromNative(unsafe.Pointer(val)))
	return &out, nil
}

func marshalValueList(p unsafe.Pointer) (interface{}, error) {
	value := glib.ValueFromNative(p)

	// must copy since we don't own the gvalue passed into this marshal
	out, err := value.Copy()

	if err != nil {
		return nil, err
	}

	return (*ValueListValue)(out), nil
}

func marshalInt64Range(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	return &Int64RangeValue{
		start: int64(C.gst_value_get_int64_range_min(v)),
		end:   int64(C.gst_value_get_int64_range_max(v)),
		step:  int64(C.gst_value_get_int64_range_step(v)),
	}, nil
}

func marshalIntRange(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	return &IntRangeValue{
		start: int(C.gst_value_get_int_range_min(v)),
		end:   int(C.gst_value_get_int_range_max(v)),
		step:  int(C.gst_value_get_int_range_step(v)),
	}, nil
}

func marshalBitmask(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	return Bitmask(C.gst_value_get_bitmask(v)), nil
}

func marshalFlagset(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	return &FlagsetValue{
		flags: uint(C.gst_value_get_flagset_flags(v)),
		mask:  uint(C.gst_value_get_flagset_mask(v)),
	}, nil
}

func marshalDoubleRange(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	return &Float64RangeValue{
		start: float64(C.gst_value_get_double_range_min(v)),
		end:   float64(C.gst_value_get_double_range_max(v)),
	}, nil
}

func marshalFraction(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	out := &FractionValue{
		num:   int(C.gst_value_get_fraction_numerator(v)),
		denom: int(C.gst_value_get_fraction_denominator(v)),
	}
	return out, nil
}

func marshalFractionRange(p unsafe.Pointer) (interface{}, error) {
	v := toGValue(p)
	start := C.gst_value_get_fraction_range_min(v)
	end := C.gst_value_get_fraction_range_max(v)
	return &FractionRangeValue{
		start: ValueGetFraction(glib.ValueFromNative(unsafe.Pointer(start))),
		end:   ValueGetFraction(glib.ValueFromNative(unsafe.Pointer(end))),
	}, nil
}

func marshalBufferingMode(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return BufferingMode(c), nil
}

func marshalFormat(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return Format(c), nil
}

func marshalMessageType(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return MessageType(c), nil
}

func marshalPadLinkReturn(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return PadLinkReturn(c), nil
}

func marshalState(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return State(c), nil
}

func marshalSeekFlags(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return SeekFlags(c), nil
}

func marshalSeekType(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return SeekType(c), nil
}

func marshalStateChangeReturn(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_enum(toGValue(p))
	return StateChangeReturn(c), nil
}

func marshalGhostPad(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapGhostPad(obj), nil
}

func marshalProxyPad(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapProxyPad(obj), nil
}

func marshalPad(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapPad(obj), nil
}

func marshalMessage(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstMessageUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalObject(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapObject(obj), nil
}

func marshalBus(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapBus(obj), nil
}

func marshalElementFactory(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapElementFactory(obj), nil
}

func marshalPipeline(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapPipeline(obj), nil
}

func marshalPluginFeature(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapPluginFeature(obj), nil
}

func marshalElement(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapElement(obj), nil
}

func marshalBin(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_object(toGValue(p))
	obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
	return wrapBin(obj), nil
}

func marshalAllocationParams(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	if c == nil {
		return nil, nil
	}
	return wrapAllocationParams(C.gst_allocation_params_copy((*C.GstAllocationParams)(unsafe.Pointer(c)))), nil
}

func marshalMemory(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstMemoryUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalBuffer(p unsafe.Pointer) (interface{}, error) {
	c := C.getBufferValue(toGValue(p))
	return FromGstBufferUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalBufferList(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstBufferListUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalCaps(p unsafe.Pointer) (interface{}, error) {
	c := C.gst_value_get_caps(toGValue(p))
	return FromGstCapsUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalCapsFeatures(p unsafe.Pointer) (interface{}, error) {
	c := C.gst_value_get_caps_features(toGValue(p))
	if c == nil {
		return nil, nil
	}
	return wrapCapsFeatures(C.gst_caps_features_copy(c)), nil
}

func marshalContext(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstContextUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalTOC(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstTOCUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalTOCEntry(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstTocEntryUnsafeNone(unsafe.Pointer(c)), nil
}

func marsalTagList(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstTagListUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalEvent(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstEventUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalSegment(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	if c == nil {
		return nil, nil
	}
	return wrapSegment(C.gst_segment_copy((*C.GstSegment)(unsafe.Pointer(c)))), nil
}

func marshalQuery(p unsafe.Pointer) (interface{}, error) {
	c := C.g_value_get_boxed(toGValue(p))
	return FromGstQueryUnsafeNone(unsafe.Pointer(c)), nil
}

func marshalSample(p unsafe.Pointer) (interface{}, error) {
	c := C.getSampleValue(toGValue(p))
	return FromGstSampleUnsafeNone(unsafe.Pointer(c)), nil
}
