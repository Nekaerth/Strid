package game

type Map struct {
	Width int
	Height int
	[][]*Field Fields
}

type Field struct {
	X int
	Y int
}

func newMap(width int, height int) *Map {
	returnMap := &Map{width, height, make([][]*Field, height)}
	
	for y := 0; y < height; y++ {
		returnMap.Fields[y] = make([]*Field, width)
		for x := 0; x < width; x++ {
			returnMap.Fields[y][x] = &Field{x, y}
		}
	}
	
	return returnMap
}