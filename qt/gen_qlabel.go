package qt

/*

#include "gen_qlabel.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QLabel struct {
	h *C.QLabel
	*QFrame
}

func (this *QLabel) cPointer() *C.QLabel {
	if this == nil {
		return nil
	}
	return this.h
}

func newQLabel(h *C.QLabel) *QLabel {
	if h == nil {
		return nil
	}
	return &QLabel{h: h, QFrame: newQFrame_U(unsafe.Pointer(h))}
}

func newQLabel_U(h unsafe.Pointer) *QLabel {
	return newQLabel((*C.QLabel)(h))
}

// NewQLabel constructs a new QLabel object.
func NewQLabel() *QLabel {
	ret := C.QLabel_new()
	return newQLabel(ret)
}

// NewQLabel2 constructs a new QLabel object.
func NewQLabel2(text string) *QLabel {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	ret := C.QLabel_new2(text_Cstring, C.size_t(len(text)))
	return newQLabel(ret)
}

// NewQLabel3 constructs a new QLabel object.
func NewQLabel3(parent *QWidget) *QLabel {
	ret := C.QLabel_new3(parent.cPointer())
	return newQLabel(ret)
}

// NewQLabel4 constructs a new QLabel object.
func NewQLabel4(parent *QWidget, f int) *QLabel {
	ret := C.QLabel_new4(parent.cPointer(), (C.int)(f))
	return newQLabel(ret)
}

// NewQLabel5 constructs a new QLabel object.
func NewQLabel5(text string, parent *QWidget) *QLabel {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	ret := C.QLabel_new5(text_Cstring, C.size_t(len(text)), parent.cPointer())
	return newQLabel(ret)
}

// NewQLabel6 constructs a new QLabel object.
func NewQLabel6(text string, parent *QWidget, f int) *QLabel {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	ret := C.QLabel_new6(text_Cstring, C.size_t(len(text)), parent.cPointer(), (C.int)(f))
	return newQLabel(ret)
}

func (this *QLabel) MetaObject() *QMetaObject {
	ret := C.QLabel_MetaObject(this.h)
	return newQMetaObject_U(unsafe.Pointer(ret))
}

func QLabel_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_Tr(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QLabel_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_TrUtf8(s_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QLabel) Text() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_Text(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QLabel) Pixmap() *QPixmap {
	ret := C.QLabel_Pixmap(this.h)
	return newQPixmap_U(unsafe.Pointer(ret))
}

func (this *QLabel) PixmapWithQtReturnByValueConstant(param1 uintptr) *QPixmap {
	ret := C.QLabel_PixmapWithQtReturnByValueConstant(this.h, (C.uintptr_t)(param1))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPixmap(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPixmap) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QLabel) Picture() *QPicture {
	ret := C.QLabel_Picture(this.h)
	return newQPicture_U(unsafe.Pointer(ret))
}

func (this *QLabel) PictureWithQtReturnByValueConstant(param1 uintptr) *QPicture {
	ret := C.QLabel_PictureWithQtReturnByValueConstant(this.h, (C.uintptr_t)(param1))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQPicture(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QPicture) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QLabel) Movie() *QMovie {
	ret := C.QLabel_Movie(this.h)
	return newQMovie_U(unsafe.Pointer(ret))
}

func (this *QLabel) TextFormat() uintptr {
	ret := C.QLabel_TextFormat(this.h)
	return (uintptr)(ret)
}

func (this *QLabel) SetTextFormat(textFormat uintptr) {
	C.QLabel_SetTextFormat(this.h, (C.uintptr_t)(textFormat))
}

func (this *QLabel) Alignment() int {
	ret := C.QLabel_Alignment(this.h)
	return (int)(ret)
}

func (this *QLabel) SetAlignment(alignment int) {
	C.QLabel_SetAlignment(this.h, (C.int)(alignment))
}

func (this *QLabel) SetWordWrap(on bool) {
	C.QLabel_SetWordWrap(this.h, (C.bool)(on))
}

func (this *QLabel) WordWrap() bool {
	ret := C.QLabel_WordWrap(this.h)
	return (bool)(ret)
}

func (this *QLabel) Indent() int {
	ret := C.QLabel_Indent(this.h)
	return (int)(ret)
}

func (this *QLabel) SetIndent(indent int) {
	C.QLabel_SetIndent(this.h, (C.int)(indent))
}

func (this *QLabel) Margin() int {
	ret := C.QLabel_Margin(this.h)
	return (int)(ret)
}

func (this *QLabel) SetMargin(margin int) {
	C.QLabel_SetMargin(this.h, (C.int)(margin))
}

func (this *QLabel) HasScaledContents() bool {
	ret := C.QLabel_HasScaledContents(this.h)
	return (bool)(ret)
}

func (this *QLabel) SetScaledContents(scaledContents bool) {
	C.QLabel_SetScaledContents(this.h, (C.bool)(scaledContents))
}

func (this *QLabel) SizeHint() *QSize {
	ret := C.QLabel_SizeHint(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQSize(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QSize) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QLabel) MinimumSizeHint() *QSize {
	ret := C.QLabel_MinimumSizeHint(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQSize(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QSize) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QLabel) SetBuddy(buddy *QWidget) {
	C.QLabel_SetBuddy(this.h, buddy.cPointer())
}

func (this *QLabel) Buddy() *QWidget {
	ret := C.QLabel_Buddy(this.h)
	return newQWidget_U(unsafe.Pointer(ret))
}

func (this *QLabel) HeightForWidth(param1 int) int {
	ret := C.QLabel_HeightForWidth(this.h, (C.int)(param1))
	return (int)(ret)
}

func (this *QLabel) OpenExternalLinks() bool {
	ret := C.QLabel_OpenExternalLinks(this.h)
	return (bool)(ret)
}

func (this *QLabel) SetOpenExternalLinks(open bool) {
	C.QLabel_SetOpenExternalLinks(this.h, (C.bool)(open))
}

func (this *QLabel) SetTextInteractionFlags(flags int) {
	C.QLabel_SetTextInteractionFlags(this.h, (C.int)(flags))
}

func (this *QLabel) TextInteractionFlags() int {
	ret := C.QLabel_TextInteractionFlags(this.h)
	return (int)(ret)
}

func (this *QLabel) SetSelection(param1 int, param2 int) {
	C.QLabel_SetSelection(this.h, (C.int)(param1), (C.int)(param2))
}

func (this *QLabel) HasSelectedText() bool {
	ret := C.QLabel_HasSelectedText(this.h)
	return (bool)(ret)
}

func (this *QLabel) SelectedText() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_SelectedText(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QLabel) SelectionStart() int {
	ret := C.QLabel_SelectionStart(this.h)
	return (int)(ret)
}

func (this *QLabel) SetText(text string) {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	C.QLabel_SetText(this.h, text_Cstring, C.size_t(len(text)))
}

func (this *QLabel) SetPixmap(pixmap *QPixmap) {
	C.QLabel_SetPixmap(this.h, pixmap.cPointer())
}

func (this *QLabel) SetPicture(picture *QPicture) {
	C.QLabel_SetPicture(this.h, picture.cPointer())
}

func (this *QLabel) SetMovie(movie *QMovie) {
	C.QLabel_SetMovie(this.h, movie.cPointer())
}

func (this *QLabel) SetNum(num int) {
	C.QLabel_SetNum(this.h, (C.int)(num))
}

func (this *QLabel) SetNumWithNum(num float64) {
	C.QLabel_SetNumWithNum(this.h, (C.double)(num))
}

func (this *QLabel) Clear() {
	C.QLabel_Clear(this.h)
}

func (this *QLabel) LinkActivated(link string) {
	link_Cstring := C.CString(link)
	defer C.free(unsafe.Pointer(link_Cstring))
	C.QLabel_LinkActivated(this.h, link_Cstring, C.size_t(len(link)))
}

func (this *QLabel) OnLinkActivated(slot func()) {
	var slotWrapper miqtCallbackFunc = func(argc C.int, args *C.void) {
		slot()
	}

	C.QLabel_connect_LinkActivated(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slotWrapper))))
}

func (this *QLabel) LinkHovered(link string) {
	link_Cstring := C.CString(link)
	defer C.free(unsafe.Pointer(link_Cstring))
	C.QLabel_LinkHovered(this.h, link_Cstring, C.size_t(len(link)))
}

func (this *QLabel) OnLinkHovered(slot func()) {
	var slotWrapper miqtCallbackFunc = func(argc C.int, args *C.void) {
		slot()
	}

	C.QLabel_connect_LinkHovered(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slotWrapper))))
}

func QLabel_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_Tr2(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QLabel_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_Tr3(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QLabel_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_TrUtf82(s_Cstring, c_Cstring, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QLabel_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QLabel_TrUtf83(s_Cstring, c_Cstring, (C.int)(n), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QLabel) Delete() {
	C.QLabel_Delete(this.h)
}
