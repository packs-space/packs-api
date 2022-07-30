package main

import (
	"fmt"

	blockFeeder "github.com/packs-space/packs-api/block_feed"
	blockProcessor "github.com/packs-space/packs-api/block_processor"
	"github.com/packs-space/packs-api/config"
)

func main() {
	apiConfig := config.NewConfig()
	apiConfig.Print()

	blockFeed := blockFeeder.NewAggregateBlockFeed(
		apiConfig.CurrentBlock,
		apiConfig.RPCEndpoints,
		apiConfig.WSEndpoints,
	)

	processor := blockProcessor.NewBlockProcessor()

	// start subscribing to block
	if apiConfig.DisableSync {
		fmt.Println("running without sync...")
		forever()
	} else if cBlockFeed, blockFeedErr := blockFeed.Subscribe(0); blockFeedErr != nil {
		panic(blockFeedErr)
	} else {
		for {
			feed := <-cBlockFeed

			fmt.Printf("new block: %v", feed)

			processor.Process(feed.Block)
		}
	}
}

func forever() {
	<-(chan int)(nil)
}
