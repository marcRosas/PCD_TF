package k_means

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	EuclidDistance = "euclid"
)

type DistanceFunc func(*Point, *Cluster) (float64, error)

func EuclidDistanceFunc(p *Point, cl *Cluster) (distance float64, err error) {
	if p == nil {
		err = errors.New("undefined point")
		return
	}

	if cl == nil {
		err = errors.New("undefined cluster")
		return
	}

	if cl.GetCentralValuesCount() != p.GetElementsCount() {
		err = errors.New("cluster and point has different number of elements to comapare")
		return
	}

	for i := int32(0); i < cl.GetCentralValuesCount(); i++ {
		distance += math.Pow(cl.GetCentralValue(i)-p.GetValue(i), float64(2.0))
	}

	distance = math.Sqrt(float64(distance))

	return
}

func getNearestCenterId(p *Point, cls Clusters, fcalc DistanceFunc) (clusterId int32, err error) {
	if len(cls) < 2 {
		err = errors.New("no clusters are available")
		return clusterId, err
	}

	clusterId = 0

	minDistance, err := fcalc(p, cls[0])
	if err != nil {
		return
	}

	for i := 0; i < len(cls); i++ {
		distance, err := fcalc(p, cls[i])
		if err != nil {
			return clusterId, err
		}

		if distance < minDistance {
			clusterId = cls[i].GetId()
			minDistance = distance
		}
	}

	return clusterId, err
}

func isPointUsed(pointId int32, ps []int32) bool {
	for i := range ps {
		if ps[i] == pointId {
			return true
		}
	}

	return false
}

func Calc(points Points, k int32, maxIterations int32, fcalc DistanceFunc) (clusters Clusters, err error) {
	var timeStart = time.Now()

	if k > int32(len(points)) {
		err = errors.New("number of clusters is bigger than number of points")
		return
	}

	var randomSource = rand.NewSource(time.Now().UnixNano())
	var randomizer = rand.New(randomSource)

	var usedPoints []int32

	for i := int32(0); i < k; i++ {
		for true {
			var randomPointId = randomizer.Int31n(k)
			if !isPointUsed(randomPointId, usedPoints) {
				clusters = append(clusters, NewCluster(i, points[randomPointId]))
				usedPoints = append(usedPoints, randomPointId)
				points[randomPointId].SetCluster(i)

				break
			}
		}
	}

	var iterationNumber = int32(1)

	for true {
		var done = true
		for _, p := range points {
			var oldClusterId = p.GetClusterId()
			var nearestClusterId, err = getNearestCenterId(p, clusters, fcalc)
			if err != nil {
				return clusters, err
			}

			if oldClusterId != nearestClusterId {
				if oldClusterId != -1 {
					clusters[oldClusterId].RemovePoint(p.GetId())
				}

				p.SetCluster(nearestClusterId)
				clusters[nearestClusterId].AddPoint(p)
			}

			done = false
		}

		for _, cl := range clusters {
			if cl.GetPointsCount() > 0 {
				var dimensionsNumber = cl.GetCentralValuesCount()
				var pointsNumber = cl.GetPointsCount()

				for d := int32(0); d < dimensionsNumber; d++ {
					var sum = float64(0)

					for p := int32(0); p < pointsNumber; p++ {
						sum += cl.GetPoint(p).GetValue(d)
					}

					cl.SetCentralValue(d, sum/float64(pointsNumber))
				}
			}
		}

		if done == true || iterationNumber == maxIterations {
			break
		} else {
			iterationNumber += 1
		}
	}

	fmt.Println("computed in " + time.Since(timeStart).String() +
		"\nclusters: " + strconv.Itoa(len(clusters)) +
		"\npoints: " + strconv.Itoa(len(points)))

	return clusters, err
}

func (cls Clusters) PrettyPrint() {
	for _, cl := range cls {
		fmt.Println("Cluster", cl.GetId(), ":")
		for i := int32(0); i < cl.GetPointsCount(); i++ {
			p := cl.GetPoint(i)
			fmt.Println("Point", p.GetId(), ":", p.GetValues())
		}
		fmt.Println("Values:", cl.GetCentralValues())
	}
}
