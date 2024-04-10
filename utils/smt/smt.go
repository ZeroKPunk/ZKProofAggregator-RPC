package smt

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: ${SRCDIR}/lib/libsparse_merkle_proof_helper.so
#include "sparse_merkle_proof.h"

// Convert Go byte slices to C-style arrays
static inline uint8_t** convertToCArray(const uint8_t *s[], int len) {
    return (uint8_t**)s;
}
*/
import "C"
import (
	"github.com/ethereum/go-ethereum/common"
	"unsafe"
)

type SparseMerkleProofInput struct {
	leavesKeys   []common.Hash
	leavesValues []common.Hash
	index        uint
}

type MerkleProof struct {
	root     common.Hash
	index    uint
	value    common.Hash
	siblings []common.Hash
}

func (smt SparseMerkleProofInput) getLeavesKeysAndValues() ([][]byte, [][]byte) {
	keys := make([][]byte, len(smt.leavesKeys))
	values := make([][]byte, len(smt.leavesValues))
	for k, v := range smt.leavesKeys {
		keys[k] = v.Bytes()
	}
	for k, v := range smt.leavesValues {
		values[k] = v.Bytes()
	}

	return keys, values
}

func (smt SparseMerkleProofInput) getMerkleProof(root, value []byte, index uint, sib [][]byte) MerkleProof {
	siblings := make([]common.Hash, len(sib))
	for k, s := range sib {
		siblings[k] = common.BytesToHash(s)
	}
	return MerkleProof{
		root:     common.BytesToHash(root),
		index:    index,
		value:    common.BytesToHash(value),
		siblings: siblings,
	}
}

func (smt SparseMerkleProofInput) GenerateSparseMerkleProof() MerkleProof {

	keys, values := smt.getLeavesKeysAndValues()
	index := C.uintptr_t(2)

	// Convert Go byte slices to C-style arrays
	cKeys := make([]*C.uint8_t, len(keys))
	cValues := make([]*C.uint8_t, len(values))
	for i, key := range keys {
		cKeys[i] = (*C.uint8_t)(unsafe.Pointer(C.CBytes(key)))
	}
	for i, value := range values {
		cValues[i] = (*C.uint8_t)(unsafe.Pointer(C.CBytes(value)))
	}

	proof := C.generate_sparse_merkle_proof(
		C.convertToCArray(&cKeys[0], C.int(len(cKeys))),
		C.uintptr_t(len(keys)),
		C.convertToCArray(&cValues[0], C.int(len(cValues))),
		C.uintptr_t(len(values)),
		index,
	)

	var valueSlice []byte
	if proof.value.data != nil {
		valueSlice = C.GoBytes(unsafe.Pointer(proof.value.data), C.int(proof.value.len))
	}

	siblings := make([][]byte, proof.siblings.len)
	siblingsPtr := (*[1 << 30]C.Bytes)(unsafe.Pointer(proof.siblings.data))[:proof.siblings.len:proof.siblings.len]
	for i := 0; i < int(proof.siblings.len); i++ {
		siblingSlice := C.GoBytes(unsafe.Pointer(siblingsPtr[i].data), C.int(siblingsPtr[i].len))
		siblings[i] = siblingSlice
	}

	var rootSlice []byte
	if proof.root.data != nil {
		rootSlice = C.GoBytes(unsafe.Pointer(proof.root.data), C.int(proof.root.len))
	}

	for _, key := range cKeys {
		C.free(unsafe.Pointer(key))
	}
	for _, value := range cValues {
		C.free(unsafe.Pointer(value))
	}

	return smt.getMerkleProof(rootSlice, valueSlice, uint(proof.index), siblings)
}
