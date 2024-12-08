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
