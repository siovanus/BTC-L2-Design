package main

import (
	"github.com/prysmaticlabs/prysm/encoding/bytesutil"
)

var (
	PREFIX = []byte{}
)

type BTCBlock struct {
	Hash []byte
}

func GetBtcBlock(h uint64) *BTCBlock {
	return &BTCBlock{}
}

func Hash(data []byte) []byte {
	return nil
}

func GetSeed(epoch uint64) []byte {
	blockHash := GetBtcBlock(epoch - 6).Hash
	seed := append(PREFIX, blockHash...)
	seed = append(seed, bytesutil.Bytes8(epoch)...)

	seed32 := Hash(seed)
	return seed32
}

func GetProposerSelectionSeed(seed []byte, slot uint64) []byte {

	data := append(PREFIX, seed...)
	data = append(data, bytesutil.Bytes8(slot)...)

	v := Hash(data)
	return v
}

func GetProposerIndex(vrf []byte, totalStake uint64) uint64 {
	h := Hash(vrf)
	hash8 := h[:8]
	hash8Int := bytesutil.FromBytes8(hash8)
	index := hash8Int % totalStake

	return index
}

type Validator struct {
	Stake uint64
}

func GetProposer(index uint64, validatorList []*Validator) *Validator {
	var sum uint64
	for _, v := range validatorList {
		sum := sum + v.Stake
		if sum <= index {
			return v
		}
	}
	panic("proposer index overflow")
}
