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

func PdfParse(file string, start, end int)  {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	cstart := C.int(start)
	cend := C.int(end)

	result := int(C.mypdf(cfile, cstart, cend))

	common.Logger.Printf("parse file %s result %d", file, result)
}