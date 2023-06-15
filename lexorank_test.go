package lexorank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRank_RankBetween(t *testing.T) {
	testCases := []struct {
		prev    string
		next    string
		between string
		err     error
	}{
		// The prev rank and next rank are same
		{"a", "a", "", ErrRanksIdentical},
		{"", "", "", ErrRanksIdentical},
		{"0", "00", "", ErrRanksConsecutive},
		{"0a", "0a0", "", ErrRanksConsecutive},
		// The prev rank is empty
		{"", "2", "1", nil},
		// The next rank is empty
		{"z", "", "1", nil},
		// Normal case
		{"1", "9", "2", nil},
		{"12", "1b", "13", nil},
		{"101", "123", "102", nil},
		{"cz", "d", "cz1", nil},
		{"d04", "d041", "d0401", nil},
		{"z1", "z1001", "z10001", nil},
		// Prev rank greater than next rank
		{"d041", "d04", "d0401", nil},
		{"z1", "z1001", "z10001", nil},
	}

	for _, tc := range testCases {
		between, err := RankBetween(tc.prev, tc.next)
		if between != tc.between {
			t.Errorf("RankBetween(%s, %s) => %s, want %s", tc.prev, tc.next, between, tc.between)
			t.Errorf(err.Error())
		}
	}
}

func TestRank_RankBetweenN(t *testing.T) {
	testCases := []struct {
		prev  string
		next  string
		count int
		err   error
	}{
		// Error cases
		{"a", "a", 10, ErrRanksIdentical},
		{"", "", 10, ErrRanksIdentical},
		{"0", "00", 10, ErrRanksConsecutive},
		{"0a", "0a0", 10, ErrRanksConsecutive},
		// Success cases
		{"1", "11", 1000, nil},
		{"1", "2", 10000, nil},
	}

	for _, tc := range testCases {
		result, err := RankBetweenN(tc.prev, tc.next, tc.count)
		if tc.err != nil {
			assert.ErrorIs(t, err, tc.err)
			continue
		}
		assert.Nil(t, err)
		assert.Equal(t, tc.count, len(result))
		for i := 0; i < tc.count; i++ {
			if i < tc.count-1 {
				assert.True(t, result[i] < result[i+1])
			}
			assert.True(t, result[i] > tc.prev)
			assert.True(t, result[i] < tc.next)
		}
	}
}

func TestRank_RankIncrease(t *testing.T) {
	// Basic case
	assert.Equal(t, "2", RankIncrease("1"))
	assert.Equal(t, "A", RankIncrease("9"))
	assert.Equal(t, "a", RankIncrease("Z"))
	assert.Equal(t, "b", RankIncrease("a"))
	assert.Equal(t, "z1", RankIncrease("z"))
	// Complex case
	assert.Equal(t, "12", RankIncrease("11"))
	assert.Equal(t, "5f", RankIncrease("5e"))
	assert.Equal(t, "dc", RankIncrease("dbz"))
	assert.Equal(t, "500A", RankIncrease("5009"))
	assert.Equal(t, "500a", RankIncrease("500Z"))
	assert.Equal(t, "b", RankIncrease("azzzzzz"))
	assert.Equal(t, "az", RankIncrease("ayzzzzzzz"))
	assert.Equal(t, "zzzzzz1", RankIncrease("zzzzzz"))
}

func TestRank_RankDecrease(t *testing.T) {
	// Basic case
	assert.Equal(t, "01", RankDecrease("1"))
	assert.Equal(t, "b", RankDecrease("c"))
	assert.Equal(t, "8", RankDecrease("9"))
	assert.Equal(t, "Z", RankDecrease("a"))
	assert.Equal(t, "9", RankDecrease("A"))
	// Complex case
	assert.Equal(t, "01", RankDecrease("11"))
	assert.Equal(t, "5d", RankDecrease("5e"))
	assert.Equal(t, "001", RankDecrease("01"))
	assert.Equal(t, "40Z", RankDecrease("40a"))
	assert.Equal(t, "409", RankDecrease("40A"))
	assert.Equal(t, "00001", RankDecrease("01001"))
}
