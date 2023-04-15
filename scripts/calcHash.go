package main


import (
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
)



func CalculateHash(str string) common.Hash {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	var hashResult common.Hash
	copy(hashResult[:], hashBytes[:])
	return hashResult
}




