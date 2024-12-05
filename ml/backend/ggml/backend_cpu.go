package ggml

// #cgo amd64,avx CPPFLAGS: -mavx
// #cgo amd64,avx2 CPPFLAGS: -mavx2 -mfma
// #cgo amd64,f16c CPPFLAGS: -mf16c
// #include "ggml-backend.h"
import "C"

func newCPUBackend() *C.struct_ggml_backend {
	return C.ggml_backend_cpu_init()
}
