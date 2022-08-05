package facetory_method

//M4a1 具体产品
type M4a1 struct {
	Gun
}

func newM4a1() IGun {
	return &M4a1{
		Gun: Gun{
			name:  "M4a1 gun",
			power: 1,
		}}
}
