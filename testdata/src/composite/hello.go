package composite

type A interface {
	Array(obj [4]string)
	Chan(obj1 chan int, obj2 <-chan string, obj3 chan<- float64)
	Func(obj1 func(), obj2 func(int, string) error)
	Map(obj map[int]string)
	Pointer(obj *int)
	Slice(obj []string)
	Struct(obj struct{ a int })
}
