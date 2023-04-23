//go:build moretests
// +build moretests

package stitch

import (
	"image"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AutoStitcher_Set2_All(t *testing.T) {
	c := Config{
		PixelsPerM:  42,
		MinSpeedKPH: 10,
		MaxSpeedKPH: 160,
		MinLengthM:  10,
	}
	r := image.Rect(0, 0, 350, 290)

	basepath := "testdata/set2/"

	vids, err := filepath.Glob(basepath + "*.mkv")
	require.NoError(t, err)

	for _, vid := range vids {
		runTestSimple(t, c, r, vid, 138)
	}
}

func Test_AutoStitcher_Set2(t *testing.T) {
	c := Config{
		PixelsPerM:  42,
		MinSpeedKPH: 10,
		MaxSpeedKPH: 160,
		MinLengthM:  10,
	}
	r := image.Rect(0, 0, 350, 290)

	basepath := "testdata/set2/"
	runTestSimple(t, c, r, basepath+"127.mkv", 500)
}
