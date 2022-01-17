package tuple

type Tuple2[T any, U any] struct {
	First  T
	Second U
}

func NewTuple2[T any, U any](f T, s U) Tuple2[T, U] {
	return Tuple2[T, U]{
		First:  f,
		Second: s,
	}
}
