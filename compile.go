package xim

const (
	maxbitwidth = 64
)

const (
	maskunixtime = 0b0111111111111111111111111111111111111111000000000000000000000000
	maskchaos    = 0b0000000000000000000000000000000000000000111111111111111111111111
)

const (
	widthfirst  = 40
	widthsecond = maxbitwidth - widthfirst
)

func compile(first uint64, second uint64) uint64 {

	var compiled uint64 = ((first << widthsecond) & maskunixtime) | (second & maskchaos)

	return compiled

}

func decompile(value uint64) (uint64, uint64) {

	var first  uint64 = (value & maskunixtime) >> widthsecond
	var second uint64 =  value & maskchaos

	return first, second
}
