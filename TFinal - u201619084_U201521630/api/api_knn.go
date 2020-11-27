package api

import (
	"sort"
)

func Definir(k, numMaximo, numValores int) (resp *KNearest) {
	return &KNearest{
		k:          k,
		numValores: numValores,
		data:       [][]float32{},
		valores:    []int{},
		numMaximo:  numMaximo,
	}
}

func (knn *KNearest) Len() int {
	return len(knn.metricas)
}

func (knn *KNearest) Less(i, j int) bool {
	return knn.metricas[i].metrica < knn.metricas[j].metrica
}

func (knn *KNearest) Swap(i, j int) {
	knn.metricas[i], knn.metricas[j] = knn.metricas[j], knn.metricas[i]
}

func (knn *KNearest) Aprender(data [][]float32, valores []int) (err error) {
	num := len(data)
	for i := 0; i < num; i++ {
		knn.data = append(knn.data, data[i])
		knn.valores = append(knn.valores, valores[i])
	}
	return
}

func (knn *KNearest) Predecir(data []float32) (valor int) {
	knn.calcularDistancias(data)
	sort.Sort(knn)
	if knn.metricas[0].metrica == float32(0) {
		idx := knn.metricas[0].idx
		valor = knn.valores[idx]
		return
	}
	labelCounts := make([]int, knn.numValores)
	for i := 0; i < knn.k; i++ {
		idx := knn.metricas[i].idx
		lbl := knn.valores[idx]
		labelCounts[lbl] = labelCounts[lbl] + 1
	}
	maxCount := 0
	for i, count := range labelCounts {
		if count > maxCount {
			valor = i
			maxCount = count
		}
	}
	return
}

func (knn *KNearest) calcularDistancias(desde []float32) (err error) {
	knn.metricas = make([]metrica, len(knn.data))
	for i, hasta := range knn.data {
		knn.metricas[i] = metrica{
			idx:     i,
			metrica: calcularDistancia(desde, hasta),
		}
	}
	return
}
