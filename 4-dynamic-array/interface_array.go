package p04dynamicarray

type ArrayItemType interface {
	int32 | int64 | float32 | float64 | string
}

type IArray[T ArrayItemType] interface {
	Size() int
	Get(int) T
	Set(int, T)
	Add(int, T)
	Remove(int) T
}
