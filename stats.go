package stats

import "math"

func Mean(x []float64) float64 {
    sum := float64(0.)
    for _, xi := range x {
        sum += xi
    }
    return sum/float64(len(x))
}

func RMS(x []float64) float64 {
    mean := Mean(x)
    sum := float64(0.)
    for _, xi := range x {
        sum += (xi-mean)*(xi-mean)
    }
    return math.Sqrt(sum/float64(len(x)))
}

