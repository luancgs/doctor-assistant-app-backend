package utils

import (
	"strconv"
)

func ConvertId(id string) (uint, error) {
	convertedId, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	return uint(convertedId), nil
}
