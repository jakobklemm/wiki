package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	db_root := "/home/jeykey/documents/projects/wiki/database/"
	files, err := ioutil.ReadDir(db_root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		path := f.Name()
		input, err := ioutil.ReadFile(db_root + path)
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
		err = ioutil.WriteFile(db_root+path, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
