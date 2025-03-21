package point

type Point struct {
	ID int
	Long string
	Lat string
}

func (p *Point) Init(id int, long string, lat string) {
	p.ID = id
	p.Long = long
	p.Lat = lat
}


type FilterPoint struct {
	ID			int
	Long		string
	Lat			string
	Active		bool
	Repair		bool
	Assigned	bool
	LongTime	bool
}