package shape

type Rectangle struct {
	Width  float32
	Length float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.Width * rectangle.Length
}

func (Rectangle *Rectangle) Perimeter() float32 {
	return 2 * (Rectangle.Width + Rectangle.Length)
}
