package hashes

const MAGIC_1 uint64 = 791853

func DJB2(str string) uint64 {
	var hash uint64 = MAGIC_1

	for _, char := range str {
		hash = ((hash << 5) + hash) + uint64(char)
	}

	return hash
}
