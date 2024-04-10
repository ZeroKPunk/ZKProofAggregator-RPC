package smt

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSparseMerkleProofInput_GenerateSparseMerkleProof(t *testing.T) {
	leavesKeys := []common.Hash{
		common.BigToHash(big.NewInt(1)),
		common.BigToHash(big.NewInt(2)),
		common.BigToHash(big.NewInt(3)),
	}
	leavesValues := []common.Hash{
		common.BigToHash(big.NewInt(2)),
		common.BigToHash(big.NewInt(10)),
		common.BigToHash(big.NewInt(51)),
	}
	input := SparseMerkleProofInput{
		leavesKeys:   leavesKeys,
		leavesValues: leavesValues,
		index:        2,
	}
	merkleProof := input.GenerateSparseMerkleProof()

	fmt.Printf("Index: %d\n", merkleProof.index)
	fmt.Printf("Value: %x\n", merkleProof.value)
	fmt.Printf("Siblings: %x\n", merkleProof.siblings)
	fmt.Printf("Root: %x\n", merkleProof.root)

	assert.Equal(t, merkleProof.index, uint(3))
	assert.Equal(t, merkleProof.value, common.HexToHash("789f6365f94749dcc6c41d74a7dd3a8ef1b2384715577c73fc52025dbc549e13"))
	assert.Equal(t, merkleProof.root, common.HexToHash("22ee6ab614c8c9cf484420c59fa9cef2bf9a8536095c5c5572ac1b6918c0e705"))
	assert.Equal(t, merkleProof.siblings, []common.Hash{common.HexToHash("71ab6f4175312c78ee0bc776ab91b0e6f86356e7f44807b6856a894231d4d98d"), common.HexToHash("9e356f158483696de5f958ab07e54c06e9d16f32eba5c9d5895f72a2d8e0dd44")})
}
