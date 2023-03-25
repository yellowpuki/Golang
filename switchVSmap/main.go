package main

type PlaintType struct {
	name string
}

var Apple, Banana, Tomatoe PlaintType
var fruits map[PlaintType]struct{} = map[PlaintType]struct{}{
	Apple:   struct{}{},
	Banana:  struct{}{},
	Tomatoe: struct{}{},
}

func isFruitSwitch(plaintType PlaintType) bool {
	switch plaintType {
	case Apple, Banana, Tomatoe:
		return true
	}
	return false
}

func isFruitMap(plaintType PlaintType) bool {
	_, ok := fruits[plaintType]
	return ok
}

func main() {

}
