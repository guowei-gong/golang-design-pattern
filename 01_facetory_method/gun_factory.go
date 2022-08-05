package facetory_method

import "fmt"

// 工厂
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "m4a1" {
		return newM4a1(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}
