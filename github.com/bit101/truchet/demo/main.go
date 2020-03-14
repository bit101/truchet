package main

import (
	"log"
	"os/exec"
)

func main() {

	patterns := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"x",
		"y",
		"z",
		"0",
	}

	for _, p := range patterns {

		cmd := exec.Command("../truch/truch",
			"-i", "_original.png",
			"-p", p,
			"-t", "10",
			"-o", "./out/"+p+".png")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := exec.Command("montage",
		"-label", "%t",
		"_original.png", "out/*.png",
		"-geometry", "200x200+5+5",
		"-shadow",
		"patterns.png")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
