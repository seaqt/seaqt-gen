package qt6

/*

#include "gen_qimage.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QImage__InvertMode int

const (
	QImage__InvertRgb  QImage__InvertMode = 0
	QImage__InvertRgba QImage__InvertMode = 1
)

type QImage__Format int

const (
	QImage__Format_Invalid                  QImage__Format = 0
	QImage__Format_Mono                     QImage__Format = 1
	QImage__Format_MonoLSB                  QImage__Format = 2
	QImage__Format_Indexed8                 QImage__Format = 3
	QImage__Format_RGB32                    QImage__Format = 4
	QImage__Format_ARGB32                   QImage__Format = 5
	QImage__Format_ARGB32_Premultiplied     QImage__Format = 6
	QImage__Format_RGB16                    QImage__Format = 7
	QImage__Format_ARGB8565_Premultiplied   QImage__Format = 8
	QImage__Format_RGB666                   QImage__Format = 9
	QImage__Format_ARGB6666_Premultiplied   QImage__Format = 10
	QImage__Format_RGB555                   QImage__Format = 11
	QImage__Format_ARGB8555_Premultiplied   QImage__Format = 12
	QImage__Format_RGB888                   QImage__Format = 13
	QImage__Format_RGB444                   QImage__Format = 14
	QImage__Format_ARGB4444_Premultiplied   QImage__Format = 15
	QImage__Format_RGBX8888                 QImage__Format = 16
	QImage__Format_RGBA8888                 QImage__Format = 17
	QImage__Format_RGBA8888_Premultiplied   QImage__Format = 18
	QImage__Format_BGR30                    QImage__Format = 19
	QImage__Format_A2BGR30_Premultiplied    QImage__Format = 20
	QImage__Format_RGB30                    QImage__Format = 21
	QImage__Format_A2RGB30_Premultiplied    QImage__Format = 22
	QImage__Format_Alpha8                   QImage__Format = 23
	QImage__Format_Grayscale8               QImage__Format = 24
	QImage__Format_RGBX64                   QImage__Format = 25
	QImage__Format_RGBA64                   QImage__Format = 26
	QImage__Format_RGBA64_Premultiplied     QImage__Format = 27
	QImage__Format_Grayscale16              QImage__Format = 28
	QImage__Format_BGR888                   QImage__Format = 29
	QImage__Format_RGBX16FPx4               QImage__Format = 30
	QImage__Format_RGBA16FPx4               QImage__Format = 31
	QImage__Format_RGBA16FPx4_Premultiplied QImage__Format = 32
	QImage__Format_RGBX32FPx4               QImage__Format = 33
	QImage__Format_RGBA32FPx4               QImage__Format = 34
	QImage__Format_RGBA32FPx4_Premultiplied QImage__Format = 35
	QImage__NImageFormats                   QImage__Format = 36
)

type QImage struct {
	h *C.QImage
	*QPaintDevice
}

func (this *QImage) cPointer() *C.QImage {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QImage) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

func newQImage(h *C.QImage) *QImage {
	if h == nil {
		return nil
	}
	return &QImage{h: h, QPaintDevice: UnsafeNewQPaintDevice(unsafe.Pointer(h))}
}

func UnsafeNewQImage(h unsafe.Pointer) *QImage {
	return newQImage((*C.QImage)(h))
}

// NewQImage constructs a new QImage object.
func NewQImage() *QImage {
	ret := C.QImage_new()
	return newQImage(ret)
}

// NewQImage2 constructs a new QImage object.
func NewQImage2(size *QSize, format QImage__Format) *QImage {
	ret := C.QImage_new2(size.cPointer(), (C.int)(format))
	return newQImage(ret)
}

// NewQImage3 constructs a new QImage object.
func NewQImage3(width int, height int, format QImage__Format) *QImage {
	ret := C.QImage_new3((C.int)(width), (C.int)(height), (C.int)(format))
	return newQImage(ret)
}

// NewQImage4 constructs a new QImage object.
func NewQImage4(data *byte, width int, height int, format QImage__Format) *QImage {
	ret := C.QImage_new4((*C.uchar)(unsafe.Pointer(data)), (C.int)(width), (C.int)(height), (C.int)(format))
	return newQImage(ret)
}

// NewQImage5 constructs a new QImage object.
func NewQImage5(data *byte, width int, height int, format QImage__Format) *QImage {
	ret := C.QImage_new5((*C.uchar)(unsafe.Pointer(data)), (C.int)(width), (C.int)(height), (C.int)(format))
	return newQImage(ret)
}

// NewQImage6 constructs a new QImage object.
func NewQImage6(data *byte, width int, height int, bytesPerLine int64, format QImage__Format) *QImage {
	ret := C.QImage_new6((*C.uchar)(unsafe.Pointer(data)), (C.int)(width), (C.int)(height), (C.ptrdiff_t)(bytesPerLine), (C.int)(format))
	return newQImage(ret)
}

// NewQImage7 constructs a new QImage object.
func NewQImage7(data *byte, width int, height int, bytesPerLine int64, format QImage__Format) *QImage {
	ret := C.QImage_new7((*C.uchar)(unsafe.Pointer(data)), (C.int)(width), (C.int)(height), (C.ptrdiff_t)(bytesPerLine), (C.int)(format))
	return newQImage(ret)
}

// NewQImage8 constructs a new QImage object.
func NewQImage8(fileName string) *QImage {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	ret := C.QImage_new8(fileName_ms)
	return newQImage(ret)
}

// NewQImage9 constructs a new QImage object.
func NewQImage9(param1 *QImage) *QImage {
	ret := C.QImage_new9(param1.cPointer())
	return newQImage(ret)
}

// NewQImage10 constructs a new QImage object.
func NewQImage10(fileName string, format string) *QImage {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	ret := C.QImage_new10(fileName_ms, format_Cstring)
	return newQImage(ret)
}

func (this *QImage) OperatorAssign(param1 *QImage) {
	C.QImage_OperatorAssign(this.h, param1.cPointer())
}

func (this *QImage) Swap(other *QImage) {
	C.QImage_Swap(this.h, other.cPointer())
}

func (this *QImage) IsNull() bool {
	return (bool)(C.QImage_IsNull(this.h))
}

func (this *QImage) DevType() int {
	return (int)(C.QImage_DevType(this.h))
}

func (this *QImage) OperatorEqual(param1 *QImage) bool {
	return (bool)(C.QImage_OperatorEqual(this.h, param1.cPointer()))
}

func (this *QImage) OperatorNotEqual(param1 *QImage) bool {
	return (bool)(C.QImage_OperatorNotEqual(this.h, param1.cPointer()))
}

func (this *QImage) Detach() {
	C.QImage_Detach(this.h)
}

func (this *QImage) IsDetached() bool {
	return (bool)(C.QImage_IsDetached(this.h))
}

func (this *QImage) Copy() *QImage {
	_ret := C.QImage_Copy(this.h)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Copy2(x int, y int, w int, h int) *QImage {
	_ret := C.QImage_Copy2(this.h, (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Format() QImage__Format {
	return (QImage__Format)(C.QImage_Format(this.h))
}

func (this *QImage) ConvertToFormat(f QImage__Format) *QImage {
	_ret := C.QImage_ConvertToFormat(this.h, (C.int)(f))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertToFormat2(f QImage__Format, colorTable []uint) *QImage {
	colorTable_CArray := (*[0xffff]C.uint)(C.malloc(C.size_t(8 * len(colorTable))))
	defer C.free(unsafe.Pointer(colorTable_CArray))
	for i := range colorTable {
		colorTable_CArray[i] = (C.uint)(colorTable[i])
	}
	colorTable_ma := C.struct_miqt_array{len: C.size_t(len(colorTable)), data: unsafe.Pointer(colorTable_CArray)}
	_ret := C.QImage_ConvertToFormat2(this.h, (C.int)(f), colorTable_ma)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ReinterpretAsFormat(f QImage__Format) bool {
	return (bool)(C.QImage_ReinterpretAsFormat(this.h, (C.int)(f)))
}

func (this *QImage) ConvertedTo(f QImage__Format) *QImage {
	_ret := C.QImage_ConvertedTo(this.h, (C.int)(f))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertTo(f QImage__Format) {
	C.QImage_ConvertTo(this.h, (C.int)(f))
}

func (this *QImage) Width() int {
	return (int)(C.QImage_Width(this.h))
}

func (this *QImage) Height() int {
	return (int)(C.QImage_Height(this.h))
}

func (this *QImage) Size() *QSize {
	_ret := C.QImage_Size(this.h)
	_goptr := newQSize(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Rect() *QRect {
	_ret := C.QImage_Rect(this.h)
	_goptr := newQRect(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Depth() int {
	return (int)(C.QImage_Depth(this.h))
}

func (this *QImage) ColorCount() int {
	return (int)(C.QImage_ColorCount(this.h))
}

func (this *QImage) BitPlaneCount() int {
	return (int)(C.QImage_BitPlaneCount(this.h))
}

func (this *QImage) Color(i int) uint {
	return (uint)(C.QImage_Color(this.h, (C.int)(i)))
}

func (this *QImage) SetColor(i int, c uint) {
	C.QImage_SetColor(this.h, (C.int)(i), (C.uint)(c))
}

func (this *QImage) SetColorCount(colorCount int) {
	C.QImage_SetColorCount(this.h, (C.int)(colorCount))
}

func (this *QImage) AllGray() bool {
	return (bool)(C.QImage_AllGray(this.h))
}

func (this *QImage) IsGrayscale() bool {
	return (bool)(C.QImage_IsGrayscale(this.h))
}

func (this *QImage) Bits() *byte {
	return (*byte)(C.QImage_Bits(this.h))
}

func (this *QImage) Bits2() *byte {
	return (*byte)(C.QImage_Bits2(this.h))
}

func (this *QImage) ConstBits() *byte {
	return (*byte)(C.QImage_ConstBits(this.h))
}

func (this *QImage) SizeInBytes() int64 {
	return (int64)(C.QImage_SizeInBytes(this.h))
}

func (this *QImage) ScanLine(param1 int) *byte {
	return (*byte)(C.QImage_ScanLine(this.h, (C.int)(param1)))
}

func (this *QImage) ScanLineWithInt(param1 int) *byte {
	return (*byte)(C.QImage_ScanLineWithInt(this.h, (C.int)(param1)))
}

func (this *QImage) ConstScanLine(param1 int) *byte {
	return (*byte)(C.QImage_ConstScanLine(this.h, (C.int)(param1)))
}

func (this *QImage) BytesPerLine() int64 {
	return (int64)(C.QImage_BytesPerLine(this.h))
}

func (this *QImage) Valid(x int, y int) bool {
	return (bool)(C.QImage_Valid(this.h, (C.int)(x), (C.int)(y)))
}

func (this *QImage) ValidWithPt(pt *QPoint) bool {
	return (bool)(C.QImage_ValidWithPt(this.h, pt.cPointer()))
}

func (this *QImage) PixelIndex(x int, y int) int {
	return (int)(C.QImage_PixelIndex(this.h, (C.int)(x), (C.int)(y)))
}

func (this *QImage) PixelIndexWithPt(pt *QPoint) int {
	return (int)(C.QImage_PixelIndexWithPt(this.h, pt.cPointer()))
}

func (this *QImage) Pixel(x int, y int) uint {
	return (uint)(C.QImage_Pixel(this.h, (C.int)(x), (C.int)(y)))
}

func (this *QImage) PixelWithPt(pt *QPoint) uint {
	return (uint)(C.QImage_PixelWithPt(this.h, pt.cPointer()))
}

func (this *QImage) SetPixel(x int, y int, index_or_rgb uint) {
	C.QImage_SetPixel(this.h, (C.int)(x), (C.int)(y), (C.uint)(index_or_rgb))
}

func (this *QImage) SetPixel2(pt *QPoint, index_or_rgb uint) {
	C.QImage_SetPixel2(this.h, pt.cPointer(), (C.uint)(index_or_rgb))
}

func (this *QImage) PixelColor(x int, y int) *QColor {
	_ret := C.QImage_PixelColor(this.h, (C.int)(x), (C.int)(y))
	_goptr := newQColor(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) PixelColorWithPt(pt *QPoint) *QColor {
	_ret := C.QImage_PixelColorWithPt(this.h, pt.cPointer())
	_goptr := newQColor(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) SetPixelColor(x int, y int, c *QColor) {
	C.QImage_SetPixelColor(this.h, (C.int)(x), (C.int)(y), c.cPointer())
}

func (this *QImage) SetPixelColor2(pt *QPoint, c *QColor) {
	C.QImage_SetPixelColor2(this.h, pt.cPointer(), c.cPointer())
}

func (this *QImage) ColorTable() []uint {
	var _ma C.struct_miqt_array = C.QImage_ColorTable(this.h)
	_ret := make([]uint, int(_ma.len))
	_outCast := (*[0xffff]C.uint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (uint)(_outCast[i])
	}
	return _ret
}

func (this *QImage) SetColorTable(colors []uint) {
	colors_CArray := (*[0xffff]C.uint)(C.malloc(C.size_t(8 * len(colors))))
	defer C.free(unsafe.Pointer(colors_CArray))
	for i := range colors {
		colors_CArray[i] = (C.uint)(colors[i])
	}
	colors_ma := C.struct_miqt_array{len: C.size_t(len(colors)), data: unsafe.Pointer(colors_CArray)}
	C.QImage_SetColorTable(this.h, colors_ma)
}

func (this *QImage) DevicePixelRatio() float64 {
	return (float64)(C.QImage_DevicePixelRatio(this.h))
}

func (this *QImage) SetDevicePixelRatio(scaleFactor float64) {
	C.QImage_SetDevicePixelRatio(this.h, (C.double)(scaleFactor))
}

func (this *QImage) DeviceIndependentSize() *QSizeF {
	_ret := C.QImage_DeviceIndependentSize(this.h)
	_goptr := newQSizeF(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Fill(pixel uint) {
	C.QImage_Fill(this.h, (C.uint)(pixel))
}

func (this *QImage) FillWithColor(color *QColor) {
	C.QImage_FillWithColor(this.h, color.cPointer())
}

func (this *QImage) Fill2(color GlobalColor) {
	C.QImage_Fill2(this.h, (C.int)(color))
}

func (this *QImage) HasAlphaChannel() bool {
	return (bool)(C.QImage_HasAlphaChannel(this.h))
}

func (this *QImage) SetAlphaChannel(alphaChannel *QImage) {
	C.QImage_SetAlphaChannel(this.h, alphaChannel.cPointer())
}

func (this *QImage) CreateAlphaMask() *QImage {
	_ret := C.QImage_CreateAlphaMask(this.h)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) CreateHeuristicMask() *QImage {
	_ret := C.QImage_CreateHeuristicMask(this.h)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) CreateMaskFromColor(color uint) *QImage {
	_ret := C.QImage_CreateMaskFromColor(this.h, (C.uint)(color))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Scaled(w int, h int) *QImage {
	_ret := C.QImage_Scaled(this.h, (C.int)(w), (C.int)(h))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ScaledWithQSize(s *QSize) *QImage {
	_ret := C.QImage_ScaledWithQSize(this.h, s.cPointer())
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ScaledToWidth(w int) *QImage {
	_ret := C.QImage_ScaledToWidth(this.h, (C.int)(w))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ScaledToHeight(h int) *QImage {
	_ret := C.QImage_ScaledToHeight(this.h, (C.int)(h))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Transformed(matrix *QTransform) *QImage {
	_ret := C.QImage_Transformed(this.h, matrix.cPointer())
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_TrueMatrix(param1 *QTransform, w int, h int) *QTransform {
	_ret := C.QImage_TrueMatrix(param1.cPointer(), (C.int)(w), (C.int)(h))
	_goptr := newQTransform(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Mirrored() *QImage {
	_ret := C.QImage_Mirrored(this.h)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) RgbSwapped() *QImage {
	_ret := C.QImage_RgbSwapped(this.h)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Mirror() {
	C.QImage_Mirror(this.h)
}

func (this *QImage) RgbSwap() {
	C.QImage_RgbSwap(this.h)
}

func (this *QImage) InvertPixels() {
	C.QImage_InvertPixels(this.h)
}

func (this *QImage) ColorSpace() *QColorSpace {
	_ret := C.QImage_ColorSpace(this.h)
	_goptr := newQColorSpace(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertedToColorSpace(param1 *QColorSpace) *QImage {
	_ret := C.QImage_ConvertedToColorSpace(this.h, param1.cPointer())
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertToColorSpace(param1 *QColorSpace) {
	C.QImage_ConvertToColorSpace(this.h, param1.cPointer())
}

func (this *QImage) SetColorSpace(colorSpace *QColorSpace) {
	C.QImage_SetColorSpace(this.h, colorSpace.cPointer())
}

func (this *QImage) ColorTransformed(transform *QColorTransform) *QImage {
	_ret := C.QImage_ColorTransformed(this.h, transform.cPointer())
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ApplyColorTransform(transform *QColorTransform) {
	C.QImage_ApplyColorTransform(this.h, transform.cPointer())
}

func (this *QImage) Load(device *QIODevice, format string) bool {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_Load(this.h, device.cPointer(), format_Cstring))
}

func (this *QImage) LoadWithFileName(fileName string) bool {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	return (bool)(C.QImage_LoadWithFileName(this.h, fileName_ms))
}

func (this *QImage) LoadFromData(data QByteArrayView) bool {
	return (bool)(C.QImage_LoadFromData(this.h, data.cPointer()))
}

func (this *QImage) LoadFromData2(buf *byte, lenVal int) bool {
	return (bool)(C.QImage_LoadFromData2(this.h, (*C.uchar)(unsafe.Pointer(buf)), (C.int)(lenVal)))
}

func (this *QImage) LoadFromDataWithData(data []byte) bool {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	return (bool)(C.QImage_LoadFromDataWithData(this.h, data_alias))
}

func (this *QImage) Save(fileName string) bool {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	return (bool)(C.QImage_Save(this.h, fileName_ms))
}

func (this *QImage) SaveWithDevice(device *QIODevice) bool {
	return (bool)(C.QImage_SaveWithDevice(this.h, device.cPointer()))
}

func QImage_FromData(data QByteArrayView) *QImage {
	_ret := C.QImage_FromData(data.cPointer())
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_FromData2(data *byte, size int) *QImage {
	_ret := C.QImage_FromData2((*C.uchar)(unsafe.Pointer(data)), (C.int)(size))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_FromDataWithData(data []byte) *QImage {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	_ret := C.QImage_FromDataWithData(data_alias)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) CacheKey() int64 {
	return (int64)(C.QImage_CacheKey(this.h))
}

func (this *QImage) PaintEngine() *QPaintEngine {
	return UnsafeNewQPaintEngine(unsafe.Pointer(C.QImage_PaintEngine(this.h)))
}

func (this *QImage) DotsPerMeterX() int {
	return (int)(C.QImage_DotsPerMeterX(this.h))
}

func (this *QImage) DotsPerMeterY() int {
	return (int)(C.QImage_DotsPerMeterY(this.h))
}

func (this *QImage) SetDotsPerMeterX(dotsPerMeterX int) {
	C.QImage_SetDotsPerMeterX(this.h, (C.int)(dotsPerMeterX))
}

func (this *QImage) SetDotsPerMeterY(dotsPerMeterY int) {
	C.QImage_SetDotsPerMeterY(this.h, (C.int)(dotsPerMeterY))
}

func (this *QImage) Offset() *QPoint {
	_ret := C.QImage_Offset(this.h)
	_goptr := newQPoint(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) SetOffset(offset *QPoint) {
	C.QImage_SetOffset(this.h, offset.cPointer())
}

func (this *QImage) TextKeys() []string {
	var _ma C.struct_miqt_array = C.QImage_TextKeys(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QImage) Text() string {
	var _ms C.struct_miqt_string = C.QImage_Text(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImage) SetText(key string, value string) {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	value_ms := C.struct_miqt_string{}
	value_ms.data = C.CString(value)
	value_ms.len = C.size_t(len(value))
	defer C.free(unsafe.Pointer(value_ms.data))
	C.QImage_SetText(this.h, key_ms, value_ms)
}

func (this *QImage) PixelFormat() *QPixelFormat {
	_ret := C.QImage_PixelFormat(this.h)
	_goptr := newQPixelFormat(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_ToPixelFormat(format QImage__Format) *QPixelFormat {
	_ret := C.QImage_ToPixelFormat((C.int)(format))
	_goptr := newQPixelFormat(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_ToImageFormat(format QPixelFormat) QImage__Format {
	return (QImage__Format)(C.QImage_ToImageFormat(format.cPointer()))
}

func (this *QImage) Copy1(rect *QRect) *QImage {
	_ret := C.QImage_Copy1(this.h, rect.cPointer())
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertToFormat22(f QImage__Format, flags ImageConversionFlag) *QImage {
	_ret := C.QImage_ConvertToFormat22(this.h, (C.int)(f), (C.int)(flags))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertToFormat3(f QImage__Format, colorTable []uint, flags ImageConversionFlag) *QImage {
	colorTable_CArray := (*[0xffff]C.uint)(C.malloc(C.size_t(8 * len(colorTable))))
	defer C.free(unsafe.Pointer(colorTable_CArray))
	for i := range colorTable {
		colorTable_CArray[i] = (C.uint)(colorTable[i])
	}
	colorTable_ma := C.struct_miqt_array{len: C.size_t(len(colorTable)), data: unsafe.Pointer(colorTable_CArray)}
	_ret := C.QImage_ConvertToFormat3(this.h, (C.int)(f), colorTable_ma, (C.int)(flags))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertedTo2(f QImage__Format, flags ImageConversionFlag) *QImage {
	_ret := C.QImage_ConvertedTo2(this.h, (C.int)(f), (C.int)(flags))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ConvertTo2(f QImage__Format, flags ImageConversionFlag) {
	C.QImage_ConvertTo2(this.h, (C.int)(f), (C.int)(flags))
}

func (this *QImage) CreateAlphaMask1(flags ImageConversionFlag) *QImage {
	_ret := C.QImage_CreateAlphaMask1(this.h, (C.int)(flags))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) CreateHeuristicMask1(clipTight bool) *QImage {
	_ret := C.QImage_CreateHeuristicMask1(this.h, (C.bool)(clipTight))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) CreateMaskFromColor2(color uint, mode MaskMode) *QImage {
	_ret := C.QImage_CreateMaskFromColor2(this.h, (C.uint)(color), (C.int)(mode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Scaled3(w int, h int, aspectMode AspectRatioMode) *QImage {
	_ret := C.QImage_Scaled3(this.h, (C.int)(w), (C.int)(h), (C.int)(aspectMode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Scaled4(w int, h int, aspectMode AspectRatioMode, mode TransformationMode) *QImage {
	_ret := C.QImage_Scaled4(this.h, (C.int)(w), (C.int)(h), (C.int)(aspectMode), (C.int)(mode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Scaled2(s *QSize, aspectMode AspectRatioMode) *QImage {
	_ret := C.QImage_Scaled2(this.h, s.cPointer(), (C.int)(aspectMode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Scaled32(s *QSize, aspectMode AspectRatioMode, mode TransformationMode) *QImage {
	_ret := C.QImage_Scaled32(this.h, s.cPointer(), (C.int)(aspectMode), (C.int)(mode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ScaledToWidth2(w int, mode TransformationMode) *QImage {
	_ret := C.QImage_ScaledToWidth2(this.h, (C.int)(w), (C.int)(mode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) ScaledToHeight2(h int, mode TransformationMode) *QImage {
	_ret := C.QImage_ScaledToHeight2(this.h, (C.int)(h), (C.int)(mode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Transformed2(matrix *QTransform, mode TransformationMode) *QImage {
	_ret := C.QImage_Transformed2(this.h, matrix.cPointer(), (C.int)(mode))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Mirrored1(horizontally bool) *QImage {
	_ret := C.QImage_Mirrored1(this.h, (C.bool)(horizontally))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Mirrored2(horizontally bool, vertically bool) *QImage {
	_ret := C.QImage_Mirrored2(this.h, (C.bool)(horizontally), (C.bool)(vertically))
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Mirror1(horizontally bool) {
	C.QImage_Mirror1(this.h, (C.bool)(horizontally))
}

func (this *QImage) Mirror2(horizontally bool, vertically bool) {
	C.QImage_Mirror2(this.h, (C.bool)(horizontally), (C.bool)(vertically))
}

func (this *QImage) InvertPixels1(param1 QImage__InvertMode) {
	C.QImage_InvertPixels1(this.h, (C.int)(param1))
}

func (this *QImage) Load2(fileName string, format string) bool {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_Load2(this.h, fileName_ms, format_Cstring))
}

func (this *QImage) LoadFromData22(data QByteArrayView, format string) bool {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_LoadFromData22(this.h, data.cPointer(), format_Cstring))
}

func (this *QImage) LoadFromData3(buf *byte, lenVal int, format string) bool {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_LoadFromData3(this.h, (*C.uchar)(unsafe.Pointer(buf)), (C.int)(lenVal), format_Cstring))
}

func (this *QImage) LoadFromData23(data []byte, format string) bool {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_LoadFromData23(this.h, data_alias, format_Cstring))
}

func (this *QImage) Save2(fileName string, format string) bool {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_Save2(this.h, fileName_ms, format_Cstring))
}

func (this *QImage) Save3(fileName string, format string, quality int) bool {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_Save3(this.h, fileName_ms, format_Cstring, (C.int)(quality)))
}

func (this *QImage) Save22(device *QIODevice, format string) bool {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_Save22(this.h, device.cPointer(), format_Cstring))
}

func (this *QImage) Save32(device *QIODevice, format string, quality int) bool {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	return (bool)(C.QImage_Save32(this.h, device.cPointer(), format_Cstring, (C.int)(quality)))
}

func QImage_FromData22(data QByteArrayView, format string) *QImage {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	_ret := C.QImage_FromData22(data.cPointer(), format_Cstring)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_FromData3(data *byte, size int, format string) *QImage {
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	_ret := C.QImage_FromData3((*C.uchar)(unsafe.Pointer(data)), (C.int)(size), format_Cstring)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImage_FromData23(data []byte, format string) *QImage {
	data_alias := C.struct_miqt_string{}
	data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	data_alias.len = C.size_t(len(data))
	format_Cstring := C.CString(format)
	defer C.free(unsafe.Pointer(format_Cstring))
	_ret := C.QImage_FromData23(data_alias, format_Cstring)
	_goptr := newQImage(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImage) Text1(key string) string {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	var _ms C.struct_miqt_string = C.QImage_Text1(this.h, key_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QImage) Delete() {
	C.QImage_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QImage) GoGC() {
	runtime.SetFinalizer(this, func(this *QImage) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
