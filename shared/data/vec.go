package data

type Vec[T float64 | int] struct {
	X T
	Y T
}

func (v Vec[T]) Add(other Vec[T]) Vec[T] {
	return Vec[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec[T]) Sub(other Vec[T]) Vec[T] {
	return Vec[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec[T]) Red() Vec[int] {
	gdc := gcd(int(v.X), int(v.Y))
	return Vec[int]{
		X: int(v.X) / gdc,
		Y: int(v.Y) / gdc,
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
