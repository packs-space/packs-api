package block_processor

import (
	tendermint "github.com/tendermint/tendermint/types"
)

type Processor struct {
}

func NewBlockProcessor() BlockProcessor {
	return &Processor{}
}

func (bp *Processor) Process(block *tendermint.Block) error {
	// @todo process block
	return nil
}
