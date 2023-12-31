package aod

import (
	"log"
	"os"
	"strings"
)

func Solve_d2() {

	score := 0
	content, err := os.ReadFile("aod/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)           // Splits by whitespace
		concatenated := strings.Join(parts, "") // Joins without spaces
		score += int((concatenated[1]-concatenated[0]-1)%3*3 + concatenated[1] - 87)

	}
	log.Println(score)
}
