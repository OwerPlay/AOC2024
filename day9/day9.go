package day9

import (
	"AOC2024/utils"
	"unicode"
)

type fileInfo struct {
	id    int
	start int
	end   int
	size  int
}

func parseDiskMap(diskMap string) ([]rune, []fileInfo) {
	var blocks []rune
	fileID := 0
	isFile := true

	for _, ch := range diskMap {
		if !unicode.IsDigit(ch) {
			panic("Invalid input character")
		}
		length := int(ch - '0')
		if length < 0 {
			panic("Negative length?")
		}

		if isFile && length > 0 {
			idChar := rune('0' + fileID)
			for j := 0; j < length; j++ {
				blocks = append(blocks, idChar)
			}
			fileID++
		} else if !isFile && length > 0 {
			for j := 0; j < length; j++ {
				blocks = append(blocks, '.')
			}
		}

		isFile = !isFile
	}

	files := []fileInfo{}
	var startPos int = -1
	var currentID int = -1
	for i, b := range blocks {
		if b != '.' {
			id := int(b - '0')
			if id != currentID {
				if currentID != -1 {
					oldFile := fileInfo{
						id:    currentID,
						start: startPos,
						end:   i - 1,
						size:  (i - startPos),
					}
					files = append(files, oldFile)
				}
				currentID = id
				startPos = i
			}
		} else {
			if currentID != -1 {
				oldFile := fileInfo{
					id:    currentID,
					start: startPos,
					end:   i - 1,
					size:  (i - startPos),
				}
				files = append(files, oldFile)
				currentID = -1
			}
		}
	}

	if currentID != -1 {
		oldFile := fileInfo{
			id:    currentID,
			start: startPos,
			end:   len(blocks) - 1,
			size:  (len(blocks) - startPos),
		}
		files = append(files, oldFile)
	}

	return blocks, files
}

func computeChecksum(blocks []rune) int {
	checksum := 0
	for i, b := range blocks {
		if b != '.' {
			id := int(b - '0')
			checksum += i * id
		}
	}
	return checksum
}

func Part1() int {
	lines, err := utils.ReadLines("input/day9.txt")
	if err != nil {
		panic(err)
	}

	if len(lines) == 0 {
		return 0
	}

	diskMap := lines[0]
	if len(diskMap) == 0 {
		return 0
	}

	blocks, _ := parseDiskMap(diskMap)

	for {
		dotIndex := -1
		for i, b := range blocks {
			if b == '.' {
				dotIndex = i
				break
			}
		}
		if dotIndex == -1 {
			break
		}

		hasFileAfter := false
		for i := dotIndex + 1; i < len(blocks); i++ {
			if blocks[i] != '.' {
				hasFileAfter = true
				break
			}
		}

		if !hasFileAfter {
			break
		}

		rightmostFileIndex := -1
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i] != '.' {
				rightmostFileIndex = i
				break
			}
		}

		if rightmostFileIndex == -1 {
			break
		}

		fileBlock := blocks[rightmostFileIndex]
		blocks[dotIndex] = fileBlock
		blocks[rightmostFileIndex] = '.'
	}

	return computeChecksum(blocks)
}

func Part2() int {
	lines, err := utils.ReadLines("input/day9.txt")
	if err != nil {
		panic(err)
	}

	if len(lines) == 0 {
		return 0
	}

	diskMap := lines[0]
	if len(diskMap) == 0 {
		return 0
	}

	blocks, files := parseDiskMap(diskMap)

	maxID := -1
	for _, f := range files {
		if f.id > maxID {
			maxID = f.id
		}
	}

	fileMap := make(map[int]*fileInfo)
	for i := range files {
		file := files[i]
		fileMap[file.id] = &files[i]
	}

	updateFilePositions := func(id int) {
		f := fileMap[id]
		f.start = -1
		f.end = -1
		count := 0
		for i, b := range blocks {
			if b != '.' {
				bid := int(b - '0')
				if bid == id {
					if f.start == -1 {
						f.start = i
					}
					f.end = i
					count++
				}
			}
		}
		f.size = f.end - f.start + 1
		if f.start == -1 {
			f.size = 0
		}
	}

	findLeftmostRun := func(fileStart, fsize int) (int, int) {
		runStart := -1
		runLength := 0
		bestStart := -1

		for i := 0; i < fileStart; i++ {
			if blocks[i] == '.' {
				if runStart == -1 {
					runStart = i
					runLength = 1
				} else {
					runLength++
				}
			} else {
				if runLength >= fsize {
					bestStart = runStart
					break
				}
				runStart = -1
				runLength = 0
			}
		}
		if runStart != -1 && runLength >= fsize && bestStart == -1 {
			bestStart = runStart
		}

		if bestStart == -1 {
			return -1, -1
		}
		return bestStart, bestStart + fsize - 1
	}

	for id := maxID; id >= 0; id-- {
		f := fileMap[id]
		if f == nil || f.size == 0 {
			continue
		}
		startRun, endRun := findLeftmostRun(f.start, f.size)
		if startRun == -1 {
			continue
		}

		fileChar := rune('0' + id)

		for i := f.start; i <= f.end; i++ {
			blocks[i] = '.'
		}
		for i := startRun; i <= endRun; i++ {
			blocks[i] = fileChar
		}

		updateFilePositions(id)
	}

	return computeChecksum(blocks)
}
