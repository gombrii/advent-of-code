package day9

import (
	"github.com/gombrii/Advent-of-code/shared/dat"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

type file struct {
	id   *int
	size int
	uid  int
}

func Part2(data []byte) any {
	in := parse.String(data)

	decompressed := decompress(in)
	optimized := optimizeWithtFrag(decompressed)
	checksum := checksum(optimized)

	return checksum
}

func optimizeWithtFrag(data []*int) []*int {
	optimized, reversed := format(data)

	for _, fil := range reversed.Iter {
		if fil.id == nil {
			continue
		}
		for _, space := range optimized.Iter {
			if space == fil {
				break
			}
			if space.id != nil {
				continue
			}
			if space.size >= fil.size {
				optimized.Insert(file{
					id:   nil,
					size: fil.size,
					uid:  fil.uid,
				}, fil)
				optimized.Insert(fil, space)
				optimized.Remove(space)
				optimized.Insert(file{
					id:   nil,
					size: space.size - fil.size,
					uid:  space.uid,
				}, fil)
				break
			}
		}
	}

	out := []*int{}
	for _, file := range optimized.Iter {
		for i := 0; i < file.size; i++ {
			out = append(out, file.id)
		}
	}
	return out
}

func format(nonFormatted []*int) (formatted dat.Linked[file], reversed dat.Linked[file]) {
	formatted = dat.NewLinked[file]()
	reversed = dat.NewLinked[file]()

	i := 0
	for i < len(nonFormatted) {
		var file file
		if nonFormatted[i] != nil {
			i, file = readFile(nonFormatted, i)
		} else {
			i, file = readSpace(nonFormatted, i)
		}

		formatted.Append(file)
		reversed.Prepend(file)
	}

	return formatted, reversed
}

func readFile(data []*int, i int) (int, file) {
	file := file{
		id:  data[i],
		uid: i,
	}
	for i < len(data) && data[i] == file.id {
		i++
	}
	file.size = i - file.uid
	return i, file
}

func readSpace(data []*int, i int) (int, file) {
	space := file{
		id:  nil,
		uid: i,
	}
	for i < len(data) && data[i] == nil {
		i++
	}
	space.size = i - space.uid
	return i, space
}
