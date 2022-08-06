package builder

// 具体生成器

type NormalBuilder struct {
	floor      int
	doorType   string
	windowType string
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) getHouse() House {
	return House{
		floor:      b.floor,
		doorType:   b.doorType,
		windowType: b.windowType,
	}
}
