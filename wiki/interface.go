package wiki

type IShape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length int
	width int
}

type Square struct {
	side int
}

func (r Rectangle) Area() float64 {
	return float64(r.length * r.width)
}

func (r Rectangle) Perimeter() float64 {
	return float64(r.length * 2 + r.width * 2)
}

func (s Square) Area() float64 {
	return float64(s.side ^ 2)
}

func (s Square) Perimeter() float64 {
	return float64(s.side ^ 2)
}
