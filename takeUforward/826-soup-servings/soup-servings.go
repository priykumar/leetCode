package main

import (
    "math"
)

func soupServings(n int) float64 {
    if n > 5000 {
        return 1.0
    }
    units := int(math.Ceil(float64(n) / 25.0))
    cache := make([][]float64, units+1)
    for i := range cache {
        cache[i] = make([]float64, units+1)
        for j := range cache[i] {
            cache[i][j] = -1.0
        }
    }

    var calcProb func(int, int) float64
    calcProb = func(soupA, soupB int) float64 {
        if soupA <= 0 && soupB <= 0 {
            return 0.5
        }
        if soupA <= 0 {
            return 1.0
        }
        if soupB <= 0 {
            return 0.0
        }
        if cache[soupA][soupB] != -1.0 {
            return cache[soupA][soupB]
        }
        prob := 0.25 * (
            calcProb(soupA-4, soupB) +
            calcProb(soupA-3, soupB-1) +
            calcProb(soupA-2, soupB-2) +
            calcProb(soupA-1, soupB-3))
        cache[soupA][soupB] = prob
        return prob
    }

    return calcProb(units, units)
}