package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }

func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (r Rectangle) Area() float64   { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Type assertion extracts the concrete type from an interface.
// Comma-ok form is safe — no panic on failure.
func describeShape(s Shape) {
	if c, ok := s.(Circle); ok {
		fmt.Printf("Circle  radius=%.2f  area=%.2f\n", c.Radius, s.Area())
		return
	}
	if r, ok := s.(Rectangle); ok {
		fmt.Printf("Rectangle %gx%g  area=%.2f  perimeter=%.2f\n",
			r.Width, r.Height, s.Area(), r.Perimeter())
		return
	}
	fmt.Printf("unknown shape  area=%.2f\n", s.Area())
}

// Type switch is cleaner when handling many concrete types
func printType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %q\n", v)
	case bool:
		fmt.Printf("bool: %v\n", v)
	case []int:
		fmt.Printf("[]int len=%d\n", len(v))
	case Shape:
		fmt.Printf("Shape area=%.2f\n", v.Area())
	default:
		fmt.Printf("unknown: %T\n", v)
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}
	for _, s := range shapes {
		describeShape(s)
	}

	fmt.Println()
	values := []any{42, "hello", true, []int{1, 2, 3}, Circle{Radius: 3}, 3.14}
	for _, v := range values {
		printType(v)
	}
}
