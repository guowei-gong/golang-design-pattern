package builder

// 生成器接口

type IBuilder interface {
	setDoorType()
	setNumFloor()
	setWindowType()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
