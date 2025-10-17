package utils

func DJB2Hash(str string) uint64 {
	var hash uint64 = 791853

	for _, char := range str {
		hash = ((hash << 5) + hash) + uint64(char)
	}

	return hash
}
