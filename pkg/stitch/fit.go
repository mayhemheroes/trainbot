package stitch

import (
	"math"
	"time"

	"github.com/jo-m/trainbot/pkg/ransac"
	"github.com/rs/zerolog/log"
)

const modelNParams = 2

func model(dx float64, ps []float64) float64 {
	v0 := ps[0]
	a := ps[1]
	return v0 + a*dx
}

// Returns fitted dx values, length will always be the same as the input.
// Also returns estimated v0 and acceleration (in pixels/s).
func fitDx(ts []time.Time, dx []int) ([]int, float64, float64, error) {
	start := time.Now()
	defer log.Trace().Dur("dur", time.Since(start)).Msg("fitDx() duration")

	// Sanity checks.
	if len(dx) != len(ts) {
		log.Panic().Msg("dx and ts do not have the same length, this should not happen")
	}
	if len(dx) < (modelNParams+1)*3 {
		return nil, 0, 0, errors.New("sequence length too short")
	}

	n := len(dx)
	t0 := ts[0]
	// Convert dx to float and calculate relative time.
	tsF := make([]float64, n) // Will contain zero-based time in seconds.
	dxF := make([]float64, n) // Will contain dx values.
	for i := range tsF {
		tsF[i] = float64(ts[i].Sub(t0).Seconds())
		dxF[i] = float64(dx[i])
	}

	params := ransac.MetaParams{
		MinModelPoints:  modelNParams + 1,
		MaxIter:         25,
		MinInliers:      len(dxF) / 2,
		InlierThreshold: 3.,
		Seed:            0,
	}
	fit, err := ransac.Ransac(tsF, dxF, model, modelNParams, params)
	if err != nil {
		return nil, 0, 0, err
	}

	dxFit := make([]int, n)
	for i, tsF := range tsF {
		dxF := model(tsF, fit.X)
		dxFit[i] = int(math.Round(dxF))
	}

	return dxFit, fit.X[0], fit.X[1], nil
}
