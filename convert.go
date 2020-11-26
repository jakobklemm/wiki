package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	path := "./database/philosophy.md"
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "](") && strings.Contains(line, ".org") {
			ending := strings.Replace(line, ".org)", ")", 1)
			lines[i] = strings.Replace(ending, "](", "](/database/", 1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
