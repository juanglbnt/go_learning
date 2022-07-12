package function

type Square struct {
    X, Y float64
}

func (s Square) Area() float64 {
    return s.X * s.Y 
}

func (s Square) Perim() float64 {
    return (2*s.X) + (2*s.Y)
}
