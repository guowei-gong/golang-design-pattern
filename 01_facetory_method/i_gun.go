package facetory_method

//IGun 产品接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
