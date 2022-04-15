package algos

// 単純移動平均線
func CalcSma(period int, closes []float64) []float64 {
	result := make([]float64, len(closes))
	periodTotal := 0.0

	for i := 0; i < period-1; i++ {
		periodTotal += closes[i]
	}

	for i := period - 1; i < len(closes); i++ {
		periodTotal += closes[i]
		result[i] = periodTotal / float64(period)
		periodTotal -= closes[i-period+1]
	}

	return result
}
