// Package day9 solves puzzle available on https://adventofcode.com/2025/day/9
package day9

import (
	"math"
	"slices"

	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)
	tiles := parseData(in)
	rects := allRects(tiles)

	return biggestNonintersecting(rects, tiles)
}

func allRects(tiles [][2]int) [][2][2]int {
	rects := make([][2][2]int, 0, len(tiles)*len(tiles))

	for i, tile := range tiles {
		for _, other := range tiles[i:] {
			rects = append(rects, [2][2]int{tile, other})
		}
	}

	slices.SortFunc(rects, func(a, b [2][2]int) int { return int(area(b[x], b[y]) - area(a[x], a[y])) })

	return rects
}

func biggestNonintersecting(rects [][2][2]int, tiles [][2]int) int {
	for _, rect := range rects {
		maxX := int(math.Max(float64(rect[a][x]), float64(rect[b][x])))
		maxY := int(math.Max(float64(rect[a][y]), float64(rect[b][y])))
		minX := int(math.Min(float64(rect[a][x]), float64(rect[b][x])))
		minY := int(math.Min(float64(rect[a][y]), float64(rect[b][y])))

		if !intersects(minX, minY, maxX, maxY, tiles) {
			return area(rect[a], rect[b])
		}
	}

	return 0
}

func intersects(minX, minY, maxX, maxY int, tiles [][2]int) bool {
	for i := range len(tiles) {
		line := [2][2]int{tiles[i], tiles[(i+1)%len(tiles)]}
		lmaxX := int(math.Max(float64(line[a][x]), float64(line[b][x])))
		lmaxY := int(math.Max(float64(line[a][y]), float64(line[b][y])))
		lminX := int(math.Min(float64(line[a][x]), float64(line[b][x])))
		lminY := int(math.Min(float64(line[a][y]), float64(line[b][y])))

		if lmaxX > minX && lminX < maxX && lmaxY > minY && lminY < maxY {
			return true
		}
	}

	return false
}
