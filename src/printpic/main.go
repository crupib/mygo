package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for y := range pic {
        pic[y] = make([]uint8, dx)
        for x := range pic[y] {
            pic[y][x] = uint8(5 * (x + y))
        }
    }
    return pic
}

func main() {
    //pic.Show(Pic(40,30)) // doesn't work, but why?
    pic.Show(Pic) // works, but why? Where are the values for dx and dy set?
}
