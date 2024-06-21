package Hash

import (
	"encoding/hex"
	"fmt"
)

func Hash(value string) string {
	encoded := hex.EncodeToString([]byte(value))
	fmt.Println("encoding", encoded)
	sumOfBytes, rounds := hashValue()
	return encoded

}

func hashValue(byteVal []int, round int) ([]int, int) {
	if len(byteVal) < 2 {
		return byteVal, round
	}
	var addedSlice []int
	length := len(byteVal) - 1
	i := 0
	for i := 0; i < len(byteVal); i = i + 2 {
		added := int(byteVal[i]) + int(byteVal[i+1])
		addedSlice = append(addedSlice, added)

	}
	if i < length {
		addedSlice = append(addedSlice, int(byteVal[length]))
	}

	return hashValue(addedSlice, round+1)

}
