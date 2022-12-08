package est

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func genTs(n int) []time.Time {
	ret := make([]time.Time, n)
	for i := 0; i < n; i++ {
		ret[i] = time.Time{}.Add(time.Second * time.Duration(i))
	}
	return ret
}

func Test_fitDx_simple(t *testing.T) {
	testData := []int{
		34, 34, 34, 34, 34, 26, 0, 34, 1, 1, 0, 0, 20, 0, 34, 34, 34, 34, 25, 34, 34,
		34, 34, 34, 34, 34, 34, 34, 22, 0, 34, 34, 28, 34, 34, 26, 27, 34, 34, 34, 34,
		34, 34, 0, 0, 34, 34, 34, 0, 34, 34, 34, 34, 1, 34, 34, 22, 34, 34, 34, 34, 0,
		34, 0, 34, 34, 26, 34, 34, 34, 3, 34, 34, 32, 34, 34, 34, 7, 0, 34, 0, 34, 1,
		34, 34, 0, 34, 34, 5, 34, 5, 34, 27, 0, 0, 34, 34, 34, 34, 34, 32, 31, 34, 34,
		29, 25, 34, 10, 0, 6, 0, 34, 0, 34, 1, 24, 34, 34, 35,
	}

	res, err := fitDx(genTs(len(testData)), testData)
	require.NoError(t, err)
	assert.Equal(t, 119, len(res))
	assert.Equal(t, 34, res[10])
}

func Test_fitDx_difficult(t *testing.T) {
	testData := []int{
		36, 36, 41, 36, 35, 0, 0, 40, 36, 37, 36, 42, 41, 17, 35, 35, 42, 41, 17, 41, 36, 41, 17, 0, 35, 36, 36,
		40, 35, 42, 41, 41, 36, 40, 36, 35, 0, 40, 36, 36, 36, 42, 35, 42, 36, 42, 36, 41, 36, 0, 41, 36, 41, 17,
		0, 36, 41, 40, 36, 41, 42, 35, 41, 41, 40, 41, 17, 41, 36, 42, 41, 36, 41, 17, 42, 17, 41, 17, 41, 18, 36,
		40, 36, 35, 0, 37, 41, 36, 36, 35, 37, 36, 36, 36, 36, 36, 40, 0, 41, 36, 41, 36, 42, 36, 42, 37, 42, 40,
		35, 41, 42, 37, 41, 35, 0, 17, 36, 36, 36, 36, 36, 0, 36, 36, 36, 36, 35, 17, 0, 36, 36, 35, 35, 41, 40,
		42, 41, 35, 36, 36, 41, 41, 36, 40, 17, 40, 36, 36, 35, 36, 17, 36, 40, 40, 36, 41, 36, 35, 0, 37, 36, 36,
		36, 40, 35, 42, 35, 41, 35, 40, 35, 41, 41, 36, 40, 17, 40, 36, 37, 40, 35, 42, 41, 35, 0, 40, 36, 36, 0,
		41, 36, 40, 40, 42, 41, 17, 41, 17, 41, 17, 41, 17, 35, 36, 41, 35, 0, 35, 40, 36, 36, 36, 42, 35, 40, 36,
		36, 35, 36, 0, 41, 36, 40, 36, 42, 35, 41, 36, 42, 36, 42, 36, 41, 36, 42, 35, 0, 17, 36, 41, 36, 36, 36,
		36, 36, 36, 36, 36, 17, 17, 0, 36, 40, 41, 35, 36, 36, 42, 41, 35, 35, 35, 42, 40, 36, 37, 0, 40, 40, 41,
		36, 36, 35, 42, 35, 37, 36, 41, 40, 17, 0, 36, 42, 41, 35, 41, 17, 41, 35, 39, 35, 40, 35, 41, 40, 36, 36,
		0, 40, 41, 36, 37, 35, 42, 40, 40, 36, 40, 36, 35, 0, 41, 36, 37, 35, 42, 36, 42, 36, 42, 40, 0, 41, 17,
		41, 36, 42, 17, 0, 36, 40, 41, 36, 36, 42, 35, 40, 41, 36, 35, 35, 17, 40, 36, 42, 36, 41, 35, 40, 35, 40,
		35, 42, 35, 41, 36, 41, 17, 0, 17, 36, 41, 36, 36, 35,
	}

	res, err := fitDx(genTs(len(testData)), testData)

	require.NoError(t, err)
	assert.Equal(t, 365, len(res))
	assert.Equal(t, 37, res[10])
}

func Test_fitDx_negative(t *testing.T) {
	testData := []int{
		-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
		-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	}

	res, err := fitDx(genTs(len(testData)), testData)
	require.NoError(t, err)
	assert.Equal(t, 20, len(res))
	assert.Equal(t, -9, res[5])
	assert.Equal(t, -10, res[len(res)-5])
}

func Test_fitDx_rounding(t *testing.T) {
	testData := []int{
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	}

	res, err := fitDx(genTs(len(testData)), testData)
	require.NoError(t, err)
	assert.Equal(t,
		[]int{
			9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 10, 9, 9, 9, 10,
			9, 10, 10, 10, 8, 10, 10, 10, 8, 10, 10, 10, 10, 10, 11,
		},
		res,
	)
}

func Test_fitDx_rounding_negative(t *testing.T) {
	testData := []int{
		-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
		-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	}

	res, err := fitDx(genTs(len(testData)), testData)
	require.NoError(t, err)
	assert.Equal(t,
		[]int{
			-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -10, -9, -9, -9, -10, -9, -10, -10, -10, -8, -10, -10, -10, -8, -10, -10, -10, -10, -10, -11,
		},
		res,
	)
}
