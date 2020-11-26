package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./database/hexen.md")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "](") && strings.Contains(line, ".org") {
			ending := strings.Replace(line, ".org)", ")", 1)
			lines[i] = strings.Replace(ending, "](", "](database/", 1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./database/hexen.md", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
