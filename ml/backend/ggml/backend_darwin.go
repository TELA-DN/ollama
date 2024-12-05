package ggml

//go:generate sh -c "echo \"// Code generated $(date). DO NOT EDIT.\n\" >ggml-metal-embed.metal"
//go:generate sh -c "sed -e '/#include \"ggml-common.h\"/r ggml-common.h' -e '/#include \"ggml-common.h\"/d' ggml-metal.metal >>ggml-metal-embed.metal"

// #cgo CPPFLAGS: -Wno-incompatible-pointer-types-discards-qualifiers
// #cgo LDFLAGS: -framework Foundation
// #cgo amd64,avx2 CPPFLAGS: -DGGML_USE_ACCELERATE -DACCELERATE_USE_LAPACK -DACCELERATE_LAPACK_ILP64
// #cgo amd64,avx2 LDFLAGS: -framework Accelerate
// #cgo arm64 CPPFLAGS: -DGGML_USE_METAL -DGGML_METAL_EMBED_LIBRARY -DGGML_USE_ACCELERATE -DACCELERATE_NEW_LAPACK -DACCELERATE_LAPACK_ILP64 -DGGML_USE_BLAS -DGGML_METAL_NDEBUG
// #cgo arm64 LDFLAGS: -framework Metal -framework MetalKit -framework Accelerate
// #include "ggml-metal.h"
import "C"

func newBackend() *C.struct_ggml_backend {
	return C.ggml_backend_metal_init()
}
