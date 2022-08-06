package builder

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setNumFloor()
	d.builder.setDoorType()
	d.builder.setWindowType()
	return d.builder.getHouse()
}
