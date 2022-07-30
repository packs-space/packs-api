package block_processor

import (
	tendermint "github.com/tendermint/tendermint/types"
)

type BlockProcessor interface {
	Process(*tendermint.Block) error
}
