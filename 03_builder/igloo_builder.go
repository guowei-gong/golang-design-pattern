package builder

// 生成器

type IglooBuilder struct {
	floor      int
	doorType   string
	windowType string
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		floor:      b.floor,
		doorType:   b.doorType,
		windowType: b.windowType,
	}
}
