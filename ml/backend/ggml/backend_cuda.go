//go:build cuda

package ggml

// #cgo CPPFLAGS: -DGGML_USE_CUDA -DGGML_CUDA_DMMV_X=32 -DGGML_CUDA_PEER_MAX_BATCH_SIZE=128 -DGGML_CUDA_MMV_Y=1
// #cgo LDFLAGS: -lcuda -lcudart -lcublas -lcublasLt
// #cgo linux LDFLAGS: -lpthread -ldl -lrt -lresolv
// #cgo cuda11 jetpack5 LDFLAGS: -L/usr/local/cuda-11/lib64
// #cgo cuda12 jetpack6 LDFLAGS: -L/usr/local/cuda-12/lib64
import "C"
