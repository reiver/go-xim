package iid

const (
	maxbitwidth = 64
)

const (
	maskfirst  = 0b0111111111111111111111111111111111111111000000000000000000000000
	masksecond = 0b0000000000000000000000000000000000000000111111111111111111111111
)

const (
	widthfirst  = 40
	widthsecond = maxbitwidth - widthfirst
)

func compile(first uint64, second uint64) uint64 {

	var compiled uint64 = ((first << widthsecond) & maskfirst) | (second & masksecond)

	return compiled

}

func decompile(value uint64) (uint64, uint64) {

	var first  uint64 = (value & maskfirst) >> widthsecond
	var second uint64 =  value & masksecond

	return first, second
}
