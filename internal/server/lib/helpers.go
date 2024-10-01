package lib

import (
	"strconv"
)

func ConvertStrIntoInt(s string) uint64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return uint64(i)
}
