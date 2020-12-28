package core

// #include <stdlib.h>
// #include "mypdf.c"
// #cgo CFLAGS: -g -Wall -I../libs -std=c99
// #cgo LDFLAGS: /usr/local/lib/libimagequant.a /usr/local/lib/libmupdf.a /usr/local/lib/libmupdf-third.a -lm
import "C"
import (
	"app/common"
	"unsafe"
)

func PngCompress(file string) int {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	return int(C.png_compress(cfile))
}

func PdfSize(file string) int {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	return int(C.mypdf_size(cfile))
}

func PdfParse(file string, zoom, start, end, compress int)  {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	czoom := C.int(zoom)
	cstart := C.int(start)
	cend := C.int(end)
	ccompress := C.int(compress)

	result := int(C.mypdf_parse(cfile, czoom, cstart, cend, ccompress))

	common.Logger.Printf("parse file %s result %d", file, result)
}