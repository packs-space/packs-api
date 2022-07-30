package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ChainID      string
	RPCEndpoints []string
	WSEndpoints  []string
	DBName       string
	CurrentBlock int64
	DisableSync  bool
}

func NewConfig() Config {
	cfg := Config{
		// ChainID sets expected chain id for this instance
		ChainID: getStrEnv("PACKS_API_CHAIN_ID"),

		// RPCEndpoints is where to pull txs from when fast-syncing
		RPCEndpoints: func() []string {
			endpoints := getStrEnv("PACKS_API_RPC_ENDPOINTS")
			return strings.Split(endpoints, ",")
		}(),

		// WSEndpoints is where to pull txs from when normal syncing
		WSEndpoints: func() []string {
			endpoints := getStrEnv("PACKS_API_WS_ENDPOINTS")
			return strings.Split(endpoints, ",")
		}(),

		// DB name
		DBName: getStrEnv("PACKS_API_DB_NAME"),

		// DisableSync sets a flag where if true service won't accept any blocks (usually for debugging)
		DisableSync: func() bool {
			disableSync := getStrEnv("PACKS_API_DISABLE_SYNC")
			return disableSync == "true"
		}(),

		// Block height to start syncing from
		CurrentBlock: int64(getIntEnv("PACKS_API_CURRENT_BLOCK")),
	}

	return cfg
}

func (cfg Config) Print() {
	fmt.Println(cfg)
}

func getStrEnv(tag string) string {
	if e := os.Getenv(tag); e == "" {
		panic(fmt.Errorf("environment variable %s not set; expected string, got %s \"\"", tag, e))
	} else {
		return e
	}
}

func getIntEnv(tag string) int {
	val := getStrEnv(tag)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("expected int env variable for %s", tag))
	}
	return ret
}
