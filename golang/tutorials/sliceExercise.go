package main

func Pic(dx, dy int) [][]uint8 {
	var xs = make([]([]uint8), dy) //missed the dy, obviously you need the length of the slice added
	for x, _ := range xs {
		xs[x] = make([]uint8, dx) //same thing as above comment
		for y, _ := range xs[x] {
			xs[x][y] = uint8((x & y) & (x & y)) //i dont really get this
		}
	}
	return xs
}
