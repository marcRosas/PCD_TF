package k_means

type Point struct {
	PointId   int32
	ClusterId int32
	Name      string
	Values    []float64
}

type Points []*Point

func NewPoint(pointId int32, values []float64, name string) *Point {
	return &Point{
		PointId: pointId,
		ClusterId: -1,
		Name: name,
		Values: values,
	}
}

func (p *Point) GetId() int32 {
	return p.PointId
}

func (p *Point) GetClusterId() int32 {
	return p.ClusterId
}

func (p *Point) GetName() string {
	return p.Name
}

func (p *Point) GetValue(index int32) float64 {
	return p.Values[index]
}

func (p *Point) GetValues() []float64 {
	return p.Values
}

func (p *Point) GetElementsCount() int32 {
	return int32(len(p.Values))
}

func (p *Point) SetCluster(clusterId int32) {
	p.ClusterId = clusterId
}