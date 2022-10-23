package footprint

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Buffer struct {
	data []byte
	indices [][]int
}

func GenerateSeed(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func GenerateDictionaryBuffer() *Buffer {
	file, err := os.ReadFile("words_alpha.txt")

	if err != nil {
		log.Fatal(err)
	}

	bom := []byte{0xEF, 0xBB, 0xBF}

	if bytes.HasPrefix(file, bom) {
		file = file[len(bom):]
	}

	buffer := Buffer{data: file, indices: make([][]int, 0)}

	for i := 0; ; {
		j := bytes.IndexByte(file[i:], '\n')

		// EOF
		if j < 0 {
			if len(file[i:]) > 0 {
				buffer.indices = append(buffer.indices, []int{i, len(file)})
			}
			break
		}

		index := []int{i}

		j += i
		i = j + 1

		if j > 0 && file[j - 1] == '\r' {
			j--
		}

		index = append(index, j)
		buffer.indices = append(buffer.indices, index)
	}

	return &buffer
}

func (buffer *Buffer) ReadLine(line int) string {
	if line < 0 || line >= len(buffer.indices) {
		panic("Generated seed out of indices range")
	}

	return string((*buffer).data[(*buffer).indices[line][0]:(*buffer).indices[line][1]])
}

func GenerateFootprint(words int) string {
	buffer := GenerateDictionaryBuffer()

	var footprint []string

	for i := 0; i < words; i++ {
		var word string

		for len(word) > 2 {
			seed := GenerateSeed(len(buffer.indices))
			word = buffer.ReadLine(seed)
		}
		footprint = append(footprint, word)
	}

	return strings.Join(footprint, "-")
}
