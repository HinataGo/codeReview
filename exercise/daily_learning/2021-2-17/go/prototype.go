package _go

type f struct {
	F func()
}
type X struct {
	f     func()
	count func()
}

func NewX(f func(x X)) X {
	count := 0
	x := X{
		count: func() {
			count += 1
		},
	}
	x.f = func() {
		f(x)
	}
	return x
}
func (x X) F() {
	x.count()
	x.f
}

func main() {
	x1 := X{}
	x1.F()
	x2 := X{}
	x2.F()
}
