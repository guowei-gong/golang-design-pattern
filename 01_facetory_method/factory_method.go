package facetory_method

import "fmt"

//IGun 产品接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

//Gun 具体产品
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

//Ak47 具体产品
type Ak47 struct {
	Gun
}

//M4a1 具体产品
type M4a1 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "Ak47 gun",
			power: 4,
		}}
}

func newM4a1() IGun {
	return &M4a1{
		Gun: Gun{
			name:  "M4a1 gun",
			power: 1,
		}}
}

//getGun 工厂
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "m4a1" {
		return newM4a1(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}
