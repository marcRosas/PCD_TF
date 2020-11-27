package api

import (
	"math"
)

func calcularDistancia(desde, hasta []float32) (respuesta float32) {
	recorre := float64(0)
	for i := 0; i < len(desde); i++ {
		diff := float64(desde[i] - hasta[i])
		recorre += diff * diff
	}
	respuesta = float32(math.Sqrt(recorre))
	return
}
