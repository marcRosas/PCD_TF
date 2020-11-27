package k_means

import "errors"

type Cluster struct {
	ClusterId     int32
	CentralValues []float64
	Points        []*Point
}

type Clusters []*Cluster

func NewCluster(clusterId int32, cm *Point) *Cluster {
	var cmValues []float64
	for i := int32(0); i < cm.GetElementsCount(); i++ {
		cmValues = append(cmValues, cm.GetValue(i))
	}

	return &Cluster{
		ClusterId:     clusterId,
		CentralValues: cmValues,
		Points:        []*Point{cm},
	}
}

func (cl *Cluster) GetId() int32 {
	return cl.ClusterId
}

func (cl *Cluster) GetCentralValue(index int32) float64 {
	return cl.CentralValues[index]
}

func (cl *Cluster) GetCentralValues() []float64 {
	return cl.CentralValues
}

func (cl *Cluster) GetCentralValuesCount() int32 {
	return int32(len(cl.CentralValues))
}

func (cl *Cluster) SetCentralValue(index int32, val float64) error {
	if index > int32(len(cl.CentralValues)) {
		return errors.New("index out of range")
	} else {
		cl.CentralValues[index] = val
	}

	return nil
}

func (cl *Cluster) GetPoint(index int32) *Point {
	return cl.Points[index]
}

func (cl *Cluster) GetPointsCount() int32 {
	return int32(len(cl.Points))
}

func (cl *Cluster) AddPoint(p *Point) {
	cl.Points = append(cl.Points, p)
}

func (cl *Cluster) RemovePoint(pointId int32) bool {
	for i, p := range cl.Points {
		if p.GetId() == pointId {
			cl.Points = append(cl.Points[:i], cl.Points[i+1:]...)
			return true
		}
	}

	return false
}
