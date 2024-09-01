package qt

/*

#include "gen_qpainterpath.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QPainterPath struct {
	h *C.QPainterPath
}

func (this *QPainterPath) cPointer() *C.QPainterPath {
	if this == nil {
		return nil
	}
	return this.h
}

func newQPainterPath(h *C.QPainterPath) *QPainterPath {
	if h == nil {
		return nil
	}
	return &QPainterPath{h: h}
}

func newQPainterPath_U(h unsafe.Pointer) *QPainterPath {
	return newQPainterPath((*C.QPainterPath)(h))
}

// NewQPainterPath constructs a new QPainterPath object.
func NewQPainterPath() *QPainterPath {
	ret := C.QPainterPath_new()
	return newQPainterPath(ret)
}

// NewQPainterPath2 constructs a new QPainterPath object.
func NewQPainterPath2(startPoint *QPointF) *QPainterPath {
	ret := C.QPainterPath_new2(startPoint.cPointer())
	return newQPainterPath(ret)
}

// NewQPainterPath3 constructs a new QPainterPath object.
func NewQPainterPath3(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_new3(other.cPointer())
	return newQPainterPath(ret)
}

func (this *QPainterPath) OperatorAssign(other *QPainterPath) {
	C.QPainterPath_OperatorAssign(this.h, other.cPointer())
}

func (this *QPainterPath) Swap(other *QPainterPath) {
	C.QPainterPath_Swap(this.h, other.cPointer())
}

func (this *QPainterPath) Clear() {
	C.QPainterPath_Clear(this.h)
}

func (this *QPainterPath) Reserve(size int) {
	C.QPainterPath_Reserve(this.h, (C.int)(size))
}

func (this *QPainterPath) Capacity() int {
	ret := C.QPainterPath_Capacity(this.h)
	return (int)(ret)
}

func (this *QPainterPath) CloseSubpath() {
	C.QPainterPath_CloseSubpath(this.h)
}

func (this *QPainterPath) MoveTo(p *QPointF) {
	C.QPainterPath_MoveTo(this.h, p.cPointer())
}

func (this *QPainterPath) MoveTo2(x float64, y float64) {
	C.QPainterPath_MoveTo2(this.h, (C.double)(x), (C.double)(y))
}

func (this *QPainterPath) LineTo(p *QPointF) {
	C.QPainterPath_LineTo(this.h, p.cPointer())
}

func (this *QPainterPath) LineTo2(x float64, y float64) {
	C.QPainterPath_LineTo2(this.h, (C.double)(x), (C.double)(y))
}

func (this *QPainterPath) ArcMoveTo(rect *QRectF, angle float64) {
	C.QPainterPath_ArcMoveTo(this.h, rect.cPointer(), (C.double)(angle))
}

func (this *QPainterPath) ArcMoveTo2(x float64, y float64, w float64, h float64, angle float64) {
	C.QPainterPath_ArcMoveTo2(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h), (C.double)(angle))
}

func (this *QPainterPath) ArcTo(rect *QRectF, startAngle float64, arcLength float64) {
	C.QPainterPath_ArcTo(this.h, rect.cPointer(), (C.double)(startAngle), (C.double)(arcLength))
}

func (this *QPainterPath) ArcTo2(x float64, y float64, w float64, h float64, startAngle float64, arcLength float64) {
	C.QPainterPath_ArcTo2(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h), (C.double)(startAngle), (C.double)(arcLength))
}

func (this *QPainterPath) CubicTo(ctrlPt1 *QPointF, ctrlPt2 *QPointF, endPt *QPointF) {
	C.QPainterPath_CubicTo(this.h, ctrlPt1.cPointer(), ctrlPt2.cPointer(), endPt.cPointer())
}

func (this *QPainterPath) CubicTo2(ctrlPt1x float64, ctrlPt1y float64, ctrlPt2x float64, ctrlPt2y float64, endPtx float64, endPty float64) {
	C.QPainterPath_CubicTo2(this.h, (C.double)(ctrlPt1x), (C.double)(ctrlPt1y), (C.double)(ctrlPt2x), (C.double)(ctrlPt2y), (C.double)(endPtx), (C.double)(endPty))
}

func (this *QPainterPath) QuadTo(ctrlPt *QPointF, endPt *QPointF) {
	C.QPainterPath_QuadTo(this.h, ctrlPt.cPointer(), endPt.cPointer())
}

func (this *QPainterPath) QuadTo2(ctrlPtx float64, ctrlPty float64, endPtx float64, endPty float64) {
	C.QPainterPath_QuadTo2(this.h, (C.double)(ctrlPtx), (C.double)(ctrlPty), (C.double)(endPtx), (C.double)(endPty))
}

func (this *QPainterPath) CurrentPosition() *QPointF {
	ret := C.QPainterPath_CurrentPosition(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPointF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPointF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) AddRect(rect *QRectF) {
	C.QPainterPath_AddRect(this.h, rect.cPointer())
}

func (this *QPainterPath) AddRect2(x float64, y float64, w float64, h float64) {
	C.QPainterPath_AddRect2(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h))
}

func (this *QPainterPath) AddEllipse(rect *QRectF) {
	C.QPainterPath_AddEllipse(this.h, rect.cPointer())
}

func (this *QPainterPath) AddEllipse2(x float64, y float64, w float64, h float64) {
	C.QPainterPath_AddEllipse2(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h))
}

func (this *QPainterPath) AddEllipse3(center *QPointF, rx float64, ry float64) {
	C.QPainterPath_AddEllipse3(this.h, center.cPointer(), (C.double)(rx), (C.double)(ry))
}

func (this *QPainterPath) AddText(point *QPointF, f *QFont, text string) {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	C.QPainterPath_AddText(this.h, point.cPointer(), f.cPointer(), text_Cstring, C.ulong(len(text)))
}

func (this *QPainterPath) AddText2(x float64, y float64, f *QFont, text string) {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	C.QPainterPath_AddText2(this.h, (C.double)(x), (C.double)(y), f.cPointer(), text_Cstring, C.ulong(len(text)))
}

func (this *QPainterPath) AddPath(path *QPainterPath) {
	C.QPainterPath_AddPath(this.h, path.cPointer())
}

func (this *QPainterPath) AddRegion(region *QRegion) {
	C.QPainterPath_AddRegion(this.h, region.cPointer())
}

func (this *QPainterPath) AddRoundedRect(rect *QRectF, xRadius float64, yRadius float64) {
	C.QPainterPath_AddRoundedRect(this.h, rect.cPointer(), (C.double)(xRadius), (C.double)(yRadius))
}

func (this *QPainterPath) AddRoundedRect2(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64) {
	C.QPainterPath_AddRoundedRect2(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h), (C.double)(xRadius), (C.double)(yRadius))
}

func (this *QPainterPath) AddRoundRect(rect *QRectF, xRnd int, yRnd int) {
	C.QPainterPath_AddRoundRect(this.h, rect.cPointer(), (C.int)(xRnd), (C.int)(yRnd))
}

func (this *QPainterPath) AddRoundRect2(x float64, y float64, w float64, h float64, xRnd int, yRnd int) {
	C.QPainterPath_AddRoundRect2(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h), (C.int)(xRnd), (C.int)(yRnd))
}

func (this *QPainterPath) AddRoundRect3(rect *QRectF, roundness int) {
	C.QPainterPath_AddRoundRect3(this.h, rect.cPointer(), (C.int)(roundness))
}

func (this *QPainterPath) AddRoundRect4(x float64, y float64, w float64, h float64, roundness int) {
	C.QPainterPath_AddRoundRect4(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h), (C.int)(roundness))
}

func (this *QPainterPath) ConnectPath(path *QPainterPath) {
	C.QPainterPath_ConnectPath(this.h, path.cPointer())
}

func (this *QPainterPath) Contains(pt *QPointF) bool {
	ret := C.QPainterPath_Contains(this.h, pt.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) ContainsWithRect(rect *QRectF) bool {
	ret := C.QPainterPath_ContainsWithRect(this.h, rect.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) Intersects(rect *QRectF) bool {
	ret := C.QPainterPath_Intersects(this.h, rect.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) Translate(dx float64, dy float64) {
	C.QPainterPath_Translate(this.h, (C.double)(dx), (C.double)(dy))
}

func (this *QPainterPath) TranslateWithOffset(offset *QPointF) {
	C.QPainterPath_TranslateWithOffset(this.h, offset.cPointer())
}

func (this *QPainterPath) Translated(dx float64, dy float64) *QPainterPath {
	ret := C.QPainterPath_Translated(this.h, (C.double)(dx), (C.double)(dy))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) TranslatedWithOffset(offset *QPointF) *QPainterPath {
	ret := C.QPainterPath_TranslatedWithOffset(this.h, offset.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) BoundingRect() *QRectF {
	ret := C.QPainterPath_BoundingRect(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRectF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRectF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) ControlPointRect() *QRectF {
	ret := C.QPainterPath_ControlPointRect(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRectF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRectF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) FillRule() uintptr {
	ret := C.QPainterPath_FillRule(this.h)
	return (uintptr)(ret)
}

func (this *QPainterPath) SetFillRule(fillRule uintptr) {
	C.QPainterPath_SetFillRule(this.h, (C.uintptr_t)(fillRule))
}

func (this *QPainterPath) IsEmpty() bool {
	ret := C.QPainterPath_IsEmpty(this.h)
	return (bool)(ret)
}

func (this *QPainterPath) ToReversed() *QPainterPath {
	ret := C.QPainterPath_ToReversed(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) ElementCount() int {
	ret := C.QPainterPath_ElementCount(this.h)
	return (int)(ret)
}

func (this *QPainterPath) ElementAt(i int) *QPainterPath__Element {
	ret := C.QPainterPath_ElementAt(this.h, (C.int)(i))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath__Element(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath__Element) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) SetElementPositionAt(i int, x float64, y float64) {
	C.QPainterPath_SetElementPositionAt(this.h, (C.int)(i), (C.double)(x), (C.double)(y))
}

func (this *QPainterPath) Length() float64 {
	ret := C.QPainterPath_Length(this.h)
	return (float64)(ret)
}

func (this *QPainterPath) PercentAtLength(t float64) float64 {
	ret := C.QPainterPath_PercentAtLength(this.h, (C.double)(t))
	return (float64)(ret)
}

func (this *QPainterPath) PointAtPercent(t float64) *QPointF {
	ret := C.QPainterPath_PointAtPercent(this.h, (C.double)(t))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPointF(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPointF) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) AngleAtPercent(t float64) float64 {
	ret := C.QPainterPath_AngleAtPercent(this.h, (C.double)(t))
	return (float64)(ret)
}

func (this *QPainterPath) SlopeAtPercent(t float64) float64 {
	ret := C.QPainterPath_SlopeAtPercent(this.h, (C.double)(t))
	return (float64)(ret)
}

func (this *QPainterPath) IntersectsWithQPainterPath(p *QPainterPath) bool {
	ret := C.QPainterPath_IntersectsWithQPainterPath(this.h, p.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) ContainsWithQPainterPath(p *QPainterPath) bool {
	ret := C.QPainterPath_ContainsWithQPainterPath(this.h, p.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) United(r *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_United(this.h, r.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) Intersected(r *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_Intersected(this.h, r.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) Subtracted(r *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_Subtracted(this.h, r.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) SubtractedInverted(r *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_SubtractedInverted(this.h, r.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) Simplified() *QPainterPath {
	ret := C.QPainterPath_Simplified(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) OperatorEqual(other *QPainterPath) bool {
	ret := C.QPainterPath_OperatorEqual(this.h, other.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) OperatorNotEqual(other *QPainterPath) bool {
	ret := C.QPainterPath_OperatorNotEqual(this.h, other.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath) OperatorBitwiseAnd(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_OperatorBitwiseAnd(this.h, other.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) OperatorBitwiseOr(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_OperatorBitwiseOr(this.h, other.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) OperatorPlus(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_OperatorPlus(this.h, other.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) OperatorMinus(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_OperatorMinus(this.h, other.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPath) OperatorBitwiseAndAssign(other *QPainterPath) {
	C.QPainterPath_OperatorBitwiseAndAssign(this.h, other.cPointer())
}

func (this *QPainterPath) OperatorBitwiseOrAssign(other *QPainterPath) {
	C.QPainterPath_OperatorBitwiseOrAssign(this.h, other.cPointer())
}

func (this *QPainterPath) OperatorPlusAssign(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_OperatorPlusAssign(this.h, other.cPointer())
	return newQPainterPath_U(unsafe.Pointer(ret))
}

func (this *QPainterPath) OperatorMinusAssign(other *QPainterPath) *QPainterPath {
	ret := C.QPainterPath_OperatorMinusAssign(this.h, other.cPointer())
	return newQPainterPath_U(unsafe.Pointer(ret))
}

func (this *QPainterPath) AddRoundedRect4(rect *QRectF, xRadius float64, yRadius float64, mode uintptr) {
	C.QPainterPath_AddRoundedRect4(this.h, rect.cPointer(), (C.double)(xRadius), (C.double)(yRadius), (C.uintptr_t)(mode))
}

func (this *QPainterPath) AddRoundedRect7(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode uintptr) {
	C.QPainterPath_AddRoundedRect7(this.h, (C.double)(x), (C.double)(y), (C.double)(w), (C.double)(h), (C.double)(xRadius), (C.double)(yRadius), (C.uintptr_t)(mode))
}

func (this *QPainterPath) Delete() {
	C.QPainterPath_Delete(this.h)
}

type QPainterPathStroker struct {
	h *C.QPainterPathStroker
}

func (this *QPainterPathStroker) cPointer() *C.QPainterPathStroker {
	if this == nil {
		return nil
	}
	return this.h
}

func newQPainterPathStroker(h *C.QPainterPathStroker) *QPainterPathStroker {
	if h == nil {
		return nil
	}
	return &QPainterPathStroker{h: h}
}

func newQPainterPathStroker_U(h unsafe.Pointer) *QPainterPathStroker {
	return newQPainterPathStroker((*C.QPainterPathStroker)(h))
}

// NewQPainterPathStroker constructs a new QPainterPathStroker object.
func NewQPainterPathStroker() *QPainterPathStroker {
	ret := C.QPainterPathStroker_new()
	return newQPainterPathStroker(ret)
}

// NewQPainterPathStroker2 constructs a new QPainterPathStroker object.
func NewQPainterPathStroker2(pen *QPen) *QPainterPathStroker {
	ret := C.QPainterPathStroker_new2(pen.cPointer())
	return newQPainterPathStroker(ret)
}

func (this *QPainterPathStroker) SetWidth(width float64) {
	C.QPainterPathStroker_SetWidth(this.h, (C.double)(width))
}

func (this *QPainterPathStroker) Width() float64 {
	ret := C.QPainterPathStroker_Width(this.h)
	return (float64)(ret)
}

func (this *QPainterPathStroker) SetCapStyle(style uintptr) {
	C.QPainterPathStroker_SetCapStyle(this.h, (C.uintptr_t)(style))
}

func (this *QPainterPathStroker) CapStyle() uintptr {
	ret := C.QPainterPathStroker_CapStyle(this.h)
	return (uintptr)(ret)
}

func (this *QPainterPathStroker) SetJoinStyle(style uintptr) {
	C.QPainterPathStroker_SetJoinStyle(this.h, (C.uintptr_t)(style))
}

func (this *QPainterPathStroker) JoinStyle() uintptr {
	ret := C.QPainterPathStroker_JoinStyle(this.h)
	return (uintptr)(ret)
}

func (this *QPainterPathStroker) SetMiterLimit(length float64) {
	C.QPainterPathStroker_SetMiterLimit(this.h, (C.double)(length))
}

func (this *QPainterPathStroker) MiterLimit() float64 {
	ret := C.QPainterPathStroker_MiterLimit(this.h)
	return (float64)(ret)
}

func (this *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	C.QPainterPathStroker_SetCurveThreshold(this.h, (C.double)(threshold))
}

func (this *QPainterPathStroker) CurveThreshold() float64 {
	ret := C.QPainterPathStroker_CurveThreshold(this.h)
	return (float64)(ret)
}

func (this *QPainterPathStroker) SetDashPattern(dashPattern uintptr) {
	C.QPainterPathStroker_SetDashPattern(this.h, (C.uintptr_t)(dashPattern))
}

func (this *QPainterPathStroker) SetDashPatternWithDashPattern(dashPattern []float64) {
	// For the C ABI, malloc a C array of raw pointers
	dashPattern_CArray := (*[0xffff]C.double)(C.malloc(C.size_t(8 * len(dashPattern))))
	defer C.free(unsafe.Pointer(dashPattern_CArray))
	for i := range dashPattern {
		dashPattern_CArray[i] = (C.double)(dashPattern[i])
	}
	C.QPainterPathStroker_SetDashPatternWithDashPattern(this.h, &dashPattern_CArray[0], C.ulong(len(dashPattern)))
}

func (this *QPainterPathStroker) DashPattern() []float64 {
	var _out *C.double = nil
	var _out_len C.size_t = 0
	C.QPainterPathStroker_DashPattern(this.h, &_out, &_out_len)
	ret := make([]float64, int(_out_len))
	_outCast := (*[0xffff]C.double)(unsafe.Pointer(_out)) // mrs jackson
	for i := 0; i < int(_out_len); i++ {
		ret[i] = (float64)(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QPainterPathStroker) SetDashOffset(offset float64) {
	C.QPainterPathStroker_SetDashOffset(this.h, (C.double)(offset))
}

func (this *QPainterPathStroker) DashOffset() float64 {
	ret := C.QPainterPathStroker_DashOffset(this.h)
	return (float64)(ret)
}

func (this *QPainterPathStroker) CreateStroke(path *QPainterPath) *QPainterPath {
	ret := C.QPainterPathStroker_CreateStroke(this.h, path.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPainterPath(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPainterPath) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QPainterPathStroker) Delete() {
	C.QPainterPathStroker_Delete(this.h)
}

type QPainterPath__Element struct {
	h *C.QPainterPath__Element
}

func (this *QPainterPath__Element) cPointer() *C.QPainterPath__Element {
	if this == nil {
		return nil
	}
	return this.h
}

func newQPainterPath__Element(h *C.QPainterPath__Element) *QPainterPath__Element {
	if h == nil {
		return nil
	}
	return &QPainterPath__Element{h: h}
}

func newQPainterPath__Element_U(h unsafe.Pointer) *QPainterPath__Element {
	return newQPainterPath__Element((*C.QPainterPath__Element)(h))
}

func (this *QPainterPath__Element) IsMoveTo() bool {
	ret := C.QPainterPath__Element_IsMoveTo(this.h)
	return (bool)(ret)
}

func (this *QPainterPath__Element) IsLineTo() bool {
	ret := C.QPainterPath__Element_IsLineTo(this.h)
	return (bool)(ret)
}

func (this *QPainterPath__Element) IsCurveTo() bool {
	ret := C.QPainterPath__Element_IsCurveTo(this.h)
	return (bool)(ret)
}

func (this *QPainterPath__Element) OperatorEqual(e *QPainterPath__Element) bool {
	ret := C.QPainterPath__Element_OperatorEqual(this.h, e.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath__Element) OperatorNotEqual(e *QPainterPath__Element) bool {
	ret := C.QPainterPath__Element_OperatorNotEqual(this.h, e.cPointer())
	return (bool)(ret)
}

func (this *QPainterPath__Element) Delete() {
	C.QPainterPath__Element_Delete(this.h)
}
