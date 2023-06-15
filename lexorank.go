package lexorank

import (
	"errors"
	"sort"
)

var ErrRanksIdentical = errors.New("ranks must be different")
var ErrRanksConsecutive = errors.New("ranks must be inconsecutive")

func RankBetween(prevValue, nextValue string) (string, error) {
	if prevValue == nextValue {
		return "", ErrRanksIdentical
	}
	if prevValue > nextValue {
		prevValue, nextValue = nextValue, prevValue
	}
	if nextValue == prevValue+"0" { // two consecutive ranks, there is no rank between
		return "", ErrRanksConsecutive
	}

	returnedValue := RankIncrease(prevValue)
	if returnedValue < nextValue {
		return returnedValue, nil
	}

	returnedValue = prevValue + "1"
	if returnedValue < nextValue {
		return returnedValue, nil
	}

	prefix := "0"
	returnedValue = prevValue + prefix + "1"
	for returnedValue >= nextValue {
		prefix += "0"
		returnedValue = prevValue + prefix + "1"
	}
	return returnedValue, nil
}

// RankBetweenN returns n ranks between prev and next
func RankBetweenN(prev, next string, n int) ([]string, error) {
	baseRank, err := RankBetween(prev, next)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, n)
	suffix := []byte{}
	for i := 0; i < n; i++ {
		suffix = incrementChars(suffix)
		res = append(res, baseRank+string(suffix))
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res, nil
}

func RankIncrease(value string) string {
	chars := []byte(value)

	for i := len(value) - 1; i >= 0; i-- {
		next, ok := incrementChar(chars[i])
		if !ok {
			continue
		}
		return string(chars[:i]) + string(next)
	}

	// Add "1" instead of "0", as suffix "0" can cause 2 consecutive ranks
	return value + "1"
}

func RankDecrease(value string) string {
	chars := []byte(value)
	length := len(chars)

	prev, ok := decrementChar(chars[length-1])
	if ok && prev != '0' {
		return string(chars[:length-1]) + string(prev)
	}

	for i := length - 2; i >= 0; i-- {
		prev, ok := decrementChar(chars[i])
		if !ok {
			continue
		}
		return string(chars[:i]) + string(prev) + string(chars[i+1:])
	}

	return "0" + value
}
