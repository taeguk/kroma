package derive

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type UserDepositSource struct {
	L1BlockHash common.Hash
	LogIndex    uint64
}

const (
	UserDepositSourceDomain      = 0
	L1InfoDepositSourceDomain    = 1
	MintTokenDepositSourceDomain = 10
)

func (dep *UserDepositSource) SourceHash() common.Hash {
	var input [32 * 2]byte
	copy(input[:32], dep.L1BlockHash[:])
	binary.BigEndian.PutUint64(input[32*2-8:], dep.LogIndex)
	depositIDHash := crypto.Keccak256Hash(input[:])
	var domainInput [32 * 2]byte
	binary.BigEndian.PutUint64(domainInput[32-8:32], UserDepositSourceDomain)
	copy(domainInput[32:], depositIDHash[:])
	return crypto.Keccak256Hash(domainInput[:])
}

type L1InfoDepositSource struct {
	L1BlockHash common.Hash
	SeqNumber   uint64
}

func (dep *L1InfoDepositSource) SourceHash() common.Hash {
	var input [32 * 2]byte
	copy(input[:32], dep.L1BlockHash[:])
	binary.BigEndian.PutUint64(input[32*2-8:], dep.SeqNumber)
	depositIDHash := crypto.Keccak256Hash(input[:])

	var domainInput [32 * 2]byte
	binary.BigEndian.PutUint64(domainInput[32-8:32], L1InfoDepositSourceDomain)
	copy(domainInput[32:], depositIDHash[:])
	return crypto.Keccak256Hash(domainInput[:])
}

type MintTokenDepositSource struct {
	L2ParentHash uint64
}

func (dep *MintTokenDepositSource) SourceHash() common.Hash {
	var input [32 * 2]byte
	// binary.BigEndian.PutUint64(input[32-8:], dep.ParentGasUsed)
	// binary.BigEndian.PutUint64(input[32*2-8:], dep.ParentGasTarget)
	depositIDHash := crypto.Keccak256Hash(input[:])

	var domainInput [32 * 2]byte
	binary.BigEndian.PutUint64(domainInput[32-8:32], MintTokenDepositSourceDomain)
	copy(domainInput[32:], depositIDHash[:])
	return crypto.Keccak256Hash(domainInput[:])
}
