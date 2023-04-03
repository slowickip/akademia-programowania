package ships

type Point struct {
	X, Y int
}

func (p Point) Add(a Point) Point {
	return Point{
		X: p.X + a.X,
		Y: p.Y + a.Y,
	}
}

type Ship []Point

func (s Ship) Size() int {
	return len(s)
}

func (s Ship) MoveTo(p Point) Ship {
	delta := Point{
		X: p.X - s[0].X,
		Y: p.Y - s[0].Y,
	}
	ns := make(Ship, s.Size())
	for i, v := range s {
		ns[i] = v.Add(delta)
	}
	return ns
}
