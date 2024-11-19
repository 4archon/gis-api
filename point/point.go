package point

type Point struct {
	Long string
	Lat string
}

func (p *Point) Init(long string, lat string) {
	p.Long = long
	p.Lat = lat
}