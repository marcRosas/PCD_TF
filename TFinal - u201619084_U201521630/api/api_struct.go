package api

type KNearest struct {
	k          int
	numValores int
	data       [][]float32
	valores    []int
	numMaximo  int
	metricas   []metrica
}

type metrica struct {
	idx     int
	metrica float32
}
