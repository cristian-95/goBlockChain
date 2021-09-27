package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Tipo Block: cada bloco contem dados que serão escritos no blockchain
type Block struct {
	Index     int    // posição do dado escrito no blockchain
	Timestamp string // o horario em que o dado é escrito
	BPM       int    // batidas de coração por minuto,neste caso a informação aser gravada
	Hash      string // Hash é um identificador SHA256 que representa a gravação dos dados
	PrevHash  string // identificador que representa a gravação anterior
}

var BlockChain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func gennerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(BlockChain) {
		BlockChain = newBlocks
	}
}

func main() {

}
