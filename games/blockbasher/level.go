package main

type Level struct {
	LevelNumber      int
	NumberOfBlocks   int
	StrengthOfBlocks int
	Blocks           []*Block
	BlockLayout      *BlockLayout
}

func NewLevel(levelNumber int) *Level {
	var (
		numberOfBlocks   = 4 + levelNumber
		strengthOfBlocks = 1
		n                int
	)
	if numberOfBlocks > 20 {
		numberOfBlocks = 20
		strengthOfBlocks += levelNumber - 15
	}
	level := &Level{
		LevelNumber:      levelNumber,
		NumberOfBlocks:   numberOfBlocks,
		StrengthOfBlocks: strengthOfBlocks,
		BlockLayout:      NewBlockLayout(strengthOfBlocks),
	}

	// likely not needed, started then turned into BlockLayout deployment
	for n < numberOfBlocks {
		level.Blocks = append(level.Blocks, NewBlock(0, 0, level.StrengthOfBlocks))
		n += 1
	}
	return level
}
