package app

import "flag"

var (
	seed      int64
	blocksLen int
	blockSize int
	commit    bool
	period    int
)

func init() {
	flag.Int64Var(&seed, "seed", 42, "simulation random seed")
	flag.IntVar(&blocksLen, "blocks_len", 500, "number of blocks")
	flag.IntVar(&blockSize, "block_size", 200, "operations per block")
	flag.BoolVar(&commit, "commit", false, "have the simulation commit")
	flag.IntVar(&period, "period", 1, "run slow invariants only once every period assertions")
}
