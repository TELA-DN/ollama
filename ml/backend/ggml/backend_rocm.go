//go:build rocm

package ggml

// #cgo CPPFLAGS: -DGGML_USE_CUDA -DGGML_USE_HIPBLAS -DGGML_CUDA_DMMV_X=32 -DGGML_CUDA_PEER_MAX_BATCH_SIZE=128 -DGGML_CUDA_MMV_Y=1
// #cgo LDFLAGS: -L/opt/rocm/lib -lhipblas -lamdhip64 -lrocblas
// #cgo linux LDFLAGS: -lpthread -ldl -lrt -lresolv
import "C"
