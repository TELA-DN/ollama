package ggml

// #cgo CPPFLAGS: -D_WIN32_WINNT=0x602
// #cgo LDFLAGS: -lmsvcrt -static -static-libgcc -static-libstdc++
// #cgo arm64 CPPFLAGS: -D__aarch64__ -D__ARM_NEON -D__ARM_FEATURE_FMA
// #include "ggml-backend.h"
import "C"

func newBackend() *C.struct_ggml_backend {
	return newCPUBackend()
}
